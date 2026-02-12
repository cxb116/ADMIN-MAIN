package ssp

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	SspMediaRouter
	SspAppRouter
	Ssp_ad_slotRouter
}

var (
	mediaApi  = api.ApiGroupApp.SspApiGroup.SspMediaApi
	AppApi    = api.ApiGroupApp.SspApiGroup.SspAppApi
	adSlotApi = api.ApiGroupApp.SspApiGroup.Ssp_ad_slotApi
)
