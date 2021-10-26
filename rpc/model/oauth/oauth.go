package oauth

import (
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type (
	UserModel struct {
		eloquent  *gorm.DB
		tableName string
	}

	User struct {
		Id             int64     `gorm:"column:id" json:"id"`                             // 编号
		Name           string    `gorm:"column:name" json:"name"`                         // 用户名
		NickName       string    `gorm:"column:nick_name" json:"nick_name"`               // 昵称
		Avatar         string    `gorm:"column:avatar" json:"avatar"`                     // 头像
		Password       string    `gorm:"column:password" json:"password"`                 // 密码
		Salt           string    `gorm:"column:salt" json:"salt"`                         // 加密盐
		Email          string    `gorm:"column:email" json:"email"`                       // 邮箱
		Mobile         string    `gorm:"column:mobile" json:"mobile"`                     // 手机号
		Status         int64     `gorm:"column:status" json:"status"`                     // 状态  0：禁用   1：正常
		DeptId         int64     `gorm:"column:dept_id" json:"dept_id"`                   // 机构ID
		JobId          int64     `gorm:"column:job_id" json:"job_id"`                     // 岗位ID
		CreateBy       string    `gorm:"column:create_by" json:"create_by"`               // 创建人
		CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`           // 创建时间
		LastUpdateBy   string    `gorm:"column:last_update_by" json:"last_update_by"`     // 更新人
		LastUpdateTime time.Time `gorm:"column:last_update_time" json:"last_update_time"` // 更新时间
		DelFlag        int64     `gorm:"column:del_flag" json:"del_flag"`                 // 是否删除  -1：已删除  0：正常

	}
)

// NewSysUserModel Separate data service
func NewSysUserModel(eloquent *gorm.DB) *UserModel {
	return &UserModel{
		eloquent:  eloquent,
		tableName: "user",
	}
}

// 查询单个
func (m *UserModel) FindOne(condition map[string]interface{}) (*User, int64, error) {

	_ = m.eloquent.Transaction(func(tx *gorm.DB) error {

		return nil
	})

	var user User
	result := m.eloquent.Table(m.tableName).Select("id", "name").Where(condition).FirstOrInit(&user)
	if result.Error != nil {
		logx.Errorf("user FindOne First Err：%s", result.Error)
		return nil, 0, result.Error
	}

	return &user, result.RowsAffected, nil
}
