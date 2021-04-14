import React from "react";
import { useHistory } from "react-router-dom";

export default function ButtonAdd(props) {
  let history = useHistory();

  function handleClick(path) {
    history.push(path);
  }

  return (
    <div className="btn__add--wrapper">
      <button onClick={() => handleClick(props.path)} className="btn btn__add">
        <svg
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M6.85714 6.85714V0H9.14286V6.85714H16V9.14286H9.14286V16H6.85714V9.14286H0V6.85714H6.85714Z"
            fill="#70777F"
          />
        </svg>
      </button>
      <p className="btn__add--label">Click here to register a new bill</p>
    </div>
  );
}
