package middleauth

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"crypto/rand"
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	middlemodel "github.com/teerachot007/gofiber-middlewere/middle-model"
)

// var (
// 	privateKey *rsa.PrivateKey
// )

func GenKey() *rsa.PrivateKey {
	rng := rand.Reader
	var err error
	pk, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}
	return pk
}

func GenerateToken(content map[string]interface{}, token1_exp, token2_exp time.Duration, signed_token, signed_refresh string) middlemodel.Token {
	var msgToken middlemodel.Token
	location, _ := time.LoadLocation("Asia/Bangkok")
	at := time.Now().In(location).Add(token1_exp).Unix()
	rt := time.Now().In(location).Add(token2_exp).Unix()
	msgToken.AccessToken = generateToken(content, signed_token, at)
	msgToken.RefreshToken = generateToken(content, signed_refresh, rt)
	return msgToken
}

func RefreshToken(refreshToken, refreshTokenSecret string, value_struct interface{}) error {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshTokenSecret), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if value_struct != nil {
			cl_marshal, _ := json.Marshal(claims)
			err := json.Unmarshal([]byte(string(cl_marshal)), &value_struct)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return err
	}
}

func generateToken(content map[string]interface{}, signed string, expire int64) string {
	// Create the Claims
	location, _ := time.LoadLocation("Asia/Bangkok")
	claims := jwt.MapClaims{
		"context": content,
		"iat":     time.Now().In(location).Unix(),
		"exp":     expire,
		"sub":     utils.UUIDv4(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Create token
	t, err := token.SignedString([]byte(signed))               // Generate encoded token and send it as response.
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		// return c.SendStatus(fiber.StatusInternalServerError)
	}
	return t
}

// https://github.com/gofiber/jwt
func MiddleWare(signed string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		// Filter:         nil,
		SuccessHandler: authSuccess,
		ErrorHandler:   authError,
		SigningKey:     []byte(signed),
		// SigningKeys:   nil,
		SigningMethod: "HS256",
		// ContextKey:    nil,
		// Claims:        nil,
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
	})
}

func authError(c *fiber.Ctx, e error) error {
	logrus.Info("Unauthorized", e.Error())
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}

func authSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func Getclaims(c *fiber.Ctx, value interface{}) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	d, _ := json.Marshal(claims)
	err := json.Unmarshal([]byte(string(d)), &value)
	if err != nil {
		panic(err)
	}
}
