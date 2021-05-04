package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets80c0d7c4b5f47513d79ddc23ee8b0f4488cb1b9b = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <title>Primes of sieve</title>\n  <link rel=\"stylesheet\" href=\"/public/index.css\">\n</head>\n<body>\n  <div id=\"root\"></div>\n  <script type=\"text/javascript\" src=\"/public/index.js\"></script>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"views"}, "/views": []string{"home.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1620145568, 1620145568996107222),
		Data:     nil,
	}, "/views": &assets.File{
		Path:     "/views",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1620137450, 1620137450901192792),
		Data:     nil,
	}, "/views/home.html": &assets.File{
		Path:     "/views/home.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1620137450, 1620137450900942329),
		Data:     []byte(_Assets80c0d7c4b5f47513d79ddc23ee8b0f4488cb1b9b),
	}}, "")
