package models

// Binding for JSON, tags for marshalling
type Endpoint struct {
	Name string `form:"name" json:"name" binding:"required" redis:"str1"`
	URL  string `form:"url" json:"url" binding:"required" redis:"str2"`
}

type Endpoints []Endpoint

type EndpointRepository interface {
	FindByName(name string) (*Endpoint, error)
	Save(endpoint Endpoint) error
}
