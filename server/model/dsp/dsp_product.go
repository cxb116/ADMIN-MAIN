// 自动生成模板DspProduct
package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 预算产品 结构体  DspProduct
type DspProduct struct {
	global.GVA_MODEL
	Name         *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`                               //产品名称
	DspCompanyId *string `json:"dsp_company_id" form:"dsp_company_id" gorm:"column:dsp_company_id;" binding:"required"` //公司
}

// TableName 预算产品 DspProduct自定义表名 dsp_product
func (DspProduct) TableName() string {
	return "dsp_product"
}
