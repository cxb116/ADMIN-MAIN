
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type DspLaunchSearch struct{
      LaunchStrategy  string `json:"launchStrategy" form:"launchStrategy"` 
    request.PageInfo
}
