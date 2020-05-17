import React, {useState, useEffect} from "react"
import {
    useParams
  } from "react-router-dom";
import axios from "axios";

function CMS() {
    let params  = useParams()
    console.log(params.name, "YEEEET")

    useEffect(() => {
        axios.get(`${process.env.REACT_APP_REST_API}/rest/api/v1/cms/companydetails?company=${params.name}`)
        .then( res => {
            console.log(res)
        })
    }, [])

    return ( 
        <div>
        {params.name} 
        </div>
    )
}

export default CMS;