package repositories

import (
	m "intab-core/models"
	srv "intab-core/services"

	valid "github.com/asaskevich/govalidator"
)

//AccountRepository 账户处理类
type AccountRepository struct {
	User *m.User
}

//InitAccountRepository 初始化一个空账户处理类实例
func InitAccountRepository() *AccountRepository {
	return &AccountRepository{User: &m.User{}}
}

//NewAccountRepositoryWithID 通过ID创建账户处理类实例
func NewAccountRepositoryWithID(userID int) *AccountRepository {
	user := m.NewUserWithID(userID)
	accountRepository := &AccountRepository{User: user}
	srv.DB().Model(&accountRepository.User).Related(&accountRepository.User.UserDetail, "UserDetail")
	return accountRepository
}

//NewAccountRepository 创建账户处理类实例
func NewAccountRepository(name string, email string, phone string, password string) (*AccountRepository, error) {
	user := m.NewUser(name, email, phone, password)

	result, err := valid.ValidateStruct(user)
	if err != nil {
		return nil, err
	} else if result {
		accountRepository := &AccountRepository{User: user}
		srv.DB().Model(&accountRepository.User).Related(&accountRepository.User.UserDetail, "UserDetail")
		return accountRepository, nil
	} else {
		return nil, nil
	}
}

//LoginOrCreate 登录或创建
func (ap *AccountRepository) LoginOrCreate() error {
	return ap.User.LoginOrCreate()
}

//Login 登录
func (ap *AccountRepository) Login() error {
	return ap.User.Login()
}

//SetUserStatusNormal 将用户状态设为普通
func (ap *AccountRepository) SetUserStatusNormal() {
	ap.User.Status = m.UserStatusNormal
}

//SetUserStatusGuest 将用户状态设为游客
func (ap *AccountRepository) SetUserStatusGuest() {
	ap.User.Status = m.UserStatusGuest
}
