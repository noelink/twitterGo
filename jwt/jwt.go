package jwt

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ngonzalezo/twitterGo/models"
	"time"
)

func GeneroJWT(ctx context.Context, t models.Usuario) (string, error) {
	fmt.Println("Entre a generacion de jwt")
	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	miClave := []byte(jwtSign)
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		fmt.Println("Hubo un error chato!!!", err.Error())
		return tokenStr, err
	}
	fmt.Println("Todo bien al parecer")
	return tokenStr, nil
}
