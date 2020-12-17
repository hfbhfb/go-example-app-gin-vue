package v1

import (
	"errors"
	menuv1 "github.com/kuops/go-example-app/server/pkg/apis/menu/v1"
	"github.com/kuops/go-example-app/server/pkg/utils/md5"
	"gorm.io/gorm"
)

type dao struct {
	db *gorm.DB
}

func (s *dao)Login(u *User) (*User,error){
	var user User
	u.Password = md5.Encrypt(u.Password)
	err := s.db.Where("user_name = ? AND password = ?", u.UserName, u.Password).First(&user).Error
	return &user,err
}

func (s *dao)Info(u *User) (*User,error){
	var user User
	err := s.db.Where("id = ? AND user_name = ?", u.ID, u.UserName).First(&user).Error
	return &user,err
}

func (s *dao)GetAllMenu() ([]menuv1.Menu,error){
	var menu []menuv1.Menu
	err := s.db.Where(&menuv1.Menu{}).Find(&menu).Error
	return menu,err
}

func (s *dao)GetMenuByUID(u *User) ([]menuv1.Menu,error){
	var menu []menuv1.Menu
	sql := `select * FROM table_menu 
	WHERE id in (select menu_id from table_role_menu WHERE role_id in (SELECT role_id FROM table_user_role WHERE user_id = ?))`
	err := s.db.Debug().Raw(sql, u.ID).Find(&menu).Error
	return menu,err
}

func (s *dao)ChangePassword(u *User) error {
	password := md5.Encrypt(u.Password)
	err := s.db.Debug().Model(User{}).Where("id = ? AND user_name = ?", u.ID, u.UserName).Update("password",password).Error
	return err
}

func (s *dao)GetUserInfoList(limit,offset int) ([]User,int64,error) {
	var users []User
	var total int64
	 _ = s.db.Model(User{}).Count(&total)
	err := s.db.Model(User{}).Limit(limit).Offset(offset).Find(&users).Error
	return users,total,err
}

func (s *dao)UpdateUserInfo(u *User) (*User,error) {
	err := s.db.Model(u).Updates(u).Error
	return u,err
}

func (s *dao)CreateUser(u *User) (*User,error) {
	var user User
	if !errors.Is(s.db.Where("user_name = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return &User{},errors.New("用户名已注册")
	}
	err := s.db.Model(&User{}).Create(u).Error
	return u,err
}
