
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

            <el-form-item label="媒体Id" prop="media_id">
  <el-input v-model.number="searchInfo.media_id" placeholder="搜索条件" />
</el-form-item>

            <el-form-item label="应用id" prop="app_id">
  <el-input v-model.number="searchInfo.app_id" placeholder="搜索条件" />
</el-form-item>

            <el-form-item label="广告位名称" prop="name">
  <el-input v-model="searchInfo.name" placeholder="搜索条件" />
</el-form-item>

            <el-form-item label="内部广告位名称" prop="name_alise">
  <el-input v-model="searchInfo.name_alise" placeholder="搜索条件" />
</el-form-item>

            <el-form-item label="广告类型" prop="scene_id">
  <el-input v-model.number="searchInfo.scene_id" placeholder="搜索条件" />
</el-form-item>

            <el-form-item label="结算方式" prop="ssp_pay_type">
    <el-tree-select v-model="searchInfo.ssp_pay_type" placeholder="请选择结算方式" :data="pay_typeOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
</el-form-item>

            <el-form-item label="是否启用" prop="enable">
    <el-tree-select v-model="searchInfo.enable" placeholder="请选择是否启用" :data="enableOptions" style="width:100%" filterable :clearable="true" check-strictly ></el-tree-select>
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
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />

          <el-table-column
              align="left"
              label="广告位名称:(ID)"
              width="250"
          >
            <template #default="{ row }">
              {{ row.name }}（{{ row.ID }}）
            </template>
          </el-table-column>

          <el-table-column align="left" label="内部广告名称" prop="name_alise" width="150" />

            <el-table-column align="left" label="应用名称" prop="app_name" width="150" />

          <el-table-column
              align="left"
              label="操作系统"
              width="120"
          >
            <template #default="{ row }">
              {{ row.app_os_type === '1' ? 'Android' : row.app_os_type === '2' ? 'iOS' : '未知' }}
            </template>
          </el-table-column>

            <el-table-column align="left" label="广告类型" width="120" prop="scene_name" />
            <el-table-column align="left" label="结算方式" prop="ssp_pay_type" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.ssp_pay_type,pay_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="分成系数" prop="ssp_deal_ratio" width="120" />

            <el-table-column align="left" label="交互类型" prop="interaction_type" width="120" />

            <el-table-column align="left" label="广告位高" prop="height" width="120" />

            <el-table-column align="left" label="广告位宽" prop="width" width="120" />

            <el-table-column align="left" label="广告位图片" prop="ad_image" width="120" />

          <el-table-column sortable align="left" label="日12期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
          </el-table-column>

            <el-table-column align="left" label="是否启用" prop="enable" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.enable,enableOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
              <el-button  type="primary" link icon="edit" class="table-button" @click="updateSsp_ad_slotFunc(scope.row)">编辑</el-button>
              <el-button  type="primary" link class="table-button table-config" @click="updateSsp_ad_slotConfig(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>配置</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSizeMax" v-model="dialogFormVisibleConfig" :show-close="false" :before-close="closeConfigDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">流量分流配置</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="saveConfig">保存配置</el-button>
                  <el-button @click="closeConfigDialog">取消</el-button>
                </div>
              </div>
            </template>

          <div class="config-container">
            <!-- 配置说明 -->
            <el-alert
              title="配置说明"
              type="info"
              :closable="false"
              show-icon
              style="margin-bottom: 20px;"
            >
              <template #default>
                <ul style="margin: 0; padding-left: 20px;">
                  <li>添加需要分流的广告位，设置每个广告位的流量权重</li>
                  <li>所有广告位的权重总和必须为 100</li>
                  <li>如果只有一个广告位，则权重自动设为 100</li>
                  <li>流量将按照权重比例分配到各个广告位</li>
                </ul>
              </template>
            </el-alert>

            <!-- 当前广告位信息 -->
            <el-card shadow="never" style="margin-bottom: 20px;">
              <template #header>
                <span style="font-weight: 600;">当前广告位</span>
              </template>
              <el-descriptions :column="3" border>
                <el-descriptions-item label="广告位名称">{{ currentSlotInfo.name }}</el-descriptions-item>
                <el-descriptions-item label="内部名称">{{ currentSlotInfo.name_alise }}</el-descriptions-item>
                <el-descriptions-item label="广告类型">{{ currentSlotInfo.scene_name }}</el-descriptions-item>
              </el-descriptions>
            </el-card>

            <!-- 分流配置列表 -->
            <el-card shadow="never">
              <template #header>
                <div class="flex justify-between items-center">
                  <span style="font-weight: 600;">分流广告位列表</span>
                  <el-button type="primary" icon="Plus" @click="addFlowSlot">添加广告位</el-button>
                </div>
              </template>

              <div v-if="flowSlots.length === 0" class="empty-state">
                <el-empty description="暂无分流配置，点击上方按钮添加广告位" />
              </div>

              <div v-else>
                <el-table
                  :data="flowSlots"
                  border
                  style="width: 100%; margin-bottom: 20px;"
                >
                  <el-table-column label="序号" width="60" align="center">
                    <template #default="scope">{{ scope.$index + 1 }}</template>
                  </el-table-column>

                  <el-table-column label="广告位" min-width="200">
                    <template #default="scope">
                      <el-select
                        v-model="scope.row.ad_slot_id"
                        placeholder="请选择广告位"
                        filterable
                        style="width: 100%;"
                        @change="handleSlotChange(scope.$index)"
                      >
                        <el-option
                          v-for="item in availableSlotOptions"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value"
                          :disabled="isSlotSelected(item.value, scope.$index)"
                        />
                      </el-select>
                    </template>
                  </el-table-column>

                  <el-table-column label="分流权重" width="150">
                    <template #default="scope">
                      <el-input-number
                        v-model="scope.row.weight"
                        :min="0"
                        :max="100"
                        :precision="0"
                        :controls="true"
                        style="width: 100%;"
                        @change="handleWeightChange"
                      />
                    </template>
                  </el-table-column>

                  <el-table-column label="占比显示" width="120" align="center">
                    <template #default="scope">
                      <el-tag :type="getWeightTagType(scope.row.weight)">{{ scope.row.weight }}%</el-tag>
                    </template>
                  </el-table-column>

                  <el-table-column label="操作" width="100" align="center">
                    <template #default="scope">
                      <el-button
                        type="danger"
                        link
                        icon="Delete"
                        @click="removeFlowSlot(scope.$index)"
                      >删除</el-button>
                    </template>
                  </el-table-column>
                </el-table>

                <!-- 权重统计 -->
                <div class="weight-summary">
                  <el-statistic group direction="horizontal">
                    <el-statistic-label>权重总和</el-statistic-label>
                    <el-statistic-number :class="{'is-error': totalWeight !== 100}">
                      {{ totalWeight }}
                    </el-statistic-number>
                    <el-statistic-number>/ 100</el-statistic-number>
                  </el-statistic>

                  <div v-if="totalWeight !== 100" class="weight-warning">
                    <el-alert
                      :title="totalWeight < 100 ? `权重不足，还差 ${100 - totalWeight}` : `权重超出，超出 ${totalWeight - 100}`"
                      type="warning"
                      :closable="false"
                      show-icon
                    />
                  </div>

                  <div v-else class="weight-success">
                    <el-alert
                      title="权重配置正确"
                      type="success"
                      :closable="false"
                      show-icon
                    />
                  </div>
                </div>

                <!-- 快速填充 -->
                <div v-if="flowSlots.length > 0" class="quick-actions">
                  <el-button
                    size="small"
                    @click="averageWeight"
                  >平均分配权重</el-button>
                  <el-button
                    size="small"
                    @click="autoFillWeight"
                  >自动补满100</el-button>
                  <el-button
                    size="small"
                    @click="resetAllWeight"
                  >重置所有权重</el-button>
                </div>
              </div>
            </el-card>
          </div>
    </el-drawer>
    <el-drawer destroy-on-close :size="appStore.drawerSizeMax" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
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
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="媒体名称:" prop="media_id">
                  <el-select
                      v-model="formData.media_id"
                      placeholder="请选择媒体"
                      filterable
                      :disabled="type === 'update'"
                      style="width:100%"
                  >
                    <el-option
                        v-for="item in mediaOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                  </el-select>
                </el-form-item>
              </el-col>

              <el-col :span="12">
                <el-form-item label="应用名称:" prop="app_id">
                  <el-select
                      v-model="formData.app_id"
                      placeholder="请选择应用"
                      filterable
                      :disabled="type === 'update'"
                      style="width:100%"
                  >
                    <el-option
                        v-for="item in appOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="广告位名称:" prop="name">
                  <el-input v-model="formData.name" :clearable="true" placeholder="请输入广告位名称" />
                </el-form-item>
              </el-col>

              <el-col :span="12">
                <el-form-item label="内部广告位名称:" prop="name_alise">
                  <el-input v-model="formData.name_alise" :clearable="true" placeholder="请输入内部广告位名称" />
                </el-form-item>
              </el-col>
            </el-row>

            <el-row :gutter="20">
              <!-- 广告类型 -->
              <el-col :span="12">
                <el-form-item label="广告类型:" prop="scene_id">
                  <el-select
                      v-model="formData.scene_id"
                      placeholder="请选择广告类型"
                      filterable
                      clearable
                      style="width:100%"
                  >
                    <el-option
                        v-for="item in sceneOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                  </el-select>
                </el-form-item>
              </el-col>

              <!-- 结算方式 + 分成系数 -->
              <el-col :span="12">
                <el-row :gutter="10">
                  <!-- 结算方式 -->
                  <el-col :span="14">
                    <el-form-item label="结算方式:" prop="ssp_pay_type">
                      <el-tree-select
                          v-model="formData.ssp_pay_type"
                          placeholder="请选择结算方式"
                          :data="pay_typeOptions"
                          style="width:100%"
                          filterable
                          clearable
                          check-strictly
                      />
                    </el-form-item>
                  </el-col>

                  <!-- 分成系数 -->
                  <el-col :span="10">
                    <el-form-item
                        label="分成系数:"
                        prop="ssp_deal_ratio"
                        :required="isSharePay"
                    >
                      <el-input
                          v-model.number="formData.ssp_deal_ratio"
                          placeholder="请输入分成系数"
                          :readonly="formData.ssp_pay_type !== '1'"
                          :class="{ 'is-disabled': formData.ssp_pay_type !== '1' }"
                      />

                    </el-form-item>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>

            <!-- ⭐⭐⭐ 交互类型（位掩码）⭐⭐⭐ -->
            <el-form-item label="交互类型" prop="interaction_type">
              <el-checkbox-group v-model="formData.interaction_type_arr">
                <el-checkbox
                    v-for="item in interactionOptions"
                    :key="item.value"
                    :value="item.value"
                    :label="item.label"
                />
              </el-checkbox-group>
            </el-form-item>

            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="广告位高:" prop="height">
                  <el-input v-model.number="formData.height" :clearable="true" placeholder="请输入广告位高" />
                </el-form-item>
              </el-col>

              <el-col :span="12">
                <el-form-item label="广告位宽:" prop="width">
                  <el-input v-model.number="formData.width" :clearable="true" placeholder="请输入广告位宽" />
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="广告位图片:" prop="ad_image">
    <el-input v-model="formData.ad_image" :clearable="true" placeholder="请输入广告位图片" />
</el-form-item>
            <el-form-item label="是否启用:" prop="enable">
    <el-tree-select v-model="formData.enable" placeholder="请选择是否启用" :data="enableOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="媒体Id">
    {{ detailForm.media_id }}
</el-descriptions-item>
                    <el-descriptions-item label="应用id">
    {{ detailForm.app_id }}
</el-descriptions-item>
                    <el-descriptions-item label="广告位名称">
    {{ detailForm.name }}
</el-descriptions-item>
                    <el-descriptions-item label="内部广告位名称">
    {{ detailForm.name_alise }}
</el-descriptions-item>
                    <el-descriptions-item label="广告类型">
    {{ detailForm.scene_id }}
</el-descriptions-item>
                    <el-descriptions-item label="结算方式">
    {{ detailForm.ssp_pay_type }}
</el-descriptions-item>
                    <el-descriptions-item label="分成系数">
    {{ detailForm.ssp_deal_ratio }}
</el-descriptions-item>
                    <el-descriptions-item label="交互类型">
    {{ detailForm.interaction_type }}
</el-descriptions-item>
                    <el-descriptions-item label="广告位高">
    {{ detailForm.height }}
</el-descriptions-item>
                    <el-descriptions-item label="广告位宽">
    {{ detailForm.width }}
</el-descriptions-item>
                    <el-descriptions-item label="广告位图片">
    {{ detailForm.ad_image }}
</el-descriptions-item>
                    <el-descriptions-item label="是否启用">
    {{ detailForm.enable }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createSsp_ad_slot,
  deleteSsp_ad_slot,
  updateSsp_ad_slot,
  findSsp_ad_slot,
  getSsp_ad_slotList
} from '@/api/ssp/sspAdSlot'

import { ref, reactive, onMounted, watch, computed, nextTick } from 'vue'
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from '@/pinia'
import { getSspMediaList } from '@/api/ssp/sspMedia'
import { getSspAppList } from '@/api/ssp/sspApp'
import { getDspAdSceneList } from '@/api/dsp/dspAdScene'
defineOptions({
  name: 'Ssp_ad_slot'
})

/* ================= 交互类型位掩码定义 ================= */
const interactionOptions = [
  { label: '打开网页', value: 1 },    // 1 << 0
  { label: 'deeplink', value: 2 },     // 1 << 1
  { label: '直接下载应用', value: 4 }, // 1 << 2
  { label: '广点通', value: 8 },
  { label: '应用商店下载', value: 16 },// 1 << 3
  { label: '小程序跳转', value: 32 },  // 1 << 4
  { label: '快应用', value: 64 },      // 1 << 6
]

/* ================= 工具函数 ================= */
const calcInteractionMask = (arr) =>
    arr.reduce((sum, v) => sum | v, 0)

const parseInteractionMask = (mask) =>
    interactionOptions
        .filter(i => (mask & i.value) === i.value)
        .map(i => i.value)




const appStore = useAppStore()
const btnLoading = ref(false)
const showAllQuery = ref(false)

// =================== 流量分流配置 ===================
const currentSlotInfo = ref({
  ID: null,
  name: '',
  name_alise: '',
  scene_name: ''
})
const flowSlots = ref([])
const allSlotOptions = ref([])

// =================== 字典 ===================
const pay_typeOptions = ref([])
const enableOptions = ref([])

// =================== 下拉数据源（新增） ===================
const mediaOptions = ref([])
const appOptions = ref([])
const sceneOptions = ref([])


// =================== 表单 ===================
const formData = ref({
  media_id: undefined,
  app_id: undefined,
  name: '',
  name_alise: '',
  scene_id: undefined,
  ssp_pay_type: '',
  ssp_deal_ratio: undefined,
  interaction_type: 0,        // ✅ 存数据库（int）
  interaction_type_arr: [],  // ✅ 前端 checkbox 用
  height: undefined,
  width: undefined,
  ad_image: '',
  enable: '',
  flow_config: []            // ✅ 流量分流配置
})

const isSharePay = computed(() => formData.value.ssp_pay_type === "1")

// 计算总权重
const totalWeight = computed(() => {
  return flowSlots.value.reduce((sum, item) => sum + (item.weight || 0), 0)
})

// 可选的广告位列表（排除当前广告位和已选广告位）
const availableSlotOptions = computed(() => {
  return allSlotOptions.value.filter(slot => {
    // 排除当前广告位
    if (slot.value === currentSlotInfo.value.ID) return false
    return true
  })
})

// 监听媒体变化 → 加载应用
watch(
    () => formData.value.media_id,
    (val) => {
      if (type.value === 'create') {
        formData.value.app_id = undefined
      }
      loadAppOptions(val)
    }
)

// 监听结算方式 → 非分成时清空分成系数
watch(
    () => formData.value.ssp_pay_type,
    (val) => {
      if (val !== '1') {
        formData.value.ssp_deal_ratio = undefined
      }
    }
)


// ⚠️ 先设 pay_type，再设 ratio
nextTick(() => {
  if (formData.value.ssp_pay_type === '1') {
    formData.value.ssp_deal_ratio = res.data.ssp_deal_ratio
  }
})

const loadMediaOptions = async () => {
  const res = await getSspMediaList({ page: 1, pageSize: 1000 })
  if (res.code === 0) {
    mediaOptions.value = res.data.list.map(item => ({
      label: item.name,
      value: item.ID
    }))
  }
}


const loadSceneOptions = async () => {
  const res = await getDspAdSceneList({ page: 1, pageSize: 1000 })
  if (res.code === 0) {
    sceneOptions.value = res.data.list.map(item => ({
      label: item.name,   // 广告类型名称
      value: item.ID     // 广告类型ID
    }))
  }
}

const loadAppOptions = async (mediaId) => {
  if (!mediaId) {
    appOptions.value = []
    return
  }

  const res = await getSspAppList({
    page: 1,
    pageSize: 1000,
    m_id: mediaId   // ⭐关键
  })

  if (res.code === 0) {
    appOptions.value = res.data.list.map(item => ({
      label: item.name,
      value: item.ID
    }))
  }
}



const rule = reactive({
  media_id: [{ required: true, message: '请选择媒体', trigger: 'change' }],
  app_id: [{ required: true, message: '请选择应用', trigger: 'change' }],
  name: [{ required: true, message: '请输入广告位名称', trigger: 'blur' }],
  scene_id: [{ required: true, message: '请选择广告类型', trigger: 'change' }],
  ssp_pay_type: [{ required: true, message: '请选择结算方式', trigger: 'change' }],
  enable: [{ required: true, message: '请选择是否启用', trigger: 'change' }],
  ssp_deal_ratio: [
    {
      validator: (_, value, callback) => {
        if (formData.value.ssp_pay_type === 1 && (value === undefined || value === null)) {
          callback(new Error('分成方式下必须填写分成系数'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =================== 表格 ===================
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})

const getTableData = async () => {
  const res = await getSsp_ad_slotList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const handleCurrentChange = val => {
  page.value = val
  getTableData()
}

const handleSizeChange = val => {
  pageSize.value = val
  getTableData()
}

// =================== 多选 ===================
const multipleSelection = ref([])
const handleSelectionChange = val => {
  multipleSelection.value = val
}

// =================== 删除 ===================
const deleteRow = row => {
  ElMessageBox.confirm('确定删除吗？', '提示').then(async () => {
    const res = await deleteSsp_ad_slot({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// =================== 弹窗 ===================
const dialogFormVisible = ref(false)
const dialogFormVisibleConfig = ref(false)
const detailShow = ref(false)
const detailForm = ref({})
const type = ref('create')

const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    media_id: undefined,
    app_id: undefined,
    name: '',
    name_alise: '',
    scene_id: undefined,
    ssp_pay_type: '',
    ssp_deal_ratio: undefined,
    interaction_type: 0,
    interaction_type_arr: [],
    height: undefined,
    width: undefined,
    ad_image: '',
    enable: '',
    flow_config: []
  }
}

const enterDialog = async () => {
  btnLoading.value = true

  elFormRef.value.validate(async valid => {
    if (!valid) {
      btnLoading.value = false
      return
    }

    // ⭐⭐⭐ 位掩码计算（核心）
    formData.value.interaction_type =
        calcInteractionMask(formData.value.interaction_type_arr)

    const res =
        type.value === 'create'
            ? await createSsp_ad_slot(formData.value)
            : await updateSsp_ad_slot(formData.value)

    btnLoading.value = false

    if (res.code === 0) {
      ElMessage.success('操作成功')
      closeDialog()
      getTableData()
    }
  })
}


// =================== 编辑 ===================
// 关闭详情drawer
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}

// =================== 编辑 ===================
const updateSsp_ad_slotFunc = async (row) => {
  const res = await findSsp_ad_slot({ ID: row.ID })
  if (res.code === 0) {
    type.value = 'update'

    formData.value = {
      ...res.data,
      interaction_type_arr: parseInteractionMask(res.data.interaction_type)
    }

    await loadAppOptions(res.data.media_id)

    nextTick(() => {
      if (formData.value.ssp_pay_type === 1) {
        formData.value.ssp_deal_ratio = res.data.ssp_deal_ratio
      }
    })

    dialogFormVisible.value = true
  }
}

// =================== 配置 ===================

// 加载所有广告位选项
const loadAllSlotOptions = async () => {
  const res = await getSsp_ad_slotList({ page: 1, pageSize: 1000 })
  if (res.code === 0) {
    allSlotOptions.value = res.data.list.map(item => ({
      label: `${item.name} (${item.name_alise})`,
      value: item.ID,
      scene_id: item.scene_id
    }))
  }
}

// 打开配置drawer
const updateSsp_ad_slotConfig = async (row) => {
  const res = await findSsp_ad_slot({ ID: row.ID })
  if (res.code === 0) {
    // 获取广告类型名称
    const sceneOption = sceneOptions.value.find(s => s.value === res.data.scene_id)

    // 设置当前广告位信息
    currentSlotInfo.value = {
      ID: res.data.ID,
      name: res.data.name,
      name_alise: res.data.name_alise,
      scene_name: sceneOption ? sceneOption.label : res.data.scene_id
    }

    // 加载所有广告位列表
    await loadAllSlotOptions()

    // 解析已有的流量分流配置
    if (res.data.flow_config && Array.isArray(res.data.flow_config)) {
      flowSlots.value = res.data.flow_config.map(item => ({
        ad_slot_id: item.ad_slot_id,
        weight: item.weight
      }))
    } else {
      flowSlots.value = []
    }

    dialogFormVisibleConfig.value = true
  }
}

// 关闭配置drawer
const closeConfigDialog = () => {
  dialogFormVisibleConfig.value = false
  flowSlots.value = []
  currentSlotInfo.value = {
    ID: null,
    name: '',
    name_alise: '',
    scene_name: ''
  }
}

// 添加流量绑定广告位
const addFlowSlot = () => {
  flowSlots.value.push({
    ad_slot_id: undefined,
    weight: 0
  })

  // 如果只有一个广告位，自动设置为100
  if (flowSlots.value.length === 1) {
    flowSlots.value[0].weight = 100
  } else {
    // 如果添加了第二个或更多广告位，需要重新分配权重
    redistributeWeight()
  }
}

// 删除流量绑定
const removeFlowSlot = (index) => {
  flowSlots.value.splice(index, 1)

  // 删除后如果只剩一个，自动设置为100
  if (flowSlots.value.length === 1) {
    flowSlots.value[0].weight = 100
  } else if (flowSlots.value.length === 0) {
    // 如果全部删除，保持空数组
  } else {
    // 重新分配剩余权重
    redistributeWeight()
  }
}

// 广告位选择变化
const handleSlotChange = (index) => {
  // 可以在这里添加验证逻辑
}

// 判断广告位是否已被选择
const isSlotSelected = (slotId, currentIndex) => {
  return flowSlots.value.some((item, index) =>
    index !== currentIndex && item.ad_slot_id === slotId
  )
}

// 权重变化处理
const handleWeightChange = () => {
  // 权重变化时的处理
}

// 重新分配权重（当添加/删除广告位时）
const redistributeWeight = () => {
  const total = totalWeight.value
  const count = flowSlots.value.length

  if (count === 0) return

  // 如果总权重不是100，按比例重新分配
  if (total !== 100 && total > 0) {
    flowSlots.value.forEach(item => {
      item.weight = Math.round((item.weight / total) * 100)
    })
  } else if (total === 0) {
    // 如果所有权重都是0，平均分配
    const avgWeight = Math.floor(100 / count)
    flowSlots.value.forEach((item, index) => {
      item.weight = index === count - 1 ? 100 - avgWeight * (count - 1) : avgWeight
    })
  }

  // 确保总和为100
  adjustTotalWeight()
}

// 调整总权重为100
const adjustTotalWeight = () => {
  const total = flowSlots.value.reduce((sum, item) => sum + (item.weight || 0), 0)
  if (total !== 100 && flowSlots.value.length > 0) {
    const diff = 100 - total
    flowSlots.value[flowSlots.value.length - 1].weight += diff
  }
}

// 平均分配权重
const averageWeight = () => {
  const count = flowSlots.value.length
  if (count === 0) return

  const avgWeight = Math.floor(100 / count)
  flowSlots.value.forEach((item, index) => {
    item.weight = index === count - 1 ? 100 - avgWeight * (count - 1) : avgWeight
  })
}

// 自动补满100
const autoFillWeight = () => {
  const currentTotal = totalWeight.value
  const diff = 100 - currentTotal

  if (diff > 0 && flowSlots.value.length > 0) {
    flowSlots.value[flowSlots.value.length - 1].weight += diff
  }
}

// 重置所有权重
const resetAllWeight = () => {
  flowSlots.value.forEach(item => {
    item.weight = 0
  })
}

// 获取权重标签类型
const getWeightTagType = (weight) => {
  if (weight === 0) return 'info'
  if (weight < 30) return ''
  if (weight < 60) return 'warning'
  if (weight < 80) return 'success'
  return 'danger'
}

// 保存配置
const saveConfig = async () => {
  // 验证至少有一个流量绑定
  if (flowSlots.value.length === 0) {
    ElMessage.warning('请至少添加一个流量绑定')
    return
  }

  // 验证所有广告位都已选择
  const hasEmptySlot = flowSlots.value.some(item => !item.ad_slot_id)
  if (hasEmptySlot) {
    ElMessage.warning('请为所有流量绑定选择广告位')
    return
  }

  // 验证总权重为100
  if (totalWeight.value !== 100) {
    ElMessage.warning('所有权重总和必须为100')
    return
  }

  btnLoading.value = true

  try {
    // 构建保存数据
    const saveData = {
      ID: currentSlotInfo.value.ID,
      flow_config: flowSlots.value.map(item => ({
        ad_slot_id: item.ad_slot_id,
        weight: item.weight
      }))
    }

    const res = await updateSsp_ad_slot(saveData)

    btnLoading.value = false

    if (res.code === 0) {
      ElMessage.success('配置保存成功')
      closeConfigDialog()
      getTableData()
    }
  } catch (error) {
    btnLoading.value = false
    ElMessage.error('保存失败：' + error.message)
  }
}

// =================== 初始化 ===================
const setOptions = async () => {
  pay_typeOptions.value = await getDictFunc('pay_type')
  enableOptions.value = await getDictFunc('enable')
}

onMounted(() => {
  setOptions()
  loadMediaOptions()
  loadSceneOptions()
  getTableData()
})


</script>
<style>
.table-config {

}
</style>