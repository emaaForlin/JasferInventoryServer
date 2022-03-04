package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	User   string
	ApiKey string
}

type AuthorizedUser struct {
	User   string
	Apikey string
}

func NewUser(apikey string) (User, error) {
	if apikey == "" {
		return User{}, fmt.Errorf("an apikey is needed")
	}
	for i, j := range apikey {
		if string(j) == string(":") {
			user := apikey[:i]
			apikey = apikey[i+1:]
			if len(user) < 1 || len(apikey) < 1 {
				return User{}, fmt.Errorf("invalid apikey")
			}
			return User{ApiKey: apikey, User: user}, nil
		}
	}
	return User{}, fmt.Errorf("invalid apikey")
}

/*func ValidateAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, ok := c.MustGet("dbConn").(*gorm.DB)
		if !ok {
			log.Println("[ERROR] Error connection to the database")
			panic("error connection DB")
		}
		apiKey := c.Request.Header.Get("apikey")
		user, err := NewUser(apiKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
		}
		var authorized AuthorizedUser
		db.Model(&AuthorizedUser{}).Find(&authorized)

		if user.ApiKey != authorized.Apikey || user.User != authorized.User {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
			return
		}
		return
	}
}*/

func ValidateAPIKey(c *gin.Context) bool {
	db, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		log.Println("[ERROR] Error connection to the database")
		panic("error connection DB")
	}
	apiKey := c.Request.Header.Get("apikey")
	user, err := NewUser(apiKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
	}
	var authorized AuthorizedUser
	db.Model(&AuthorizedUser{}).Limit(5).Find(&authorized)
	if user.ApiKey == authorized.Apikey && user.User == authorized.User {
		return true
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
		return false
	}
}
