
package dsp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
)

type DspAdSceneService struct {}
// CreateDspAdScene 创建广告类型记录
// Author [yourname](https://github.com/yourname)
func (adSceneService *DspAdSceneService) CreateDspAdScene(ctx context.Context, adScene *dsp.DspAdScene) (err error) {
	err = global.GVA_DB.Create(adScene).Error
	return err
}

// DeleteDspAdScene 删除广告类型记录
// Author [yourname](https://github.com/yourname)
func (adSceneService *DspAdSceneService)DeleteDspAdScene(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&dsp.DspAdScene{},"id = ?",ID).Error
	return err
}

// DeleteDspAdSceneByIds 批量删除广告类型记录
// Author [yourname](https://github.com/yourname)
func (adSceneService *DspAdSceneService)DeleteDspAdSceneByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dsp.DspAdScene{},"id in ?",IDs).Error
	return err
}

// UpdateDspAdScene 更新广告类型记录
// Author [yourname](https://github.com/yourname)
func (adSceneService *DspAdSceneService)UpdateDspAdScene(ctx context.Context, adScene dsp.DspAdScene) (err error) {
	err = global.GVA_DB.Model(&dsp.DspAdScene{}).Where("id = ?",adScene.ID).Updates(&adScene).Error
	return err
}

// GetDspAdScene 根据ID获取广告类型记录
// Author [yourname](https://github.com/yourname)
func (adSceneService *DspAdSceneService)GetDspAdScene(ctx context.Context, ID string) (adScene dsp.DspAdScene, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&adScene).Error
	return
}
// GetDspAdSceneInfoList 分页获取广告类型记录
// Author [yourname](https://github.com/yourname)
func (adSceneService *DspAdSceneService)GetDspAdSceneInfoList(ctx context.Context, info dspReq.DspAdSceneSearch) (list []dsp.DspAdScene, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dsp.DspAdScene{})
    var adScenes []dsp.DspAdScene
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?", *info.Name)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&adScenes).Error
	return  adScenes, total, err
}
func (adSceneService *DspAdSceneService)GetDspAdScenePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}

// GetDictionaryTreeListByType 参考数据字典协议，查询广告场景
// Author [yourname](https://github.com/yourname)
func (adSceneService *DspAdSceneService)GetDictionaryTreeListByType(ctx context.Context) (list []dsp.DspAdScene, err error) {
	// 查询所有启用的广告场景数据
	err = global.GVA_DB.Model(&dsp.DspAdScene{}).Where("enable = ?", "1").Find(&list).Error
	return list, err
}


