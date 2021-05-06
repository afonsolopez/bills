// Import React and React Hooks
import React, { useState, useEffect } from "react";

import ".././App.css";
import TotalSpent from "../molecules/home/totalSpent";
import RemainingDays from "../molecules/home/remainingDays";
import MonthChart from "../molecules/month/monthChart";

function MonthTrack() {
  // Define storage for data
  const [bigNumbers, setBigNumbers] = useState([
    {
      total: 0.0,
      remainingDays: 0,
    },
  ]);
  const [state, setState] = useState({
    time_stamp: "",
    total: 0.0,
  });
  const [entries, setEntries] = useState([
    {
      id: 0,
      title: "",
      company: "",
      price: 0.0,
      tag: "",
      time_stamp: "",
    },
  ]);
  // Handle when to show or not the modal
  const [modalOpen, setModalOpen] = useState(false);

  // Change the display mode of the modal based on it state
  const checkModal = (modalOpen) => {
    if (modalOpen) {
      return { display: "block" };
    }
    return { display: "none" };
  };

  let modalStyle = checkModal(modalOpen);

  // Close and reset form button handler
  const openModal = (e, id) => {
    e.preventDefault();
    console.log(id);
    setModalOpen(true);
  };

  // Close and reset form button handler
  const closeModal = (e) => {
    e.preventDefault();
    setModalOpen(false);
  };

    // Fetch data for the month graph
    useEffect(() => {
      let mounted = true;
  
      fetch(`${process.env.REACT_APP_API_PATH || ""}/by-tag`, {
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

  // Fetch entries data
  useEffect(() => {
    let mounted = true;
    fetch(`${process.env.REACT_APP_API_PATH || ""}/month`, {
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

  let rows;
  if (entries) {
    let billsList = entries;
    let billsTable = billsList.map((b, index) => (
      <tr key={index}>
        <td className="rowTitle">
          <p>{b.company}</p>
        </td>
        <td className="rowTitle size16">
          <p>{b.title}</p>
        </td>
        <td className="rowPrice">
          <p>{b.price.toFixed(2)}</p>
        </td>
        <td className="rowTitle">
          <p>{b.tag}</p>
        </td>
        <td className="rowTitle rowDel">
          <svg
            className="deleteBucket"
            onClick={(e) => openModal(e, b.id)}
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M17 6H22V8H20V21C20 21.2652 19.8946 21.5196 19.7071 21.7071C19.5196 21.8946 19.2652 22 19 22H5C4.73478 22 4.48043 21.8946 4.29289 21.7071C4.10536 21.5196 4 21.2652 4 21V8H2V6H7V3C7 2.73478 7.10536 2.48043 7.29289 2.29289C7.48043 2.10536 7.73478 2 8 2H16C16.2652 2 16.5196 2.10536 16.7071 2.29289C16.8946 2.48043 17 2.73478 17 3V6ZM18 8H6V20H18V8ZM13.414 14L15.182 15.768L13.768 17.182L12 15.414L10.232 17.182L8.818 15.768L10.586 14L8.818 12.232L10.232 10.818L12 12.586L13.768 10.818L15.182 12.232L13.414 14ZM9 4V6H15V4H9Z"
              fill="var(--text-muted)"
            />
          </svg>
        </td>
      </tr>
    ));
    rows = billsTable;
  } else {
    rows = (
      <tr>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
      </tr>
    );
  }
  return (
    <>
      <div className="grid-container--box">
        <div className="grid-container--box--top__2">
            <MonthChart bills={state} />
          <div className="grid-container--box--top--double">
          <TotalSpent total={bigNumbers[0].total} />
          <RemainingDays days={bigNumbers[0].remainingDays} />
          </div>
        </div>
        <div className="grid-container--box--bottom">
          <div className="card">
            <p className="card--title">Bills throught the month</p>
            <br />
            <div className="tableWrapper">
              <table>
                <thead>
                  <tr>
                    <th className="rowHeader">Company</th>
                    <th className="rowHeader">Title</th>
                    <th className="rowHeader">Price ($)</th>
                    <th className="rowHeader">Tag</th>
                    <th className="rowHeader rowDel">Delete</th>
                  </tr>
                </thead>
                <tbody>{rows}</tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <div id="myModal" className="modal" style={modalStyle}>
        <div className="modal-content">
          <span onClick={(e) => closeModal(e)} className="close">
            &times;
          </span>
          <p>Modal working!</p>
          <br />
          <button className="btn btn__solid" onClick={(e) => closeModal(e)}>
            Return
          </button>
        </div>
      </div>
    </>
  );
}

export default MonthTrack;
