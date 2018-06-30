package manager

//import "time"

type User struct {
	Id			uint		`gorm:"AUTO_INCREMENT;primary_key"`
	Name		string		`gorm:"not null;unique;column:username"`
	Passwd		string		`gorm:"not null"`
	Email		string		`gorm:"column:email"`
	//Created		time.Time	`gorm:"column:createdat"`
	//Updated		time.Time	`gorm:"column:updatedat"`
	//Deleted		*time.Time	`gorm:"column:deletedat"`
}

func (User) TableName() string {
	return "users"
}