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
	app.FireStore.FirestoreClient.Collection("users").Add(ctx, user)
	res.Send(w, 201, "Created", user)
}
