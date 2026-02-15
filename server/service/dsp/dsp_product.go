
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

// GetDictionaryTreeListByType 参考数据字典协议，查询产品
// 当 dspCompanyId 为空时，查询所有公司的产品；当指定 dspCompanyId 时，只查询该公司的产品
// Author [yourname](https://github.com/yourname)
func (dProductService *DspProductService) GetDictionaryTreeListByType(ctx context.Context, dspCompanyId *string) (list []dsp.DspProduct, err error) {
	db := global.GVA_DB.Model(&dsp.DspProduct{})

	// 如果指定了公司ID，则添加过滤条件
	if dspCompanyId != nil && *dspCompanyId != "" {
		db = db.Where("dsp_company_id = ?", *dspCompanyId)
	}

	err = db.Find(&list).Error
	return list, err
}



// CascaderResult SQL 查询结果结构体
type CascaderResult struct {
	DspProductId   string `gorm:"column:dsp_product_id"`
	DspProductName string `gorm:"column:dsp_product_name"`
	DspCompanyId   string `gorm:"column:dsp_company_id"`
	DspCompanyName string `gorm:"column:dsp_company_name"`
}

// Cascader 根据公司选择产品，第一级是公司，第二级是产品
// Author [yourname](https://github.com/yourname)
func (dProductService *DspProductService) Cascader(ctx context.Context, result *[]map[string]interface{}) (err error) {
	sql := `
		SELECT dp.id AS dsp_product_id,
		       dp.name AS dsp_product_name,
		       dp.dsp_company_id,
		       dc.name AS dsp_company_name
		FROM dsp_product dp
		LEFT JOIN dsp_company dc ON dp.dsp_company_id = dc.id
		ORDER BY dp.dsp_company_id, dp.id
	`

	var results []CascaderResult
	err = global.GVA_DB.Raw(sql).Scan(&results).Error
	if err != nil {
		return err
	}

	// 按 dsp_company_id 分组
	companyMap := make(map[string][]CascaderResult)
	for _, item := range results {
		companyMap[item.DspCompanyId] = append(companyMap[item.DspCompanyId], item)
	}

	// 构建返回数据
	var cascaderData []map[string]interface{}
	for companyId, products := range companyMap {
		if len(products) == 0 {
			continue
		}

		// 构建 children
		var children []map[string]interface{}
		for _, product := range products {
			children = append(children, map[string]interface{}{
				"value": product.DspProductId,
				"label": product.DspProductName,
			})
		}

		// 构建公司级数据
		cascaderData = append(cascaderData, map[string]interface{}{
			"value":    companyId,
			"label":    products[0].DspCompanyName,
			"children": children,
		})
	}

	*result = cascaderData
	return nil
}


