import React from "react";
import ButtonText from "../../atoms/btn-text";

function LatestBills(props) {
  console.log(props.bills);

  let rows;
  if (props.bills) {
    let billsList = props.bills.slice(0, 4);
    let billsTable = billsList.map((b, index) => (
      <tr key={index}>
        <td className="rowTitle">
          <p>{b.company}</p>
        </td>
        <td className="rowPrice">
          <p>{b.price.toFixed(2)}</p>
        </td>
      </tr>
    ));
    rows = billsTable;
  } else {
    rows = (
      <tr>
        <td></td>
        <td></td>
      </tr>
    );
  }

  return (
    <div className="card">
      <p className="card--title">Latest bills</p>
      <table>
        <thead>
          <tr>
            <th className="rowHeader">Company</th>
            <th className="rowHeader">Price ($)</th>
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </table>
      <ButtonText title="See more" path="/details" />
    </div>
  );
}

export default LatestBills;
