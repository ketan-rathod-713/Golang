package userapi

import (
	"fmt"
	"net/http"
)

func (a *Api) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all users but first authorize if given user is admin or not")
}
