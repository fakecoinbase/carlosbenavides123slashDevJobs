
const pushServerPublicKey =
  "AAAAwRI-MiY:APA91bHExWXJyxdLKzDo-w7iS79nKs2e5QiYr4xIYBMzc4uHvExfXfkuHfrktpAkFEqTRp-uw5h5PBVQLwHmhlMkpF7Ub6SsFG_zXUODCp3OXe5LTB4w-aBsFiGzlWOqhKz_jf6izL-f";

function isPushNotificationSupported() {
  return "serviceWorker" in navigator && "PushManager" in window;
}

async function askUserPermission() {
  return await Notification.requestPermission();
}

function registerServiceWorker() {
  return navigator.serviceWorker.register("/sw.js");
}

async function createNotificationSubscription() {
  //wait for service worker installation to be ready
  const serviceWorker = await navigator.serviceWorker.ready;
  // subscribe and return the subscription
  return await serviceWorker.pushManager.subscribe({
    userVisibleOnly: true,
    applicationServerKey: pushServerPublicKey
  });
}

function sendNotification() {
  console.log("send notif method");
}

function getUserSubscription() {
  //wait for service worker installation to be ready, and then
  return navigator.serviceWorker.ready
    .then(function(serviceWorker) {
      return serviceWorker.pushManager.getSubscription();
    })
    .then(function(pushSubscription) {
      return pushSubscription;
    });
}

export {
  isPushNotificationSupported,
  askUserPermission,
  registerServiceWorker,
  sendNotification,
  createNotificationSubscription,
  getUserSubscription
};
