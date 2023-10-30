package user

import (
	"github.com/dstgo/wilson/internal/data"
	"github.com/dstgo/wilson/internal/data/entity"
	"github.com/dstgo/wilson/internal/pkg/utils/cp"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/dstgo/wilson/internal/types/user"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserIdGenerator interface {
	Generate() string
}

var generator UserIdGenerator = Sha256UUIDGen{}

type Sha256UUIDGen struct{}

func (s Sha256UUIDGen) Generate() string {
	return cryptor.Sha256(uuid.NewString())
}

func GenerateUserId() string {
	return generator.Generate()
}

func NewUserInfo(ds *data.DataSource) UserInfo {
	return UserInfo{
		ds: ds,
	}
}

type UserInfo struct {
	ds *data.DataSource
}

func (u UserInfo) GetUserInfoByEmail(email string) (user.Info, error) {
	userEntity, found, err := GetUserByEmail(u.ds.ORM(), email)
	if err != nil {
		return user.Info{}, err
	} else if !found {
		return user.Info{}, user.ErrUserNotFound
	}
	return u.GetUserInfoById(userEntity.Id)
}

func (u UserInfo) GetUserInfoByUUID(uuid string) (user.Info, error) {
	userEntity, found, err := GetUserByUUID(u.ds.ORM(), uuid)
	if err != nil {
		return user.Info{}, err
	} else if !found {
		return user.Info{}, user.ErrUserNotFound
	}

	return u.GetUserInfoById(userEntity.Id)
}

func (u UserInfo) GetUserInfoByName(name string) (user.Info, error) {
	userEntity, found, err := GetUserByName(u.ds.ORM(), name)
	if err != nil {
		return user.Info{}, err
	} else if !found {
		return user.Info{}, user.ErrUserNotFound
	}

	return u.GetUserInfoById(userEntity.Id)
}

func (u UserInfo) GetUserInfoById(userId uint) (user.Info, error) {
	userEntity, b, err := GetUserById(u.ds.ORM(), userId)
	if err != nil {
		return user.Info{}, err
	} else if !b {
		return user.Info{}, user.ErrUserNotFound
	}

	userInfo := user.Info{
		UUID:      userEntity.UUID,
		Username:  userEntity.Username,
		Email:     userEntity.Email,
		CreatedAt: userEntity.CreatedAt,
	}

	roles, err := u.GetUserRoles(userEntity.UUID)
	if err != nil {
		return user.Info{}, err
	}
	userInfo.Roles = roles

	return userInfo, nil
}

func (u UserInfo) GetUserInfoList(opt user.PageOption) ([]user.Info, error) {
	pageUser, err := ListByPage(u.ds.ORM(), opt)
	if err != nil {
		return []user.Info{}, err
	}

	userInfoList := make([]user.Info, 0, len(pageUser))

	if err := cp.Copy(&pageUser, &userInfoList); err != nil {
		return []user.Info{}, system.ErrProgram.Wrap(err)
	}

	return userInfoList, err
}

func GetUserById(db *gorm.DB, id uint) (entity.User, bool, error) {
	findUser := entity.User{}
	result := db.Where("id = ?", id).Find(&findUser)

	found, err := data.HasRecordFound(result)
	if err != nil {
		return findUser, false, system.ErrDatabase.Wrap(err)
	} else if !found {
		return findUser, false, nil
	}
	return findUser, true, nil
}

func GetUserByName(db *gorm.DB, username string) (entity.User, bool, error) {
	findUser := entity.User{}
	result := db.Where("username = ?", username).First(&findUser)
	found, err := data.HasRecordFound(result)
	if err != nil {
		return findUser, false, system.ErrDatabase.Wrap(err)
	} else if !found {
		return findUser, false, nil
	}
	return findUser, true, nil
}

func GetUserByUUID(db *gorm.DB, uuid string) (entity.User, bool, error) {
	findUser := entity.User{}
	result := db.Where("uuid =?", uuid).First(&findUser)
	found, err := data.HasRecordFound(result)
	if err != nil {
		return findUser, false, system.ErrDatabase.Wrap(err)
	} else if !found {
		return findUser, false, nil
	}
	return findUser, true, nil
}

func GetUserByEmail(db *gorm.DB, email string) (entity.User, bool, error) {
	findUser := entity.User{}
	result := db.Model(findUser).Where("email =?", email).First(&findUser)
	found, err := data.HasRecordFound(result)
	if err != nil {
		return findUser, false, system.ErrDatabase.Wrap(err)
	} else if !found {
		return findUser, false, nil
	}
	return findUser, true, nil
}

func ListByPage(db *gorm.DB, pageOpt user.PageOption) ([]entity.User, error) {
	pageDB := db
	pageDB.Scopes(data.Pages(pageOpt.Page, pageOpt.Size))
	if len(pageOpt.Order) > 0 {
		pageDB.Scopes(data.Order(pageOpt.Order, pageOpt.Desc))
	}

	var (
		users []entity.User
	)

	if len(pageOpt.Search) > 0 {
		query := "%" + pageOpt.Search + "%"
		pageDB = pageDB.Where("username LIKE ? OR email LIKE ?", query, query)
	}

	err := pageDB.Preload("Roles").Find(&users).Error
	if err != nil {
		return users, system.ErrDatabase.Wrap(err)
	}
	return users, nil
}

// ListAllUsers
// in most time, you should use ListByPage
func ListAllUsers(db *gorm.DB) ([]entity.User, error) {
	var users []entity.User
	err := db.Find(&users).Error
	if err != nil {
		return users, system.ErrDatabase.Wrap(err)
	}
	return users, err
}

func Count(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(entity.User{}).Count(&count).Error
	if err != nil {
		return count, system.ErrDatabase.Wrap(err)
	}
	return count, err
}
