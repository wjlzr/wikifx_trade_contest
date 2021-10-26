package responseerror

import (
	"gorm.io/gorm"
)

type (
	SysErrorModel struct {
		Eloquent  *gorm.DB
		tableName string
	}

	Error struct {
		Code string `json:"code"`
		ZhCN string `json:"zh-CN"`
		ZhHK string `json:"zh-HK"`
		ZhTW string `json:"zh-TW"`
		En   string `json:"en"`
		Vi   string `json:"vi"`
		Th   string `json:"th"`
		Fr   string `json:"fr"`
		Id   string `json:"id"`
		Es   string `json:"es"`
		Ru   string `json:"ru"`
		De   string `json:"de"`
		Tl   string `json:"tl"`
		It   string `json:"it"`
		Hi   string `json:"hi"`
		Ja   string `json:"ja"`
		Ko   string `json:"ko"`
		Pt   string `json:"pt"`
	}
)

// NewSysUserModel Separate data service
func NewSysUserModel(eloquent *gorm.DB) *SysErrorModel {
	return &SysErrorModel{
		Eloquent:  eloquent,
		tableName: "user",
	}
}

func (m *SysErrorModel) GetError(code, lang string) string {

	var results []map[string]string

	m.Eloquent.Table(m.tableName).Where("code = ?", code).Find(&results)
	for _, val := range results {
		for k, v := range val {
			if k == lang {
				return v
			}
		}
	}
	return ""
}
