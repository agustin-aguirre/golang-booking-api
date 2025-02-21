package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretlyconfidentialkey123$"

func GenerateToken(userId int64, email string) (string, error) {
	signingMethod := jwt.SigningMethodHS256
	claims := jwt.MapClaims{
		"email":      "",
		"userId":     "",
		"expiration": time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(signingMethod, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	keyFunc := func(jwtToken *jwt.Token) (any, error) {
		// verificar si el token está firmado con el mismo método que esperamos
		_, signingMethodMatches := jwtToken.Method.(*jwt.SigningMethodHMAC) //sintaxis de "type assertion"

		// de Copilot:
		// 	Una type assertion en Go se usa para convertir un valor de un tipo interface{}
		// a un tipo más específico.
		// 	En este caso, estamos tratando de convertir el valor de jwtToken.Method
		// al tipo *jwt.SigningMethodHMAC.
		// 	La sintaxis .(*TipoEspecifico) se usa para realizar esta conversión.
		// 	Si el valor en jwtToken.Method no es de tipo *jwt.SigningMethodHMAC,
		// se producirá un pánico (panic) en tiempo de ejecución.

		if !signingMethodMatches {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	}
	parsedToken, err := jwt.Parse(token, keyFunc)

	if err != nil {
		return errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return errors.New("token is not valid")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("invalid token claim")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
