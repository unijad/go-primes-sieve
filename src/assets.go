package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets80c0d7c4b5f47513d79ddc23ee8b0f4488cb1b9b = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <title>Primes of sieve</title>\n</head>\n<body>\n  <h1>Hello Home</h1>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"views"}, "/views": []string{"home.html"}}, map[string]*assets.File{
	"/views": &assets.File{
		Path:     "/views",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1620134708, 1620134708265303344),
		Data:     nil,
	}, "/views/home.html": &assets.File{
		Path:     "/views/home.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1620134708, 1620134708265166301),
		Data:     []byte(_Assets80c0d7c4b5f47513d79ddc23ee8b0f4488cb1b9b),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1620134651, 1620134651667462515),
		Data:     nil,
	}}, "")
