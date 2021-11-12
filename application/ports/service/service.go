package service

type (
	Default interface {
		Create(domain interface{}) (interface{}, error)
		Update(id string, domain interface{}) (interface{}, error)
		GetOne(id string) (interface{}, error)
		List(query string) ([]interface{}, error)
		Delete(id string) (interface{}, error)
	}
)
