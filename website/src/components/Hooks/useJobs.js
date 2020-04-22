import { useEffect, useState } from "react";
import axios from "axios";

export function useJobs() {
  const [jobs, setJobs] = useState([]);
  const [filteredJobs, setFilteredJobs] = useState([]);
  const [nonFilteredJobs, setNonFilteredJobs] = useState([]);

  const [location, setLocation] = useState("");
  const [experience, setExperience] = useState("");
  const [company, setCompany] = useState("");

  const [companyDropdown, setCompanyDropdown] = useState([])
  const [companyUUID, setCompanyUUID] = useState(new Map())

  const [homePage, setHomePage] = useState([])
  const [companyPage, setCompanyPage] = useState([])

  const [loading, setLoading] = useState(false)


  useEffect(() => {
    axios.get("http://localhost:8080/rest/api/v1/jobs/company/list/").then(res => {
      let temp = []
      var myMap = new Map();
      for(var obj of res.data) {
        temp.push({value:obj["company_name"], label:obj["company_name"]})
        myMap.set(obj["company_name"], obj["company_uuid"])
      }
      setCompanyDropdown(temp)
      setCompanyUUID(myMap)
    })
  }, [])


  useEffect(() => {
    setLoading(true)
    axios.get("http://localhost:8080/rest/api/v1/jobs/").then(res => {
      setJobs(res.data);
      setHomePage(res.data)
      setLoading(false)
    });
  }, []);

  useEffect(() => {
    if (location === "" && experience === "" && company == "") {
      setJobs(homePage)
      setCompanyPage([])
      return
    }
    // rest api call if companyPage doesn't equal to company name
    if (company !== "" && experience === "" && location === "") {
      if(companyPage.length != 0) {
        if(companyPage[0].company_name === company) {
          setJobs(companyPage)
          return
        }
      }
      let uuid = companyUUID.get(company)
      setLoading(true)
      axios.get(`http://localhost:8080/rest/api/v1/jobs/company/search/${uuid}`)
        .then((res) =>{
          setJobs(res.data)
          setCompanyPage(res.data)
          setLoading(false)
        })
    } else if (company !== "" && experience !== "" && location === "") {
      setJobs(filterByExperience(companyPage, experience))
    } else if (company !== "" && experience === "" && location !== "") {
      setJobs(filterByLocation(companyPage, location))
    } else if (company !== "" && experience !== "" && location !== "") {
      setJobs(filterByExperienceAndLocation(companyPage, experience, location))
    }
    // var filtered = [];

    // otherwise filter as much w/o calling api
    // assume companyPage state is empty
    if (company === "" && experience !== "" && location === "") {
      setJobs(filterByExperience(homePage, experience))
    } else if (company === "" && experience === "" && location !== "") {
      setJobs(filterByLocation(homePage, location))
    } else if(company === "" && experience !== "" && location !== "") {
      setJobs(filterByExperienceAndLocation(homePage, experience, location))
    }
  }, [company, location, experience]);

  function filterByExperience(data, experience) {
    return data.filter(i => i.level === experience)
  }

  function filterByLocation(data, location) {
    return data.filter(i => i.job_location.includes(location))
  }

  function filterByExperienceAndLocation(data, experience, location) {
    let filtered = []
    filtered = data.filter(i => i.level === experience)
    filtered = filtered.filter(i => i.job_location.includes(location))
    return filtered
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
    setLoading
  };
}