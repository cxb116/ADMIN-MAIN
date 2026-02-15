
// 自动生成模板DspLaunch
package dsp
import (
	"time"
	"gorm.io/datatypes"
)

// 预算投放表 结构体  DspLaunch
type DspLaunch struct {
  Id  *int32 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //id字段
  CreatedAt  *int32 `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`  //创建时间
  UpdatedAt  *int32 `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`  //更新时间
  DeletedAt  *int32 `json:"deletedAt" form:"deletedAt" gorm:"comment:删除时间;column:deleted_at;"`  //删除时间
  SspSlotId  *int32 `json:"sspSlotId" form:"sspSlotId" gorm:"comment:流量广告位id;column:ssp_slot_id;"`  //流量广告位id
  DspSlotId  *int32 `json:"dspSlotId" form:"dspSlotId" gorm:"comment:预算广告位id;column:dsp_slot_id;"`  //预算广告位id
  TrafficWeight  *string `json:"trafficWeight" form:"trafficWeight" gorm:"comment:流量权重，例如 100，一个流量绑定了2个预算，一个预算是30，一个预算是70;column:traffic_weight;size:255;"`  //流量权重
  LaunchStrategy  datatypes.JSON `json:"launchStrategy" form:"launchStrategy" gorm:"comment:投放策略: 1 对接第三方，2: 自主投放;column:launch_strategy;size:10;" swaggertype:"array,object"`  //投放策略
  FloorPrice  *int32 `json:"floorPrice" form:"floorPrice" gorm:"comment:底价（给上游媒体底价）;column:floor_price;"`  //底价
  IpLimit  *int32 `json:"ipLimit" form:"ipLimit" gorm:"comment:ip限流次数;column:ip_limit;"`  //ip限流次数
  LogCaptureAt  *int32 `json:"logCaptureAt" form:"logCaptureAt" gorm:"default:300;comment:捕获日志时长，默认5分钟（在当前时间戳的基础上加的秒数）;column:log_capture_at;"`  //捕获日志时长
  TrackSchwarz  *string `json:"trackSchwarz" form:"trackSchwarz" gorm:"comment:上报黑名单;column:track_schwarz;size:255;"`  //上报黑名单
  Req  *int32 `json:"req" form:"req" gorm:"comment:请求次数(次/天);column:req;"`  //请求次数(次/天)
  Ims  *int32 `json:"ims" form:"ims" gorm:"comment:展现次数;column:ims;"`  //展现次数
  Clk  *int32 `json:"clk" form:"clk" gorm:"comment:点击次数;column:clk;"`  //点击次数
  LaunchTime  *time.Time `json:"launchTime" form:"launchTime" gorm:"comment:投放时段 1 全时段 2 自定义;column:launch_time;size:10;"`  //投放时段
  CrowdDirection  datatypes.JSON `json:"crowdDirection" form:"crowdDirection" gorm:"comment:人群定向 1 不限制，2 定向，3排除;column:crowd_direction;size:10;" swaggertype:"array,object"`  //人群定向
  RegionDirection  datatypes.JSON `json:"regionDirection" form:"regionDirection" gorm:"comment:地域定向 1 不限制，2 定向，3排除;column:region_direction;size:10;" swaggertype:"array,object"`  //地域定向
  BrandDirection  datatypes.JSON `json:"brandDirection" form:"brandDirection" gorm:"comment:品牌定向 1 不限制，2 定向，3排除;column:brand_direction;size:10;" swaggertype:"array,object"`  //品牌定向
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:191;"`  //备注
}


// TableName 预算投放表 DspLaunch自定义表名 dsp_launch
func (DspLaunch) TableName() string {
    return "dsp_launch"
}





