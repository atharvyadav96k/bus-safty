package user_register

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"errors"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/user/register/applyaer"
	"github.com/atharvyadav96k/gcp/common/entity"
	"github.com/atharvyadav96k/gcp/common/res"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res.BadRequest(w, []error{errors.New("Method not allowed")})
		return
	}
	app := applyaer.Init()
	defer app.Close()
	var user database_models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	ctx := context.Background()
	errs := entity.ValidateStruct(&user)
	if len(errs) != 0 {
		res.BadRequest(w, errs)
		return
	}
	docRef := app.FireStore.FirestoreClient.Collection("users").Doc(user.WhiteListedEmailID.String())
	_, err := docRef.Create(ctx, user)
	fmt.Println(err)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			err = errors.New("user already exits with this email")
		}
		res.BadRequest(w, []error{err})
		return
	}
	res.Send(w, 201, "Created", user)
}
