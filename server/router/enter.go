package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/ssp"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Dsp     dsp.RouterGroup
	Ssp     ssp.RouterGroup
}
