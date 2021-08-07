package models

type Auth struct {
	ID       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth Auth

	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}
