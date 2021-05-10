import React from "react";

function TotalSpent(props) {
  return (
    <div className="card card__color">
      <p className="card--title">Total spent this month</p>
      <p className="card--totalSpent">$ {props.total.toFixed(2)}</p>
    </div>
  );
}

export default TotalSpent;
