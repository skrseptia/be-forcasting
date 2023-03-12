package listing

import "food_delivery_api/pkg/util"

type User struct {
	util.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ImageURL string `json:"image_url"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	UserType string `json:"user_type"`
}
