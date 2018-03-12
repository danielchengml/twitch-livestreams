# Twitch Live Streaming App (Backend Web Server)

This Mobile Application will allow users to view Top Live Streams that are streaming on twitch.tv. This application is build using Go and React Native.

## User Stories
1. As a User, I want to see the Top 100 Games that are currently live. Upon selecting the game, I want to see the top 100 streams currently streaming for that game. Upon selecting the game, I want to watch the livestream in my browser.

2. As a User, I want to see the Top 100 Streams. I can then select the stream that I want to watch.


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

```
$ go build
```

### Installing

A step by step series of examples that tell you have to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc

## Sources: Twitch API (https://dev.twitch.tv/docs/api)

- API for Top-100 Games that are Streaming Live:
https://api.twitch.tv/kraken/games/top?limit=100&client_id={API_KEY}
- API for Top-100 Live Streams based on Viewership:
https://api.twitch.tv/kraken/streams?limit=100&client_id={API_KEY}
- API for Top-100 Live Streams based for Selected Game: https://api.twitch.tv/kraken/streams?game={GAME_NAME}&limit=100&client_id={API_KEY}
