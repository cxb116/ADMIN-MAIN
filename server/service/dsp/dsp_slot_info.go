
package dsp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
)

type DspSlotInfoService struct {}
// CreateDspSlotInfo 创建预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService) CreateDspSlotInfo(ctx context.Context, dspSlotInfo *dsp.DspSlotInfo) (err error) {
	err = global.GVA_DB.Create(dspSlotInfo).Error
	return err
}

// DeleteDspSlotInfo 删除预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService)DeleteDspSlotInfo(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&dsp.DspSlotInfo{},"id = ?",ID).Error
	return err
}

// DeleteDspSlotInfoByIds 批量删除预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService)DeleteDspSlotInfoByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dsp.DspSlotInfo{},"id in ?",IDs).Error
	return err
}

// UpdateDspSlotInfo 更新预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService)UpdateDspSlotInfo(ctx context.Context, dspSlotInfo dsp.DspSlotInfo) (err error) {
	err = global.GVA_DB.Model(&dsp.DspSlotInfo{}).Where("id = ?",dspSlotInfo.ID).Updates(&dspSlotInfo).Error
	return err
}

// GetDspSlotInfo 根据ID获取预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService)GetDspSlotInfo(ctx context.Context, ID string) (dspSlotInfo dsp.DspSlotInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dspSlotInfo).Error
	return
}
// GetDspSlotInfoInfoList 分页获取预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService)GetDspSlotInfoInfoList(ctx context.Context, info dspReq.DspSlotInfoSearch) (list []dsp.DspSlotInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dsp.DspSlotInfo{})
    var dspSlotInfos []dsp.DspSlotInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.SceneId != nil {
        db = db.Where("scene_id = ?", *info.SceneId)
    }
    if info.DspSlotCode != nil && *info.DspSlotCode != "" {
        db = db.Where("dsp_slot_code = ?", *info.DspSlotCode)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&dspSlotInfos).Error
	return  dspSlotInfos, total, err
}
func (dspSlotInfoService *DspSlotInfoService)GetDspSlotInfoPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
