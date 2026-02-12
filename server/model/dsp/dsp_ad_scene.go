
// 自动生成模板DspAdScene
package dsp
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 广告类型 结构体  DspAdScene
type DspAdScene struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`  //广告类型
  Enable  *string `json:"enable" form:"enable" gorm:"column:enable;" binding:"required"`  //是否启用
}


// TableName 广告类型 DspAdScene自定义表名 dsp_ad_scene
func (DspAdScene) TableName() string {
    return "dsp_ad_scene"
}





