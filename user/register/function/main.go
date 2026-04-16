package user_register

import (
	"context"
	"encoding/json"
	"net/http"

	database_models "github.com/atharvyadav96k/bus-safty-app/database/models"
	"github.com/atharvyadav96k/bus-safty/user/register/applyaer"
	"github.com/atharvyadav96k/gcp/common/entity"
	"github.com/atharvyadav96k/gcp/common/res"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	app := applyaer.Init()
	defer app.Close()
	var user database_models.User
	json.NewDecoder(r.Body).Decode(&user)
	ctx := context.Background()
	errs := entity.ValidateStruct(user)
	if len(errs) != 0 {
		res.BadRequest(w, errs)
	}
	app.FireStore.FirestoreClient.Collection("users").Add(ctx, user)
	res.Send(w, 201, "Created", user)
}
