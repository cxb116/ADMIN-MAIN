package dsp

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DspAdSceneApi
	DspCompanyApi
	DspProductApi
	DspSlotInfoApi
}

var (
	adSceneService     = service.ServiceGroupApp.DspServiceGroup.DspAdSceneService
	dspCompanyService  = service.ServiceGroupApp.DspServiceGroup.DspCompanyService
	dProductService    = service.ServiceGroupApp.DspServiceGroup.DspProductService
	dspSlotInfoService = service.ServiceGroupApp.DspServiceGroup.DspSlotInfoService
)

// DictionaryItem 数据字典项
type DictionaryItem struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

// DictionaryResponse 数据字典响应
type DictionaryResponse struct {
	List []DictionaryItem `json:"list"`
}
