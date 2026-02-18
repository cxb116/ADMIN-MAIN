
// 自动生成模板DspSlotInfo
package dsp
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 预算位管理 结构体  DspSlotInfo
type DspSlotInfo struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`  //预算位名称
  SceneId  *int64 `json:"scene_id" form:"scene_id" gorm:"column:scene_id;" binding:"required"`  //广告类型
  DspSlotCode  *string `json:"dsp_slot_code" form:"dsp_slot_code" gorm:"column:dsp_slot_code;" binding:"required"`  //预算方广告位
  OsType  *int64 `json:"os_type" form:"os_type" gorm:"column:os_type;"`  //操作系统类型
  DspAppKey  *string `json:"dsp_app_key" form:"dsp_app_key" gorm:"column:dsp_app_key;"`  //预算方APPKEY
  DspAppSercet  *string `json:"dsp_app_secret" form:"dsp_app_secret" gorm:"column:dsp_app_secret;"`  //预算方APPSECRET
  DspAppId  *string `json:"dsp_app_id" form:"dsp_app_id" gorm:"column:dsp_app_id;"`  //预算APPID
  DspAppPkg  *string `json:"dsp_app_pkg" form:"dsp_app_pkg" gorm:"column:dsp_app_pkg;"`  //预算方应用包名
  DspAppVer  *string `json:"dsp_app_ver" form:"dsp_app_ver" gorm:"column:dsp_app_ver;"`  //应用版本号
  DspAppStoreVer  *string `json:"dsp_app_store_ver" form:"dsp_app_store_ver" gorm:"column:dsp_app_store_ver;"`  //应用商店版本号
  DspAppStoreLink  *string `json:"dsp_app_store_link" form:"dsp_app_store_link" gorm:"column:dsp_app_store_link;"`  //应用商店地址
  DspPayType  *string `json:"dsp_pay_type" form:"dsp_pay_type" gorm:"column:dsp_pay_type;"`  //结算方式
  DspDealRatio  *int64 `json:"dsp_deal_ratio" form:"dsp_deal_ratio" gorm:"column:dsp_deal_ratio;"`  //成交价系数
  DspCompanyId  *int64 `json:"dsp_company_id" form:"dsp_company_id" gorm:"column:dsp_company_id;" binding:"required"`  //公司id
  DspProductId  *int64 `json:"dsp_product_id" form:"dsp_product_id" gorm:"column:dsp_product_id;" binding:"required"`  //产品id
  Remark  *string `json:"remark" form:"remark" gorm:"column:remark;type:text;"`  //备注
}


// TableName 预算位管理 DspSlotInfo自定义表名 dsp_slot_info
func (DspSlotInfo) TableName() string {
    return "dsp_slot_info"
}





