package service

import "douyin/entity"

type ResponseMessageChat struct {
	*ResponseCommon
	Data *[]entity.Message
}
