import React, { useState, useEffect } from "react";
import "./billForm.css";

function BillForm() {
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");
  const [company, setCompany] = useState("");
  const [tag, setTag] = useState("");
  const [date, setDate] = useState("");
  const [companies] = React.useState([
    {
      label: "Luke Skywalker",
      value: "Luke Skywalker",
    },
    { label: "C-3PO", value: "C-3PO" },
    { label: "R2-D2", value: "R2-D2" },
  ]);
  const [tags] = React.useState([
    {
      label: "Luke Skywalker",
      value: "Luke Skywalker",
    },
    { label: "C-3PO", value: "C-3PO" },
    { label: "R2-D2", value: "R2-D2" },
  ]);

  const [isNewCompany, setIsNewCompany] = useState(false);
  const [companyInput, setCompanyInput] = useState(<></>);
  const [isNewTag, setIsNewTag] = useState(false);
  const [tagInput, setTagInput] = useState(<></>);

  const handleClick = () => {
    setCompany("");
    setIsNewCompany(!isNewCompany);
  }
    const handleClick2 = () => {
    setTag("");
    setIsNewTag(!isNewTag);
  }
  const formId = "something";

  const handleSubmit = (evt) => {
    evt.preventDefault();

    fetch(`${process.env.REACT_APP_API_PATH || ""}/createNewBill`, {
      method: "post",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: JSON.stringify({
        title: title,
        price: price,
        isNewCompany: isNewCompany,
        company: company,
        isNewTag: isNewTag,
        tag: tag,
        date: date,
      }),
    });
    console.log("bill data sent");
  };

  useEffect(() => {
    let mounted = true;
    let thisCompanyInput;
    let thisTagInput;
    if (mounted) {
      if (isNewCompany) {
        thisCompanyInput = (
          <>
            <label htmlFor="newCompany">Register a new company</label>
            <br />
            <input
              placeholder="New company name..."
              type="text"
              id="newCompany"
              name="newCompany"
              value={company}
              onChange={(e) => setCompany(e.target.value)}
            />
          </>
        );
        setCompanyInput(thisCompanyInput);
      } else {
        thisCompanyInput = (
          <>
            <label htmlFor="company">Select company from list</label>
            <br />
            <select 
            name="company"
            // value={company}
            value={"DEFAULT"}
            onChange={e => setCompany(e.currentTarget.value)}
            >
              <option value="DEFAULT" disabled hidden>Please Choose...</option>
              {companies.map((c) => (
                <option key={c.value} value={c.value}>
                  {c.label}
                </option>
              ))}
            </select>
          </>
        );
        setCompanyInput(thisCompanyInput);
      }

      if (isNewTag) {
        thisTagInput = (
          <>
            <label htmlFor="newTag">Register a new tag</label>
            <br />
            <input
              placeholder="New tag name..."
              type="text"
              id="newTag"
              name="newTag"
              value={tag}
              onChange={(e) => setTag(e.target.value)}
            />
          </>
        );
        setTagInput(thisTagInput);
      } else {
        thisTagInput = (
          <>
            <label htmlFor="newTag">Select a new tag</label>
            <br />
            <select 
            name="newTag"
            // value={tag}
            value={"DEFAULT"}
            onChange={e => setTag(e.currentTarget.value)}
            >
               <option value="DEFAULT" disabled hidden>Please Choose...</option>
              {tags.map((t) => (
                <option key={t.value} value={t.value}>
                  {t.label}
                </option>
              ))}
            </select>
          </>
        );
        setTagInput(thisTagInput);
      }
    }
    return () => {
      mounted = false;
    };
  }, [company, tag, isNewCompany, isNewTag, tags, companies]);

  return (
    <div className="grid-container--form--box">
      <div className="grid-container--form">
        <div className="card">
          <form id={formId} onSubmit={handleSubmit}>
            <div>
              <div>
                <p className="card--title">Bill info</p>
              </div>
              <div>
                <div>
                  <label htmlFor="title">Title</label>
                  <br />
                  <input
                    placeholder="Give it a custom title..."
                    type="text"
                    id="title"
                    name="title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                  />
                </div>
                <div>
                  <label htmlFor="price">Price (R$)</label>
                  <br />
                  <input
                    placeholder="0.00"
                    type="text"
                    id="price"
                    name="price"
                    value={price}
                    onChange={(e) => setPrice(e.target.value)}
                  />
                </div>
              </div>
            </div>

            <div>
              <div>
                <p className="card--title">Company</p>
              </div>
              <div>
                <div>
                <p>Create new company</p>
                  <label className="switch">
                    <input
                    type="checkbox"
                    id="isNewContact"
                    name="isNewContact"
                    onClick={handleClick}
                    checked={isNewCompany}
                    onChange={(e) => setIsNewCompany(e.target.checked)}
                    />
                    <span className="slider round"></span>
                  </label>
                </div>
                <div>{companyInput}</div>
              </div>
            </div>

            <div>
              <div>
                <p className="card--title">Tags</p>
              </div>
              <div>
                <div>
                  <p>Create new tag</p>
                <label className="switch">
                    <input
                    type="checkbox"
                    id="isNewTag"
                    name="isNewTag"
                    onClick={handleClick2}
                    checked={isNewTag}
                    onChange={(e) => setIsNewTag(e.target.checked)}
                    />
                    <span className="slider round"></span>
                  </label>
                </div>
                <div>{tagInput}</div>
              </div>
            </div>

            <div>
              <div className="">
                <p className="card--title">Payment date</p>
              </div>
              <div>
                <div>
                  <label htmlFor="date">Select the payment date</label>
                  <br />
                  <input
                    type="text"
                    id="date"
                    value={date}
                    onChange={(e) => setDate(e.target.value)}
                  />
                </div>
              </div>
            </div>
          </form>
        </div>

        <div className="card">
          <div>
            <p className="card--title">Checkout</p>
            <ul>
              <li>{title}</li>
              <li>{price}</li>
              <li>{company}</li>
              <li>{tag}</li>
              <li>{date}</li>
            </ul>
          </div>
          <button
            type="submit"
            form={formId}
            className="btn btn__solid btn__full"
          >
            Register new bill
          </button>
        </div>
      </div>
    </div>
  );
}

export default BillForm;
