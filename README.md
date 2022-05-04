# Checkout System 1.0

## Product list preview
![image](https://user-images.githubusercontent.com/32862937/166623305-2b5aa5ec-b84f-4ba9-8502-3b04f25df0c8.png)
![image](https://user-images.githubusercontent.com/32862937/166621498-11f52ec3-5125-4b2c-986e-ac9d67bc74e3.png)

## Promotion features
Current version 1.0
Support 3 rules of promotion
- Each sale of a MacBook Pro comes with a free Raspberry Pi B
- Buy 3 Google Homes for the price of 2
- Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers

## Checkout result preview
![image](https://user-images.githubusercontent.com/32862937/166621829-0b075d39-fb6d-427b-a486-d5095653d6b3.png)

## Layout
```tree
├── .github
│   └── workflows
│       └── audit.yaml
├── .gitignore
├── graph
│   └── schema.graphqls
├── model
│   └── product.go
├── service
│   ├── cart_test.go
│   ├── cart.go
│   ├── product.go
│   ├── promotion_test.go
│   └── promotion.go
├── go.mod
├── go.sum
├── README.md
├── LICENSE
├── main.go
```

## Prerequisites
- Programming Language: Golang (latest version, https://go.dev/dl/)

## Trial palyground in main file preview
![image](https://user-images.githubusercontent.com/32862937/166622537-5ede0736-e4b2-40b6-adb1-6b8304cd4f7e.png)

## How to run the project

1. Install dependencies 
```go mod download```

2. Run the application 
```go run main.go```
