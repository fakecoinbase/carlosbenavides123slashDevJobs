import "./Home.scss";
import React, { useEffect } from "react";
import { css } from "@emotion/core";
import ClipLoader from "react-spinners/ClipLoader";
import InfiniteScroll from "react-infinite-scroller";

import Card from "../Card/Card";
const override = css`
  margin: 0 auto;
  border-color: red;
  text-align: center;
  display: flex;
  justify-content: center;
`;
function Home({ joblist, loading }) {
  function fetchMoreData() {
    console.log("haha");
  }

  return (
    <>
      <InfiniteScroll
        dataLength={joblist.length}
        next={fetchMoreData}
        hasMore={loading}
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
        {!loading && (
          <div className="cards">
            {joblist && joblist.map(item => <Card key={item.uuid} {...item} />)}
          </div>
        )}
      </InfiniteScroll>
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
