import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Sidebar from "./molecules/sidebar";
import BillForm from "./templates/billForm";
import MonthTrack from "./templates/monthTrack";
import Info from "./templates/info";

ReactDOM.render(
  <React.StrictMode>
    <Router>
      <div className="grid-container">
        <Sidebar />

        <Switch>
        <Route path="/info">
            <Info />
          </Route>
        <Route path="/month">
            <MonthTrack />
          </Route>
          <Route path="/new">
            <BillForm />
          </Route>
          <Route path="/">
            <App />
          </Route>
        </Switch>
      </div>
    </Router>
  </React.StrictMode>,
  document.getElementById("root")
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
