package models

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Tickets []int  `json:"tickets"`
}

func ExistsByID(id int) bool {
	for _, user := range fakeUserDB() {
		if user.Id == id {
			return true
		}
	}
	return false
}

func GetUser(id int) *User {
	for _, user := range fakeUserDB() {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func GetUsersByID(ids []int) []User {
	temp := []User{}

	for _, user := range fakeUserDB() {
		if inIdList(ids, user.Id) {
			temp = append(temp, user)
		}
	}
	return temp
}

func GetAllUsers() []User {
	return fakeUserDB()
}

func inIdList(list []int, val int) bool {
	for _, list_val := range list {
		if list_val == val {
			return true
		}
	}
	return false
}

func fakeUserDB() []User {
	return []User{
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
}
