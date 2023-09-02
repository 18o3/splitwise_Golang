package Handlers

type User struct {
	name         string
	transactions []string
	UserID       int
	balance      int
	UPI          string
}

var countID int
var NameToID map[string]int
var Users []User

func Init() {
	countID = 1
	NameToID = make(map[string]int)
}
