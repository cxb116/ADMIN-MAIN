
// 自动生成模板SspMedia
package ssp
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 媒体 结构体  SspMedia
type SspMedia struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`  //媒体名称
  Account  *string `json:"account" form:"account" gorm:"column:account;" binding:"required"`  //媒体账号
  Password  *string `json:"password" form:"password" gorm:"column:password;" binding:"required"`  //密码
  MediaCompanyName  *string `json:"media_company_name" form:"media_company_name" gorm:"column:media_company_name;" binding:"required"`  //公司名称
  MediaCompanyShort  *string `json:"media_company_short" form:"media_company_short" gorm:"column:media_company_short;"`  //公司简称
  MediaCompanyCode  *string `json:"media_company_code" form:"media_company_code" gorm:"column:media_company_code;"`  //社会信用代码
  MediaCompanyLicense  *string `json:"media_company_license" form:"media_company_license" gorm:"column:media_company_license;"`  //营业执照
  MediaCompanyAddress  *string `json:"media_company_address" form:"media_company_address" gorm:"column:media_company_address;"`  //公司地址
  MediaOwnerName  *string `json:"media_owner_name" form:"media_owner_name" gorm:"column:media_owner_name;" binding:"required"`  //法人姓名
  ContactName  *string `json:"contact_name" form:"contact_name" gorm:"column:contact_name;"`  //联系人
  ContactPhone  *string `json:"contact_phone" form:"contact_phone" gorm:"column:contact_phone;"`  //联系电话
  ContactEmail  *string `json:"contact_email" form:"contact_email" gorm:"column:contact_email;"`  //联系邮箱
  Enable  *string `json:"enable" form:"enable" gorm:"column:enable;" binding:"required"`  //是否启用
}


// TableName 媒体 SspMedia自定义表名 ssp_media
func (SspMedia) TableName() string {
    return "ssp_media"
}





