import React, { useState, useEffect } from "react";
import "./Notifications.scss";
import Button from "../Button/Button";
import axios from "axios";
import Switch from "react-switch";

function Notifications({ sw }) {
  const [companies, setCompanies] = useState([]);
  const [expLevel, setExpLevel] = useState([
    "Intern",
    "Entry",
    "Mid",
    "Senior",
    "Manager"
  ]);
  const [companyChosen, setCompanyChosen] = useState([]);
  const [expChosen, setExpChosen] = useState([]);
  const [historyCompanies, setHistoryCompanies] = useState([]);
  const [historyExp, setHistoryExp] = useState([]);
  const [apiCalled, setApiCalled] = useState(false);

  useEffect(() => {
    if (sw.fcmID && apiCalled == false) {
      setApiCalled(true);
      axios
        .get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/company/list/`)
        .then(res => {
          setCompanies(res.data);
        });
      axios
        .get(
          `${process.env.REACT_APP_REST_API}/rest/api/v1/notifications/preferences?deviceUUID=${sw.fcmID}`
        )
        .then(res => {
          let json = res.data;
          if(json["companies_uuid"] === null) {
            return
          }
          setHistoryCompanies(json["companies_uuid"]);
          setCompanyChosen(json["companies_uuid"]);
          for (let item in json) {
            if (json[item] === true) {
              let key = item;
              key = key[0].toUpperCase() + key.substring(1);
              setHistoryExp(oldArray => [...oldArray, key]);
              setExpChosen(oldArray => [...oldArray, key]);
            }
          }
          setApiCalled(false);
        })
        .catch(err => console.log(err));
    }
  }, [sw.fcmID]);

  function handleCompanyRadioButton(chk, e, id) {
    if (companyChosen.indexOf(id) > -1) {
      let arrFilter = companyChosen.filter(val => val !== id);
      setCompanyChosen(arrFilter);
    } else {
      setCompanyChosen(oldArray => [...oldArray, id]);
    }
  }

  function handleExpRadioButton(chk, e, id) {
    if (expChosen.indexOf(id) > -1) {
      let arrFilter = expChosen.filter(val => val !== id);
      setExpChosen(arrFilter);
    } else {
      setExpChosen(oldArray => [...oldArray, id]);
    }
  }

  function handleJobSubmit() {
    if(expChosen.length === 0 || companyChosen.length === 0) {
      console.log("please select something!")
      return
    }
    let data = {
      device_uuid: sw.fcmID,
      companies_uuid: companyChosen,
      intern: expChosen.includes("Intern"),
      entry: expChosen.includes("Entry"),
      mid: expChosen.includes("Mid"),
      senior: expChosen.includes("Senior"),
      manager: expChosen.includes("Manager")
    };

    if(historyExp.length === 0 || historyCompanies.length === 0) {
      axios({
        method: "post",
        url: `http://localhost:8080/rest/api/v1/notifications/preferences/create`,
        data: JSON.stringify(data)
      })
        .then(function(response) {
          //handle success
          console.log(response);
        })
        .catch(function(response) {
          //handle error
          console.log(response);
        });
    } else if (historyExp.length > 0 || historyCompanies.length > 0) {

      let removeCompanies = historyCompanies.filter(compuuid => !companyChosen.includes(compuuid))
      data = Object.assign({remove_comp_uuid: removeCompanies}, data)
      console.log(removeCompanies, data, companyChosen)
      axios({
        method: "post",
        url: `http://localhost:8080/rest/api/v1/notifications/preferences/update`,
        data: JSON.stringify(data)
      })
        .then(function(response) {
          //handle success
          console.log(response);
        })
        .catch(function(response) {
          //handle error
          console.log(response);
        });
    }
  }

  return (
    <>
      {!sw.fcmID ||
        (sw.fcmID.length === 0 && (
          <div>
            <h2>Please allow Notifications!</h2>
            <button onClick={sw.onClickAskUserPermission}>yeet</button>
            <button onClick={sw.onClickSubscribeToPushNotification}>uhh</button>
          </div>
        ))}
      {sw.fcmID &&
        companies.length != 0 &&
        companies.map(item => (
          <>
            <label>
              <span>{item.company_name}</span>
              <Switch
                key={item.companies_uuid}
                id={item.company_uuid}
                onChange={handleCompanyRadioButton}
                checked={companyChosen.includes(item.company_uuid)}
              />
            </label>
            <br />
          </>
        ))}
      <br />
      {sw.fcmID &&
        expLevel.map(item => (
          <>
            <label>
              <span>{item}</span>
              <Switch
                key={item}
                id={item}
                onChange={handleExpRadioButton}
                checked={expChosen.includes(item)}
              />
            </label>
            <br />
          </>
        ))}

      {sw.fcmID && (
        <Button btype="dropdown" onClick={handleJobSubmit} text={"Submit!"} />
      )}
    </>
  );
}

export default Notifications;
