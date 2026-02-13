// 自动生成模板Ssp_ad_slot
package ssp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 媒体广告位 结构体  Ssp_ad_slot
type Ssp_ad_slot struct {
	global.GVA_MODEL
	MediaId         *int64  `json:"media_id" form:"media_id" gorm:"column:media_id;" binding:"required"`             //媒体Id
	AppId           *int64  `json:"app_id" form:"app_id" gorm:"column:app_id;" binding:"required"`                   //应用id
	Name            *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`                         //广告位名称
	NameAlise       *string `json:"name_alise" form:"name_alise" gorm:"column:name_alise;"`                          //内部广告位名称
	SceneId         *int64  `json:"scene_id" form:"scene_id" gorm:"column:scene_id;" binding:"required"`             //广告类型
	SspPayType      *string `json:"ssp_pay_type" form:"ssp_pay_type" gorm:"column:ssp_pay_type;" binding:"required"` //结算方式
	SspDealRatio    *int64  `json:"ssp_deal_ratio" form:"ssp_deal_ratio" gorm:"column:ssp_deal_ratio;"`              //分成系数
	InteractionType *int64  `json:"interaction_type" form:"interaction_type" gorm:"column:interaction_type;"`        //交互类型
	Height          *int64  `json:"height" form:"height" gorm:"column:height;"`                                      //广告位高
	Width           *int64  `json:"width" form:"width" gorm:"column:width;"`                                         //广告位宽
	AdImage         *string `json:"ad_image" form:"ad_image" gorm:"column:ad_image;"`                                //广告位图片
	Enable          *string `json:"enable" form:"enable" gorm:"column:enable;" binding:"required"`                   //是否启用
}

type Ssp_ad_slot_res struct {
	global.GVA_MODEL
	MediaId         *int64  `json:"media_id" form:"media_id" gorm:"column:media_id;" binding:"required"`             //媒体Id
	AppId           *int64  `json:"app_id" form:"app_id" gorm:"column:app_id;" binding:"required"`                   //应用id
	Name            *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`                         //广告位名称
	NameAlise       *string `json:"name_alise" form:"name_alise" gorm:"column:name_alise;"`                          //内部广告位名称
	SceneId         *int64  `json:"scene_id" form:"scene_id" gorm:"column:scene_id;" binding:"required"`             //广告类型
	SspPayType      *string `json:"ssp_pay_type" form:"ssp_pay_type" gorm:"column:ssp_pay_type;" binding:"required"` //结算方式
	SspDealRatio    *int64  `json:"ssp_deal_ratio" form:"ssp_deal_ratio" gorm:"column:ssp_deal_ratio;"`              //分成系数
	InteractionType *int64  `json:"interaction_type" form:"interaction_type" gorm:"column:interaction_type;"`        //交互类型
	Height          *int64  `json:"height" form:"height" gorm:"column:height;"`                                      //广告位高
	Width           *int64  `json:"width" form:"width" gorm:"column:width;"`                                         //广告位宽
	AdImage         *string `json:"ad_image" form:"ad_image" gorm:"column:ad_image;"`                                //广告位图片
	Enable          *string `json:"enable" form:"enable" gorm:"column:enable;" binding:"required"`                   //是否启用
	AppName         *string `json:"app_name" gorm:"app_name;"`
	AppOsType       *string `json:"app_os_type" gorm:"app_os_type;"`
	SceneName       *string `json:"scene_name" gorm:"scene_name;"`
}

// TableName 媒体广告位 Ssp_ad_slot自定义表名 ssp_ad_slot
func (Ssp_ad_slot) TableName() string {
	return "ssp_ad_slot"
}
