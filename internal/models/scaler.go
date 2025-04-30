package models

type EnlistmentRequest struct {
	Enlistment string `json:"enlistment"`
	Scale      string `json:"scale"`
}

type EnlistmentResponse struct {
	Scaled []string `json:"scaled"`
}
