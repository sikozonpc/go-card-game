# Golang card game implementation

Client-server project built in Golang and Typescript.

The server is constructed of a REST API and a Webscoket route that serves the real time matches and the rest api is used for operations outside of the match (create a game, join a game, etc...)

On the server runs the whole game engine, the client only renders the UI and calls the actions, while in the server all of the logic behind  the game is taken care of.
