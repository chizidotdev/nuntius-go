# Nuntius Go

Nuntius is an interactive anonymous messaging webapp with a dare game. Built with Go, HTMX, Templ and PostgresQL. This is a clone of the [original nuntius](https://github.com/chizidotdev/nuntius) built with NextJS...
Goal of this project was to experiment with htmx and templ.

### Explored Features of this application
- Send and receive anonymous messages
- Google oauth
- Session storage with gin and postgres
- Clean architecture

## Development
#### Prerequisites
Ensure you have these installed
- [templ cli](https://templ.guide/quick-start/installation)
- [golang air](https://github.com/cosmtrek/air)

Run the following command to start the server
```bash
make server
```

### TODO
- Containerize the app
- Websockets implementation
 
