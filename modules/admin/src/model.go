package admin

type User struct {
	Id			uint		`gorm:"AUTO_INCREMENT;primary_key;column:Id"`
	Name		string		`gorm:"not null;unique;column:Name"`
	Role		string		`gorm:"not null;column:Role"`
	Passwd		string		`gorm:"not null;column:Passwd"`
	Email		string		`gorm:"column:Email"`
	Created		int64	`gorm:"column:CreatedAt"`
	Updated		int64	`gorm:"column:UpdatedAt"`
	Deleted		string	`gorm:"column:DeletedAt"`
}

func (this *User)Enforce(name string, cz string) bool {
	return Rbac.Enforce(this.Role, name, cz)
}