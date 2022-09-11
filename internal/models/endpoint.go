package models

// Binding for JSON, tags for marshalling
type Endpoint struct {
	Name string `form:"name" json:"name" binding:"required" redis:"str1"`
	URL  string `form:"url" json:"url" binding:"required" redis:"str2"`
}

type EndpointRepository interface {
	FindByName(name string) (*Endpoint, error)
	FindAll() ([]*Endpoint, error)
	FindIfExists(name string) (bool, error)
	FindAllKeys() ([]string, error)
	Save(endpoint Endpoint) error
}
