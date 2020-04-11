import { useEffect, useState } from "react"
import axios from "axios"

export function useJobs() {
    const [jobs, setJobs] = useState([])

    useEffect(() => {
        axios.get('http://localhost:8080/rest/api/v1/jobs/')
        .then( res => {
            setJobs([res.data])
        })
    }, [])

    return {
        jobs,
        setJobs
    };
}