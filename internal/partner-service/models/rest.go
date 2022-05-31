package models

type CreatePartnerInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type CreatePartnerOutput struct {
	Id string `json:"id"`
}

type DeletePartnerInput struct {
	Id string `json:"id"`
}
