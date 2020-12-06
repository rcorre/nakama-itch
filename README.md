# Nakama Itch

This is a [nakama](https://github.com/heroiclabs/nakama) module that uses [custom authentication](https://heroiclabs.com/docs/authentication/#custom) to authenticate a user via [itch.io](https://itch.io/).

# Usage

1. Create an [itch app manifest](https://itch.io/docs/itch/integrating/manifest.html) for your game
2. Run your game from the itch.io app
3. Extract the API token from the `ITCHIO_API_KEY` environment variable
4. Pass the API token as the ID to `authenticate_custom`

Godot example:

```
var token := OS.get_environment("ITCHIO_API_KEY")
var session: NakamaSession = yield(client.authenticate_custom_async(token), "completed")
```

# Build and run

```
build.sh && docker-compose up
```
