package user

type UserSerialize struct {
	// This is not the model, more like a serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func SerializeAll(users *[]User) []UserSerialize{
	responseUsers := []UserSerialize{}
	for _, user := range *users {
		responseUser := Serialize(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return responseUsers
}

func Serialize(user User) UserSerialize {
	return UserSerialize{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}