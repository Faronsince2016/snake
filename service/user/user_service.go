package user

import (
	"github.com/1024casts/snake/model"
	"github.com/1024casts/snake/repository"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Service 用户服务接口定义
// 使用大写的service对外保留方法
type Service interface {
	CreateUser(user model.UserModel) (id uint64, err error)
	UpdateUser(userMap map[string]interface{}, id uint64) error
	GetUserByID(id uint64) (*model.UserModel, error)
	GetUserListByIds(id []uint64) (map[uint64]*model.UserModel, error)
	GetUserByPhone(phone int) (*model.UserModel, error)
	GetUserByEmail(email string) (*model.UserModel, error)
}

// UserSvc 直接初始化，可以避免在使用时再实例化
var UserSvc = NewUserService()

// 用小写的 service 实现接口中定义的方法
type userService struct {
	userRepo repository.UserRepo
}

// NewUserService 实例化一个userService
// 通过 NewService 函数初始化 Service 接口
// 依赖接口，不要依赖实现，面向接口编程
func NewUserService() Service {
	return &userService{
		userRepo: repository.NewUserRepo(),
	}
}

func (srv *userService) CreateUser(user model.UserModel) (id uint64, err error) {
	id, err = srv.userRepo.CreateUser(model.GetDB(), user)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (srv *userService) UpdateUser(userMap map[string]interface{}, id uint64) error {
	err := srv.userRepo.Update(userMap, id)

	if err != nil {
		return err
	}

	return nil
}

func (srv *userService) GetUserByID(id uint64) (*model.UserModel, error) {
	userModel, err := srv.userRepo.GetUserByID(id)
	if err != nil {
		return userModel, errors.Wrapf(err, "get user info err from db by id: %d", id)
	}

	return userModel, nil
}

// 批量获取
func (srv *userService) GetUserListByIds(id []uint64) (map[uint64]*model.UserModel, error) {
	userModels, err := srv.userRepo.GetUsersByIds(id)
	retMap := make(map[uint64]*model.UserModel)

	if err != nil {
		return retMap, errors.Wrapf(err, "get user model err from db by id: %v", id)
	}

	for _, v := range userModels {
		retMap[v.ID] = v
	}

	return retMap, nil
}

func (srv *userService) GetUserByPhone(phone int) (*model.UserModel, error) {
	userModel, err := srv.userRepo.GetUserByPhone(phone)
	if err != nil || gorm.IsRecordNotFoundError(err) {
		return userModel, errors.Wrapf(err, "get user info err from db by phone: %d", phone)
	}

	return userModel, nil
}

func (srv *userService) GetUserByEmail(email string) (*model.UserModel, error) {
	userModel, err := srv.userRepo.GetUserByEmail(email)
	if err != nil || gorm.IsRecordNotFoundError(err) {
		return userModel, errors.Wrapf(err, "get user info err from db by email: %s", email)
	}

	return userModel, nil
}
