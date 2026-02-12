package ssp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ssp"
	sspReq "github.com/flipped-aurora/gin-vue-admin/server/model/ssp/request"
)

type SspMediaService struct{}

// CreateSspMedia 创建媒体记录
// Author [yourname](https://github.com/yourname)
func (mediaService *SspMediaService) CreateSspMedia(ctx context.Context, media *ssp.SspMedia) (err error) {
	err = global.GVA_DB.Create(media).Error
	return err
}

// DeleteSspMedia 删除媒体记录
// Author [yourname](https://github.com/yourname)
func (mediaService *SspMediaService) DeleteSspMedia(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&ssp.SspMedia{}, "id = ?", ID).Error
	return err
}

// DeleteSspMediaByIds 批量删除媒体记录
// Author [yourname](https://github.com/yourname)
func (mediaService *SspMediaService) DeleteSspMediaByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]ssp.SspMedia{}, "id in ?", IDs).Error
	return err
}

// UpdateSspMedia 更新媒体记录
// Author [yourname](https://github.com/yourname)
func (mediaService *SspMediaService) UpdateSspMedia(ctx context.Context, media ssp.SspMedia) (err error) {
	err = global.GVA_DB.Model(&ssp.SspMedia{}).Where("id = ?", media.ID).Updates(&media).Error
	return err
}

// GetSspMedia 根据ID获取媒体记录
// Author [yourname](https://github.com/yourname)
func (mediaService *SspMediaService) GetSspMedia(ctx context.Context, ID string) (media ssp.SspMedia, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&media).Error
	return
}

// GetSspMediaInfoList 分页获取媒体记录
// Author [yourname](https://github.com/yourname)
func (mediaService *SspMediaService) GetSspMediaInfoList(ctx context.Context, info sspReq.SspMediaSearch) (list []ssp.SspMedia, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ssp.SspMedia{})
	var medias []ssp.SspMedia
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name = ?", *info.Name)
	}
	if info.MediaCompanyAddress != nil && *info.MediaCompanyAddress != "" {
		db = db.Where("media_company_address LIKE ?", "%"+*info.MediaCompanyAddress+"%")
	}
	if info.MediaOwnerName != nil && *info.MediaOwnerName != "" {
		db = db.Where("media_owner_name = ?", *info.MediaOwnerName)
	}
	if info.ContactName != nil && *info.ContactName != "" {
		db = db.Where("contact_name = ?", *info.ContactName)
	}
	if info.ContactPhone != nil && *info.ContactPhone != "" {
		db = db.Where("contact_phone = ?", *info.ContactPhone)
	}
	if info.Enable != nil && *info.Enable != "" {
		db = db.Where("enable = ?", *info.Enable)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&medias).Error
	return medias, total, err
}

func (mediaService *SspMediaService) GetAllSspMediaInfoList(ctx context.Context, info sspReq.SspMediaSearch) (list []ssp.SspMedia, total int64, err error) {

	// 创建db
	db := global.GVA_DB.Model(&ssp.SspMedia{})
	var medias []ssp.SspMedia
	if info.Enable != nil && *info.Enable != "" {
		db = db.Where("enable = ?", *info.Enable)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&medias).Error
	return medias, total, err
}

func (mediaService *SspMediaService) GetSspMediaPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
