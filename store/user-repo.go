package store

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (repo UserRepository) Exists(username string) (bool, error) {
	var count int64
	err := db.Model(&User{}).Where("username = ?", username).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo UserRepository) FindByUsername(username string) (*User, error) {
	var user User
	err := db.First(&user, "username = ?", username).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo UserRepository) Create(user *User) error {
	return db.Create(&user).Error
}
