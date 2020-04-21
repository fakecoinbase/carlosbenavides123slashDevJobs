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
  const [experienceDropdown, setExperienceDropdown] = useState([])

  const [companyUUID, setCompanyUUID] = useState(new Map())

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
      console.log(myMap)
    })
  }, [])


  useEffect(() => {
    axios.get("http://localhost:8080/rest/api/v1/jobs/").then(res => {
      setJobs(res.data);
    });
  }, []);

  useEffect(() => {
    // rest api call
    if (company !== "") {
      console.log(companyUUID, company, "TEST")
      let uuid = companyUUID.get(company)
      axios.get(`http://localhost:8080/rest/api/v1/jobs/company/search/${uuid}`)
      .then(res => setJobs(res.data))
    }
    var filtered = [];

    // otherwise filter as much w/o calling api
    if (company === "" && (experience !== "" || location !== "")) {
      setNonFilteredJobs(jobs);

      filtered = jobs.filter(i => i.level == experience);

      if (location !== "") {
        // filtered = jobs.filter( i => i.location )
      }
      console.log(filtered);
      setJobs(filtered);
    } else if (experience === "" && location === "") {
      setJobs(nonFilteredJobs);
    }
  }, [company, location, experience]);

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
    setCompanyDropdown
  };
}
