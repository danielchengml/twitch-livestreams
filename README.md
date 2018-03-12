# Twitch Live Streaming App

This Mobile Application will allow users to view Top Live Streams that are streaming on [twitch.tv](https://www.twitch.tv/). This application is build using React Native and a Go API.

## User Stories
1. As a User, I want to search for live creators currently streaming on Twitch.

1. As a User, I want to __see a list of streams__ currently streaming based on my search request.

1. As a User, I want to be able to __select a particular live stream__ and view it immediately.

1. As a User, I want to be able to __return to where I left off__ and select another stream to watch.


## Getting Started

These instructions will describe the __two different components__ of the project and direct users to setup the respective components.

_Note: The two components (Frontend & Backend) should be __setup separately__ but __run simultaneously.__ This can be done by using 2 separate Command Terminals._

It is recommended to __run__ the [Backend component](./Backend/README.md) __first__ followed by the [Frontend Component](./Frontend/README.md).


### Back End API / Server

 The Back End of this application will __serve information__ to the mobile Front End based on the search query. The backend component also acts as a gateway that has the capacity to manage multiple APIs. The current API for this application uses _[Twitch Developers API](https://dev.twitch.tv/api)_.

 The backend of this application was written in [Go](https://golang.org/) which is an extremely fast and robust web server. The Go API will serve and handle request from the front end via a local port ie. `localhost:8080`.

 Detailed setup instructions can be found in `twitch-livestreams` > `Backend` > `README.md`
  or [here](./Backend/README.md).

### Front End Mobile Application
The Front End of this application will allow users to search for live streams on Twitch.tv using a _keyword search query_. This application was written in React Native which is a popular language for mobile applications.

<img src="./Frontend/images/mobile-app-ui.jpeg" height="400">

The Front End will receive information from the Backend through API GET Requests. The server is currently designed to run locally and will serve on `localhost:8080`.

Detailed setup instructions can be found in `twitch-livestreams` > `Frontend` > `README.md`
 or [here](./Frontend/README.md).

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
