# Nuntius Go

Nuntius is an interactive anonymous messaging webapp with a dare game. Built with Go, HTMX, Templ and PostgresQL. This is a clone of the [original nuntius](https://github.com/chizidotdev/nuntius) built with NextJS...
Goal of this project was to experiment with the golang, htmx and templ stack.

### Explored Features
- Templ + HTMX for server side rendering and client side interactivity
- Google oauth2
- Session storage with gin and postgres
- Send and receive anonymous messages
- Clean architecture pattern

## Development
Ensure you have [docker](https://docs.docker.com/engine/install) and the [compose plugin](https://docs.docker.com/compose/install) installed and setup on your machine

Run the following command to start the server
```bash
docker compose up
```
Navigate to `localhost:8080` on your browser to view the app

### Architecture Overview
```bash
cmd/ # Main application entry point
config/ # Application configuration ( env variables )
internal/
   |-- app/
   |   |-- components/ # Htmx + Templ components
   |   |-- db/ # Database connection
   |   |-- drivers/ # External services ( HTTP client )
   |-- core/
   |   |-- domain/ # Core domain data models
   |   |-- service/ # Core business logic
```

### TODO
- [x] Containerize with docker
- [ ] Websockets implementation
 
### Contact
- 𝕏 [@chizidotdev](https://twitter.com/chizidotdev)