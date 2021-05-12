import React from "react";

import ".././App.css";

function Info() {
  return (
    <div className="grid-container--box--info">
      <div className="card card">
        <p className="card--title">Application info</p>
        <br />
        <div className="align-card">
            <b>Technical info:</b>
            <p>Software version: <b>0.1a</b></p>
            <p>License: GNU General Public License v3.0</p>

            <br/>
            <b>About me</b>
            <p>Follow me on Twitter as <strong>@hiafonsolopez</strong>, or GitHub as <strong>@afonsolopez</strong>.</p>
            <p>But if you whant to find more information about me, check out <strong>afonsolopez.com</strong>.</p>
            <br/>
            
            <p>Designed and developed with ❤️ in São Paulo/Brazil</p>
        </div>
      </div>
    </div>
  );
}

export default Info;
