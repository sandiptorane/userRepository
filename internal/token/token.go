package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
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

//Store the created token in tokenStore
var  tokenStore = make(map[string]string)

//CreateToken will generate the  token for signed user
func CreateToken(userName string, w http.ResponseWriter,r *http.Request) (string, error) {
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
		w.WriteHeader(http.StatusInternalServerError)
		return "", err
	}
	bearer := "Bearer "+tokenString
	w.Header().Set("Authorization",bearer)  //set token in Authorization Bearer Token header
	r.Header.Set("Authorization",bearer)
	tokenStore[tokenString]=claims.Username  //also store token to tokenStore
	return tokenString, nil
}

//IsAuthorized authorise protected endpoints
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("Authenticating user")
		auth := req.Header.Get("Authorization")  //jwt-token stored in Bearer Token Authorization header
		//log.Println("auth token:",auth)
		tokenPart := strings.Split(auth," ")
		if len(tokenPart)!=2{
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w,"Not Authorized")
			return
		}
		tokenString := tokenPart[1]  //tokenPart[1] contains actual token string

		if _,found := tokenStore[tokenString];!found{  //check user token present in tokenStore
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w,"Not Authorized")
			return
		}

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
	auth := req.Header.Get("Authorization")  //jwt-token stored in Bearer Token Authorization header
	//log.Println("auth token:",auth)
	tokenPart := strings.Split(auth," ")
	if len(tokenPart)!=2{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w,"Not Authorized")
		return
	}
	tokenString := tokenPart[1]  //tokenPart[1] contains actual token string
	delete(tokenStore,tokenString)  //delete tokenString from tokenStore
	w.Header().Set("Authorization","Bearer ")
	req.Header.Set("Authorization","Bearer ")
	fmt.Fprintln(w, "signed out successfully")
	log.Println("signed out successfully")
}

//GetUsername will return user if signed in
func GetUserName(w http.ResponseWriter,req *http.Request) (userName string) {
	auth := req.Header.Get("Authorization")  //jwt-token stored in Bearer Token Authorization header
	//log.Println("auth token:",auth)
	tokenPart := strings.Split(auth," ")
	if len(tokenPart)!=2{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w,"Not Authorized")
		return
	}
	tokenString := tokenPart[1]  //tokenPart[1] contains actual token string

	if user,found := tokenStore[tokenString];found{  //check user token present in tokenStore
		return user
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintln(w,"Not Authorized")
	return
}
