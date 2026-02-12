
package ssp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ssp"
    sspReq "github.com/flipped-aurora/gin-vue-admin/server/model/ssp/request"
)

type Ssp_ad_slotService struct {}
// CreateSsp_ad_slot 创建媒体广告位记录
// Author [yourname](https://github.com/yourname)
func (adSlotService *Ssp_ad_slotService) CreateSsp_ad_slot(ctx context.Context, adSlot *ssp.Ssp_ad_slot) (err error) {
	err = global.GVA_DB.Create(adSlot).Error
	return err
}

// DeleteSsp_ad_slot 删除媒体广告位记录
// Author [yourname](https://github.com/yourname)
func (adSlotService *Ssp_ad_slotService)DeleteSsp_ad_slot(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&ssp.Ssp_ad_slot{},"id = ?",ID).Error
	return err
}

// DeleteSsp_ad_slotByIds 批量删除媒体广告位记录
// Author [yourname](https://github.com/yourname)
func (adSlotService *Ssp_ad_slotService)DeleteSsp_ad_slotByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]ssp.Ssp_ad_slot{},"id in ?",IDs).Error
	return err
}

// UpdateSsp_ad_slot 更新媒体广告位记录
// Author [yourname](https://github.com/yourname)
func (adSlotService *Ssp_ad_slotService)UpdateSsp_ad_slot(ctx context.Context, adSlot ssp.Ssp_ad_slot) (err error) {
	err = global.GVA_DB.Model(&ssp.Ssp_ad_slot{}).Where("id = ?",adSlot.ID).Updates(&adSlot).Error
	return err
}

// GetSsp_ad_slot 根据ID获取媒体广告位记录
// Author [yourname](https://github.com/yourname)
func (adSlotService *Ssp_ad_slotService)GetSsp_ad_slot(ctx context.Context, ID string) (adSlot ssp.Ssp_ad_slot, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&adSlot).Error
	return
}
// GetSsp_ad_slotInfoList 分页获取媒体广告位记录
// Author [yourname](https://github.com/yourname)
func (adSlotService *Ssp_ad_slotService)GetSsp_ad_slotInfoList(ctx context.Context, info sspReq.Ssp_ad_slotSearch) (list []ssp.Ssp_ad_slot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ssp.Ssp_ad_slot{})
    var adSlots []ssp.Ssp_ad_slot
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.MediaId != nil {
        db = db.Where("media_id = ?", *info.MediaId)
    }
    if info.AppId != nil {
        db = db.Where("app_id = ?", *info.AppId)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?", *info.Name)
    }
    if info.NameAlise != nil && *info.NameAlise != "" {
        db = db.Where("name_alise = ?", *info.NameAlise)
    }
    if info.SceneId != nil {
        db = db.Where("scene_id = ?", *info.SceneId)
    }
    if info.SspPayType != nil && *info.SspPayType != "" {
        db = db.Where("ssp_pay_type = ?", *info.SspPayType)
    }
    if info.Enable != nil && *info.Enable != "" {
        db = db.Where("enable = ?", *info.Enable)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&adSlots).Error
	return  adSlots, total, err
}
func (adSlotService *Ssp_ad_slotService)GetSsp_ad_slotPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
