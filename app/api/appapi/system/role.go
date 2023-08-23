package system

import "github.com/dstgo/wilson/app/logic/systemLogic"

func NewRoleApi(roleLogic systemLogic.RoleLogic) RoleApi {
	return RoleApi{roleLogic}
}

type RoleApi struct {
	RoleLogic systemLogic.RoleLogic
}
