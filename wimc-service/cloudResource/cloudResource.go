package cloudResource

type CloudResource struct {
	CloudResourceId int    `json:"cloudResourceId"`
	CloudId         string `json:"cloudId"`
	Location        string `json:"location"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Notes           string `json:"notes"`
}
