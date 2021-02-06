package models

type Address struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	ZipCode      string `json:"zipCode"`
	City         string `json:"city"`
	StateCountry string `json:"stateCountry"`
	Country      string `json:"country"`
	Email        string `json:"email"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

//func (a *Address)  {
//}