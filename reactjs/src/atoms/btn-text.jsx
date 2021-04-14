import React from "react";
import { useHistory } from "react-router-dom";

export default function ButtonText(props) {
  let history = useHistory();

  function handleClick(path) {
    history.push(path);
  }


  return (
    <button className="btn btn__text" onClick={() => handleClick(props.path)} >
      {props.title}
    </button>
  );
}
