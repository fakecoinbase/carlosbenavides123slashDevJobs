import { useEffect, useState } from "react";
import axios from "axios";

export function useJobs() {
  const [jobs, setJobs] = useState([]);
  const [filteredJobs, setFilteredJobs] = useState([]);
  const [nonFilteredJobs, setNonFilteredJobs] = useState([]);

  const [location, setLocation] = useState("");
  const [experience, setExperience] = useState("");
  const [company, setCompany] = useState("");

  useEffect(() => {
    axios.get("http://localhost:8080/rest/api/v1/jobs/").then(res => {
      setJobs(res.data);
    });
  }, []);

  useEffect(() => {
    // rest api call
    if (company !== "") {
      console.log("1");
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
    setExperience
  };
}
