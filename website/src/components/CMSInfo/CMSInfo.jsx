import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import { css } from "@emotion/core";
import ClipLoader from "react-spinners/ClipLoader";
import Button from "../Button/Button";

const override = css`
  margin: 0 auto;
  border-color: red;
  text-align: center;
  display: flex;
  justify-content: center;
`;
function CMS() {
  let params = useParams();
  const [companyName, setCompanyName] = useState("");
  const [companyWebsite, setCompanyWebsite] = useState("");
  const [greenhouse, setGreenhouse] = useState(false);
  const [lever, setLever] = useState(false);
  const [other, setOther] = useState(false);
  const [wantedDepartments, setWantedDepartments] = useState("");
  const [wantedLocations, setWantedLocations] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setLoading(true);
    axios
      .get(
        `${process.env.REACT_APP_REST_API}/rest/api/v1/cms/companydetails?company=${params.name}`
      )
      .then(res => {
        var json = res.data;
        setCompanyName(json["company_name"]);
        setCompanyWebsite(json["company_website"]);
        setGreenhouse(json["greenhouse"]);
        setLever(json["lever"]);
        setOther(json["other"]);
        setWantedDepartments(json["wanted_departments"]);
        setWantedLocations(json["wanted_locations"]);
        setLoading(false);
      });
  }, []);

  function handleJobSubmit(e) {
    e.preventDefault();
    var data = JSON.stringify({
      company_uuid: params.uuid,
      company_name: companyName,
      company_website: companyWebsite,
      greenhouse: greenhouse,
      lever: lever,
      other: other,
      wanted_departments: wantedDepartments,
      wanted_locations: wantedLocations
    });
    axios.post(
      `${process.env.REACT_APP_REST_API}/rest/api/v1/cms/companydetails/update`,
      data
    );
  }
  function handleButton(e) {
    var radioValue = e.target.value;
    if (radioValue === "greenhouse") {
      setGreenhouse(true);
      setLever(false);
      setOther(false);
    } else if (radioValue === "lever") {
      setGreenhouse(false);
      setLever(true);
      setOther(false);
    } else {
      setGreenhouse(false);
      setLever(false);
      setOther(true);
    }
  }

  return (
    <div>
      {loading && (
        <div className="sweet-loading">
          <ClipLoader
            css={override}
            size={150}
            color={"#123abc"}
            loading={loading}
          />
        </div>
      )}
      {!loading && (
        <div style={{ display: "flex", justifyContent: "center" }}>
          <div className="card">
            <h3>Make changes to {params.name}.</h3>
            <input
              className="JobFormInput"
              placeholder="Company Name"
              onChange={e => setCompanyName(e.target.value)}
              value={companyName}
            />
            <input
              className="JobFormInput"
              placeholder="Career Page"
              onChange={e => setCompanyWebsite(e.target.value)}
              value={companyWebsite}
            />
            <input
              type="radio"
              name="GreenHouse"
              value="greenhouse"
              onChange={handleButton}
              checked={greenhouse}
            />
            GreenHouse
            <br />
            <input
              type="radio"
              name="Lever"
              value="lever"
              onChange={handleButton}
              checked={lever}
            />
            Lever
            <br />
            <input
              type="radio"
              name="Other"
              value="other"
              onChange={handleButton}
              checked={other}
            />
            Other
            <br />
            <input
              className="JobFormInput"
              placeholder={`Add departments (comma seperated) for ${params.name}`}
              onChange={e => setWantedDepartments(e.target.value)}
              value={wantedDepartments}
            />
            <input
              className="JobFormInput"
              placeholder={`Add locations (comma seperated) for ${params.name}`}
              onChange={e => setWantedLocations(e.target.value)}
              value={wantedLocations}
            />
            <Button
              btype="joblink"
              onClick={handleJobSubmit}
              text={"Submit!"}
            />
          </div>
        </div>
      )}
    </div>
  );
}

export default CMS;
