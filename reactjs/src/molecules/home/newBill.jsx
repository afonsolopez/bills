import React from "react";
import ButtonAdd from "../../atoms/btn-add";

function NewBill() {
  const flexCenter = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    flexDirection: "column",
    height: "100%",
    width: "100%",
  };

  return (
    <div className="card">
      <p className="card--title">Insert new bill</p>
      <div style={flexCenter}>
        <ButtonAdd path="/new" />
      </div>
    </div>
  );
}

export default NewBill;
