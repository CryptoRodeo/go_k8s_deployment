package data

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var USERS = []User{
	{
		Id:   1,
		Name: "Joey",
		Role: "Manager",
	},
	{
		Id:   2,
		Name: "Sandy",
		Role: "Developer",
	},
	{
		Id:   3,
		Name: "Anthony",
		Role: "IT",
	},
}
