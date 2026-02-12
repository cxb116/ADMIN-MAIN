
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SspAppSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      MId  *int `json:"m_id" form:"m_id"` 
      Name  *string `json:"name" form:"name"` 
      OsType  *string `json:"os_type" form:"os_type"` 
      AccessType  *string `json:"access_type" form:"access_type"` 
      Enable  *string `json:"enable" form:"enable"` 
    request.PageInfo
}
