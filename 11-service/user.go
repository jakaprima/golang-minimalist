package service

type User struct {
	Nama string
}

type UserService struct {
	db []*User
}

type UserInterface interface {
	// must implement
	// return string
	Register(u *User) string
	// 1. Register(u *User) string -> function yang nerima data user, simpen datanya di slice/map, dan cetak nama user berhasil didaftarkan
	// 2. GetUser() []*User -> function yang return semua data user di map kita
	GetUser() []*User
}

func NewUserService(db []*User) UserInterface {
	//
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + " berhasil register"
}

func (u *UserService) GetUser() []*User {
	return u.db
}
