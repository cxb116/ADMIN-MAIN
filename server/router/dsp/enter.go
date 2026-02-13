package dsp

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	DspAdSceneRouter
	DspCompanyRouter
	DspProductRouter
	DspSlotInfoRouter
}

var (
	adSceneApi     = api.ApiGroupApp.DspApiGroup.DspAdSceneApi
	dspCompanyApi  = api.ApiGroupApp.DspApiGroup.DspCompanyApi
	dProductApi    = api.ApiGroupApp.DspApiGroup.DspProductApi
	dspSlotInfoApi = api.ApiGroupApp.DspApiGroup.DspSlotInfoApi
)
