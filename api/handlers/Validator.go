package handlers

import (
	"errors"
	"net/http"
)

type ValidatorInterface interface {
	Validate(r *http.Request) bool
}

//AddNewItemRequestBody - request
type AddNewItemRequestBody struct {
	name        string
	description string
	amount      string
	timeOfDay   string
}

//Validate - validate the request
func (item *AddNewItemRequestBody) Validate(r *http.Request) error {
	name := r.FormValue("name")
	if name == "" {
		return errors.New("Valid `name` must be provided`")
	}
	item.name = name

	amount := r.FormValue("amount")
	if amount == "" {
		return errors.New("Valid `amount` must be provided`")
	}
	item.amount = amount

	description := r.FormValue("description")
	if description == "" {
		return errors.New("argument `description` must be provided`")
	}
	item.description = description

	timeOfDay := r.FormValue("timeOfDay")
	if timeOfDay == "" {
		return errors.New("argument `timeOfDay` must be provided`")
	}

	item.timeOfDay = timeOfDay

	return nil
}
