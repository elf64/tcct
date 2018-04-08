## Twitch Chat Client for Terminal made in go
Screenshots:

![alt text](https://github.com/pukapy/tcct/blob/master/screenshots/gitt.png)

## Usage
1. Edit conf.json with your username, channel, api_key(oauth2)
```json
{
  "Name": "--",
  "Server": "irc.twitch.tv",
  "Api_key": "--",
  "Channel": "--",
  "Port": "6667"
}
```
2. go run start.go

### Other stuff
* I'm using the [gocui](https://github.com/jroimartin/gocui) package for Command Line GUI
* Currently it only supports connecting to 1 twitch channel
