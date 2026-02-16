package dsp

import (
	"context"
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
	"gorm.io/gorm"
)

type DspLaunchService struct{}

// CreateDspLaunch 创建预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) CreateDspLaunch(ctx context.Context, dspLaunch *dsp.DspLaunch) (err error) {
	err = global.GVA_DB.Create(dspLaunch).Error
	return err
}

// DeleteDspLaunch 删除预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) DeleteDspLaunch(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&dsp.DspLaunch{}, "id = ?", id).Error
	return err
}

// DeleteDspLaunchByIds 批量删除预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) DeleteDspLaunchByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]dsp.DspLaunch{}, "id in ?", ids).Error
	return err
}

// UpdateDspLaunch 更新预算投放表记录
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) UpdateDspLaunch(ctx context.Context, dspLaunch dsp.DspLaunch) (err error) {
	err = global.GVA_DB.Model(&dsp.DspLaunch{}).Where("id = ?", dspLaunch.Id).Updates(&dspLaunch).Error
	return err
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
			// traffic_weight 是字符串，需要转换
			weight := 0.0
			if launch.TrafficWeight != nil {
				fmt.Sscanf(*launch.TrafficWeight, "%f", &weight)
			}
			totalWeight += weight
		}

		if totalWeight > 100 {
			return errors.New(fmt.Sprintf("预算位 %d 的流量权重总和 %.2f 超过100", dspSlotId, totalWeight))
		}

		if totalWeight < 100 {
			return errors.New(fmt.Sprintf("预算位 %d 的流量权重总和 %.2f 小于100", dspSlotId, totalWeight))
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

	return err
}

// GetDspLaunchByDspSlotId 根据预算位ID获取所有配置
// Author [yourname](https://github.com/yourname)
func (dspLaunchService *DspLaunchService) GetDspLaunchByDspSlotId(ctx context.Context, dspSlotId int32) (list []dsp.DspLaunch, err error) {
	err = global.GVA_DB.Where("dsp_slot_id = ? AND deleted_at IS NULL", dspSlotId).Find(&list).Error
	return list, err
}
