package constant

type UserRole string

const (
	UserRoleIT    = "it"
	UserRoleNurse = "nurse"
)

var StrUserRoleIT = UserRole(UserRoleIT)
var StrUserRoleNurse = UserRole(UserRoleIT)
