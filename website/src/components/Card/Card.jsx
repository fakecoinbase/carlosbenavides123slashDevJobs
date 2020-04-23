import React from "react";
import "./Card.scss";
import Button from "../Button/Button"
import moment from 'moment'

function Card({job_uuid, job_title, job_link, job_location, job_posted, job_found, company_name, cloudinary, level}) {
  function handleJobClick(e) {
    window.open(job_link, '_blank')
  }
  return (
      <div className="card" onClick={handleJobClick}>
        <span>
          <img
            src={
              cloudinary
            }
            className="icon"
            alt="company"
          />
          {job_location}
        </span>
        <br />
        <div className="info">
          {company_name}
          <br></br>
          {job_title}
        </div>
 
        <span className="apply_info">
        {job_posted != 0 && <>Posted on: {moment.unix(job_posted).format("MM/DD/YYYY")}</>}
        {job_posted == 0 && <span>Job Found on: {moment.unix(job_found).format("MM/DD/YYYY")}</span>}
        <Button btype="joblink" text={"Apply"}/>

        </span>


      </div>
  );
}

export default Card;
