
package dsp

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
    "go.uber.org/zap"
)

type DspCompanyService struct {}

// syncToEtcd 同步到 etcd
func (s *DspCompanyService) syncToEtcd(ctx context.Context, company *dsp.DspCompany, action string) error {
	etcdSync := NewEtcdSyncService()

	if !etcdSync.IsEnabled() {
		return nil // etcd 未启用，跳过同步
	}

	// 生成 key (ID 是 uint 类型)
	key := etcdSync.generateKey("company", action, int64(company.ID))

	// 删除操作只需要发送事件
	if action == "delete" {
		if err := etcdSync.Delete(ctx, key); err != nil {
			global.GVA_LOG.Error("同步删除到 etcd 失败",
				zap.Uint("id", company.ID),
				zap.Error(err))
			return err
		}
		return nil
	}

	// 添加/更新操作，发送完整数据
	if err := etcdSync.Put(ctx, key, company); err != nil {
		global.GVA_LOG.Error("同步到 etcd 失败",
			zap.Uint("id", company.ID),
			zap.String("action", action),
			zap.Error(err))
		return err
	}

	return nil
}

// CreateDspCompany 创建公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService) CreateDspCompany(ctx context.Context, dspCompany *dsp.DspCompany) (err error) {
	// 1. 写入数据库
	err = global.GVA_DB.Create(dspCompany).Error
	if err != nil {
		return err
	}

	// 2. 同步到 etcd
	if err := dspCompanyService.syncToEtcd(ctx, dspCompany, "add"); err != nil {
		// 记录日志但不回滚数据库
		global.GVA_LOG.Warn("同步到 etcd 失败，但数据库已写入")
	}

	return nil
}

// DeleteDspCompany 删除公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)DeleteDspCompany(ctx context.Context, ID string) (err error) {
	// 1. 先查询获取完整数据（etcd 需要）
	var company dsp.DspCompany
	err = global.GVA_DB.Where("id = ?", ID).First(&company).Error
	if err != nil {
		return err
	}

	// 2. 软删除数据库
	err = global.GVA_DB.Delete(&dsp.DspCompany{}, "id = ?", ID).Error
	if err != nil {
		return err
	}

	// 3. 同步到 etcd
	if err := dspCompanyService.syncToEtcd(ctx, &company, "delete"); err != nil {
		global.GVA_LOG.Warn("同步删除到 etcd 失败，但数据库已删除")
	}

	return nil
}

// DeleteDspCompanyByIds 批量删除公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)DeleteDspCompanyByIds(ctx context.Context, IDs []string) (err error) {
	// 批量删除时，循环同步到 etcd
	for _, id := range IDs {
		var company dsp.DspCompany
		if global.GVA_DB.Where("id = ?", id).First(&company).Error == nil {
			dspCompanyService.syncToEtcd(ctx, &company, "delete")
		}
	}

	err = global.GVA_DB.Delete(&[]dsp.DspCompany{}, "id in ?", IDs).Error
	return err
}

// UpdateDspCompany 更新公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)UpdateDspCompany(ctx context.Context, dspCompany dsp.DspCompany) (err error) {
	// 1. 更新数据库
	err = global.GVA_DB.Model(&dsp.DspCompany{}).Where("id = ?", dspCompany.ID).Updates(&dspCompany).Error
	if err != nil {
		return err
	}

	// 2. 同步到 etcd
	if err := dspCompanyService.syncToEtcd(ctx, &dspCompany, "update"); err != nil {
		global.GVA_LOG.Warn("同步到 etcd 失败，但数据库已更新")
	}

	return nil
}

// GetDspCompany 根据ID获取公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)GetDspCompany(ctx context.Context, ID string) (dspCompany dsp.DspCompany, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dspCompany).Error
	return
}
// GetDspCompanyInfoList 分页获取公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)GetDspCompanyInfoList(ctx context.Context, info dspReq.DspCompanySearch) (list []dsp.DspCompany, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dsp.DspCompany{})
    var dspCompanys []dsp.DspCompany
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?", *info.Name)
    }
    if info.DspCode != nil {
        db = db.Where("dsp_code = ?", *info.DspCode)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&dspCompanys).Error
	return  dspCompanys, total, err
}
func (dspCompanyService *DspCompanyService)GetDspCompanyPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}

// GetDictionaryTreeListByType 参考数据字典协议，查询业务类
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)GetDictionaryTreeListByType(ctx context.Context) (list []dsp.DspCompany, err error) {
	// 查询所有公司数据
	err = global.GVA_DB.Model(&dsp.DspCompany{}).Find(&list).Error
	return list, err
}


