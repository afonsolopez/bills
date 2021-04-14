// Import React and React Hooks
import React, { useState, useEffect } from "react";

import "./App.css";

import RemainingDays from "./molecules/home/remainingDays";
import TotalSpent from "./molecules/home/totalSpent";
import LatestBills from "./molecules/home/latestBills";
import NewBill from "./molecules/home/newBill";
import MonthChart from "./molecules/home/monthChart";

function App() {
  // Define storage for data
  const [state, setState] = useState({
    bills: [{ company: "", price: 0.0, year: 0, month: 0, day: 0 }],
    total: 0.0,
    remainingDays: 0,
  });

  useEffect(() => {
    let mounted = true;

    fetch(`${process.env.REACT_APP_API_PATH || ""}/last`, {
      // mode: 'no-cors',
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    }).then((response) => {
      if (mounted) {
        if (response.ok) {
          response.json().then((json) => {
            setState(json);
            console.log(json);
          });
        }
      }
    });

    return function cleanup() {
      mounted = false;
    };
  }, []);

  return (
    <div className="grid-container--box">
      <div className="grid-container--box--top">
        <NewBill />
        <LatestBills bills={state.bills} />
        <div className="grid-container--box--top--double">
          <TotalSpent total={state.total} />
          <RemainingDays days={state.remainingDays} />
        </div>
      </div>
      <div className="grid-container--box--bottom">
        <MonthChart bills={state.bills} />
      </div>
    </div>
  );
}

export default App;
