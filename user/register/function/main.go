package user_register

import (
	"context"
	"net/http"

	"github.com/atharvyadav96k/bus-safty/user/register/applyaer"
	"github.com/atharvyadav96k/gcp/common/res"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	app := applyaer.Init()
	defer app.Close()
	ctx := context.Background()
	user := struct {
		Name string `json:"name"`
	}{
		Name: "Atharv",
	}
	app.FireStore.FirestoreClient.Collection("users").Add(ctx, user)
	res.Send(w, "201", "Created", user)
}
