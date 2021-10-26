package user

import (
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

const (
	STATUS_TRUE    = 2  // 启用
	STATUS_FALSE   = 1  // 禁用
	DEL_FLAG_TRUE  = 1  // 未删除
	DEL_FLAG_FALSE = -1 // 已删除
)

type (
	SysUserModel struct {
		Eloquent  *gorm.DB
		tableName string
	}

	User struct {
		Id           int64     `gorm:"column:id" json:"id"`                         // 编号
		Name         string    `gorm:"column:name" json:"name"`                     // 用户名
		NickName     string    `gorm:"column:nick_name" json:"nick_name"`           // 昵称
		Avatar       string    `gorm:"column:avatar" json:"avatar"`                 // 头像
		Password     string    `gorm:"column:password" json:"password"`             // 密码
		Salt         string    `gorm:"column:salt" json:"salt"`                     // 加密盐
		Email        string    `gorm:"column:email" json:"email"`                   // 邮箱
		Mobile       string    `gorm:"column:mobile" json:"mobile"`                 // 手机号
		Status       int64     `gorm:"column:status" json:"status"`                 // 状态  0：禁用   1：正常
		DeptId       int64     `gorm:"column:dept_id" json:"dept_id"`               // 机构ID
		JobId        int64     `gorm:"column:job_id" json:"job_id"`                 // 岗位ID
		CreateBy     string    `gorm:"column:create_by" json:"create_by"`           // 创建人
		LastUpdateBy string    `gorm:"column:last_update_by" json:"last_update_by"` // 更新人
		DelFlag      int64     `gorm:"column:del_flag;default:1" json:"del_flag"`   // 是否删除  -1：已删除  0：正常
		IsAdmin      int64     `gorm:"column:is_admin" json:"is_admin"`             // 是否是管理员 1 是 -1 否
		UpdateAt     time.Time `gorm:"autoUpdateTime" json:"update_at"`
		CreateAt     time.Time `gorm:"autoCreateTime" json:"create_at"`
	}
)

// NewSysUserModel Separate data service
func NewSysUserModel(eloquent *gorm.DB) *SysUserModel {
	return &SysUserModel{
		Eloquent:  eloquent,
		tableName: "user",
	}
}

// 查询单个
func (m *SysUserModel) FindOne(condition map[string]interface{}) (*User, int64, error) {

	var user User
	result := m.Eloquent.Table(m.tableName).Select("id", "name").Where(condition).FirstOrInit(&user)
	if result.Error != nil {
		logx.Errorf("user FindOne First Err：%s", result.Error)
		return nil, 0, result.Error
	}

	return &user, result.RowsAffected, nil
}

// Create 新增
func (m *SysUserModel) Create(user User) (*User, int64, error) {

	result := m.Eloquent.Table(m.tableName).Create(&user)
	if result.Error != nil {
		logx.Errorf("user Create Create Err：%s", result.Error)
		return nil, 0, result.Error
	}

	return &user, result.RowsAffected, nil
}

// Update 更新/删除
func (m *SysUserModel) Update(user User) (*User, int64, error) {

	result := m.Eloquent.Table(m.tableName).Updates(&user)
	if result.Error != nil {
		logx.Errorf("user Update Updates Err：%s", result.Error)
		return nil, 0, result.Error
	}

	return &user, result.RowsAffected, nil
}

// QueryList 列表
func (m *SysUserModel) QueryList(pageIndex, pageSize int, name string) (users []User, count int64, err error) {

	table := m.Eloquent.Table(m.tableName).Where("status = ? and del_flag = ?", STATUS_TRUE, DEL_FLAG_TRUE)

	if name != "" {
		table = table.Where("name LIKE ?", "%"+name+"%")
	}

	result := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&users)
	if result.Error != nil {
		logx.Errorf("user QueryList Find Err：%s", result.Error)
		return nil, 0, result.Error
	}

	return users, result.RowsAffected, nil
}
