# Simple Auth
This is simple auth build with golang and VueJs

# Golang API
## Installation
go get github.com/rorikurniadi/simple-auth

## Config
modify config in configs directory. (configs/config)

## Run Services
```
go run main.go
```

# Front End Strategy with VueJs
## Installation
```
cd views
npm install
```

## Config
You need modify endpoint golang in views/src/services/main.VueJs

## Run
```
npm run dev
```

## To Do
- Forgot Password
- Send Email (Forgot Password / Activation)