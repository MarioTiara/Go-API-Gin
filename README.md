# Golang Web API using Gin

### Project Features
* Using Gin for CRUD Demonstration
* Using JWT as Token via [jwt-go package](https://github.com/dgrijalva/jwt-go)
* Using Reposotiry Pattern to communicate to DB

### Project Dependencies
* Install and Run Postgress service in your localhost for storing data

### Aplication LifeCycle
- Install node modules
  ```
  $ go get . || go mod || make goinstall
  ```
- Build application
  ```
  $ go build -o main || make goprod
  ```
- Start application in development
  ```
 $ go run main.go | make godev
  ```
**Author**
[Mario Tiara Pratama](https://github.com/MarioTiara)