module.exports = {
    "globDirectory": "./src/public",
    "globPatterns": [
        "**/*.{js,css,png,svg,jpg,gif,json,woff,woff2,eot,ttf,ico,webmanifest,map}"
    ],
    "swDest": "./src/public/sw.js",
    // Define runtime caching rules.
    "runtimeCaching": [{
        // Match any request that ends with .png, .jpg, .jpeg or .svg.
        "urlPattern": /\.(?:png|jpg|jpeg|svg|woff2)$/,

        // Apply a cache-first strategy.
        "handler": 'CacheFirst',

        "options": {
            // Use a custom cache name.
            "cacheName": 'images',

            // Only cache 50 images.
            "expiration": {
                "maxEntries": 50,
            },
        },
    }],
};