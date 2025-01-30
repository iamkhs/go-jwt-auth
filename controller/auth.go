package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-jwt-auth/config"
	"go-jwt-auth/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

// HashPassword hashes a given plaintext password using bcrypt with a cost factor of 14.
// It returns the hashed password as a string and an error if the hashing process fails.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// GenerateToken creates a new JWT token for the given user.
// The token includes the user's username and email as claims,
// along with an expiration time of 24 hours from creation.
// It returns the signed token as a string and an error if signing fails.
func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	return token.SignedString(secretKey)
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, _ := HashPassword(user.Password)
	user.Password = hashPassword

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var requestUser models.User
	var user models.User

	// Bind JSON request body
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	if err := config.DB.Where("email = ?", requestUser.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	fmt.Println("Stored Hashed Password:", user.Password)
	fmt.Println("Entered Raw Password:", requestUser.Password)

	// Compare password
	isPasswordValid := CheckPassword(user.Password, requestUser.Password)
	if !isPasswordValid {
		fmt.Println("Password comparison failed")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}

	// Generate JWT token
	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Successful login response
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Error comparing passwords:", err)
	}
	return err == nil
}
