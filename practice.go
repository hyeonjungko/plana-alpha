// main.go
package main

import (
  // "context"
  "fmt"
  "net/http"
  "time"
  jwtmiddleware  "github.com/auth0/go-jwt-middleware"
  "github.com/urfave/negroni"
  jwt "github.com/dgrijalva/jwt-go"
  "github.com/gorilla/mux"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  user := r.Context().Value("user");
  fmt.Fprintf(w, "This is an authenticated request")
  fmt.Fprintf(w, "Claim content:\n")
  for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
    fmt.Fprintf(w, "%s :\t%#v\n", k, v)
  }
})

type UserClaim struct {
	Id       int
	Username string
	jwt.StandardClaims
}

var expirationTime = time.Hour * 24

func main() {
	expirationTime := time.Now().Add(expirationTime)
	claims := UserClaim{
		Id:       1,
		Username: "felipe",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "test",
		},
	}
	// Sign and get the complete encoded token as a string using the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("5920d81c931950445c36b20fe3fbe90bce12b40ccff047ab1a634ee188a7d31883229e72f3c6f48cd64170299be039b884af2a5f33677b8ca06240cdb6852b61"))
	fmt.Println(tokenString, err)
	
  r := mux.NewRouter()

  jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
      return []byte("5920d81c931950445c36b20fe3fbe90bce12b40ccff047ab1a634ee188a7d31883229e72f3c6f48cd64170299be039b884af2a5f33677b8ca06240cdb6852b61"), nil
    },
    // When set, the middleware verifies that tokens are signed with the specific signing algorithm
    // If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
    // Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
    SigningMethod: jwt.SigningMethodHS256,
  })

  r.Handle("/ping", negroni.New(
    negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
    negroni.Wrap(myHandler),
  ))
  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}