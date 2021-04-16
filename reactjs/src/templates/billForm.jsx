import React, { useState, useEffect } from "react";
import "./billForm.css";

function BillForm() {
  
  // Here we handle all inputs data
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");
  const [company, setCompany] = useState("");
  const [tag, setTag] = useState("");
  const [date, setDate] = useState("");
  const [companies, setCompanies] = useState([
    {
      label: "",
      value: "",
    },
  ]);
  const [tags, setTags] = useState([
    {
      label: "",
      value: "",
    },
  ]);

  // Those hooks handles a bool value fow new "tags" or "companies" inputs
  const [isNewCompany, setIsNewCompany] = useState(false);
  const [isNewTag, setIsNewTag] = useState(false);

  // Stores the rendered options from each select input
  const [companyInput, setCompanyInput] = useState(<></>);
  const [tagInput, setTagInput] = useState(<></>);
  
  // Handles a bool value if the for is ready to be send
  const [isDone, setIsDone] = useState(true);

  // 
  const formId = "something";

  // Functions to turn the "new company" and "new tag" inputs on/off
  const handleClick = () => {
    setCompany("");
    setIsNewCompany(!isNewCompany);
  };
  const handleClick2 = () => {
    setTag("");
    setIsNewTag(!isNewTag);
  };

  

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

    fetch(`${process.env.REACT_APP_API_PATH || ""}/getAllCompanies`, {
      // mode: 'no-cors',
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    }).then((response) => {
      if (mounted) {
        if (response.ok) {
          response.json().then((json) => {
            setCompanies(json);
            console.log(json);
          });
        }
      }
    });

    return function cleanup() {
      mounted = false;
    };
  }, []);

  useEffect(() => {
    let mounted = true;

    fetch(`${process.env.REACT_APP_API_PATH || ""}/getAllTags`, {
      // mode: 'no-cors',
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    }).then((response) => {
      if (mounted) {
        if (response.ok) {
          response.json().then((json) => {
            setTags(json);
            console.log(json);
          });
        }
      }
    });

    return function cleanup() {
      mounted = false;
    };
  }, []);

  useEffect(() => {
    let mounted = true;
    let thisCompanyInput;
    let thisTagInput;
    if (mounted) {
      if (isNewCompany) {
        thisCompanyInput = (
          <div>
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
          </div>
        );
        setCompanyInput(thisCompanyInput);
      } else {
        thisCompanyInput = (
          <div>
            <label htmlFor="company">Select company from list</label>
            <br />
            <select
              name="company"
              // value={company}
              value={company || "DEFAULT"}
              onChange={(e) => setCompany(e.target.value)}
            >
              <option value="DEFAULT" disabled hidden>
                Please Choose...
              </option>
              {companies.map((c) => (
                <option key={c.value} value={c.value}>
                  {c.label}
                </option>
              ))}
            </select>
          </div>
        );
        setCompanyInput(thisCompanyInput);
      }

      if (isNewTag) {
        thisTagInput = (
          <div>
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
          </div>
        );
        setTagInput(thisTagInput);
      } else {
        thisTagInput = (
          <div>
            <label htmlFor="newTag">Select a new tag</label>
            <br />
            <select
              name="newTag"
              value={tag || "DEFAULT"}
              onChange={(e) => setTag(e.target.value)}
            >
              <option value="DEFAULT" disabled hidden>
                Please Choose...
              </option>
              {tags.map((t) => (
                <option key={t.value} value={t.value}>
                  {t.label}
                </option>
              ))}
            </select>
          </div>
        );
        setTagInput(thisTagInput);
      }
    }
    return () => {
      mounted = false;
    };
  }, [company, tag, isNewCompany, isNewTag, tags, companies]);

  const regxPrice = /^[0-9]+(\.[0-9]{1,2})?$/gm;
  const regxDate = /([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))/;

  let checkTitleInput = title.length <= 0;
  let checkPriceinput = !regxPrice.test(price.replace(",", "."));
  let checkDateInput = !regxDate.test(date);
  let checkCompanyInput = !company;
  let checkTagInput = !tag;
  function colorValidator(condition) {
    if (condition) {
      return {
        color: "var(--text-muted)",
      };
    } else {
      return {
        color: "var(--text-main)",
      };
    }
  }

  let titleValidator = colorValidator(checkTitleInput);
  let priceValidator = colorValidator(checkPriceinput);
  let dateValidator = colorValidator(checkDateInput);
  let companyValidator = colorValidator(checkCompanyInput);
  let tagValidator = colorValidator(checkTagInput);

  useEffect(() => {
    let mounted = true;
    if (mounted) {
      if (
        !checkTitleInput &&
        !checkPriceinput &&
        !checkDateInput &&
        !checkCompanyInput &&
        !checkTagInput
      ) {
        setIsDone(false);
      }
    }

    return function cleanup() {
      mounted = false;
    };
  }, [
    checkTitleInput,
    checkPriceinput,
    checkDateInput,
    checkCompanyInput,
    checkTagInput,
  ]);

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
                  <label htmlFor="price">Price ($)</label>
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
                <div className="card--switch-wrapper">
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
                <div className="card--switch-wrapper">
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
                    placeholder="yyyy-mm-dd"
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
            <hr />
            <br />
            <p style={titleValidator}>
              <b>Title:</b> <br />
              {title}
            </p>
            <br />
            <p style={priceValidator}>
              <b>Price:</b>
              <br /> $ {price.replace(",", ".")}
            </p>
            <br />
            <p style={companyValidator}>
              <b>Company:</b> <br />
              {company}
            </p>
            <br />
            <p style={tagValidator}>
              <b>Tags:</b> <br />
              {tag}
            </p>
            <br />
            <p style={dateValidator}>
              <b>Date:</b> <br /> {date.toString().replace(/-/g, "/")}
            </p>
          </div>
          <button
            type="submit"
            form={formId}
            disabled={isDone}
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
