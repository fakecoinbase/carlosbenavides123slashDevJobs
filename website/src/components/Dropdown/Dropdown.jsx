import React, { useState, useEffect, useRef } from "react";
import Button from "../Button/Button";
import "./Dropdown.scss";
// import Dropdown from "react-dropdown";
import "react-dropdown/style.css";
import Select from "react-select";

function NavDropdown({ setCompany, setLocation, setExperience, companyDropdown }) {

  console.log(companyDropdown)
  function handleLocation(e) {
    console.log("location event", e)
    if (e === null) {
      setLocation('')
      console.log("location event 1")
    }
    if (e !== null) {
      setLocation(e["value"])
      console.log("location event 2")
    }
  }

  function handleExperience(e) {
    console.log("event")
    console.log(e)
    if (e === null) {
      console.log("null")
      setExperience('')
    }
    if (e !== null) {
      console.log(e["value"], "experience event 2")
      setExperience(e["value"])
    }
  }

  function handleCompany(e) {
    console.log("event")
    console.log(e)
    if (e === null) {
      console.log("null")
      setCompany('')
    }
    if (e !== null) {
      console.log(e["value"])
      setCompany(e["value"])
    }
  }

  let location = [{ value: "Los Angeles", label: "Los Angeles"}, { value:"San Francisco", label:"San Francisco"}];
  let experience = [
    { value: "Intern", label: "Intern" },
    { value: "Entry", label: "Entry" },
    { value: "Mid", label: "Mid" },
    { value: "Senior", label: "Senior" },
    { value: "Manager", label: "Manager" },

  ];

  return (
    <div className="navDropdown">
      <Select
        name="colors"
        options={location}
        className="dropdown"
        classNamePrefix="select"
        placeholder="Location"
        onChange={handleLocation}
        isClearable={true}
      />
      <Select
        options={experience}
        className="dropdown"
        classNamePrefix="select"
        placeholder="Experience"
        onChange={handleExperience}
        isClearable={true}
      />
      <Select
        options={companyDropdown}
        className="dropdown"
        classNamePrefix="select"
        placeholder="Companies"
        onChange={handleCompany}
        isClearable={true}
      />
    </div>
  );
}

export default NavDropdown;
