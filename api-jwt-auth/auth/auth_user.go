package auth

import (
	"fmt"
	"net/http"

	"strconv"
	"strings"
	"time"

	"github.com/austineyoyogie/go-hardware-store/configs"
	"github.com/austineyoyogie/go-hardware-store/utils"
)

var c = configs.LoadConfigs()

func GenerateJWT(user *User) (string, error) {
	claims := jwt.StandardClaims{
		Issuer: string(user.Email), //-> good to use as well
		//Issuer: strconv.Itoa(int(user.ID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.JWT.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func RefreshJWT(user *User) (string, error) {
	claims := jwt.StandardClaims{
		//Issuer: string(user.Email), -> good to use as well
		Issuer:    strconv.Itoa(int(user.ID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshString, err := refresh.SignedString([]byte(c.JWT.SecretKey))
	if err != nil {
		return "", err
	}
	return refreshString, nil
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func IsAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authToken := ExtractToken(r)

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(c.JWT.SecretKey), nil
		})

		if err != nil {
			utils.ResponseWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			utils.ResponseWithError(w, http.StatusUnauthorized, "Not Authorizated")
		}
	})
}
