package dsp

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ DspAdSceneApi }

var adSceneService = service.ServiceGroupApp.DspServiceGroup.DspAdSceneService
