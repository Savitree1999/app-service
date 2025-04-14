package repository

import (
	"github.com/Savitree1999/app-service/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db  *gorm.DB
	log *zap.SugaredLogger
}

func NewUserRepository(db *gorm.DB, log *zap.SugaredLogger) *UserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

func (r *UserRepository) Create(user *model.User) error {
	// Log the query or the user being inserted
	r.log.Infof("Inserting user into database: %v", user)

	// Perform the insert operation
	if err := r.db.Create(user).Error; err != nil {
		r.log.Errorf("Failed to insert user into database: %v", err)
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		r.log.Errorf("Error finding user by email: %v", err)
		return nil, err
	}
	r.log.Infof("Found user: %v", user)
	return &user, nil
}
