package user

type (
	User struct {
		ID   string `json:"id" form:"id"`
		Name string `json:"name" form:"name"`
	}
)
