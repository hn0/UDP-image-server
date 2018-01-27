UDP image server
================

Is a simple application for serving images on UDP port 20000 designed to be used on local networks with good connectivity.
A goal of the application is to optimize image scaling and serving for processing and delivery speeds. As compiled application image scale performance, in respect to scripting languages like php is significantly better, especially in environments where, in size, large images need to be server from a server that doesn't serve web pages (eg. file server).

In such scenarios, UDP communication between servers present itself as a logical choice for performance boost.

dependencies
------------

ImageMagick 7.x libs must be installed on the system in order to use application [url:'https://github.com/gographics/imagick']


build
-----

After dependencies are present on the system a go tool can be used for building binaries.

Eg.
```code
go build -v -o bin/img_server app
```

usage
-----

A root folder containing images to be served must be passed to the application as first argument. When making requests from the client, a relative path from the root folder must be used for getting the images.

In addition to serving images application performs "on the fly" image scaling where size arguments are provided as a part of the request. Image paths sent to image server must be relative in respect to image root path on the server. A simple test scripts are provided with details on UDP request structure.

Eg.
```code
./img_server /path/to/img_root
```

php bindings
------------

A simple class for php binding is provided that can be used for obtaining base64 encoded image representation. Class provides a simple abstraction of communication with a image server and can be easily incorporated in other php projects as needed.


performance
-----------

TODO
