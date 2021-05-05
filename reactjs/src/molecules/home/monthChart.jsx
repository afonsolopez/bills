import React from "react";
import {
  LineChart,
  Line,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

function MonthChart(props) {
console.log(props.bills);
  return (
    <div className="card">
      <p className="card--title">Bills throught the month</p>
      <br />
      <ResponsiveContainer width="100%" height="100%">
        <LineChart
          data={props.bills}
          margin={{ top: 5, right: 20, bottom: 5, left: 0 }}
        >
          <Line
            type="monotone"
            dataKey="total"
            stroke="var(--variable)"
            strokeWidth={3}
          />
          <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
          <XAxis dataKey="time_stamp" />
          <YAxis />
          <Tooltip />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}

export default MonthChart;
