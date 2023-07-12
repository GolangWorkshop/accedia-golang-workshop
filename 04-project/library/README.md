golang workshop project

# Starting the project

1. Run 
```sh 
go run ./cmd/main.go
``` 
1. If you want to have hot reloading, install `nodemon` (`node` is required):
```sh
npm install nodemon -g
```
and run 
```sh 
nodemon --signal SIGTERM --exec go run ./cmd/main.go
```
