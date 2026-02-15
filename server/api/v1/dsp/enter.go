package dsp

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DspAdSceneApi
	DspCompanyApi
	DspProductApi
	DspSlotInfoApi
	DspLaunchApi
}

var (
	adSceneService     = service.ServiceGroupApp.DspServiceGroup.DspAdSceneService
	dspCompanyService  = service.ServiceGroupApp.DspServiceGroup.DspCompanyService
	dProductService    = service.ServiceGroupApp.DspServiceGroup.DspProductService
	dspSlotInfoService = service.ServiceGroupApp.DspServiceGroup.DspSlotInfoService
	dspLaunchService   = service.ServiceGroupApp.DspServiceGroup.DspLaunchService
)
