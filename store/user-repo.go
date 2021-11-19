package store

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (repo UserRepository) AddUser(user *User) error {
	return db.Create(&user).Error
}
