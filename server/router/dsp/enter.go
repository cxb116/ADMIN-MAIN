package dsp

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ DspAdSceneRouter }

var adSceneApi = api.ApiGroupApp.DspApiGroup.DspAdSceneApi
