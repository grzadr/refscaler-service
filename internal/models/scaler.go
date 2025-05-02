package models

// EnlistmentRequest contains the input for scaling an enlistment
// @Description Request parameters for scaling an enlistment
type EnlistmentRequest struct {
	Enlistment string `json:"enlistment" example:"Item 1: 1 year\nItem2: 1 month"`
	Scale      string `json:"scale" example:"1 minute"`
}

// EnlistmentResponse contains the scaled results
// @Description Response containing scaled enlistment values
type EnlistmentResponse struct {
	Scaled []string `json:"scaled" example:"[\"Item 1: 1 minute\",\"Item 2: 5 seconds\"]"`
}
