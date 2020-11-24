package main

//go:generate swag init


// nolint:lll
//go:generate mockgen -destination ./internal/app/payment/database/mock/database.go -package mockDatabase -source ./internal/app/payment/database/database.go Database


// @title Payment System API
// @version 1.0
// @description Backend API Payment System
// @termsOfService http://swagger.io/terms/

// @x-extension-openapi {"example": "value on a json format"}

func main() {}
