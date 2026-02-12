
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SspMediaSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Name  *string `json:"name" form:"name"` 
      MediaCompanyAddress  *string `json:"media_company_address" form:"media_company_address"` 
      MediaOwnerName  *string `json:"media_owner_name" form:"media_owner_name"` 
      ContactName  *string `json:"contact_name" form:"contact_name"` 
      ContactPhone  *string `json:"contact_phone" form:"contact_phone"` 
      Enable  *string `json:"enable" form:"enable"` 
    request.PageInfo
}
