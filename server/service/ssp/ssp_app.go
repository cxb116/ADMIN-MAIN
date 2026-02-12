package ssp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ssp"
	sspReq "github.com/flipped-aurora/gin-vue-admin/server/model/ssp/request"
)

type SspAppService struct{}

// CreateSspApp 创建媒体应用记录
// Author [yourname](https://github.com/yourname)
func (AppService *SspAppService) CreateSspApp(ctx context.Context, App *ssp.SspApp) (err error) {
	err = global.GVA_DB.Create(App).Error
	return err
}

// DeleteSspApp 删除媒体应用记录
// Author [yourname](https://github.com/yourname)
func (AppService *SspAppService) DeleteSspApp(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&ssp.SspApp{}, "id = ?", ID).Error
	return err
}

// DeleteSspAppByIds 批量删除媒体应用记录
// Author [yourname](https://github.com/yourname)
func (AppService *SspAppService) DeleteSspAppByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]ssp.SspApp{}, "id in ?", IDs).Error
	return err
}

// UpdateSspApp 更新媒体应用记录
// Author [yourname](https://github.com/yourname)
func (AppService *SspAppService) UpdateSspApp(ctx context.Context, App ssp.SspApp) (err error) {
	err = global.GVA_DB.Model(&ssp.SspApp{}).Where("id = ?", App.ID).Updates(&App).Error
	return err
}

// GetSspApp 根据ID获取媒体应用记录
// Author [yourname](https://github.com/yourname)
func (AppService *SspAppService) GetSspApp(ctx context.Context, ID string) (App ssp.SspApp, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&App).Error
	return
}

// GetSspAppInfoList 分页获取媒体应用记录
// Author [yourname](https://github.com/yourname)
func (AppService *SspAppService) GetSspAppInfoList(ctx context.Context, info sspReq.SspAppSearch) (list []ssp.SspAppModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Debug().Table("ssp_app As app").Select("app.*,m.name AS media_name").
		Joins("left join ssp_media AS m on m.id = app.m_id")
	var Apps []ssp.SspAppModel
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("app.created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.MId != nil {
		db = db.Where("app.m_id = ?", *info.MId)
	}
	if info.Name != nil && *info.Name != "" {
		db = db.Where("app.name = ?", *info.Name)
	}
	if info.OsType != nil && *info.OsType != "" {
		db = db.Where("app.os_type = ?", *info.OsType)
	}
	if info.AccessType != nil && *info.AccessType != "" {
		db = db.Where("app.access_type = ?", *info.AccessType)
	}
	if info.Enable != nil && *info.Enable != "" {
		db = db.Where("app.enable = ?", *info.Enable)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Apps).Error
	return Apps, total, err
}
func (AppService *SspAppService) GetSspAppPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
