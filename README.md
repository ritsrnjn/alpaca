# alpaca

Note: Alpaca will act as an backend server for the game.

Steps to run on local
1. Install go
2. cd to home directory of repo
3. go run cmd/main.go

Ideally the server should start running on :8080
Then use one of these 2 curl command to verify
1. curl --location --request GET 'http://localhost:8080/'
2. curl --location --request GET 'http://localhost:8080/game'


Also, current version is already running on heroku, use thse curl to verify.
1. curl --location --request GET 'https://codeingerman-alpaca.herokuapp.com'
2. curl --location --request GET 'https://codeingerman-alpaca.herokuapp.com/game'



Project structure guidline: https://github.com/golang-standards/project-layout
Golang naming guidline: https://go.dev/doc/effective_go