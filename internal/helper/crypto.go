package helper

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken is a function for generate token using JWT
func GenerateToken(claim any) (token string, err error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	jwtClaim := jwt.MapClaims{}
	encodedClaim, err := json.Marshal(claim)
	if err != nil {
		fmt.Println("Error encoding claim: ", err)
		return "", err
	}

	err = json.Unmarshal(encodedClaim, &jwtClaim)
	if err != nil {
		fmt.Println("Error decoding claim: ", err)
		return "", err
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwtClaim)

	token, err = parseToken.SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Println("Error generating token: ", err)
		return "", err
	}

	return token, nil
}