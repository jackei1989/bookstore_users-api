package users

import "encoding/json"

type PublicUser struct {
	Id        int64  `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_At"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:        user.Id,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	_ = json.Unmarshal(userJson, &privateUser)
	return privateUser
}
