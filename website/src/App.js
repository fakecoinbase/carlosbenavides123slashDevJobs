import React from "react";
import NavBar from "./components/NavBar/Navbar";
import {useJobs} from "./components/Hooks/useJobs"

export default function App() {
  const joblist = useJobs();

  return (
    <>
      <NavBar joblist={joblist.jobs} />
    </>
  );
}
