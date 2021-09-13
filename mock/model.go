package mock

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

//Base 基础默认字段
type Base struct {
	// stupid id type
	ID string `gorm:"type:varchar(36);primary_key" json:"id"`

	//UnixTime
	//TODO: using gorm hook to update it automatically
	Created uint `gorm:"type:int(10);not null" json:"created"`
	Updated uint `gorm:"type:int(10);default:0" json:"-"`

	//TODO: REMOVE below
	Deleted  uint   `gorm:"type:int(10);default:0" json:"-"`
	Handlers string `gorm:"type:varchar(36);default:''" json:"-"`
	State    uint8  `gorm:"type:tinyint(1);default:1" json:"-"`
}

type Attachment struct {
	FileName  string `json:"fileName"`
	FileType  string `json:"fileType"`
	FileUrl   string `json:"fileUrl"`
	RoleClass *int   `json:"roleClass"`
}
type Attachments []Attachment

// marshal multiple struct to string
func ValueAny(a interface{}) (driver.Value, error) {
	bytes, err := json.Marshal(a)
	return string(bytes), err
}

func ScanAny(src interface{}, dst interface{}) error {
	switch value := src.(type) {
	case string:
		return json.Unmarshal([]byte(value), dst)
	case []byte:
		return json.Unmarshal(value, dst)
	default:
		return errors.New("not supported")
	}
}

func (a Attachments) Value() (driver.Value, error) { return ValueAny(a) }
func (a *Attachments) Scan(src interface{}) error  { return ScanAny(src, a) }

type TestTable struct {
	Base
	Name        string       `gorm:"type:varchar(128);not null" json:"name" binding:"required"` //名称
	Address     string       `gorm:"type:varchar(256);not null" json:"address"`                 //地址
	Content     string       `gorm:"type:text;not null" json:"content" binding:"required"`      //内容
	Contact     string       `gorm:"type:varchar(64);not null" json:"contact"`                  //联系人
	Email       string       `gorm:"type:varchar(128);not null" json:"email"`                   //邮箱
	Tag         string       `gorm:"type:varchar(256);not null" json:"tag"`
	Attachments *Attachments `gorm:"type:text" json:"attachments"`
}
