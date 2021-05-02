import React from "react";
import {
  ScatterChart,
  Scatter,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
} from "recharts";

import "./monthChart.css";

function MonthPie(props) {
  console.log(props.bills);
  return (
    <div className="card">
      <p className="card--title">Tags</p>
      <br />
      <div className="pieChart">
        <ScatterChart
          width={432}
          height={180}
          margin={{ top: 20, right: 20, bottom: 10, left: 10 }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="day" name="day" unit="" />
          <YAxis dataKey="price" name="price" unit="" />
          <Tooltip cursor={{ strokeDasharray: "3 3" }} />
          <Legend />
          <Scatter name="Bill" data={props.bills} fill="var(--variable)" />
        </ScatterChart>
      </div>
    </div>
  );
}

export default MonthPie;
