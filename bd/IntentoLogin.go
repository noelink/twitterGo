package bd

import (
	"fmt"
	"github.com/ngonzalezo/twitterGo/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	fmt.Println("Ingresamos a intento login")
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		fmt.Println("VLV ocurrio un error", err.Error())
		return usu, false
	}

	return usu, true
}
