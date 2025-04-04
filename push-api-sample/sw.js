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
