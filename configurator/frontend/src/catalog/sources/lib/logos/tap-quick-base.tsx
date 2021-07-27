import * as React from "react";

function Svg(props) {
  return (
    <svg
      id="Layer_1"
      xmlns="http://www.w3.org/2000/svg"
      height="100%"
      width="100%"
      xmlnsXlink="http://www.w3.org/1999/xlink"
      x="0px"
      y="0px"
      viewBox="0 0 200 200"
      enableBackground="new 0 0 200 200"
      xmlSpace="preserve"
      {...props}
    >
      <rect fill="#75459F" width={200} height={200} />
      <g>
        <path
          fill="#FFFFFF"
          d="M147.7,147.2c-0.4,0.6-1.1,1-1.8,1.1h-23.2c-0.4,0.1-0.7-0.2-0.8-0.5c0-0.2,0-0.4,0.1-0.6l24.4-43.2 c0.3-0.7,0.3-1.5,0-2.1l-24.5-45c-0.2-0.3-0.2-0.7,0.1-0.9c0.2-0.1,0.4-0.2,0.5-0.1h23.3c0.7,0,1.4,0.4,1.8,1.1l24.5,45 c0.3,0.7,0.3,1.5,0,2.1L147.7,147.2z"
        />
        <path
          fill="#FFFFFF"
          d="M108.1,134.7c8.7-9.2,14.7-20.1,14.7-34.4c0.4-26.4-20.6-48.1-47-48.6c-26.4-0.4-48.1,20.6-48.6,47 c0,0.5,0,1.1,0,1.6c0,25.6,20,47.5,45.4,47.9h42.9c0.4,0,0.6-0.4,0.6-0.8c0-0.1,0-0.2-0.1-0.3L108.1,134.7z M75,125.6 c-13.9,0-25.2-11.3-25.2-25.2S61.1,75.2,75,75.2s25.2,11.3,25.2,25.2S88.9,125.6,75,125.6L75,125.6z"
        />
        <path
          fill="#FFFFFF"
          d="M167.2,144.8h-1.3v3.5H165v-3.5h-1.3v-0.7h3.4L167.2,144.8L167.2,144.8z M172.7,148.2h-0.8v-3.1l0,0 l-1.2,3.1h-0.6l-1.2-3.1l0,0v3.1h-0.8v-4.2h1.2l1.2,3l1.2-3h1.2L172.7,148.2L172.7,148.2z"
        />
      </g>
    </svg>
  );
}

export default Svg;