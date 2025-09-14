package auth

import (
	"log"
	"regexp"
	"strings"

	"github.com/MegaBytee/micro-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/google/uuid"
)

type KeyAuth struct {
	key           string
	protectedURLs []*regexp.Regexp
	keyLookup     string
}

func NewKeyAuth(keyLookup string) *KeyAuth {
	return &KeyAuth{
		key:       uuid.NewString(),
		keyLookup: keyLookup,
	}
}

func (k *KeyAuth) Setup(app *fiber.App, protectedURLs []*regexp.Regexp) {
	k.protectedURLs = protectedURLs
	app.Use(keyauth.New(keyauth.Config{
		Next:         k.authFilter,
		KeyLookup:    k.keyLookup,
		Validator:    k.validator,
		ErrorHandler: k.errorHandler,
	}))
}

func (k *KeyAuth) ApiKeyLog() {
	log.Println("KeyAuth=", k.key)
}
func (k *KeyAuth) authFilter(c *fiber.Ctx) bool {
	originalURL := strings.ToLower(c.OriginalURL())

	//fmt.Println("originalURL:", originalURL)
	for _, pattern := range k.protectedURLs {
		//fmt.Println("pattern:", pattern)
		if pattern.MatchString(originalURL) {
			return false
		}
	}
	return true
}

func (k *KeyAuth) validator(c *fiber.Ctx, key string) (bool, error) {

	if key == k.key {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func (k *KeyAuth) errorHandler(c *fiber.Ctx, err error) error {
	if err == keyauth.ErrMissingOrMalformedAPIKey {
		return c.Status(fiber.StatusUnauthorized).JSON(routes.ResponseHTTP{
			Success: false,
			Message: "error",
			Data:    err.Error(),
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(routes.ResponseHTTP{
		Success: false,
		Message: "Unauthorized",
		Data:    "Invalid or expired API Key",
	})

}
