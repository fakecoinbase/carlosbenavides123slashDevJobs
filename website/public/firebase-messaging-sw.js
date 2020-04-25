importScripts("https://www.gstatic.com/firebasejs/5.9.4/firebase-app.js");
importScripts("https://www.gstatic.com/firebasejs/5.9.4/firebase-messaging.js");
firebase.initializeApp({
  // Project Settings => Add Firebase to your web app
  messagingSenderId: "829234754086"
});

const messaging = firebase.messaging();

messaging.setBackgroundMessageHandler(function(payload) {
  console.log("log background message payload:", payload)
  const promiseChain = clients
    .matchAll({
      type: "window",
      includeUncontrolled: true
    })
    .then(windowClients => {
      for (let i = 0; i < windowClients.length; i++) {
        const windowClient = windowClients[i];
        windowClient.postMessage(payload);
      }
    })
    .then(() => {
      // return registration.showNotification("my notification title");
    });
  return promiseChain;
});

function receivePushNotification(event) {
  console.log("[SW] Push Received");
  console.log("push message event:", event.data.json(), "1323");
  const obj = event.data.json()
  const title = obj["data"]["title"]
  const url = obj["data"]["url"]
  const text = obj["data"]["body"]
  const options = {
    data: url,
    body: text,
    vibrate: [200, 100, 200],
    actions: [{ action: "Detail", title: "View", icon: "https://via.placeholder.com/128/ff0000" }]
  }
  event.waitUntil(self.registration.showNotification(title, options));
}

function openPushNotification(event) {
  console.log("[SW] Notif Clicked", event.notification.data);
  console.log(event)
  event.notification.close();
  event.waitUntil(clients.openWindow(event.notification.data));
}

self.addEventListener("push", receivePushNotification);
self.addEventListener("notificationclick", openPushNotification);
