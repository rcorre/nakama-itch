docker run --rm -w "/builder" -v "${PWD}:/builder" heroiclabs/nakama-pluginbuilder:2.15.0 build -buildmode=plugin -trimpath -o ./modules/nakama_itch.so
