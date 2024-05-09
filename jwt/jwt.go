package jwt

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ngonzalezo/twitterGo/models"
	"log"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	fmt.Println("Clave generada que es segun invalida: ")
	fmt.Printf("%v %v", tokenStr, err)
	fmt.Println("")
	if err != nil {
		fmt.Println("Hubo un error chato!!!", err.Error())
		return createRandomKey()
	}
	fmt.Println("Todo bien al parecer")
	return tokenStr, nil
}

func createRandomKey() (string, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tokenString)
	return tokenString, nil
}
