import * as firebase from "firebase/app";
import "firebase/messaging";

const initializedFirebaseApp = firebase.initializeApp({
  // Project Settings => Add Firebase to your web app
  apiKey: "AIzaSyB__WduG4Oj43ILbyY8uIa6VKxF4xBs4fA",
  authDomain: "devjobs-9e2a2.firebaseapp.com",
  databaseURL: "https://devjobs-9e2a2.firebaseio.com",
  projectId: "devjobs-9e2a2",
  storageBucket: "devjobs-9e2a2.appspot.com",
  messagingSenderId: "829234754086",
  appId: "1:829234754086:web:62830db5cbfbb54d916ce7",
  measurementId: "G-7WW07KGCDQ"
});

const messaging = initializedFirebaseApp.messaging();

messaging.usePublicVapidKey(
  // Project Settings => Cloud Messaging => Web Push certificates
  "BGBUmCTuDo9rIjwiiaXU-VeHQFoBhLzeW4pvY4IQyoabQyow0OAvtwTJmjo0erWl-8W5h6pt7BLYxqjOYc1lpJA"
);

export { messaging };
