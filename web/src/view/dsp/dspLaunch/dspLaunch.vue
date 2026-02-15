
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
            <el-form-item label="投放策略" prop="launchStrategy">
    <el-tree-select v-model="searchInfo.launchStrategy" placeholder="请选择投放策略" :data="launch_strategyOptions" style="width:100%" filterable :clearable="true" check-strictly multiple ></el-tree-select>
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
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="id字段" prop="id" width="120" />

            <el-table-column align="left" label="创建时间" prop="createdAt" width="120" />

            <el-table-column align="left" label="更新时间" prop="updatedAt" width="120" />

            <el-table-column align="left" label="删除时间" prop="deletedAt" width="120" />

            <el-table-column align="left" label="流量广告位id" prop="sspSlotId" width="120" />

            <el-table-column align="left" label="流量权重" prop="trafficWeight" width="120" />

            <el-table-column align="left" label="投放策略" prop="launchStrategy" width="120">
    <template #default="scope">
    <el-tag class="mr-1" v-for="item in scope.row.launchStrategy" :key="item"> {{ filterDict(item,launch_strategyOptions) }}</el-tag>
    </template>
</el-table-column>
            <el-table-column align="left" label="底价" prop="floorPrice" width="120" />

            <el-table-column align="left" label="ip限流次数" prop="ipLimit" width="120" />

            <el-table-column align="left" label="捕获日志时长" prop="logCaptureAt" width="120" />

            <el-table-column align="left" label="上报黑名单" prop="trackSchwarz" width="120" />

            <el-table-column align="left" label="请求次数(次/天)" prop="req" width="120" />

            <el-table-column align="left" label="展现次数" prop="ims" width="120" />

            <el-table-column align="left" label="点击次数" prop="clk" width="120" />

            <el-table-column align="left" label="投放时段" prop="launchTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.launchTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="人群定向" prop="crowdDirection" width="120">
    <template #default="scope">
    <el-tag class="mr-1" v-for="item in scope.row.crowdDirection" :key="item"> {{ filterDict(item,direction_typeOptions) }}</el-tag>
    </template>
</el-table-column>
            <el-table-column align="left" label="地域定向" prop="regionDirection" width="120">
    <template #default="scope">
    <el-tag class="mr-1" v-for="item in scope.row.regionDirection" :key="item"> {{ filterDict(item,direction_typeOptions) }}</el-tag>
    </template>
</el-table-column>
            <el-table-column align="left" label="品牌定向" prop="brandDirection" width="120">
    <template #default="scope">
    <el-tag class="mr-1" v-for="item in scope.row.brandDirection" :key="item"> {{ filterDict(item,direction_typeOptions) }}</el-tag>
    </template>
</el-table-column>
            <el-table-column align="left" label="备注" prop="remark" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateDspLaunchFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="流量权重:" prop="trafficWeight">
    <el-input v-model="formData.trafficWeight" :clearable="true" placeholder="请输入流量权重" />
</el-form-item>
            <el-form-item label="投放策略:" prop="launchStrategy">
    <el-select multiple v-model="formData.launchStrategy" placeholder="请选择投放策略" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in launch_strategyOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="底价:" prop="floorPrice">
    <el-input v-model.number="formData.floorPrice" :clearable="true" placeholder="请输入底价" />
</el-form-item>
            <el-form-item label="ip限流次数:" prop="ipLimit">
    <el-input v-model.number="formData.ipLimit" :clearable="true" placeholder="请输入ip限流次数" />
</el-form-item>
            <el-form-item label="捕获日志时长:" prop="logCaptureAt">
    <el-input v-model.number="formData.logCaptureAt" :clearable="true" placeholder="请输入捕获日志时长" />
</el-form-item>
            <el-form-item label="上报黑名单:" prop="trackSchwarz">
    <el-input v-model="formData.trackSchwarz" :clearable="true" placeholder="请输入上报黑名单" />
</el-form-item>
            <el-form-item label="请求次数(次/天):" prop="req">
    <el-input v-model.number="formData.req" :clearable="true" placeholder="请输入请求次数(次/天)" />
</el-form-item>
            <el-form-item label="展现次数:" prop="ims">
    <el-input v-model.number="formData.ims" :clearable="true" placeholder="请输入展现次数" />
</el-form-item>
            <el-form-item label="点击次数:" prop="clk">
    <el-input v-model.number="formData.clk" :clearable="true" placeholder="请输入点击次数" />
</el-form-item>
            <el-form-item label="投放时段:" prop="launchTime">
    <el-date-picker v-model="formData.launchTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="人群定向:" prop="crowdDirection">
    <el-select multiple v-model="formData.crowdDirection" placeholder="请选择人群定向" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in direction_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="地域定向:" prop="regionDirection">
    <el-select multiple v-model="formData.regionDirection" placeholder="请选择地域定向" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in direction_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="品牌定向:" prop="brandDirection">
    <el-select multiple v-model="formData.brandDirection" placeholder="请选择品牌定向" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in direction_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="创建时间">
    {{ detailForm.createdAt }}
</el-descriptions-item>
                    <el-descriptions-item label="更新时间">
    {{ detailForm.updatedAt }}
</el-descriptions-item>
                    <el-descriptions-item label="删除时间">
    {{ detailForm.deletedAt }}
</el-descriptions-item>
                    <el-descriptions-item label="流量权重">
    {{ detailForm.trafficWeight }}
</el-descriptions-item>
                    <el-descriptions-item label="投放策略">
    <ArrayCtrl v-model="detailForm.launchStrategy"/>
</el-descriptions-item>
                    <el-descriptions-item label="底价">
    {{ detailForm.floorPrice }}
</el-descriptions-item>
                    <el-descriptions-item label="ip限流次数">
    {{ detailForm.ipLimit }}
</el-descriptions-item>
                    <el-descriptions-item label="捕获日志时长">
    {{ detailForm.logCaptureAt }}
</el-descriptions-item>
                    <el-descriptions-item label="上报黑名单">
    {{ detailForm.trackSchwarz }}
</el-descriptions-item>
                    <el-descriptions-item label="请求次数(次/天)">
    {{ detailForm.req }}
</el-descriptions-item>
                    <el-descriptions-item label="展现次数">
    {{ detailForm.ims }}
</el-descriptions-item>
                    <el-descriptions-item label="点击次数">
    {{ detailForm.clk }}
</el-descriptions-item>
                    <el-descriptions-item label="投放时段">
    {{ detailForm.launchTime }}
</el-descriptions-item>
                    <el-descriptions-item label="人群定向">
    <ArrayCtrl v-model="detailForm.crowdDirection"/>
</el-descriptions-item>
                    <el-descriptions-item label="地域定向">
    <ArrayCtrl v-model="detailForm.regionDirection"/>
</el-descriptions-item>
                    <el-descriptions-item label="品牌定向">
    <ArrayCtrl v-model="detailForm.brandDirection"/>
</el-descriptions-item>
                    <el-descriptions-item label="备注">
    {{ detailForm.remark }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createDspLaunch,
  deleteDspLaunch,
  deleteDspLaunchByIds,
  updateDspLaunch,
  findDspLaunch,
  getDspLaunchList
} from '@/api/dsp/dspLaunch'
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'DspLaunch'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const launch_strategyOptions = ref([])
const direction_typeOptions = ref([])
const formData = ref({
            trafficWeight: '',
            launchStrategy: [],
            floorPrice: undefined,
            ipLimit: undefined,
            logCaptureAt: undefined,
            trackSchwarz: '',
            req: undefined,
            ims: undefined,
            clk: undefined,
            launchTime: new Date(),
            crowdDirection: [],
            regionDirection: [],
            brandDirection: [],
            remark: '',
        })



// 验证规则
const rule = reactive({
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

// 查询
const getTableData = async() => {
  const table = await getDspLaunchList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    launch_strategyOptions.value = await getDictFunc('launch_strategy')
    direction_typeOptions.value = await getDictFunc('direction_type')
}

// 获取需要的字典 可能为空 按需保留
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
            deleteDspLaunchFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteDspLaunchByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateDspLaunchFunc = async(row) => {
    const res = await findDspLaunch({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteDspLaunchFunc = async (row) => {
    const res = await deleteDspLaunch({ id: row.id })
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
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        trafficWeight: '',
        launchStrategy: [],
        floorPrice: undefined,
        ipLimit: undefined,
        logCaptureAt: undefined,
        trackSchwarz: '',
        req: undefined,
        ims: undefined,
        clk: undefined,
        launchTime: new Date(),
        crowdDirection: [],
        regionDirection: [],
        brandDirection: [],
        remark: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createDspLaunch(formData.value)
                  break
                case 'update':
                  res = await updateDspLaunch(formData.value)
                  break
                default:
                  res = await createDspLaunch(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
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
  const res = await findDspLaunch({ id: row.id })
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


</script>

<style>

</style>
