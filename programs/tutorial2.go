package programs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	users map[int]*User
	mutex sync.Mutex
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*User),
	}
}

func (us *UserService) CreateUser(name, email string) *User {
	us.mutex.Lock()
	defer us.mutex.Unlock()
	fmt.Println(len(us.users))
	id := len(us.users) + 1
	user := &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	us.users[id] = user
	return user
}

func (us *UserService) GetUser(id int) *User {
	us.mutex.Lock()
	defer us.mutex.Unlock()
	return us.users[id]
}

func (us *UserService) GetAllUsers() []*User {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	var users []*User
	for _, user := range us.users {
		users = append(users, user)
	}
	return users
}

func (us *UserService) DeleteUser(id int) {
	us.mutex.Lock()
	defer us.mutex.Unlock()
	delete(us.users, id)
}

// HTTP Handlers
func (us *UserService) createUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	us.CreateUser(name, email)
	json.NewEncoder(w).Encode(us.GetAllUsers())
}

func (us *UserService) getUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	user := us.GetUser(id)
	json.NewEncoder(w).Encode(user)
}

func (us *UserService) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := us.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func (us *UserService) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	us.DeleteUser(id)
	json.NewEncoder(w).Encode(us.GetAllUsers())

	w.WriteHeader(http.StatusOK)
}
