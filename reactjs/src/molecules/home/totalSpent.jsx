import React from "react";

function TotalSpent(props) {
  return (
    <div className="card card__color">
      <p className="card--title">Total spent this month <br/> {props.monthName || ""} </p>
      <p className="card--totalSpent">$ {props.total.toFixed(2)}</p>
    </div>
  );
}

export default TotalSpent;
