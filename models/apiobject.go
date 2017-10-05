package models

// APIObject Structure provides the structure for call to an API
type APIObject struct{
	Name	string	`json:"Name"`
	URI	string	`json:"URI"`
	Method	string	`json:"Method"`
}
