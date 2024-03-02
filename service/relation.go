package service

import "douyin/entity"

type ResonseRelation struct {
	*ResponseCommon
}

func (r *ResonseRelation) Follow(userId, toUserId int64) error {
	err := entity.NewRelationDaoInstance().Follow(userId, toUserId)
	*r = ResonseRelation{
		&ResponseCommon{
			0, "关注成功",
		},
	}
	return err
}

func (r *ResonseRelation) UnFollow(userId, toUserId int64) error {
	err := entity.NewRelationDaoInstance().UnFollow(userId, toUserId)
	*r = ResonseRelation{
		&ResponseCommon{
			0, "取消关注",
		},
	}
	return err
}

type ResonseRelationList struct {
	*ResponseCommon
	UserList *[]entity.User `json:"user_list"`
}

func (r *ResonseRelationList) FollowList(id int64) error {
	users, err := entity.NewRelationDaoInstance().QueryFollowList(id)
	if err != nil {
		return err
	}
	*r = ResonseRelationList{
		&ResponseCommon{0, "成功获取关注列表"},
		users,
	}
	return nil
}

func (r *ResonseRelationList) FollowerList(id int64) error {
	users, err := entity.NewRelationDaoInstance().QueryFollowerList(id)
	if err != nil {
		return err
	}
	*r = ResonseRelationList{
		&ResponseCommon{0, "成功获取粉丝列表"},
		users,
	}
	return nil
}

func (r *ResonseRelationList) FriendList(id int64) error {
	users, err := entity.NewRelationDaoInstance().QueryFriendList(id)
	if err != nil {
		return err
	}
	*r = ResonseRelationList{
		&ResponseCommon{0, "成功获取好友列表"},
		users,
	}
	return nil
}
