package mock

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

//Base 基础默认字段
type Base struct {
	// stupid id type
	ID      string `gorm:"type:varchar(36);primary_key;comment:编号" json:"id"`
	Created uint   `gorm:"type:int(10);not null;comment:创建时间" json:"created"`
	Updated uint   `gorm:"type:int(10);default:0;comment:更新时间" json:"updated"`
	Deleted uint   `gorm:"type:int(10);default:0;comment:删除时间" json:"deleted"`
	State   uint8  `gorm:"type:tinyint(1);default:1;comment:状态" json:"state"`
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
	Name        string       `gorm:"type:varchar(128);not null;comment:名称" json:"name" binding:"required"` //名称
	Address     string       `gorm:"type:varchar(256);not null;comment:地址" json:"address"`                 //地址
	Content     string       `gorm:"type:text;not null;comment:内容" json:"content" binding:"required"`      //内容
	Contact     string       `gorm:"type:varchar(64);not null;comment:联系人" json:"contact"`                 //联系人
	Email       string       `gorm:"type:varchar(128);not null;comment:邮箱" json:"email"`                   //邮箱
	Tag         string       `gorm:"type:varchar(256);not null;comment:标签" json:"tag"`
	Attachments *Attachments `gorm:"type:text;comment:附件" json:"attachments"`
}
