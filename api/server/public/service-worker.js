// Name our cache
var CACHE_NAME = 'popcorn-pwa-cache';

// Delete old caches that are not our current one!
self.addEventListener("activate", event => {
  console.log("activate is fired");
  const cacheWhitelist = [CACHE_NAME];
  event.waitUntil(
    caches.keys()
      .then(keyList =>
        Promise.all(keyList.map(key => {
          if (!cacheWhitelist.includes(key)) {
            console.log('Deleting cache: ' + key)
            return caches.delete(key);
          }
        }))
      )
  );
});

// The first time the user starts up the PWA, 'install' is triggered.
self.addEventListener('install', function(event) {
  console.log("install is fired");
  const p = caches.open(CACHE_NAME).then(function(cache) {
    const urlsToCache = ["/index.html", "/index.js", "/icon.png"];
    cache.addAll(urlsToCache);
    console.log('cached', urlsToCache);
  })
  event.waitUntil(p);
});

// When the webpage goes to fetch files, we intercept that request and serve up the matching files
// if we have them
self.addEventListener('fetch', function(event) {
  console.log("fetch is fired");
  event.respondWith(
    caches.match(event.request).then(function(response) {
      return response || fetch(event.request);
    })
  );
});