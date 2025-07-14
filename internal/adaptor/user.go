package adaptor

import (
	"fmt"
	"go-33/internal/data/entity"
	"go-33/internal/usecase"
	"go-33/pkg/codes"
	"go-33/pkg/response"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type HandlerUser struct {
	User   usecase.UserService
	Logger *zap.Logger
}

func NewHandlerUser(user usecase.UserService, logger *zap.Logger) HandlerUser {
	return HandlerUser{
		User:   user,
		Logger: logger,
	}
}

func (h *HandlerUser) Register(w http.ResponseWriter, r *http.Request) {
	// Baca form multipart
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		response.ResponseBadRequest(w, http.StatusBadRequest, "failed to parse form-data")
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Ambil file photo
	file, header, err := r.FormFile("photo")
	if err != nil {
		response.ResponseBadRequest(w, http.StatusBadRequest, "photo is required")
		return
	}
	defer file.Close()

	// Simpan file ke server (misalnya ke folder uploads/)
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)

	hashedPassword, err := codes.GeneratePassword(password)
	if err != nil {
		response.ResponseBadRequest(w, http.StatusBadRequest, "failed to hash password")
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Password: *hashedPassword,
		Photo:    filename, // simpan nama file saja ke DB
	}

	// simpan ke DB
	err = h.User.Create(&user, file)
	if err != nil {
		response.ResponseBadRequest(w, http.StatusBadRequest, "Register failed")
		return
	}

	// kosongkan password di response
	user.Password = ""

	response.ResponseSuccess(w, http.StatusCreated, "success register", user)
}

func (h *HandlerUser) ListUser(w http.ResponseWriter, r *http.Request) {
	// type key string
	ctx := r.Context()
	userid := ctx.Value("userid")
	fmt.Println("user id on handler", userid)
	users, err := h.User.List(ctx)
	if err != nil {
		response.ResponseBadRequest(w, http.StatusBadRequest, "not found")
		return
	}
	response.ResponseSuccess(w, http.StatusOK, "success", users)
}
