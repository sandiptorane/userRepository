package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"userRepository/internal/vipers"
)

// jwt secretKey
var jwtKey = vipers.GetJwtKey()

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//SetToken will set the  token for signed user
func CreateToken(userName string, w http.ResponseWriter) (string, error) {
	log.Println("set token for signed user")
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return "", err
	}
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path:    "/",
		Expires: expirationTime,
	})
	return tokenString, nil
}

//IsAuthorized authorise protected endpoints
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		c, err := req.Cookie("token")
		if err != nil {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Not Authorized")
			return
		}
		tokenString := c.Value
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		endpoint(w, req)

	}
}

//ClearToken delete the token of signed user while sign out
func ClearToken(w http.ResponseWriter,req *http.Request) {
	c, err := req.Cookie("token")
	if err != nil {
		// If the cookie is not set, return an unauthorized status
		w.WriteHeader(http.StatusUnauthorized)
		log.Errorln("cookie error:", err)
		return
	}
	c.Value = ""
	c.Name = "token"
	c.MaxAge = -1
	http.SetCookie(w, c)
	fmt.Fprintln(w, "signed out successfully")
	log.Println("signed out successfully")
}

//GetUsername will return user if signed in
func GetUserName(w http.ResponseWriter,req *http.Request) (userName string) {
	c, err := req.Cookie("token")
	if err != nil {
		// If the cookie is not set, return an unauthorized status
		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString := c.Value
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	return claims.Username
}
