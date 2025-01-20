package user

import (
	"github.com/fennay/tgo/internal/model"
	"github.com/fennay/tgo/internal/store"
	"github.com/fennay/tgo/internal/vo/resp"
)

func PageList(page int, pageSize int) (list []*model.User, pagination resp.Pagination) {
	total := int64(0)
	store.Master().
		Model(&model.User{}).
		Count(&total)

	store.Master().
		Model(&model.User{}).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&list)

	pagination.Total = total
	pagination.Page = page
	pagination.PageSize = pageSize

	return list, pagination
}

func Add(user *model.User) {
	store.Master().Model(&model.User{}).Create(user)
}

func Save(user *model.User) {
	store.Master().Model(&model.User{}).Save(user)
}

func Delete(user *model.User) {
	store.Master().Model(&model.User{}).Delete(user)
}

func Detail(userId int) (user *model.User) {
	store.Master().Model(&model.User{}).Where("id=?", userId).First(user)
	return user
}
