package request

type ListUserQuery struct {
	UserID    *string
	Limit     *int
	Offset    *int
	Name      *string
	NIP       *string
	Role      *string
	CreatedAt *string
}
