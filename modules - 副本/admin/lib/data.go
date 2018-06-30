package admin

type Template struct {
	Id			uint		`gorm:"AUTO_INCREMENT;primary_key"`
	Name		string		`gorm:"not null;column:Name"`
	Description		string		`gorm:"not null;column:Description"`
	Version		string	`gorm:"column:Version"`
	Author		string	`gorm:"column:Author"`
	Dir		string	`gorm:"column:Dir"`
}
func (Template) TableName() string {
	return "template"
}
