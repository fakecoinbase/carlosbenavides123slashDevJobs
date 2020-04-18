import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
// import Modal from "react-modal";
import "./Navbar.scss";
import Home from "../Home/Home";
import JobForm from "../JobForm/JobForm";
import NavDropdown from "../Dropdown/Dropdown";
function Navbar({ joblist, sw }) {
  console.log(sw);
  return (
    <>
      <Router>
        <header>
          <span className="logo">Logo</span>
          <nav>
            <ul className="nav__links">
              <li>
                <Link to="/">Home</Link>
              </li>
              <li>
                <Link to="/about">About</Link>
              </li>
              <li>
                <Link to="/jobform">Add A Company</Link>
              </li>
            </ul>
          </nav>
        </header>
        <hr className="solid" />

        {/* A <Switch> looks through its children <Route>s and
                renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/about">
            <About sw={sw} />
          </Route>
          <Route path="/users">
            <Users />
          </Route>
          <Route path="/jobform">
            <JobForm />
          </Route>
          <Route path="/">
            <NavDropdown {...joblist}/>
            <hr className="solid" />
            <Home joblist={joblist.jobs} />
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

function About({ sw }) {
  return (
    <>
      <h2>About</h2>
      <button onClick={sw.onClickAskUserPermission}>yeet</button>
      <button onClick={sw.onClickSubscribeToPushNotification}>uhh</button>
      lol
      {sw.isPushNotificationSupported}
      {/* <button onClick={sw.onClickSubscribeToPushNotification}>yeah</button> */}
    </>
  );
}

function Users() {
  return <h2>Users</h2>;
}

export default Navbar;
