package auth

import (
	"net/http"
)

type Handlers interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	Logout() http.HandlerFunc
	GetByID() http.HandlerFunc
	FindByUsername() http.HandlerFunc
	Delete() http.HandlerFunc
}
