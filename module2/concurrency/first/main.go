package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var actions = []string{
	"logged in",
	"logged out", 
	"delete record",
	"create record",
	"update record",
}

type logItem struct {
	action string
	timestamp time.Time
}

type User struct{
	id int
	email string
	logs []logItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("Id: %d | Email: %s\nActivity log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.timestamp)
	}
	return out
}

func main(){
	rand.Seed(time.Now().Unix())

	users := generateUsers(10)

	for _, user := range users {
		saveUserInfo(user)
	}
}


func saveUserInfo(user User) error {
	fmt.Printf("writing file for user id: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil{
		return err
	}

	_, err = file.WriteString(user.getActivityInfo())
	return err
}


func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id: i+1,
			email: fmt.Sprintf("user%d@gmail.com", i+1),
			logs: generateLogs(rand.Intn(10)),
		}
	}
	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i:=0; i<count; i++ {
		logs[i] = logItem{
			timestamp: time.Now(),
			action: actions[rand.Intn(len(actions) - 1)],
		}
	}
	return logs
}