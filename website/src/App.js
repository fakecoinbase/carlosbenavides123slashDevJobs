import React, {useEffect} from "react";
import NavBar from "./components/NavBar/Navbar";
import {useJobs} from "./components/Hooks/useJobs"
import {useServiceWorker} from "./components/Hooks/useServiceWorker"
export default function App() {
  const joblist = useJobs();
  const sw = useServiceWorker();

  return (
    <>
      <NavBar joblist={joblist} sw={sw} />
    </>
  );
}
