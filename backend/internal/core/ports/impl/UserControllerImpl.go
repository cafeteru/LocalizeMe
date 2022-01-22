package impl

import (
	"encoding/json"
	"github.com/go-chi/chi"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserControllerImpl struct {
	service service.UserService
}

func CreateUserController(u service.UserService) *UserControllerImpl {
	return &UserControllerImpl{u}
}

func (u UserControllerImpl) CreateUserRoutes(r *chi.Mux) {
	path := "/users"
	r.Post(path, u.Create())
	r.Get(path, u.FindAll())
}

// swagger:route POST /users Users CreateUser
// Create a new user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
func (u UserControllerImpl) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debugf("%s: start", tools.GetCurrentFuncName())
		var request dto.UserRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			ports.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
			return
		}
		user, err := u.service.Create(request)
		if err != nil {
			ports.CreateErrorResponse(w, err, http.StatusBadRequest)
			return
		}
		ports.CreateResponse(w, http.StatusCreated, user)
		slog.Debugf("%s: end", tools.GetCurrentFuncName())
	}
}

// swagger:route GET /users Users FindAll
// Return all users.
//
// Responses:
// - 200: []User
func (u UserControllerImpl) FindAll() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debugf("%s: start", tools.GetCurrentFuncName())
		users := [2]domain.User{}
		users[0] = domain.User{
			ID:       primitive.ObjectID{},
			Email:    "",
			Password: "",
			IsAdmin:  false,
			IsActive: false,
		}
		users[1] = domain.User{
			ID:       primitive.ObjectID{},
			Email:    "",
			Password: "",
			IsAdmin:  true,
			IsActive: true,
		}
		ports.CreateResponse(w, http.StatusOK, users)
		slog.Debugf("%s: end", tools.GetCurrentFuncName())
	}
}
