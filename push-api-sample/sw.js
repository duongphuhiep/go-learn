// Service Worker installation
self.addEventListener('install', (event) => {
  console.log('Service Worker installed');
  /*
  Default Service Worker Behavior: When a new Service Worker is installed, it waits until all tabs using the old Service Worker are closed before activating
  If you reload the page after updating sw.js, the old Service Worker might still control the page, the "Ping" message might not reach the new Service Worker.
  Here The new Service Worker takes control right after installation, even if other tabs are using the old one. This is useful during development to test updates quickly.
  */
  event.waitUntil(self.skipWaiting()); // Activate immediately
});

// Service Worker activation
self.addEventListener('activate', (event) => {
  console.log('Service Worker activated');
  /* While skipWaiting() activates the Service Worker, clients.claim() ensures it takes control of existing tabs */
  event.waitUntil(self.clients.claim()); // Take control of all clients
});

self.addEventListener("push", (event) => {
  const data = event.data?.json() || {
    title: "New Notification",
    body: "You have a new message!",
  };

  // Send data to the frontend (for DOM append)
  console.log("Sending data to all clients");
  self.clients.matchAll().then((clients) => {
    clients.forEach((client) => {
      console.log("Sending data to client:", client);
      client.postMessage({
        type: "NOTIFICATION",
        payload: data,
      });
    });
  });

  event.waitUntil(
    self.registration.showNotification(data.title, {
      body: data.body,
      icon: "/star-fruit.svg",
    })
  );
});

// Handle Ping messages from clients
self.addEventListener('message', (event) => {
  console.log('Service Worker received message:', event.data);
});
