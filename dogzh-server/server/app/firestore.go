package app

import (
	"cloud.google.com/go/firestore"
)

type App struct {
	FireStore *firestore.Client
}
