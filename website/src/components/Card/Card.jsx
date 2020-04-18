import React from "react";
import "./Card.scss";
import Button from "../Button/Button"
import moment from 'moment'

function Card({job_uuid, job_title, job_link, job_posted, job_found, company_name, cloudinary, level}) {

  var res = job_uuid.split("_%_")
  var location = res[2].replace(/%/g, ' ')

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
          {location}
        </span>
        <br />
        <div className="info">
          {company_name}
          <br></br>
          {job_title}
        </div>
 
        <span className="apply_info">
        {job_posted && <>Posted on: {moment.unix(job_posted).format("MM/DD/YYYY")}</>
        }
        {!job_posted && <span>Job Found on: {moment.unix(job_found)}</span>}
        <Button btype="joblink" text={"Apply"}/>

        </span>


      </div>
  );
}

export default Card;
