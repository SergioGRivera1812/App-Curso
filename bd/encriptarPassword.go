package bd

import "golang.org/x/crypto/bcrypt"

//sirve para encriptar las password
func EncriptarPassword(pass string) (string, error) {
	costo := 8 // Es la cantidad elevada al cuadrado da pasadas para encriptar mientras mas alto mejor
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
