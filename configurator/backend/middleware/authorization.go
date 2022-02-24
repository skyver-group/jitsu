package middleware

import (
	"context"

	"github.com/jitsucom/jitsu/server/logging"

	"github.com/gin-gonic/gin"
	"github.com/jitsucom/jitsu/configurator/common"
	"github.com/jitsucom/jitsu/configurator/openapi"
	"github.com/jitsucom/jitsu/configurator/storages"
	"github.com/pkg/errors"
)

var (
	errUnauthorized        = errors.New("unauthorized")
	errServerTokenMismatch = errors.New("server token mismatch")
	errTokenMismatch       = errors.New("token mismatch")
)

const (
	authorityKey = "__authority"
)

type Authority struct {
	Token      string
	UserID     string
	IsAdmin    bool
	ProjectIDs common.StringSet
}

func (a *Authority) IsAnonymous() bool {
	return a.UserID == ""
}

func (a *Authority) Allow(projectID string) bool {
	return a.IsAdmin || a.ProjectIDs[projectID]
}

type Authorizator interface {
	Authorize(ctx context.Context, token string) (*Authority, error)
	FindAnyUserID(ctx context.Context) (string, error)
}

type AuthorizationInterceptor struct {
	ServerToken    string
	Authorizator   Authorizator
	Configurations *storages.ConfigurationsService
	IsSelfHosted   bool
}

func (i *AuthorizationInterceptor) Intercept(ctx *gin.Context) {
	_, clusterAdminScope := ctx.Get(openapi.ClusterAdminAuthScopes)
	_, managementScope := ctx.Get(openapi.ConfigurationManagementAuthScopes)
	if !clusterAdminScope && !managementScope {
		return
	}

	var authority *Authority
	if token := GetToken(ctx); ctx.IsAborted() {
		return
	} else if i.ServerToken == token {
		authority = &Authority{
			Token:   token,
			IsAdmin: true,
		}

		if !managementScope {
			ctx.Set(authorityKey, authority)
			return
		}

		if userID, err := i.Authorizator.FindAnyUserID(ctx); err != nil {
			invalidToken(ctx, errTokenMismatch)
			return
		} else {
			authority.UserID = userID
		}
	} else if clusterAdminScope {
		logging.SystemErrorf("server request [%s] with [%s] token has been denied: token mismatch", ctx.Request.URL.String(), token)
		invalidToken(ctx, errServerTokenMismatch)
		return
	} else if auth, err := i.Authorizator.Authorize(ctx, token); err != nil {
		logging.Errorf("failed to authenticate with token %s: %s", token, err)
		invalidToken(ctx, errTokenMismatch)
		return
	} else {
		authority = auth
		authority.Token = token
	}

	authority.ProjectIDs = make(common.StringSet)

	if managementScope {
		if !authority.IsAnonymous() {
			if projectIDs, err := i.Configurations.GetUserProjects(authority.UserID); err != nil {
				logging.Warnf("failed to get projects for user %s: %s", authority.UserID, err)
			} else {
				authority.ProjectIDs.AddAll(projectIDs...)
			}
		}

		if i.IsSelfHosted {
			authority.ProjectIDs.Add(storages.TelemetryGlobalID)
		}
	}

	ctx.Set(authorityKey, authority)
}

func (i *AuthorizationInterceptor) ManagementWrapper(body gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(openapi.ConfigurationManagementAuthScopes, "")
		i.Intercept(ctx)
		if !ctx.IsAborted() {
			body(ctx)
		}
	}
}

func GetAuthority(ctx *gin.Context) (*Authority, error) {
	if value, ok := ctx.Get(authorityKey); !ok {
		return nil, errUnauthorized
	} else if authority, ok := value.(*Authority); !ok {
		return nil, errors.Errorf("unexpected authority %+v", authority)
	} else {
		return authority, nil
	}
}
