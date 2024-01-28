package common

// Output error object.
type OutError struct {
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#kody-oshibok5
	// https://developers.sber.ru/docs/ru/gigachat/api/authorization#ispolzovanie-tokena2
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}
