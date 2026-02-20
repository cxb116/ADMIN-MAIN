
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Ssp_ad_slotSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      MediaId  *int `json:"media_id" form:"media_id"`
      AppId  *int `json:"app_id" form:"app_id"`
      Name  *string `json:"name" form:"name"`
      NameAlise  *string `json:"name_alise" form:"name_alise"`
      SceneId  *int `json:"scene_id" form:"scene_id"`
      SspPayType  *string `json:"ssp_pay_type" form:"ssp_pay_type"`
      Enable  *string `json:"enable" form:"enable"`
    request.PageInfo
}

// Ssp_ad_slotUpdate 更新媒体广告位的请求结构体（不需要 required 验证）
type Ssp_ad_slotUpdate struct {
	global.GVA_MODEL
	MediaId         *int64  `json:"media_id" form:"media_id" gorm:"column:media_id;"`
	AppId           *int64  `json:"app_id" form:"app_id" gorm:"column:app_id;"`
	Name            *string `json:"name" form:"name" gorm:"column:name;"`
	NameAlise       *string `json:"name_alise" form:"name_alise" gorm:"column:name_alise;"`
	SceneId         *int64  `json:"scene_id" form:"scene_id" gorm:"column:scene_id;"`
	SspPayType      *string `json:"ssp_pay_type" form:"ssp_pay_type" gorm:"column:ssp_pay_type;"`
	SspDealRatio    *int64  `json:"ssp_deal_ratio" form:"ssp_deal_ratio" gorm:"column:ssp_deal_ratio;"`
	InteractionType *int64  `json:"interaction_type" form:"interaction_type" gorm:"column:interaction_type;"`
	Height          *int64  `json:"height" form:"height" gorm:"column:height;"`
	Width           *int64  `json:"width" form:"width" gorm:"column:width;"`
	AdImage         *string `json:"ad_image" form:"ad_image" gorm:"column:ad_image;"`
	Enable          *string `json:"enable" form:"enable" gorm:"column:enable;"`
	// FlowConfig 流量配置（JSON 字段）
	FlowConfig      interface{} `json:"flow_config" form:"flow_config" gorm:"column:flow_config;"`
}
