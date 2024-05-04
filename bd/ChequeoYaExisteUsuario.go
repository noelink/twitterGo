package bd

import (
	"context"
	"fmt"
	"github.com/ngonzalezo/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	fmt.Println("Ingreso a chequeoUsuario existente con email: ", email)
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	fmt.Println("resultado es : ", resultado)
	return resultado, true, ID
}
