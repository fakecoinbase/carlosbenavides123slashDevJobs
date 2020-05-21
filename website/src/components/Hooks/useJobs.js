import { useEffect, useState } from "react";
import axios from "axios";
var https = require('https')
export function useJobs() {
  const [jobs, setJobs] = useState([]);
  const [filteredJobs, setFilteredJobs] = useState([]);
  const [nonFilteredJobs, setNonFilteredJobs] = useState([]);

  const [location, setLocation] = useState("");
  const [locationDropDown, setLocationDropDown] = useState([])
  const [homePageLocationDropDown, sethomePageLocationDropDown] = useState([])

  const [experience, setExperience] = useState("");

  const [company, setCompany] = useState("");
  const [companyDropdown, setCompanyDropdown] = useState([]);
  const [homePageCompanyDropdown, sethomePageCompanyDropdown] = useState([])

  const [companyUUID, setCompanyUUID] = useState(new Map());

  const [homePage, setHomePage] = useState([]);
  const [companyPage, setCompanyPage] = useState([]);

  const [loading, setLoading] = useState(false);

  const [cursor, setCursor] = useState(0)
  const [locCursor, setLocCursor] = useState("")
  const [expCursor, setExpCursor] = useState("")

  const [apiCalled, setApiCalled] = useState(true)

  const httsAgent = new https.Agent({ rejectUnauthorized: false });


  useEffect(() => {
    axios
      .get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/company/list/`, {httsAgent})
      .then(res => {
        let temp = [];
        var myMap = new Map();
        var json = res.data;
        json = json.sort(function(a, b) {
          if (a.company_name < b.company_name) {
            return -1;
          }
          if (a.company_name > b.company_name) {
            return 1;
          }
          return 0;
        });
        for (var obj of json) {
          temp.push({ value: obj["company_name"], label: obj["company_name"] });
          myMap.set(obj["company_name"], obj["company_uuid"]);
        }
        setCompanyDropdown(temp);
        sethomePageCompanyDropdown(temp)
        setCompanyUUID(myMap);
        setLocationDropDown([{ value: "Los Angeles", label: "Los Angeles"}, { value: "Venice", label: "Venice"}, { value:"San Francisco", label:"San Francisco"}, { value:"New York", label:"New York"}, { value:"Denver", label:"Denver"}, { value:"Seattle", label:"Seattle"}, { value:"Bellevue", label:"Bellevue" }, { value:"Boulder", label:"Boulder" }])
        sethomePageLocationDropDown([{ value: "Los Angeles", label: "Los Angeles"}, { value: "Venice", label: "Venice"}, { value:"San Francisco", label:"San Francisco"}, { value:"New York", label:"New York"}, { value:"Denver", label:"Denver"}, { value:"Seattle", label:"Seattle"}, { value:"Bellevue", label:"Bellevue" }, { value:"Boulder", label:"Boulder" }])
      });
  }, []);

  useEffect(() => {
    if(company !== "") {
      axios.get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/company/company/${company}`)
      .then( res => {
        var json = res.data
        json = json.sort(function(a, b) {
          if (a.location < b.location) {
            return -1;
          }
          if (a.location > b.location) {
            return 1;
          }
          return 0;
        });
        let temp = []
        for (var obj of json) {
          temp.push({ value: obj["location"], label: obj["location"] });
        }
        setLocationDropDown(temp)
      })
    } else {
      setLocationDropDown(homePageLocationDropDown)
    }
  }, [company])

  useEffect(() => {
    if(location !== "") {
      axios.get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/company/location/${location}`)
      .then( res => {
        var json = res.data
        json = json.sort(function(a, b) {
          if (a.company_name < b.company_name) {
            return -1;
          }
          if (a.company_name > b.company_name) {
            return 1;
          }
          return 0;
        });
        let temp = []
        for (var obj of json) {
          temp.push({ value: obj["company_name"], label: obj["company_name"] });
        }
        setCompanyDropdown(temp)
      })
    } else {
      setCompanyDropdown(homePageCompanyDropdown)
    }
  }, [location])

  // useEffect(() => {
  //   if(experience === "") {
  //     setFilteredJobs([])
  //   }
  // }, [experience])

  useEffect(() => {
    homepageCall()
  }, []);

  function homepageCall() {
    setLoading(true);
    setApiCalled(true)
    axios.get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/index?timestamp=`).then(res => {
      var json = res.data
      setJobs(json["Job"]);
      setHomePage(json["Job"]);
      setCursor(json["Cursor"]["next_cursor"])
      console.log(json["Cursor"]["next_cursor"], "NEXT CURSOR")
      setLoading(false);
      setApiCalled(false)
    });
  }

  useEffect(() => {
    if (location === "" && experience === "" && company == "") {
      homepageCall()
      setLocCursor("")
      setExpCursor("")
      setFilteredJobs([])
      return;
    }
    // rest api call if companyPage doesn't equal to company name
    if (company !== "" && experience === "" && location === "") {
      if (companyPage.length != 0) {
        if (companyPage[0].company_name === company) {
          setJobs(companyPage);
          return;
        }
      }
      setCursor(undefined)
      setLocCursor(undefined)
      setApiCalled(true)
      let uuid = companyUUID.get(company);
      setLoading(true);
      axios
        .get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/company/search/${uuid}`)
        .then(res => {
          setJobs(res.data);
          setCompanyPage(res.data);
          setLoading(false);
          setApiCalled(false)
        });
    } else if (company !== "" && experience !== "" && location === "") {
      checkCompanyPage(company);
      setJobs(filterByExperience(companyPage, experience));
    } else if (company !== "" && experience === "" && location !== "") {
      checkCompanyPage(company);
      setJobs(filterByLocation(companyPage, location));
    } else if (company !== "" && experience !== "" && location !== "") {
      checkCompanyPage(company);
      setJobs(filterByExperienceAndLocation(companyPage, experience, location));
    }

    if(company === "" && location !== "" && experience === "") {
      console.log(checkMemoryJobs(), "CHECK MEMORY JOBS", filteredJobs)
      if(checkMemoryJobs()) {
        apiJobsByLocation()
      } else {
        setJobs(filteredJobs)
      }
      console.log("HEEEEEEEEEEEEEEEREEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEeee")
    } else if(company === "" && location !== "" && experience !== "" && typeof(locCursor) === 'number') {
      setFilteredJobs(jobs)
      setJobs(filterByExperience(jobs, experience))
    } else if(company === "" && location === "" && experience !== "") {
      if (checkMemoryJobs()) {
        apiJobsByExperience()
      } else {
        setJobs(filteredJobs)
      }
    } else if(location !== "" && experience !== "" && typeof(expCursor) === 'number') {
      setFilteredJobs(jobs)
      setJobs(filterByLocation(jobs, location))
    }

  }, [company, location, experience]);

  function checkMemoryJobs() {
    if (
      filteredJobs === undefined ||
      filteredJobs.length === 0
    ) {
      return true
    }
    return false
  }

  function apiJobsByLocation() {
    setLoading(true);
    setCursor(undefined)
    setApiCalled(true)
    axios.get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/search/location?location=${location}&cursor=${locCursor}`).then(res => {
      var json = res.data
      console.log(json, "JOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBSJOBS")
      setJobs(json["Job"]);
      setFilteredJobs(json["Job"]);
      setLocCursor(json["Cursor"]["next_cursor"])
      console.log(json["Cursor"]["next_cursor"], "NEXT CURSOR")
      setLoading(false);
      setApiCalled(false)
    });
  }

  function apiJobsByExperience() {
    setLoading(true)
    setCursor(undefined)
    setLocCursor(undefined)
    setApiCalled(true)
    console.log(experience)
    axios.get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/search/experience?experience=${experience}&cursor=${expCursor}`)
    .then(res => {
      var json = res.data;
      setJobs(json["Job"])
      setFilteredJobs(json["Job"])
      setExpCursor(json["Cursor"]["next_cursor"])
      setLoading(false)
      setApiCalled(false)
    }) 
  }  

  function filterByExperience(data, experience) {
    return data.filter(i => i.level === experience);
  }

  function checkCompanyPage(company) {
    if (
      companyPage === undefined ||
      companyPage.length === 0 ||
      companyPage[0].company_name !== company
    ) {
      let uuid = companyUUID.get(company);
      setLoading(true);
      setApiCalled(true)
      axios
        .get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/company/search/${uuid}`)
        .then(res => {
          setJobs(res.data);
          setCompanyPage(res.data);
          setLoading(false);
          setApiCalled(false)
        });
    }
  }

  useEffect(() => {
    if (company !== "" && experience === "" && location === "") {
      return;
    } else if (company !== "" && experience !== "" && location === "") {
      setJobs(filterByExperience(companyPage, experience));
    } else if (company !== "" && experience === "" && location !== "") {
      setJobs(filterByLocation(companyPage, location));
    } else if (company !== "" && experience !== "" && location !== "") {
      setJobs(filterByExperienceAndLocation(companyPage, experience, location));
    }
  }, [companyPage]);

  function filterByLocation(data, location) {
    return data.filter(i => i.job_location.includes(location));
  }

  function filterByExperienceAndLocation(data, experience, location) {
    let filtered = [];
    filtered = data.filter(i => i.level === experience);
    filtered = filtered.filter(i => i.job_location.includes(location));
    return filtered;
  }

  return {
    jobs,
    setJobs,
    company,
    setCompany,
    location,
    setLocation,
    experience,
    setExperience,
    companyDropdown,
    setCompanyDropdown,
    loading,
    setLoading,
    locationDropDown,
    setHomePage,
    cursor,
    setCursor,
    locCursor,
    setLocCursor,
    setApiCalled,
    apiCalled,
    expCursor,
    experience,
    setExpCursor
  };
}
