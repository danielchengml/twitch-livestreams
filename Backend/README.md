# Twitch Live Streaming App (Backend API Server)

This Mobile Application will allow users to view Top Live Streams that are streaming on [twitch.tv](https://www.twitch.tv/). The Backend of this application is built using Go.

This document is a step-by-step guide to setup the Backend Development Environment using either macOS or Windows.


## Getting Started

These instructions will get you a copy of the Backend API/  Server up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

 - [The Go Programming Language](https://golang.org/dl/)
 - Any Web Browser (Chrome, Edge, Firefox)
 - (optional) [Postman](https://www.getpostman.com/)

### Installing

Here are the instructions to get the program up and running on the machine. *Do ensure that prerequisites are installed on your machine*

1. __Open the terminal__ and navigate to the "twitch-livestreams/Backend" Folder
```
$ cd twitch-livestreams
$ cd Backend
```
2. __Get and install__ the gorilla/mux Dependency
```
$ go get -u github.com/gorilla/mux
```
3. __Build the project__:
```
$ go build
```
4. __Run the project__:
```
$ go run main.go
```
The project will now run on `localhost:8080`.

## Running the tests

Once the project is running, the connection can be tested using Postman API or a Web Browser using the following url:
```
http://localhost:8080/livestreams/overwatch
```
*Note: ".../overwatch" is just a search example, you could use any other word instead*
Here is an example of the return you would expect to get:
```
{  
   "_total":1424,
   "streams":[  
      {  
         "_id":27906100720,
         "game":"Overwatch",
         "viewers":15092,
         "video_height":900,
         "average_fps":60.037525,
         "delay":0,
         "created_at":"2018-03-12T14:04:33Z",
         "is_playlist":false,
         "stream_type":"live",
         "preview":{  
            "small":"https://static-cdn.jtvnw.net/previews-ttv/live_user_timthetatman-80x45.jpg",
            "medium":"https://static-cdn.jtvnw.net/previews-ttv/live_user_timthetatman-320x180.jpg",
            "large":"https://static-cdn.jtvnw.net/previews-ttv/live_user_timthetatman-640x360.jpg",
            "template":"https://static-cdn.jtvnw.net/previews-ttv/live_user_timthetatman-{width}x{height}.jpg"
         },
         ...
```


## Setting up the Front End Environment

Once the backend server is running and returning the a valid response, the front end development environment can now be setup. The front end application is set to connect to `localhost:8080` on the backend.

The instructions to deploy the front end module can be found under the folder `twitch-livestreams` > `Frontend` > `README.md`

## Built With

* [Visual Studio Code](https://code.visualstudio.com/) - The IDE used
* [Go Programming Language](https://golang.org/) - The Programming Language Used
* [gorilla/mux](https://github.com/gorilla/mux) - URL router and dispatcher for Go


## Versioning

Version 1.0.0

## Authors

* **Daniel Cheng** -  [github.com/danielchengml](https://github.com/danielchengml)

## Acknowledgments

* Tomasz Wiszkowski, Google
* David Chang, ReactJS Developer, Airbnb
* Anbiniyar Muniandy, Full Stack Developer, Microsoft

## Sources: Twitch API (https://dev.twitch.tv/docs/api)

- API for Top-100 Live Streams based on Search Query Viewership:
https://api.twitch.tv/kraken/search/streams?limit=100&client_id={API_KEY}
