package http

import (
	"io"
	"log"
	"net/http"

	"board-games/internal/auth/delivery/mappers"
	usecase "board-games/internal/auth/usecase/service"

	"github.com/google/uuid"
)

type AuthHandlers struct {
	userManager *usecase.UserManager
}

func NewAuthHandlers(userManager *usecase.UserManager) *AuthHandlers {
	return &AuthHandlers{
		userManager: userManager,
	}
}

func (h *AuthHandlers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusBadRequest)
			return
		}
		input, err := mappers.ToRegisterRequest(data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusBadRequest)
			return
		}
		user := mappers.FromRegisterRequestToUser(input)
		if err := h.userManager.Register(user); err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		response, err := mappers.ToRegisterResponse(user.UserID.String())
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (h *AuthHandlers) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusBadRequest)
			return
		}
		input, err := mappers.ToLoginRequest(data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusBadRequest)
			return
		}
		user := mappers.FromLoginRequestToUser(input)
		if err := h.userManager.Login(user); err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		response, err := mappers.ToLoginResponse(user.UserID.String())
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (h *AuthHandlers) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not implemented", http.StatusNotImplemented)
	}
}

func (h *AuthHandlers) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusBadRequest)
			return
		}
		user, err := h.userManager.GetByID(userID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		response, err := mappers.ToGetByIDResponse(user)
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (h *AuthHandlers) FindByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.PathValue("username")
		if username == "" {
			log.Println("username is empty")
			http.Error(w, "Auth error:", http.StatusBadRequest)
			return
		}
		user, err := h.userManager.FindByUsername(username)
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		response, err := mappers.ToFindByUsernameResponse(user)
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (h *AuthHandlers) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusBadRequest)
			return
		}
		if err := h.userManager.Delete(userID); err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		response, err := mappers.ToDeleteResponse(userID.String())
		if err != nil {
			log.Println(err)
			http.Error(w, "Auth error:", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
