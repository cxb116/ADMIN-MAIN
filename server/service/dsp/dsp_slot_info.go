
package dsp

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
    "go.uber.org/zap"
)

type DspSlotInfoService struct{}

// syncToEtcd 同步到 etcd
func (s *DspSlotInfoService) syncToEtcd(ctx context.Context, slotInfo *dsp.DspSlotInfo, action string) error {
	etcdSync := NewEtcdSyncService()

	if !etcdSync.IsEnabled() {
		return nil // etcd 未启用，跳过同步
	}

	// 生成 key (ID 是 uint 类型)
	key := etcdSync.generateKey("dsp", action, int64(slotInfo.ID))

	// 删除操作只需要发送事件
	if action == "delete" {
		if err := etcdSync.Delete(ctx, key); err != nil {
			global.GVA_LOG.Error("同步删除到 etcd 失败",
				zap.Uint("id", slotInfo.ID),
				zap.Error(err))
			return err
		}
		return nil
	}

	// 添加/更新操作，发送完整数据
	if err := etcdSync.Put(ctx, key, slotInfo); err != nil {
		global.GVA_LOG.Error("同步到 etcd 失败",
			zap.Uint("id", slotInfo.ID),
			zap.String("action", action),
			zap.Error(err))
		return err
	}

	return nil
}

// CreateDspSlotInfo 创建预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService) CreateDspSlotInfo(ctx context.Context, dspSlotInfo *dsp.DspSlotInfo) (err error) {
	// 1. 写入数据库
	err = global.GVA_DB.Create(dspSlotInfo).Error
	if err != nil {
		return err
	}

	// 2. 同步到 etcd
	if err := dspSlotInfoService.syncToEtcd(ctx, dspSlotInfo, "add"); err != nil {
		// 记录日志但不回滚数据库
		global.GVA_LOG.Warn("同步到 etcd 失败，但数据库已写入")
	}

	return nil
}

// DeleteDspSlotInfo 删除预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService)DeleteDspSlotInfo(ctx context.Context, ID string) (err error) {
	// 1. 先查询获取完整数据（etcd 需要）
	var slotInfo dsp.DspSlotInfo
	err = global.GVA_DB.Where("id = ?", ID).First(&slotInfo).Error
	if err != nil {
		return err
	}

	// 2. 软删除数据库
	err = global.GVA_DB.Delete(&dsp.DspSlotInfo{}, "id = ?", ID).Error
	if err != nil {
		return err
	}

	// 3. 同步到 etcd
	if err := dspSlotInfoService.syncToEtcd(ctx, &slotInfo, "delete"); err != nil {
		global.GVA_LOG.Warn("同步删除到 etcd 失败，但数据库已删除")
	}

	return nil
}

// DeleteDspSlotInfoByIds 批量删除预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService)DeleteDspSlotInfoByIds(ctx context.Context, IDs []string) (err error) {
	// 批量删除时，循环同步到 etcd
	for _, id := range IDs {
		var slotInfo dsp.DspSlotInfo
		if global.GVA_DB.Where("id = ?", id).First(&slotInfo).Error == nil {
			dspSlotInfoService.syncToEtcd(ctx, &slotInfo, "delete")
		}
	}

	err = global.GVA_DB.Delete(&[]dsp.DspSlotInfo{}, "id in ?", IDs).Error
	return err
}

// UpdateDspSlotInfo 更新预算位管理记录
// Author [yourname](https://github.com/yourname)
func (dspSlotInfoService *DspSlotInfoService)UpdateDspSlotInfo(ctx context.Context, dspSlotInfo dsp.DspSlotInfo) (err error) {
	// 1. 更新数据库
	err = global.GVA_DB.Model(&dsp.DspSlotInfo{}).Where("id = ?",dspSlotInfo.ID).Updates(&dspSlotInfo).Error
	if err != nil {
		return err
	}

	// 2. 同步到 etcd
	if err := dspSlotInfoService.syncToEtcd(ctx, &dspSlotInfo, "update"); err != nil {
		global.GVA_LOG.Warn("同步到 etcd 失败，但数据库已更新")
	}

	return nil
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
