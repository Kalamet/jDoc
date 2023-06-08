package dto

import "github.com/kalamet/jdoc/model"

type UserDto struct {
	Phone  string `json:"phone"`
	UserId int64  `json:"user_id"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Phone:  user.Phone,
		UserId: user.ID,
	}
}
