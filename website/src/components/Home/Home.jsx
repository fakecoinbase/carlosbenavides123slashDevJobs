import "./Home.scss";
import React, {useEffect} from "react";
// import Honey from "./unnamed.png";
import Card from "../Card/Card";
function Home({joblist}) {
  return (
    <div className="cards">

      {joblist && 
        joblist.map((item) => 
          <Card key={item.uuid} {...item} />
        )
      }
    </div>
  );
}

export default Home;
