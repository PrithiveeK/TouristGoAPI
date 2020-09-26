package middleware

import (
	"log"
	"strings"

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
		log.Printf("Error decoding token: %s\n", err)
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
	if c.Locals("user_id") == nil {
		log.Println("Cannot find user_id")
		return c.Status(401).JSON(&fiber.Map{
			"code": 401,
			"msg":  "Unauthorized access",
		})
	}
	var roleID int
	if err := dbMod.DB.QueryRow(
		`SELECT role_id FROM users WHERE id = $1`,
		c.Locals("user_id"),
	).Scan(&roleID); err == nil {
		c.Locals("role_id", roleID)
	}

	return c.Next()
}

//RouteAccess verifies if the user has access for search/view or add/edit the database
func RouteAccess(c *fiber.Ctx) error {
	if c.Locals("user_id") == nil {
		log.Println("Cannot find user_id")
		return c.Status(401).JSON(&fiber.Map{
			"code": 401,
			"msg":  "Unauthorized access",
		})
	}
	if c.Locals("role_id") != nil {
		return c.Next()
	}
	accessID := -1
	routerArray := strings.Split(c.Path(), "/")
	if strings.Contains(routerArray[3], "services") {
		if routerArray[2] == "AGA" {
			accessID = 11
		} else if routerArray[2] == "BGA" {
			accessID = 1
		}
	} else if strings.Contains(routerArray[3], "agents") {
		accessID = 3
	} else if routerArray[3] == "suppliers" {
		accessID = 13
	}
	switch strings.ToLower(string(c.Request().Header.Method())) {
	case "post":
		accessID++
	case "put":
		accessID++
	case "delete":
		accessID++
	}
	queryAccess := `
		SELECT count(true) FROM user_access_mappings where user_id = $1 and access_id = $2
	`
	var haveAccess int64
	if err := dbMod.DB.QueryRow(queryAccess, c.Locals("user_id"), accessID).Scan(&haveAccess); err != nil {
		log.Printf("Cannot fetch user access data: %s\n", err)
		return c.Status(500).JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something Went wrong",
		})
	}
	if haveAccess == 0 {
		log.Println("Access is not permitted for this user")
		return c.Status(401).JSON(&fiber.Map{
			"code": 401,
			"msg":  "Unauthorized access",
		})
	}
	return c.Next()
}
