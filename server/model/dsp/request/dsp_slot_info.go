
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DspSlotInfoSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      SceneId  *int `json:"scene_id" form:"scene_id"`
      DspSlotCode  *string `json:"dsp_slot_code" form:"dsp_slot_code"`
      OsType  *int `json:"os_type" form:"os_type"`
    request.PageInfo
}
