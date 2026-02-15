
package dsp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
)

type DspLaunchService struct {}
// CreateDspLaunch 创建预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) CreateDspLaunch(ctx context.Context, dspLaunch *dsp.DspLaunch) (err error) {
	err = global.GVA_DB.Create(dspLaunch).Error
	return err
}

// DeleteDspLaunch 删除预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService)DeleteDspLaunch(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&dsp.DspLaunch{},"id = ?",id).Error
	return err
}

// DeleteDspLaunchByIds 批量删除预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService)DeleteDspLaunchByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]dsp.DspLaunch{},"id in ?",ids).Error
	return err
}

// UpdateDspLaunch 更新预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService)UpdateDspLaunch(ctx context.Context, dspLaunch dsp.DspLaunch) (err error) {
	err = global.GVA_DB.Model(&dsp.DspLaunch{}).Where("id = ?",dspLaunch.Id).Updates(&dspLaunch).Error
	return err
}

// GetDspLaunch 根据id获取预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService)GetDspLaunch(ctx context.Context, id string) (dspLaunch dsp.DspLaunch, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dspLaunch).Error
	return
}
// GetDspLaunchInfoList 分页获取预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService)GetDspLaunchInfoList(ctx context.Context, info dspReq.DspLaunchSearch) (list []dsp.DspLaunch, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dsp.DspLaunch{})
    var dspLaunchs []dsp.DspLaunch
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.LaunchStrategy != "" {
        // TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&dspLaunchs).Error
	return  dspLaunchs, total, err
}
func (dspLaunchService *DspLaunchService)GetDspLaunchPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
