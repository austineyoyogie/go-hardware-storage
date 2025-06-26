package auth_repository

import (
	//"fmt"
	"time"

	//"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/messages"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/models"
	"github.com/austineyoyogie/go-hardware-store/utils"
	"github.com/jinzhu/gorm"
)

type UsersRepository interface {
	Save(*models.User) (*models.User, error)
	Verify(string, string) (*models.User, error)
	Find(uint64) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(*models.User) error
	Delete(user_id uint64) error
	FindEmail(string) (*models.User, error)
	ResetToken(user *models.User) error
	FindPasswordResetUser(string, string) (*models.User, error)
	UpdateNewPasswordUser(user *models.User) error
}

type usersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *usersRepositoryImpl {
	return &usersRepositoryImpl{db}
}

func (u *usersRepositoryImpl) Save(user *models.User) (*models.User, error) {
	tx := u.db.Begin()

	// set name to first uppercase and email to lowercase
	user.FirstName = utils.IsTitle(user.FirstName)
	user.LastName = utils.IsTitle(user.LastName)
	user.Email = utils.IsToLower(user.Email)

	hash, _ := utils.BcryptHash(user.Password)
	user.Password = string(hash)
	//token := utils.RandomString(10)
	//user.Token = string(token)
	
	err := tx.Debug().Model(&models.User{}).Create(user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//subject := "Register Email Verify"
	//msg := messages.Deliver([]string{user.Email}, subject)
	//activate := fmt.Sprintf("http://localhost:8000/verify?email=%s&token=%s", user.Email, user.Token)
	//msg.EmailTemplate("api-jwt-auth/messages/verifys.html", activate)
	return user, tx.Commit().Error
}

func (u *usersRepositoryImpl) Verify(email string, token string) (*models.User, error) {
	tx := u.db.Begin()
	user := &models.User{}

	err := u.db.Debug().Model(&models.User{}).Where("email = ?", email).Find(user).Error
	if err != nil {
		return nil, err
	} else {
		err := u.db.Debug().Model(&models.User{}).Where("token = ?", token).Find(user).Error
		if err != nil {
			return nil, err
		} else {	
			columns := map[string]interface{}{
				"token": "",
				"active": true,
				"verify": true,
			}
			err := tx.Debug().Model(&models.User{}).Where("email = ?", user.Email).Or("token = ?", user.Token).UpdateColumns(columns).Error
			if err != nil {
				tx.Rollback()
				return nil, err
			}
			return nil, tx.Commit().Error
		}	
	}
}

func (u *usersRepositoryImpl) Find(user_id uint64) (*models.User, error) {
	user := &models.User{}
	err := u.db.Debug().Model(&models.User{}).Where("id = ?", user_id).Find(user).Error
	return user, err
}

func (u *usersRepositoryImpl) FindAll() ([]*models.User, error) {
	users := []*models.User{}
	err := u.db.Debug().Model(&models.User{}).Find(&users).Error
	return users, err
}

func (u *usersRepositoryImpl) Update(user *models.User) error {
	tx := u.db.Begin()

	columns := map[string]interface{}{
		"telephone":   user.Telephone,
		"permission_id": user.PermissionID,
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.User{}).Where("id = ?", user.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *usersRepositoryImpl) Delete(user_id uint64) error {
	tx := u.db.Begin()

	err := tx.Debug().Model(&models.User{}).Where("id = ?", user_id).Delete(&models.User{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *usersRepositoryImpl) FindEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Debug().Model(&models.User{}).Where("email = ?", email).Find(user).Error
	return user, err
}

func (u *usersRepositoryImpl) ResetToken(user *models.User) error {
	tx := u.db.Begin()

	user.Email = utils.IsToLower(user.Email)
	token := utils.RandomString(250)
	user.Token = string(token)

	columns := map[string]interface{}{
		"token":   user.Token,
		"updated_at":  time.Now(),
	}

	err := tx.Debug().Model(&models.User{}).Where("email = ?", user.Email).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//subject := "Reset Your Password"
	//msg := messages.Deliver([]string{user.Email}, subject)
	//activate := fmt.Sprintf("http://localhost:8000/resetpassword?email=%s&token=%s", user.Email, user.Token)
	//msg.EmailTemplate("api-jwt-auth/messages/resets.html", activate)
	
	return tx.Commit().Error
}

func (u *usersRepositoryImpl) FindPasswordResetUser(email string, token string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Debug().Model(&models.User{}).Where("email = ?", email).Or("token = ?", token).Find(user).Error
	return user, err
}

func (u *usersRepositoryImpl) UpdateNewPasswordUser(user *models.User) error {
	tx := u.db.Begin()
	
	hash, _ := utils.BcryptHash(user.Password)
	user.Password = string(hash)

	columns := map[string]interface{}{
		"password": user.Password,
		"token": "",
		"updated_at": time.Now(),
	}
	err := tx.Debug().Model(&models.User{}).Where("email = ?", user.Email).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
	 	return  err
	}

	//subject := "Your password has be changed"
	//msg := messages.Deliver([]string{user.Email}, subject)
	//msg.EmailTemplate("api-jwt-auth/messages/delivers.html", nil)
	return tx.Commit().Error
}
