import "./Home.scss";
import React, { useEffect } from "react";
import { css } from "@emotion/core";
import ClipLoader from "react-spinners/ClipLoader";
import Card from "../Card/Card";
const override = css`
  margin: 0 auto;
  border-color: red;
  text-align: center;
  display: flex;
  justify-content: center
`;
function Home({ joblist, loading }) {
  return (
    <>
      {loading && (
        <div className="sweet-loading">
          <ClipLoader
            css={override}
            size={150}
            color={"#123abc"}
            loading={loading}
          />
        </div>
      )}

      {!loading && (
        <div className="cards">
          {joblist && joblist.map(item => <Card key={item.uuid} {...item} />)}
        </div>
      )}
    </>
  );
}

export default Home;
