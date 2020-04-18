import { useEffect, useState } from "react";
import {
  isPushNotificationSupported,
  askUserPermission,
  registerServiceWorker,
  createNotificationSubscription,
  getUserSubscription
} from "../../push-notif";
import { messaging } from "../../firebase";

import {useLocalStorage} from "./useLocalStorage"

const pushNotificationSupported = isPushNotificationSupported();

export function useServiceWorker() {
  const [userConsent, setUserConsent] = useState(Notification.permission);
  const [fcmID, setFCMID] = useLocalStorage('fcmID', '')
  const [pushServiceSubID, setPushServiceSubID] = useState();
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (pushNotificationSupported && userConsent === "granted") {
      if ("serviceWorker" in navigator) {
        navigator.serviceWorker
          .register("./firebase-messaging-sw.js")
          .then(_ => {
            if (fcmID == ""){
              messaging
              .requestPermission()
              .then(async function() {
                const token = await messaging.getToken();
                setFCMID(token);
                console.log(token, "here");
              })
              .catch(function(err) {
                console.log("Unable to get permission to notify.", err);
              });
            }
          });
      }
    }
  }, []);

  // useEffect(() => {
  //     setLoading(true);
  //     setError(false);
  //     const getExistingSubscription = async () => {
  //         const existingSubscription = await getUserSubscription()
  //         setUserSubscription(existingSubscription)
  //         setLoading(false)
  //     };
  //     getExistingSubscription()
  // }, [])

  const onClickAskUserPermission = () => {
    setLoading(true);
    setError(false);
    askUserPermission().then(consent => {
      setUserConsent(consent);
      if (consent !== "granted") {
        console.log("consent denied");
      }
    });
    setLoading(false);
  };

  // const onClickSubscribeToPushNotification = () => {
  //     setLoading(true);
  //     setError(false);
  //     createNotificationSubscription()
  //     .then(function(subscription) {
  //         setUserSubscription(subscription)
  //     })
  //     .catch((err) => {
  //         console.log(err)
  //     })
  //     setLoading(false)
  // }

  return {
    onClickAskUserPermission,
    // onClickSubscribeToPushNotification,
    loading,
    error,
    pushServiceSubID,
    // userSubscription,
    userConsent
  };
}
