// 自动生成模板SspApp
package ssp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 媒体应用 结构体  SspApp
type SspApp struct {
	global.GVA_MODEL
	MId         *int64  `json:"m_id" form:"m_id" gorm:"column:m_id;" binding:"required"`                      //媒体id
	Name        *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`                      //应用名称
	OsType      *string `json:"os_type" form:"os_type" gorm:"column:os_type;" binding:"required"`             //操作系统
	AccessType  *string `json:"access_type" form:"access_type" gorm:"column:access_type;" binding:"required"` //接入方式
	Pkg         *string `json:"pkg" form:"pkg" gorm:"column:pkg;"`                                            //包名
	DownloadUrl *string `json:"download_url" form:"download_url" gorm:"column:download_url;"`                 //下载地址
	Enable      *string `json:"enable" form:"enable" gorm:"column:enable;" binding:"required"`                //是否启用

}

type SspAppModel struct {
	global.GVA_MODEL
	MId         *int64  `json:"m_id" form:"m_id" gorm:"column:m_id;" binding:"required"`                      //媒体id
	Name        *string `json:"name" form:"name" gorm:"column:name;" binding:"required"`                      //应用名称
	OsType      *string `json:"os_type" form:"os_type" gorm:"column:os_type;" binding:"required"`             //操作系统
	AccessType  *string `json:"access_type" form:"access_type" gorm:"column:access_type;" binding:"required"` //接入方式
	Pkg         *string `json:"pkg" form:"pkg" gorm:"column:pkg;"`                                            //包名
	DownloadUrl *string `json:"download_url" form:"download_url" gorm:"column:download_url;"`                 //下载地址
	Enable      *string `json:"enable" form:"enable" gorm:"column:enable;" binding:"required"`                //是否启用
	MediaName   *string `json:"media_name" gorm:"media_name"`
}

// TableName 媒体应用 SspApp自定义表名 ssp_app
func (SspApp) TableName() string {
	return "ssp_app"
}
