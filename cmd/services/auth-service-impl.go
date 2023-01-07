package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"idnatiya.com/golang-app/cmd/models"
	"idnatiya.com/golang-app/cmd/types"
)

type AuthServiceImpl struct {
}

func (authServiceImpl AuthServiceImpl) GetUserByEmail(email *string) (*models.User, error) {
	var user models.User
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (authServiceImpl AuthServiceImpl) Login(user *models.User) (string, error) {
	var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour

	claims := types.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "APPLICATION NAME",
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		ID:        &user.ID,
		Email:     &user.Email,
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	var JWT_SIGNATURE_KEY = []byte("@B0g0r123")
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		panic(err.Error())
	}

	return signedToken, nil
}

func (authServiceImpl AuthServiceImpl) Attemp(loginType *types.LoginType) (string, error) {
	var user models.User
	if err := models.DB.Where("email = ?", loginType.Email).First(&user).Error; err != nil {
		return "", errors.New("email or password is wrong")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginType.Password))
	if err != nil {
		return "", errors.New("email or password is wrong")
	}

	token, err := authServiceImpl.Login(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (authServiceImpl AuthServiceImpl) Register(registerType *types.RegisterType) (*models.User, error) {
	var user models.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerType.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error hash password")
	}

	user.FirstName = registerType.FirstName
	user.LastName = registerType.LastName
	user.Email = registerType.Email
	user.Password = string(hashPassword)
	user.CreatedAt = time.Now()

	if err := models.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
