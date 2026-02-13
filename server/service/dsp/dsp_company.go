
package dsp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
)

type DspCompanyService struct {}
// CreateDspCompany 创建公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService) CreateDspCompany(ctx context.Context, dspCompany *dsp.DspCompany) (err error) {
	err = global.GVA_DB.Create(dspCompany).Error
	return err
}

// DeleteDspCompany 删除公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)DeleteDspCompany(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&dsp.DspCompany{},"id = ?",ID).Error
	return err
}

// DeleteDspCompanyByIds 批量删除公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)DeleteDspCompanyByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dsp.DspCompany{},"id in ?",IDs).Error
	return err
}

// UpdateDspCompany 更新公司管理记录
// Author [yourname](https://github.com/yourname)
func (dspCompanyService *DspCompanyService)UpdateDspCompany(ctx context.Context, dspCompany dsp.DspCompany) (err error) {
	err = global.GVA_DB.Model(&dsp.DspCompany{}).Where("id = ?",dspCompany.ID).Updates(&dspCompany).Error
	return err
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
