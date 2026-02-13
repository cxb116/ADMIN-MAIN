
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DspProductSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      DspCompanyId  *int `json:"dsp_company_id" form:"dsp_company_id"` 
    request.PageInfo
}
