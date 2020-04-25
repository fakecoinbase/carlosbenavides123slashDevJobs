import React, { useState, useEffect, useRef } from "react";
import Button from "../Button/Button";
import {isMobile} from 'react-device-detect';
import "./Dropdown.scss";
// import Dropdown from "react-dropdown";
import "react-dropdown/style.css";
import Select from "react-select";

function NavDropdown({ setCompany, setLocation, setExperience, companyDropdown }) {

  const [isDeviceMobile, setIsDeviceMobile] = useState(false)
  
  useEffect(() => {
    setIsDeviceMobile(isMobile)
  }, [isMobile])

  function handleLocation(e) {
    if (e === null) {
      setLocation('')
    }
    if (e !== null) {
      setLocation(e["value"])
    }
  }

  function handleExperience(e) {
    if (e === null) {
      setExperience('')
    }
    if (e !== null) {
      setExperience(e["value"])
    }
  }

  function handleCompany(e) {
    if (e === null) {
      setCompany('')
    }
    if (e !== null) {
      setCompany(e["value"])
    }
  }

  let location = [{ value: "Los Angeles", label: "Los Angeles"}, { value: "Venice", label: "Venice"}, { value:"San Francisco", label:"San Francisco"}, { value:"New York", label:"New York"}, { value:"Denver", label:"Denver"}, { value:"Seattle", label:"Seattle"}, { value:"Bellevue", label:"Bellevue" }, { value:"Boulder", label:"Boulder" }];
  let experience = [
    { value: "Intern", label: "Intern" },
    { value: "Entry", label: "Entry" },
    { value: "Mid", label: "Mid" },
    { value: "Senior", label: "Senior" },
    { value: "Manager", label: "Manager" },
  ];

  return (
    <div className={isDeviceMobile ? "" : "navDropdown"}>
      <Select
        name="colors"
        options={location}
        className={isDeviceMobile ? "" : "navDropdownRegular"}
        classNamePrefix="select"
        placeholder="Location"
        onChange={handleLocation}
        isClearable={true}
      />
      <Select
        options={experience}
        className={isDeviceMobile ? "" : "navDropdownRegular"}
        classNamePrefix="select"
        placeholder="Experience"
        onChange={handleExperience}
        isClearable={true}
      />
      <Select
        options={companyDropdown}
        className={isDeviceMobile ? "" : "navDropdownRegular"}
        classNamePrefix="select"
        placeholder="Companies"
        onChange={handleCompany}
        isClearable={true}
      />
    </div>
  );
}

export default NavDropdown;
