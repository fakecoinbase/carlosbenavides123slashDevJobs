import React, { useState } from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
// import Modal from "react-modal";
import { useParams } from "react-router";

import "./Navbar.scss";
import Home from "../Home/Home";
import JobForm from "../JobForm/JobForm";
import NavDropdown from "../Dropdown/Dropdown";
import CMS from "../CMS/CMS";
import CMSInfo from "../CMSInfo/CMSInfo";
import Notifications from "../Notifications/Notifications"
import DevJobs from './DevJobs.svg' ;

function Navbar({ joblist, sw }) {
  const [hamburgerClicked, setHamburgerClicked] = useState(false);
  function handleHamburger(e) {
    setHamburgerClicked(!hamburgerClicked);
  }
  return (
    <>
      <Router>
        <div
          className={hamburgerClicked ? "topnav responsive" : "topnav"}
          id="topnav"
        >
          <Link to="/">
          <span className="logo">DevJobs</span>

            {/* <span className="logo"><img src={DevJobs} className="devjobslogo"/></span> */}
          </Link>
          <span className="icon" onClick={handleHamburger}>
            &#9776;
          </span>

          <div className="dropdown">
            <div className="dropdown-content">
              <li>
                <Link to="/">Home</Link>
              </li>
              <li>
                <Link to="/notifications">Notifications</Link>
              </li>
              <li>
                <Link to="/jobform">Suggest Companies</Link>
              </li>
            </div>
          </div>
        </div>
        <hr className="solid" />

        {/* A <Switch> looks through its children <Route>s and
                renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/notifications">
            <Notifications sw={sw} />
          </Route>
          <Route path="/users">
            <Users />
          </Route>
          <Route path="/jobform">
            <JobForm />
          </Route>

          <Route path="/cms/:name/:uuid" component={CMSInfo} />

          <Route path="/cms">
            <CMS />
          </Route>
          <Route path="/">
            <NavDropdown {...joblist} />
            <hr className="solid" />
            <Home {...joblist} />
          </Route>
        </Switch>
      </Router>
    </>
  );
}

// function Test() {
//   return <h2>YTes</h2>
// }
// function Home() {
//   const customStyles = {
//     content: {
//       top: "50%",
//       left: "50%",
//       right: "auto",
//       bottom: "auto",
//       marginRight: "-50%",
//       transform: "translate(-50%, -50%)"
//     }
//   };
//   // Modal.setAppElement("#yourAppElement");

//   var subtitle;
//   const [modalIsOpen, setIsOpen] = React.useState(false);
//   function openModal() {
//     setIsOpen(true);
//   }

//   function afterOpenModal() {
//     // references are now sync'd and can be accessed.
//     subtitle.style.color = "#f00";
//   }

//   function closeModal() {
//     setIsOpen(false);
//   }

//   return (
//     <div>
//       <button onClick={openModal}>Open Modal</button>
//       <Modal
//         isOpen={modalIsOpen}
//         onAfterOpen={afterOpenModal}
//         onRequestClose={closeModal}
//         style={customStyles}
//         contentLabel="Example Modal"
//       >
//         <h2 ref={_subtitle => (subtitle = _subtitle)}>Hello</h2>
//         <button onClick={closeModal}>close</button>
//         <div>I am a modal</div>
//         <form>
//           <input />
//           <button>tab navigation</button>
//           <button>stays</button>
//           <button>inside</button>
//           <button>the modal</button>
//         </form>
//       </Modal>
//     </div>
//   );
// }

function Users() {
  return <h2>Users</h2>;
}

export default Navbar;
