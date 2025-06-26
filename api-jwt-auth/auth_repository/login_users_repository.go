package auth_repository

import (
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/models"
	"github.com/jinzhu/gorm"
)

type LoginRepository interface {
	FindByEmail(string) (*models.User, error)
}

type loginRepositoryImpl struct {
	db *gorm.DB
}

func UserLoginRepository(db *gorm.DB) *loginRepositoryImpl {
	return &loginRepositoryImpl{db}
}

func (u *loginRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Debug().Model(&models.User{}).Where("email = ?", email).Take(user).Error 		
	return user, err
}

// NEED TO CHECK IF THE USER IS ACTIVATED

// Here

/*
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}
*/


// JWT Authorization | Angular Router Guards | Token Refresh
// https://www.youtube.com/watch?v=F1GUjHPpCLA

// Implementing Golang JWT Authentication and Authorization
// https://www.bacancytechnology.com/blog/golang-jwt

// Not sure
// https://dev.to/techschoolguru/how-to-create-and-verify-jwt-paseto-token-in-golang-1l5j

// https://codewithmukesh.com/blog/jwt-authentication-in-golang/