package handlers

import (
	"net/http"
)

// AddNewItem - handle add new item
func (env Env) AddNewItem(w http.ResponseWriter, r *http.Request) {
	topupRequest := &AddNewItemRequestBody{}
	if err := topupRequest.Validate(r); err != nil {
		HandleFailedResponse(err.Error(), w)
		return
	}

	JSONResponse("SUCCESS", "Your card has been top-up successfully", w)
}

// EditItem - handle editing a new item
func (env Env) EditItem(w http.ResponseWriter, r *http.Request) {
	JSONResponse("SUCCESS", "Implement to edit an item", w)
}

// GetAllItems - get all food items
func (env Env) GetAllItems(w http.ResponseWriter, r *http.Request) {
	JSONResponse("SUCCESS", "Implement to return food item list", w)
}
