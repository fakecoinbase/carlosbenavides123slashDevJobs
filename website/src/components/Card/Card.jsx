import React from "react";
import "./Card.scss";
import Button from "../Button/Button"
import moment from 'moment'

function Card({uuid, company_name, job_link, job_posted, job_found}) {

  var res = uuid.split("_%_")
  var job_title = res[1].replace(/%/g, ' ')
  var location = res[2].replace(/%/g, ' ')

  function handleJobClick(e) {
    window.open(job_link, '_blank')
  }
  return (
      <div className="card">
        <span>
          <img
            src={
              "https://res.cloudinary.com/dhxwdb3jl/image/upload/v1586121171/unnamed_wqeqel.png"
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
        <Button btype="joblink" text={"Apply"} onClick={handleJobClick}/>

        </span>


      </div>
  );
}

export default Card;
