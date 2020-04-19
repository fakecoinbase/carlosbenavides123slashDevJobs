import React, { useState, useEffect, useRef } from "react";
import Button from "../Button/Button";
import "./Dropdown.scss";
// import Dropdown from "react-dropdown";
import "react-dropdown/style.css";
import Select from "react-select";

function NavDropdown({ setCompany, setLocation, setExperience }) {

  function handleLocation(e) {
    if (e === null) {
      setLocation('')
    }
    if (e !== null && e[0]) {
      setLocation(e[0]["value"])
    }
  }

  function handleExperience(e) {
    console.log("event")
    console.log(e)
    if (e === null) {
      console.log("null")
      setExperience('')
    }
    if (e !== null && e) {
      console.log(e["value"])
      setExperience(e["value"])
    }
  }

  function handleCompany(e) {
    console.log("event")
    console.log(e)
    if (e === null) {
      console.log("null")
      setCompany('')

      console.log("ye")
    }
    if (e !== null && e) {
      console.log(e["value"])
      setCompany(e["value"])
    }
  }

  let location = [{ value: "Los Angeles", label: "Los Angeles" }];
  let companies = [
    { value: "Honey", label: "Honey" },
    { value: "Pinterest", label: "Pinterest" }
  ];
  let experience = [
    { value: "Intern", label: "Internship" },
    { value: "Entry", label: "Entry" }
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
        options={companies}
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
