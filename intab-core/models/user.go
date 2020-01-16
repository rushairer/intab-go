package models

import (
	//mysql driver
	_ "github.com/go-sql-driver/mysql"

	h "intab-core/helpers"
	srv "intab-core/services"
	"time"

	"gopkg.in/oauth2.v3/errors"
)

const (
	//UserStatusGuest 用户状态为游客
	UserStatusGuest = 0
	//UserStatusNormal 用户状态为普通
	UserStatusNormal = 1
)

//User 用户类
type User struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Name       string `gorm:"type:varchar(128);not null;unique" valid:"required"`
	Email      string `gorm:"type:varchar(128);index" valid:"email,optional"`
	Phone      string `gorm:"type:varchar(128);index" valid:"optional"`
	Password   string `gorm:"type:varchar(255)"`
	Status     int
	UserDetail UserDetail
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
}

//UserDetail 用户信息类
type UserDetail struct {
	ID           int `gorm:"primary_key" json:"id"`
	UserID       int
	Nickname     string
	Avatar       string
	ContactEmail string
	Bio          string
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
}

//NewUser 创建用户类实例
func NewUser(name string, email string, phone string, password string) *User {
	return &User{Name: name, Email: email, Phone: phone, Password: password}
}

//NewUserWithID 通过ID创建用户类实例
func NewUserWithID(userID int) *User {
	var user User
	srv.DB().First(&user, userID)
	return &user
}

/*
func (u *User) Create() (created bool) {
    if (srv.DB().NewRecord(u)) {
        //TODO name/email/phone重复的判断和返回错误码
        errors := srv.DB().Create(&u).GetErrors()
        if (len(errors) > 0) {
            for _, err := range errors {
                log.Println(err)
            }
            return false
        } else {
            return true
        }
    } else {
        log.Println("Error")
        return false
    }
}
*/

//LoginOrCreate 登录或者创建新用户
func (u *User) LoginOrCreate() error {
	password := u.Password

	pass, err := h.GenerateFromPassword([]byte(password))
	if err == nil {
		u.Password = string(pass)
		srv.DB().
			Where(User{Name: u.Name}).
			Or(User{Email: u.Email}).
			Or(User{Phone: u.Phone}).
			FirstOrCreate(&u)

		err = h.CompareHashAndPassword([]byte(u.Password), []byte(password))
		u.Password = ""
	}
	return err
}

//Login 登录
func (u *User) Login() error {
	password := u.Password

	pass, err := h.GenerateFromPassword([]byte(password))
	if err == nil {
		u.Password = string(pass)
		srv.DB().
			Where(User{Name: u.Name}).
			Or(User{Email: u.Email}).
			Or(User{Phone: u.Phone}).
			First(&u)
		if u.ID > 0 {
			err = h.CompareHashAndPassword([]byte(u.Password), []byte(password))
			u.Password = ""
		} else {
			err = errors.ErrAccessDenied
		}
	}
	return err
}
