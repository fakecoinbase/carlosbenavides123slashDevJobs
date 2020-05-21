import React, { useState } from "react";
import "./JobForm.scss";
import Button from "../Button/Button";
import axios from "axios";

function JobForm() {
  const [company, setCompany] = useState("");
  const [careerPage, setCareerPage] = useState("");
  const [cloudinary, setCloudinary] = useState("");

  function handleJobSubmit(e) {
    var data = JSON.stringify({
      company_name: company,
      company_website: careerPage,
      cloudinary: cloudinary
    });

    axios
      .post(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/`, data)
      .then(res => {
        setCompany("");
        setCareerPage("");
        setCloudinary("");
      })
      .catch(err => {
        console.log(err);
      });
  }

  return (
    <div style={{ display: "flex", justifyContent: "center" }}>
      <div className="card">
        <h3>Add a Company and I'll try my best to get data!</h3>

        <input
          className="JobFormInput"
          placeholder="Company Name"
          onChange={e => setCompany(e.target.value)}
          value={company}
        />

        <input
          className="JobFormInput"
          placeholder="Career Page"
          onChange={e => setCareerPage(e.target.value)}
          value={careerPage}
        />

        <input
          className="JobFormInput"
          placeholder="Cloudinary Link"
          onChange={e => setCloudinary(e.target.value)}
          value={cloudinary}
        />
        <Button btype="joblink" onClick={handleJobSubmit} text={"Submit!"} />
      </div>
    </div>
  );
}

export default JobForm;
