package dsp

import (
	"context"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DspLaunchService struct{}

// syncToEtcd 同步到 etcd
func (s *DspLaunchService) syncToEtcd(ctx context.Context, launch *dsp.DspLaunch, action string) error {
	etcdSync := NewEtcdSyncService()

	if !etcdSync.IsEnabled() {
		return nil // etcd 未启用，跳过同步
	}

	// 生成 key (Id 是 *uint 类型)
	id := int64(0)
	if launch.Id != nil {
		id = int64(*launch.Id)
	}
	key := etcdSync.generateKey("launch", action, id)

	// 删除操作只需要发送事件
	if action == "delete" {
		if err := etcdSync.Delete(ctx, key); err != nil {
			global.GVA_LOG.Error("同步删除到 etcd 失败",
				zap.Any("id", launch.Id),
				zap.Error(err))
			return err
		}
		return nil
	}

	// 添加/更新操作，发送完整数据
	if err := etcdSync.Put(ctx, key, launch); err != nil {
		global.GVA_LOG.Error("同步到 etcd 失败",
			zap.Any("id", launch.Id),
			zap.String("action", action),
			zap.Error(err))
		return err
	}

	return nil
}

// CreateDspLaunch 创建预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) CreateDspLaunch(ctx context.Context, dspLaunch *dsp.DspLaunch) (err error) {
	// 1. 写入数据库
	err = global.GVA_DB.Create(dspLaunch).Error
	if err != nil {
		return err
	}

	// 2. 同步到 etcd
	if err := dspLaunchService.syncToEtcd(ctx, dspLaunch, "add"); err != nil {
		// 记录日志但不回滚数据库
		global.GVA_LOG.Warn("同步到 etcd 失败，但数据库已写入")
	}

	return nil
}

// DeleteDspLaunch 删除预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) DeleteDspLaunch(ctx context.Context, id string) (err error) {
	// 1. 先查询获取完整数据（etcd 需要）
	var launch dsp.DspLaunch
	err = global.GVA_DB.Where("id = ?", id).First(&launch).Error
	if err != nil {
		return err
	}

	// 2. 软删除数据库
	err = global.GVA_DB.Delete(&dsp.DspLaunch{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	// 3. 同步到 etcd
	if err := dspLaunchService.syncToEtcd(ctx, &launch, "delete"); err != nil {
		global.GVA_LOG.Warn("同步删除到 etcd 失败，但数据库已删除")
	}

	return nil
}

// DeleteDspLaunchByIds 批量删除预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) DeleteDspLaunchByIds(ctx context.Context, ids []string) (err error) {
	// 批量删除时，循环同步到 etcd
	for _, id := range ids {
		var launch dsp.DspLaunch
		if global.GVA_DB.Where("id = ?", id).First(&launch).Error == nil {
			dspLaunchService.syncToEtcd(ctx, &launch, "delete")
		}
	}

	err = global.GVA_DB.Delete(&[]dsp.DspLaunch{}, "id in ?", ids).Error
	return err
}

// UpdateDspLaunch 更新预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) UpdateDspLaunch(ctx context.Context, dspLaunch dsp.DspLaunch) (err error) {
	// 1. 更新数据库
	err = global.GVA_DB.Model(&dsp.DspLaunch{}).Where("id = ?", dspLaunch.Id).Updates(&dspLaunch).Error
	if err != nil {
		return err
	}

	// 2. 同步到 etcd
	if err := dspLaunchService.syncToEtcd(ctx, &dspLaunch, "update"); err != nil {
		global.GVA_LOG.Warn("同步到 etcd 失败，但数据库已更新")
	}

	return nil
}

// GetDspLaunch 根据id获取预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) GetDspLaunch(ctx context.Context, id string) (dspLaunch dsp.DspLaunch, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dspLaunch).Error
	return
}

// GetDspLaunchInfoList 分页获取预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) GetDspLaunchInfoList(ctx context.Context, info dspReq.DspLaunchSearch) (list []dsp.DspLaunch, total int64, err error) {
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
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&dspLaunchs).Error
	return dspLaunchs, total, err
}
func (dspLaunchService *DspLaunchService) GetDspLaunchPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// BatchSaveDspLaunch 批量保存预算投放配置
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) BatchSaveDspLaunch(ctx context.Context, launchList []dsp.DspLaunch) (err error) {
	if len(launchList) == 0 {
		return errors.New("配置列表不能为空")
	}

	// 按 dsp_slot_id 分组
	groupMap := make(map[int32][]dsp.DspLaunch)
	for _, launch := range launchList {
		if launch.DspSlotId == nil {
			return errors.New("预算位ID不能为空")
		}
		groupMap[*launch.DspSlotId] = append(groupMap[*launch.DspSlotId], launch)
	}

	// 校验每个预算位的权重总和
	for dspSlotId, launches := range groupMap {
		var totalWeight float64
		for _, launch := range launches {
			// traffic_weight 是 int32 类型
			if launch.TrafficWeight != nil {
				totalWeight += float64(*launch.TrafficWeight)
			}
		}

		if totalWeight > 100 {
			return errors.New(fmt.Sprintf("预算位 %d 的流量权重总和 %.2f 超过100", dspSlotId, totalWeight))
		}

		if totalWeight < 100 {
			return errors.New(fmt.Sprintf("预算位 %d 的流量权重总和 %.2f 小于100", dspSlotId, totalWeight))
		}
	}

	// 在事务前查询现有配置，用于 etcd 删除通知
	oldLaunchesMap := make(map[int32][]dsp.DspLaunch)
	for dspSlotId := range groupMap {
		var oldLaunches []dsp.DspLaunch
		if err := global.GVA_DB.Where("dsp_slot_id = ? AND deleted_at IS NULL", dspSlotId).Find(&oldLaunches).Error; err != nil {
			global.GVA_LOG.Warn("查询旧配置失败，继续执行", zap.Error(err))
		} else {
			oldLaunchesMap[dspSlotId] = oldLaunches
		}
	}

	// 使用事务批量保存
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for dspSlotId, launches := range groupMap {
			// 删除该预算位的旧配置（软删除）
			if err := tx.Where("dsp_slot_id = ?", dspSlotId).Delete(&dsp.DspLaunch{}).Error; err != nil {
				return err
			}

			// 清除时间字段和ID，让数据库自动处理
			for i := range launches {
				launches[i].CreatedAt = nil
				launches[i].UpdatedAt = nil
				launches[i].DeletedAt = nil
				launches[i].Id = nil // 清除ID，让数据库自动生成
			}

			// 批量插入新配置，使用 Omit 忽略时间字段
			if err := tx.Omit("CreatedAt", "UpdatedAt", "DeletedAt").Create(&launches).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	// 事务成功后，同步到 etcd
	// 注意：launches 的 Id 字段已在插入后被 GORM 填充

	// 先发送旧配置的删除事件
	for dspSlotId, oldLaunches := range oldLaunchesMap {
		for i := range oldLaunches {
			// 发送删除事件到 etcd（value 为空）
			if syncErr := dspLaunchService.syncToEtcd(ctx, &oldLaunches[i], "delete"); syncErr != nil {
				global.GVA_LOG.Warn("同步删除到 etcd 失败",
					zap.Any("id", oldLaunches[i].Id),
					zap.Int32("dsp_slot_id", dspSlotId),
					zap.Error(syncErr))
			}
		}
	}

	for _, launches := range groupMap {
		for i := range launches {
			// 同步到 etcd，使用 "add" 操作
			if syncErr := dspLaunchService.syncToEtcd(ctx, &launches[i], "add"); syncErr != nil {
				global.GVA_LOG.Warn("同步到 etcd 失败，但数据库已写入",
					zap.Any("dsp_slot_id", launches[i].DspSlotId),
					zap.Error(syncErr))
			}
		}
	}

	return nil
}

// GetDspLaunchByDspSlotId 根据预算位ID获取所有配置
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) GetDspLaunchByDspSlotId(ctx context.Context, dspSlotId int32) (list []dsp.DspLaunch, err error) {
	err = global.GVA_DB.Where("dsp_slot_id = ? AND deleted_at IS NULL", dspSlotId).Find(&list).Error
	return list, err
}
