package ssp

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SspMediaApi
	SspAppApi
	Ssp_ad_slotApi
}

var (
	mediaService  = service.ServiceGroupApp.SspServiceGroup.SspMediaService
	AppService    = service.ServiceGroupApp.SspServiceGroup.SspAppService
	adSlotService = service.ServiceGroupApp.SspServiceGroup.Ssp_ad_slotService
)
