package middleware

import (
	"log"
	"net/http"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/response"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/service"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := jwtService.ValidateToken(c)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])

			user := jwtService.GetUser(c)
			// Check role
			exists, _ := in_array(user.Role, roles)
			if len(roles) > 0 {
				if !exists {
					response := response.BuildErrorResponse("Error", "Permission not allowed", nil)
					c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				}
			}
		} else {
			response := response.BuildErrorResponse("Error", "Your token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}

// helper for check in_arary like PHP
func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
