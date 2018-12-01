package service

import (
	"net/http"
)

// Route defines a single route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the type Routes which is just
// an array of Route struct
type Routes []Route

var routes = Routes{
	Route{
		Name:        "GetAccount",
		Method:      http.MethodGet,
		Pattern:     "/accounts/{accountID}",
		HandlerFunc: GetAccount,
	},
}
