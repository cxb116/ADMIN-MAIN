
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="!w-380px"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      
            <el-form-item label="广告类型" prop="scene_id">
  <el-select v-model="searchInfo.scene_id" placeholder="请选择广告类型" :clearable="true" style="width: 100%">
    <el-option
      v-for="item in sceneOptions.list"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    />
  </el-select>
</el-form-item>
            
            <el-form-item label="预算方广告位" prop="dsp_slot_code">
  <el-input v-model="searchInfo.dsp_slot_code" placeholder="搜索条件" />
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        :key="cascaderOptions.length"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="预算位名称" prop="name" width="120" />

            <el-table-column align="left" label="广告类型" prop="scene_name" width="120" />

            <el-table-column align="left" label="预算方广告位" prop="dsp_slot_code" width="120" />

            <el-table-column align="left" label="操作系统类型" prop="os_type" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.os_type, osTypeOptions) }}
    </template>
</el-table-column>

            <el-table-column align="left" label="预算方APPKEY" prop="dsp_app_key" width="120" />

            <el-table-column align="left" label="预算方APPSECRET" prop="dsp_app_secret" width="120" />

            <el-table-column align="left" label="预算APPID" prop="dsp_app_id" width="120" />

            <el-table-column align="left" label="预算方应用包名" prop="dsp_app_pkg" width="120" />

            <el-table-column align="left" label="应用版本号" prop="dsp_app_ver" width="120" />

            <el-table-column align="left" label="应用商店版本号" prop="dsp_app_store_ver" width="120" />

            <el-table-column align="left" label="应用商店地址" prop="dsp_app_store_link" width="120" />

            <el-table-column align="left" label="结算方式" prop="dsp_pay_type" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.dsp_pay_type,pay_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="成交价系数" prop="dsp_deal_ratio" width="120" />

            <el-table-column align="left" label="公司" prop="company_name" width="120">
    <template #default="scope">
    {{ scope.row.company_name || scope.row.dsp_company_id }}
    </template>
</el-table-column>

            <el-table-column align="left" label="产品" prop="product_name" width="120">
    <template #default="scope">
    {{ scope.row.product_name || scope.row.dsp_product_id }}
    </template>
</el-table-column>

            <el-table-column label="备注" prop="remark" width="200">
   <template #default="scope">
      [富文本内容]
   </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateDspSlotInfoFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="预算位名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入预算位名称" />
</el-form-item>
            <el-form-item label="预算方广告位:" prop="dsp_slot_code">
    <el-input v-model="formData.dsp_slot_code" :clearable="true" placeholder="请输入预算方广告位" />
</el-form-item>
            <el-form-item label="操作系统类型:" prop="os_type">
    <el-select v-model="formData.os_type" placeholder="请选择操作系统类型" :clearable="true" :disabled="type === 'update'" style="width: 100%">
      <el-option
        v-for="item in osTypeOptions"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
</el-form-item>
            <el-form-item label="广告类型:" prop="scene_id">
    <el-select v-model="formData.scene_id" placeholder="请选择广告类型" :clearable="true" :disabled="type === 'update'" style="width: 100%">
      <el-option
        v-for="item in sceneOptions.list"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
</el-form-item>
            <el-form-item label="公司产品:" prop="cascaderValue">
    <el-cascader
      v-model="formData.cascaderValue"
      :options="cascaderOptions"
      :props="{
        expandTrigger: 'hover',
        emitPath: true,
        value: 'value',
        label: 'label',
        children: 'children'
      }"
      placeholder="请选择公司产品"
      :disabled="type === 'update'"
      style="width: 100%"
      clearable
      filterable
    />
</el-form-item>
            <el-form-item label="预算方APPKEY:" prop="dsp_app_key">
    <el-input v-model="formData.dsp_app_key" :clearable="true" placeholder="请输入预算方APPKEY" />
</el-form-item>
            <el-form-item label="预算方APPSECRET:" prop="dsp_app_secret">
    <el-input v-model="formData.dsp_app_secret" :clearable="true" placeholder="请输入预算方APPSECRET" />
</el-form-item>
            <el-form-item label="预算APPID:" prop="dsp_app_id">
    <el-input v-model="formData.dsp_app_id" :clearable="true" placeholder="请输入预算APPID" />
</el-form-item>
            <el-form-item label="预算方应用包名:" prop="dsp_app_pkg">
    <el-input v-model="formData.dsp_app_pkg" :clearable="true" placeholder="请输入预算方应用包名" />
</el-form-item>
            <el-form-item label="应用版本号:" prop="dsp_app_ver">
    <el-input v-model="formData.dsp_app_ver" :clearable="true" placeholder="请输入应用版本号" />
</el-form-item>
            <el-form-item label="应用商店版本号:" prop="dsp_app_store_ver">
    <el-input v-model="formData.dsp_app_store_ver" :clearable="true" placeholder="请输入应用商店版本号" />
</el-form-item>
            <el-form-item label="应用商店地址:" prop="dsp_app_store_link">
    <el-input v-model="formData.dsp_app_store_link" :clearable="true" placeholder="请输入应用商店地址" />
</el-form-item>
            <el-form-item label="结算方式:" prop="dsp_pay_type">
    <el-tree-select v-model="formData.dsp_pay_type" placeholder="请选择结算方式" :data="pay_typeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
            <el-form-item label="成交价系数:" prop="dsp_deal_ratio">
    <el-input v-model.number="formData.dsp_deal_ratio" :clearable="true" placeholder="请输入成交价系数" />
</el-form-item>

            <!-- 预算配置 -->
            <el-divider content-position="left">预算配置</el-divider>
            <el-alert
              :type="totalWeight === 100 ? 'success' : 'warning'"
              :closable="false"
              show-icon
              style="margin-bottom: 16px"
            >
              <template #title>
                <span>当前流量权重总和: <strong>{{ totalWeight }}</strong> / 100</span>
              </template>
            </el-alert>

            <div style="margin-bottom: 16px">
              <el-button type="primary" @click="addLaunch">
                <el-icon><Plus /></el-icon> 添加预算广告位
              </el-button>
            </div>

            <el-table
              :data="launchList"
              border
              style="width: 100%; table-layout: fixed;"
              max-height="500"
              :expand-row-keys="expandedRowKeys"
              @expand-change="handleExpandChange"
              row-key="id"
            >
              <el-table-column type="expand" label="详细配置" width="80">
                <template #default="scope">
                  <div style="padding: 20px; background-color: #f9f9f9;">
                    <el-form :model="scope.row" label-width="120px" @click.stop>
                      <el-row :gutter="20">
                        <el-col :span="8">
                          <el-form-item label="投放策略">
                            <el-select
                              v-model="scope.row.launchStrategy"
                              placeholder="请选择投放策略"
                              clearable
                              style="width: 100%"
                            >
                              <el-option
                                v-for="item in launchStrategyOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                              />
                            </el-select>
                          </el-form-item>
                        </el-col>
                        <el-col :span="8">
                          <el-form-item label="IP限流次数">
                            <el-input-number
                              v-model="scope.row.ipLimit"
                              :min="0"
                              placeholder="不限流"
                              style="width: 100%"
                            />
                          </el-form-item>
                        </el-col>
                        <el-col :span="8">
                          <el-form-item label="捕获日志时长">
                            <el-input-number
                              v-model="scope.row.logCaptureAt"
                              :min="0"
                              placeholder="默认300"
                              style="width: 100%"
                            />
                          </el-form-item>
                        </el-col>
                      </el-row>

                      <el-row :gutter="20">
                        <el-col :span="8">
                          <el-form-item label="上报黑名单">
                            <el-input
                              v-model="scope.row.trackSchwarz"
                              placeholder="黑名单IP或域名"
                            />
                          </el-form-item>
                        </el-col>
                        <el-col :span="8">
                          <el-form-item label="请求次数(次/天)">
                            <el-input-number
                              v-model="scope.row.req"
                              :min="0"
                              placeholder="不限"
                              style="width: 100%"
                            />
                          </el-form-item>
                        </el-col>
                        <el-col :span="8">
                          <el-form-item label="展现次数">
                            <el-input-number
                              v-model="scope.row.ims"
                              :min="0"
                              placeholder="不限"
                              style="width: 100%"
                            />
                          </el-form-item>
                        </el-col>
                      </el-row>

                      <el-row :gutter="20">
                        <el-col :span="8">
                          <el-form-item label="点击次数">
                            <el-input-number
                              v-model="scope.row.clk"
                              :min="0"
                              placeholder="不限"
                              style="width: 100%"
                            />
                          </el-form-item>
                        </el-col>
                        <el-col :span="16">
                          <el-form-item label="投放时段">
                            <el-date-picker
                              v-model="scope.row.launchTime"
                              type="datetimerange"
                              range-separator="至"
                              start-placeholder="开始时间"
                              end-placeholder="结束时间"
                              style="width: 100%"
                            />
                          </el-form-item>
                        </el-col>
                      </el-row>

                      <el-row :gutter="20">
                        <el-col :span="8">
                          <el-form-item label="人群定向">
                            <el-select
                              v-model="scope.row.crowdDirection"
                              placeholder="请选择人群定向"
                              clearable
                              style="width: 100%"
                            >
                              <el-option
                                v-for="item in directionTypeOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                              />
                            </el-select>
                          </el-form-item>
                        </el-col>
                        <el-col :span="8">
                          <el-form-item label="地域定向">
                            <el-select
                              v-model="scope.row.regionDirection"
                              placeholder="请选择地域定向"
                              clearable
                              style="width: 100%"
                            >
                              <el-option
                                v-for="item in directionTypeOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                              />
                            </el-select>
                          </el-form-item>
                        </el-col>
                        <el-col :span="8">
                          <el-form-item label="品牌定向">
                            <el-select
                              v-model="scope.row.brandDirection"
                              placeholder="请选择品牌定向"
                              clearable
                              style="width: 100%"
                            >
                              <el-option
                                v-for="item in directionTypeOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                              />
                            </el-select>
                          </el-form-item>
                        </el-col>
                      </el-row>

                      <el-row :gutter="20">
                        <el-col :span="24">
                          <el-form-item label="备注">
                            <el-input
                              v-model="scope.row.remark"
                              type="textarea"
                              :rows="2"
                              placeholder="备注信息"
                            />
                          </el-form-item>
                        </el-col>
                      </el-row>
                    </el-form>
                  </div>
                </template>
              </el-table-column>

              <el-table-column type="index" label="序号" width="60" />

              <el-table-column label="预算广告位" min-width="300">
                <template #default="scope">
                  <el-select
                    v-model="scope.row.sspSlotId"
                    placeholder="请选择预算广告位"
                    filterable
                    style="width: 100%"
                  >
                    <el-option
                      v-for="slot in allSspSlots"
                      :key="slot.ID"
                      :label="`${slot.app_name} - ${slot.name}`"
                      :value="slot.ID"
                      :disabled="isSlotSelected(slot.ID, scope.$index)"
                    >
                      <div style="display: flex; justify-content: space-between; font-size: 12px">
                        <span>{{ slot.app_name }} - {{ slot.name }}</span>
                        <span style="color: #8492a6">
                          尺寸: {{ slot.width }}x{{ slot.height }}
                        </span>
                      </div>
                    </el-option>
                  </el-select>
                </template>
              </el-table-column>

              <el-table-column v-if="formData.dsp_pay_type === '2'" label="底价" min-width="150">
                <template #default="scope">
                  <el-input-number
                    v-model="scope.row.floorPrice"
                    :min="0"
                    :step="0.01"
                    :precision="2"
                    placeholder="请输入底价"
                    style="width: 120px"
                  />
                </template>
              </el-table-column>

              <el-table-column label="流量权重" min-width="200">
                <template #default="scope">
                  <el-input-number
                    v-model="scope.row.trafficWeight"
                    :min="0"
                    :max="100"
                    :step="1"
                    style="width: 120px"
                  />
                </template>
              </el-table-column>

              <el-table-column label="操作" width="100">
                <template #default="scope">
                  <el-button type="danger" link @click="removeLaunch(scope.$index)">
                    <el-icon><Delete /></el-icon> 删除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>

            <el-form-item label="备注:" prop="remark">
    <RichEdit v-model="formData.remark"/>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="预算位名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="广告类型">
    {{ sceneMap.value[String(detailForm.scene_id)] || detailForm.scene_id }}
</el-descriptions-item>
                    <el-descriptions-item label="预算方广告位">
    {{ detailForm.dsp_slot_code }}
</el-descriptions-item>
                    <el-descriptions-item label="操作系统类型">
    {{ filterDict(detailForm.os_type, osTypeOptions) }}
</el-descriptions-item>
                    <el-descriptions-item label="预算方APPKEY">
    {{ detailForm.dsp_app_key }}
</el-descriptions-item>
                    <el-descriptions-item label="预算方APPSECRET">
    {{ detailForm.dsp_app_secret }}
</el-descriptions-item>
                    <el-descriptions-item label="预算APPID">
    {{ detailForm.dsp_app_id }}
</el-descriptions-item>
                    <el-descriptions-item label="预算方应用包名">
    {{ detailForm.dsp_app_pkg }}
</el-descriptions-item>
                    <el-descriptions-item label="应用版本号">
    {{ detailForm.dsp_app_ver }}
</el-descriptions-item>
                    <el-descriptions-item label="应用商店版本号">
    {{ detailForm.dsp_app_store_ver }}
</el-descriptions-item>
                    <el-descriptions-item label="应用商店地址">
    {{ detailForm.dsp_app_store_link }}
</el-descriptions-item>
                    <el-descriptions-item label="结算方式">
    {{ detailForm.dsp_pay_type }}
</el-descriptions-item>
                    <el-descriptions-item label="成交价系数">
    {{ detailForm.dsp_deal_ratio }}
</el-descriptions-item>
                    <el-descriptions-item label="公司id">
    {{ detailForm.dsp_company_id }}
</el-descriptions-item>
                    <el-descriptions-item label="产品id">
    {{ detailForm.dsp_product_id }}
</el-descriptions-item>
                    <el-descriptions-item label="备注">
    <RichView v-model="detailForm.remark" />
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createDspSlotInfo,
  deleteDspSlotInfo,
  deleteDspSlotInfoByIds,
  updateDspSlotInfo,
  findDspSlotInfo,
  getDspSlotInfoList
} from '@/api/dsp/dspSlotInfo'
import {
  Cascader
} from '@/api/dsp/dspProduct'
import {getDictionaryTreeListByType as dspAdSceneList} from '@/api/dsp/dspAdScene'
import { batchSaveDspLaunch, getDspLaunchByDspSlotId } from '@/api/dsp/dspLaunch'
import { getSsp_ad_slotList } from '@/api/ssp/sspAdSlot'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch, computed } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'DspSlotInfo'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const pay_typeOptions = ref([])
const sceneOptions = ref([])
const osTypeOptions = ref([])              // 操作系统类型选项
const launchStrategyOptions = ref([])      // 投放策略选项
const directionTypeOptions = ref([])       // 定向类型选项（人群、地域、品牌）
const formData = ref({
            name: '',
            scene_id: undefined,
            dsp_slot_code: '',
            os_type: 0,
            dsp_app_key: '',
            dsp_app_secret: '',
            dsp_app_id: '',
            dsp_app_pkg: '',
            dsp_app_ver: '',
            dsp_app_store_ver: '',
            dsp_app_store_link: '',
            dsp_pay_type: '',
            dsp_deal_ratio: undefined,
            dsp_company_id: undefined,
            dsp_product_id: undefined,
            cascaderValue: [],
            remark: '',
        })

// 级联选择器选项数据
const cascaderOptions = ref([])

// 创建场景ID到名称的映射
const sceneMap = computed(() => {
  const map = {}
  if (sceneOptions.value && sceneOptions.value.list) {
    sceneOptions.value.list.forEach(scene => {
      map[String(scene.value)] = scene.label
    })
  }
  return map
})

// 创建公司ID到名称的映射
const companyMap = computed(() => {
  const map = {}
  if (cascaderOptions.value) {
    cascaderOptions.value.forEach(company => {
      map[String(company.value)] = company.label
    })
  }
  return map
})

// 创建产品ID到名称的映射（需要公司ID）
const productMap = computed(() => {
  const map = {}
  if (cascaderOptions.value) {
    cascaderOptions.value.forEach(company => {
      if (company.children) {
        company.children.forEach(product => {
          // 使用 companyId_productId 作为复合键
          map[`${company.value}_${product.value}`] = product.label
        })
      }
    })
  }
  return map
})

// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               scene_id : [{
                   required: true,
                   message: '请选择广告类型',
                   trigger: ['change','blur'],
               },
              ],
               dsp_slot_code : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               cascaderValue : [{
                   validator: (rule, value, callback) => {
                     if (!value || value.length !== 2) {
                       callback(new Error('请选择公司产品'))
                     } else {
                       callback()
                     }
                   }
               },
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可为空 按需保留
const setOptions = async () =>{
    pay_typeOptions.value = await getDictFunc('pay_type')
    // 获取广告场景数据
    const sceneRes = await dspAdSceneList()
    if (sceneRes.code === 0) {
        sceneOptions.value = sceneRes.data
    }

    // 获取操作系统类型数据字典
    osTypeOptions.value = (await getDictFunc('os_type')).map(x=>({...x,value:+x.value}))

    // 获取级联选择器数据
    const cascaderRes = await Cascader()
    if (cascaderRes.code === 0) {
        cascaderOptions.value = cascaderRes.data
    }

    // 加载投放策略数据字典
    launchStrategyOptions.value = (await getDictFunc('launch_strategy')).map(x=>({...x,value:+x.value}))
    // 加载定向类型数据字典（人群、地域、品牌定向共享）
    directionTypeOptions.value = (await getDictFunc('direction_type')).map(x=>({...x,value:+x.value}))

    // 数据加载完成后，再加载表格数据
    await getTableData()

    // 强制触发更新
    await new Promise(resolve => setTimeout(resolve, 100))
    if (tableData.value.length > 0) {
        // 触发响应式更新
        tableData.value = [...tableData.value]
    }
}

// 查询
const getTableData = async() => {
  const table = await getDspSlotInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    // 为每条数据添加场景名称、公司名称和产品名称
    tableData.value = table.data.list.map(row => {
      const company = cascaderOptions.value.find(c => c.value === String(row.dsp_company_id))
      const product = company?.children?.find(p => p.value === String(row.dsp_product_id))

      return {
        ...row,
        scene_name: sceneMap.value[String(row.scene_id)] || row.scene_id,
        company_name: company?.label || row.dsp_company_id,
        product_name: product?.label || row.dsp_product_id
      }
    })
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 初始化：先加载选项数据，再加载表格
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteDspSlotInfoFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteDspSlotInfoByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateDspSlotInfoFunc = async(row) => {
    const res = await findDspSlotInfo({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        // 编辑模式：将后端返回的 company_id 和 product_id 转换为级联选择器格式
        if (res.data.dsp_company_id && res.data.dsp_product_id) {
          formData.value.cascaderValue = [
            String(res.data.dsp_company_id),
            String(res.data.dsp_product_id)
          ]
        }
        dialogFormVisible.value = true

        // 加载预算配置数据
        currentDspSlot.value = row
        launchList.value = []
        expandedRowKeys.value = []

        try {
          const [launchRes, slotRes] = await Promise.all([
            getDspLaunchByDspSlotId({ dspSlotId: row.ID }),
            getSsp_ad_slotList({ page: 1, pageSize: 9999, enable: 1 })
          ])

          if (launchRes.code === 0 && launchRes.data) {
            launchList.value = launchRes.data.map(item => ({
              id: item.id || Date.now() + Math.random(),
              dspSlotId: row.ID,
              sspSlotId: item.sspSlotId,
              trafficWeight: item.trafficWeight ? Number(item.trafficWeight) : 0,
              floorPrice: item.floorPrice ? item.floorPrice / 100 : 0,  // 后端存储的是分，前端显示元
              launchStrategy: item.launchStrategy ? Number(item.launchStrategy) : undefined,
              ipLimit: item.ipLimit || undefined,
              logCaptureAt: item.logCaptureAt || 300,
              trackSchwarz: item.trackSchwarz || '',
              req: item.req || undefined,
              ims: item.ims || undefined,
              clk: item.clk || undefined,
              launchTime: item.launchTime || null,
              crowdDirection: item.crowdDirection || '',
              regionDirection: item.regionDirection || '',
              brandDirection: item.brandDirection || '',
              remark: item.remark || ''
            }))
          }

          if (slotRes.code === 0 && slotRes.data.list) {
            allSspSlots.value = slotRes.data.list
          }

          matchSspSlotInfo()
        } catch (error) {
          ElMessage.error('加载预算配置失败')
        }
    }
}


// 删除行
const deleteDspSlotInfoFunc = async (row) => {
    const res = await deleteDspSlotInfo({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = async () => {
    type.value = 'create'
    dialogFormVisible.value = true

    // 加载预算广告位列表
    try {
      const slotRes = await getSsp_ad_slotList({ page: 1, pageSize: 9999, enable: 1 })
      if (slotRes.code === 0 && slotRes.data.list) {
        allSspSlots.value = slotRes.data.list
      }
    } catch (error) {
      ElMessage.error('加载预算广告位失败')
    }

    // 初始化空的预算配置列表
    launchList.value = []
    expandedRowKeys.value = []
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        name: '',
        scene_id: undefined,
        dsp_slot_code: '',
        os_type: 0,
        dsp_app_key: '',
        dsp_app_secret: '',
        dsp_app_id: '',
        dsp_app_pkg: '',
        dsp_app_ver: '',
        dsp_app_store_ver: '',
        dsp_app_store_link: '',
        dsp_pay_type: '',
        dsp_deal_ratio: undefined,
        dsp_company_id: undefined,
        dsp_product_id: undefined,
        cascaderValue: [],
        remark: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false

             // 将级联选择器的值转换为后端需要的格式
             if (formData.value.cascaderValue && formData.value.cascaderValue.length === 2) {
               formData.value.dsp_company_id = Number(formData.value.cascaderValue[0])
               formData.value.dsp_product_id = Number(formData.value.cascaderValue[1])
             }

             // 将os_type转换为数字类型
             if (formData.value.os_type !== undefined && formData.value.os_type !== null) {
               formData.value.os_type = Number(formData.value.os_type)
             }

              let res
              const isNew = type.value === 'create'

              switch (type.value) {
                case 'create':
                  res = await createDspSlotInfo(formData.value)
                  break
                case 'update':
                  res = await updateDspSlotInfo(formData.value)
                  break
                default:
                  res = await createDspSlotInfo(formData.value)
                  break
              }

              if (res.code === 0) {
                // 如果是新增，获取新创建的ID
                const dspSlotId = isNew ? res.data.ID : formData.value.ID

                // 保存预算配置
                if (launchList.value.length > 0) {
                  const saveData = launchList.value.map(item => {
                    const saveItem = {
                      dspSlotId: dspSlotId,
                      sspSlotId: item.sspSlotId,
                      trafficWeight: Number(item.trafficWeight),
                      floorPrice: Math.round(item.floorPrice * 100)  // 前端是元，后端存储分
                    }

                    // 添加可选字段（只添加有值的字段）
                    if (item.launchStrategy) saveItem.launchStrategy = Number(item.launchStrategy)
                    if (item.ipLimit !== undefined && item.ipLimit !== null) saveItem.ipLimit = item.ipLimit
                    if (item.logCaptureAt !== undefined && item.logCaptureAt !== null) saveItem.logCaptureAt = item.logCaptureAt
                    if (item.trackSchwarz) saveItem.trackSchwarz = item.trackSchwarz
                    if (item.req !== undefined && item.req !== null) saveItem.req = item.req
                    if (item.ims !== undefined && item.ims !== null) saveItem.ims = item.ims
                    if (item.clk !== undefined && item.clk !== null) saveItem.clk = item.clk
                    if (item.launchTime) saveItem.launchTime = item.launchTime
                    if (item.crowdDirection) saveItem.crowdDirection = item.crowdDirection
                    if (item.regionDirection) saveItem.regionDirection = item.regionDirection
                    if (item.brandDirection) saveItem.brandDirection = item.brandDirection
                    if (item.remark) saveItem.remark = item.remark

                    return saveItem
                  })

                  try {
                    const launchRes = await batchSaveDspLaunch(saveData)
                    if (launchRes.code !== 0) {
                      ElMessage.error('预算配置保存失败：' + launchRes.msg)
                      btnLoading.value = false
                      return
                    }
                  } catch (error) {
                    ElMessage.error('预算配置保存失败：' + error.message)
                    btnLoading.value = false
                    return
                  }
                }

                btnLoading.value = false
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              } else {
                btnLoading.value = false
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findDspSlotInfo({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}

// 获取公司名称
const getCompanyName = (companyId) => {
  if (!companyId) return ''
  return companyMap.value[String(companyId)] || companyId
}

// 获取产品名称
const getProductName = (companyId, productId) => {
  if (!companyId || !productId) return productId || ''
  const key = `${String(companyId)}_${String(productId)}`
  return productMap.value[key] || productId
}

// ========== 预算配置功能 ==========
// 预算配置对话框控制
const currentDspSlot = ref({})
const launchList = ref([])
const allSspSlots = ref([])

// 展开行状态管理
const expandedRowKeys = ref([])

// 处理展开/折叠变化
const handleExpandChange = (row, expandedRows) => {
  // 根据展开的行生成 keys 列表
  expandedRowKeys.value = expandedRows.map(r => r.id || row.id)
}

// 计算权重总和
const totalWeight = computed(() => {
  return launchList.value.reduce((sum, item) => {
    return sum + Number(item.trafficWeight || 0)
  }, 0)
})

// 判断广告位是否已被选择
const isSlotSelected = (slotId, currentIndex) => {
  return launchList.value.some((item, index) => {
    return index !== currentIndex && item.sspSlotId === slotId
  })
}

// 匹配预算广告位信息
const matchSspSlotInfo = () => {
  launchList.value.forEach(launch => {
    const slot = allSspSlots.value.find(s => s.ID === launch.sspSlotId)
    if (slot) {
      launch.ssp_slot_name = slot.name
      launch.app_name = slot.app_name
    }
  })
}

// 添加配置行
const addLaunch = () => {
  const newId = Date.now() + Math.random() // 生成临时唯一ID
  launchList.value.push({
    id: newId,
    dspSlotId: currentDspSlot.value.ID,
    sspSlotId: undefined,
    trafficWeight: 0,
    floorPrice: 0,
    launchStrategy: '',
    ipLimit: undefined,
    logCaptureAt: 300,
    trackSchwarz: '',
    req: undefined,
    ims: undefined,
    clk: undefined,
    launchTime: null,
    crowdDirection: '',
    regionDirection: '',
    brandDirection: '',
    remark: ''
  })
}

// 删除配置行
const removeLaunch = (index) => {
  const deletedRow = launchList.value[index]
  launchList.value.splice(index, 1)

  // 从展开列表中移除被删除的行
  if (deletedRow && deletedRow.id) {
    expandedRowKeys.value = expandedRowKeys.value.filter(key => key !== deletedRow.id)
  }
}


</script>

<style>

</style>
