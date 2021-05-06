import React from "react";
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';

import "./monthChart.css";

function MonthPie(props) {
  console.log(props.bills);
  return (
    <div className="card">
      <p className="card--title">Tags</p>
      <br />
      <div className="pieChart">
        <BarChart
          width={426}
          height={192}
          data={props.bills}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="tag" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Bar dataKey="total" fill="var(--variable)" />
        </BarChart>
      </div>
    </div>
  );
}

export default MonthPie;
