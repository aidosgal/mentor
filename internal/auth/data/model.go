package data

type UserModel struct {
	ID          int64
	LastName    string
	FirstName   string
	UserName    string
	ChatID      string
	Role        string
	Description *string
	Phone       *string
	CreatedAt   string
	UpdatedAt   string
}
