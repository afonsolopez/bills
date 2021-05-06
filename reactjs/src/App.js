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
    time_stamp: "",
    total: 0.0,
  });
  const [entries, setEntries] = useState([
    {
      company: "",
      price: 0.0,
    },
  ]);
  const [bigNumbers, setBigNumbers] = useState([
    {
      total: 0.0,
      remainingDays: 0,
    },
  ]);

  // Fetch data for the month graph
  useEffect(() => {
    let mounted = true;

    fetch(`${process.env.REACT_APP_API_PATH || ""}/by-day`, {
      // mode: 'no-cors',
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    }).then((response) => {
      if (mounted) {
        if (response.ok) {
          response.json().then((json) => {
            if (json !== null) {
              setState(json);
            }
            console.log(json);
          });
        }
      }
    });

    return function cleanup() {
      mounted = false;
    };
  }, []);

  // Fetch latest bills entries
  useEffect(() => {
    let mounted = true;

    fetch(`${process.env.REACT_APP_API_PATH || ""}/latest`, {
      // mode: 'no-cors',
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    }).then((response) => {
      if (mounted) {
        if (response.ok) {
          response.json().then((json) => {
            if (json !== null) {
              setEntries(json);
            }
            console.log(json);
          });
        }
      }
    });

    return function cleanup() {
      mounted = false;
    };
  }, []);

  // Fetch big numbers data
  useEffect(() => {
    let mounted = true;

    fetch(`${process.env.REACT_APP_API_PATH || ""}/big-numbers`, {
      // mode: 'no-cors',
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    }).then((response) => {
      if (mounted) {
        if (response.ok) {
          response.json().then((json) => {
            if (json !== null) {
              setBigNumbers(json);
            }
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
        <LatestBills bills={entries} />
        <div className="grid-container--box--top--double">
          <TotalSpent total={bigNumbers[0].total} />
          <RemainingDays days={bigNumbers[0].remainingDays} />
        </div>
      </div>
      <div className="grid-container--box--bottom">
        <MonthChart bills={state} />
      </div>
    </div>
  );
}

export default App;
