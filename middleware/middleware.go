package middleware

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	dbMod "touristapp.com/db"
)

//Claims from the jwt token
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

//Authentication verifies the token from the user
func Authentication(c *fiber.Ctx) error {
	tokenString := c.Get("token")
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("BuvanPrithivee"), nil
	})
	if err != nil {
		return c.Status(401).JSON(&fiber.Map{
			"code": 401,
			"msg":  "Unauthorized access",
		})
	}

	c.Locals("user_id", claims.UserID)
	return c.Next()
}

//AccessAuthorization verifies the users access
func AccessAuthorization(c *fiber.Ctx) error {
	var roleID int
	if err := dbMod.DB.QueryRow(
		`SELECT role_id FROM users WHERE id = $1`,
		c.Locals("user_id"),
	).Scan(&roleID); err != nil {
		log.Println(err)
		return c.SendStatus(401)
	}
	c.Locals("role_id", roleID)
	return c.Next()
}
