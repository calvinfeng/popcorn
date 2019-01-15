// Name the cache
var CACHE_NAME = 'popcorn-pwa-cache-v0.0.1';

// The first time the user starts up the PWA, 'install' is triggered.
self.addEventListener('install', function(event) {
  console.log("Installing service worker", CACHE_NAME);
  const p = caches.open(CACHE_NAME).then(function(cache) {
    const urlsToCache = ["/index.html", "/index.js", "/android-chrome-512x512.png"];
    cache.addAll(urlsToCache);
    console.log('cached', urlsToCache);
  })
  event.waitUntil(p);
});

// Delete old caches that are not our current one!
self.addEventListener("activate", (event) => {
  console.log("Service worker has been activated, ready to fetch");
  const cacheWhitelist = [CACHE_NAME];
  event.waitUntil(
    caches.keys().then((keyList) => {
      const deletes = keyList.map((key) => {
        if (!cacheWhitelist.includes(key)) {
          console.log('Deleting cache: ' + key)
          return caches.delete(key);
        }
      });

      return Promise.all(deletes);
    })
  );
});

// When the webpage goes to fetch files, we intercept that request and serve up the matching files
// if we have them
self.addEventListener('fetch', function(event) {
  console.log("Service worker is intercepting fetch events");
  event.respondWith(
    caches.match(event.request).then(function(response) {
      return response || fetch(event.request);
    })
  );
});