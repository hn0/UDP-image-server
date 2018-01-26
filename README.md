UDP image server
================

Is a simple application for serving images on UDP port 20000 designed to be used on local networks with good connectivity.
A goal of the application is to optimize image scaling and serving for processing and delivery speeds. As compiled application image scale performance, in respect to scripting languages like php is significantly better, especially in environments where, in size, large images need to be server from a server that doesn't serve web pages (eg. file server).

In such scenarios, UDP communication between servers present itself as a logical choice for performance boost.

dependencies
------------

ImageMagick 7.x libs must be installed on the system in order to use application


usage
-----

A root folder containing images to be served must be passed to the application as first argument. When making requests from the client, a relative path from the root folder must be used for getting the images.

In addition to serving images application performs "on the fly" image scaling where size arguments are provided as a part of the request.


php bindings
------------

A simple class for php binding is provided that can be used for obtaining base64 encoded image representation.

Class covers communication between php web server and ....
