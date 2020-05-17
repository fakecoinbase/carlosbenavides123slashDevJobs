import React, {useState, useEffect} from "react"
import axios from 'axios'
import './CMS.scss'
import CMSInfo from "../CMSInfo/CMSInfo"

import { Link } from "react-router-dom";

function CMS() {
    const [companies, setCompanies] = useState([])
    useEffect(() => {
        axios.get(`${process.env.REACT_APP_REST_API}/rest/api/v1/jobs/company/list/`).then(res => {
            let json = res.data
            setCompanies(json)
        })
    }, [])

    return ( 
    <div className="cms__container">
        {companies != undefined && companies.map(company =>
        <li>
            <Link to={`/cms/${company.company_name}/${company.company_uuid}`}>{company.company_name}</Link>
        </li>
        )}
    </div>
    )
}

export default CMS;