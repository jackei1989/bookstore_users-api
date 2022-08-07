package bootstrap

import "bookstoreUsersApi/controllers/users"

func mapUrls() {
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
