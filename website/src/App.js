import React, {useEffect} from "react";
import NavBar from "./components/NavBar/Navbar";
import {useJobs} from "./components/Hooks/useJobs"
import {useServiceWorker} from "./components/Hooks/useServiceWorker"
export default function App() {
  const joblist = useJobs();
  const sw = useServiceWorker();

  // useEffect(() => {
  //   messaging.requestPermission()
  //   .then(async function() {
  //       const token = await messaging.getToken();
  //       console.log(token)
  //   })
  //   .catch(function(err) {
  //     console.log("Unable to get permission to notify.", err);
  //   });
  // })

  return (
    <>
      <NavBar joblist={joblist.jobs} sw={sw} />
    </>
  );
}
