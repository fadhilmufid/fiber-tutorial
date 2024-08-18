package user

import "encoding/json"

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

func SerializeJSON(user User) map[string]interface{} {
    return map[string]interface{}{
        "id":   user.ID,
        "first_name": user.FirstName,
        "last_name": user.LastName,

    }
}

func SerializeAllJSON(users *[]User) string{
	responseUsers, err := json.Marshal(users)
	if err != nil {
		return ""
	}
	jsonString := string(responseUsers)
	return jsonString
}