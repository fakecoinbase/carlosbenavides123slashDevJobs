import "./Home.scss";
import React, { useEffect, useState, useRef } from "react";
import { css } from "@emotion/core";
import ClipLoader from "react-spinners/ClipLoader";
import InfiniteScroll from "react-infinite-scroller";
import axios from "axios";
import Card from "../Card/Card";

const override = css`
  margin: 0 auto;
  border-color: red;
  text-align: center;
  display: flex;
  justify-content: center;
`;
function Home({
  jobs,
  loading,
  location,
  scrollMore,
  cursor,
  locCursor,
  setLocCursor,
  setJobs,
  setHomePage,
  setLoading,
  setCursor,
  setApiCalled,
  apiCalled
}) {
  var hasMore;

  function fetchMoreData() {
    if (cursor != undefined && cursor != 0 ) {
      console.log("fetch more data normal")
      setApiCalled(true)
        setLoading(true);
        axios.get(`http://localhost:8080/rest/api/v1/jobs/index?timestamp=${cursor}`).then(res => {
          var json = res.data
          setJobs(jobs => jobs.concat(json["Job"]));
          setHomePage(jobs => jobs.concat(json["Job"]));
          setLoading(false);
          setCursor(json["Cursor"]["next_cursor"])
          setApiCalled(false)
        });
    } else if (locCursor != undefined && cursor != 0) {
      setLoading(true);
      console.log("fetch more data location")
      setApiCalled(true)
      axios.get(`http://localhost:8080/rest/api/v1/jobs/search/location?location=${location}&cursor=${locCursor}`).then(res => {
        var json = res.data
        setJobs(jobs => jobs.concat(json["Job"]));
        setHomePage(jobs => jobs.concat(json["Job"]));
        setLoading(false);
        setLocCursor(json["Cursor"]["next_cursor"])
        setApiCalled(false)
      });
      
    }
    //   if (cursor != undefined && cursor != 0) {
    //     console.log(cursor)
    //     hasMore = false
    //     setLoading(true);
    //     axios.get(`http://localhost:8080/rest/api/v1/jobs/index?timestamp=${cursor}`).then(res => {
    //       var json = res.data
    //       setJobs(jobs => jobs.concat(json["Job"]));
    //       setHomePage(jobs => jobs.concat(json["Job"]));
    //       setLoading(false);
    //       setCursor(json["Cursor"]["next_cursor"])
    //     });
    //   }
  }

  function hasMore() {
    if(apiCalled === true) return false
    console.log("##############HAS MOREEEEEEEEEEE", locCursor, cursor)
    if(locCursor !== undefined && locCursor !== 0) return true
    if (cursor !== undefined && cursor !== 0) return true
    return false
  }


  return (
    <>
      {jobs && jobs !== [] && (
        <InfiniteScroll
          loadMore={fetchMoreData}
          hasMore={hasMore()}
          loader={
            <div className="sweet-loading">
              <ClipLoader
                css={override}
                size={150}
                color={"#123abc"}
                loading={loading}
              />
            </div>
          }
        >
          <div className="cards">
            {jobs && jobs.map(item => <Card key={item.uuid} {...item} />)}
          </div>
        </InfiniteScroll>
      )}
    </>
  );
}

// {loading && (
// <div className="sweet-loading">
//   <ClipLoader
//     css={override}
//     size={150}
//     color={"#123abc"}
//     loading={loading}
//   />
// </div>
// )}

export default Home;
