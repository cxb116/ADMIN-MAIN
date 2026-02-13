
package dsp

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
)

type DspProductService struct {}
// CreateDspProduct 创建预算产品记录
// Author [yourname](https://github.com/yourname)
func (dProductService *DspProductService) CreateDspProduct(ctx context.Context, dProduct *dsp.DspProduct) (err error) {
	err = global.GVA_DB.Create(dProduct).Error
	return err
}

// DeleteDspProduct 删除预算产品记录
// Author [yourname](https://github.com/yourname)
func (dProductService *DspProductService)DeleteDspProduct(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&dsp.DspProduct{},"id = ?",ID).Error
	return err
}

// DeleteDspProductByIds 批量删除预算产品记录
// Author [yourname](https://github.com/yourname)
func (dProductService *DspProductService)DeleteDspProductByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dsp.DspProduct{},"id in ?",IDs).Error
	return err
}

// UpdateDspProduct 更新预算产品记录
// Author [yourname](https://github.com/yourname)
func (dProductService *DspProductService)UpdateDspProduct(ctx context.Context, dProduct dsp.DspProduct) (err error) {
	err = global.GVA_DB.Model(&dsp.DspProduct{}).Where("id = ?",dProduct.ID).Updates(&dProduct).Error
	return err
}

// GetDspProduct 根据ID获取预算产品记录
// Author [yourname](https://github.com/yourname)
func (dProductService *DspProductService)GetDspProduct(ctx context.Context, ID string) (dProduct dsp.DspProduct, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dProduct).Error
	return
}
// GetDspProductInfoList 分页获取预算产品记录
// Author [yourname](https://github.com/yourname)
func (dProductService *DspProductService)GetDspProductInfoList(ctx context.Context, info dspReq.DspProductSearch) (list []dsp.DspProduct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dsp.DspProduct{})
    var dProducts []dsp.DspProduct
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?", *info.Name)
    }
    if info.DspCompanyId != nil {
        db = db.Where("dsp_company_id = ?", *info.DspCompanyId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&dProducts).Error
	return  dProducts, total, err
}
func (dProductService *DspProductService)GetDspProductPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
