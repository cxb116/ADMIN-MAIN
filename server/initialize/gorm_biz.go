package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ssp"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(dsp.DspAdScene{}, ssp.SspMedia{}, ssp.SspApp{}, ssp.Ssp_ad_slot{}, dsp.DspCompany{}, dsp.DspProduct{}, dsp.DspSlotInfo{}, dsp.DspLaunch{})
	if err != nil {
		return err
	}
	return nil
}
