
package request

import (
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
