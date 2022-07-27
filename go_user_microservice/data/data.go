package data

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Tickets []int  `json:"tickets"`
}

var USERS = []User{
	{
		Id:      1,
		Name:    "Joey",
		Role:    "Manager",
		Tickets: []int{1800},
	},
	{
		Id:      2,
		Name:    "Sandy",
		Role:    "Developer",
		Tickets: []int{1500},
	},
	{
		Id:      3,
		Name:    "Anthony",
		Role:    "IT",
		Tickets: []int{2000, 2001},
	},
}
