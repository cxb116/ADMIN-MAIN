
// 自动生成模板DspCompany
package dsp
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 公司管理 结构体  DspCompany
type DspCompany struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`  //公司名称
  DspCode  *int64 `json:"dsp_code" form:"dsp_code" gorm:"column:dsp_code;" binding:"required"`  //预算隐射
  Url  *string `json:"url" form:"url" gorm:"column:url;" binding:"required"`  //请求地址
  Method  *string `json:"method" form:"method" gorm:"column:method;" binding:"required"`  //请求方法
  Timeout  *int64 `json:"timeout" form:"timeout" gorm:"column:timeout;"`  //超时时间
}


// TableName 公司管理 DspCompany自定义表名 dsp_company
func (DspCompany) TableName() string {
    return "dsp_company"
}





