package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/common/response"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/config"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"gorm.io/gorm"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(ctx *gin.Context) *jwt.Token
	GetUserId(ctx *gin.Context) string
	GetUser(ctx *gin.Context) entity.User
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "admin",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "system"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(ctx *gin.Context) *jwt.Token {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		response := response.BuildErrorResponse("Failed to process request", "No token provided", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	reqToken := strings.TrimSpace(splitToken[1])

	if len(splitToken) != 2 {
		response := response.BuildErrorResponse("Failed to process request", "Bearer token not in proper format", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	t, err := jwt.Parse(reqToken, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil
	}

	return t

}

func (j *jwtService) GetUserId(ctx *gin.Context) string {
	token := j.ValidateToken(ctx)

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		id := fmt.Sprintf("%v", claims["user_id"])

		return id
	} else {
		response := response.BuildErrorResponse("Error", "Your token is not valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	return ""
}

func (j *jwtService) GetUser(ctx *gin.Context) entity.User {
	var db *gorm.DB = config.SetupDatabaseConnection()
	var userRepo repo.UserRepository = repo.NewUserRepo(db)

	id := j.GetUserId(ctx)
	user, _ := userRepo.FindByUserID(id)
	return user
}
