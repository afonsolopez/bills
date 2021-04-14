import React from "react";

function RemainingDays(props) {
  console.log(props.days);
  return (
    <div className="card">
      <p className="card--title">Remaining days</p>
      <p className="card--remainingDays">
        <span>{props.days}</span> days left
      </p>
    </div>
  );
}

export default RemainingDays;
