<template>
  <div>
    <!-- ================= 搜索区域 ================= -->
    <div class="gva-search-box">
      <el-form
          ref="elSearchFormRef"
          :inline="true"
          :model="searchInfo"
          @keyup.enter="onSubmit"
      >
        <!-- 创建时间 -->
        <el-form-item prop="createdAtRange">
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
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始时间"
              end-placeholder="结束时间"
              class="!w-380px"
          />
        </el-form-item>

        <!-- 媒体（和表单一致） -->
        <el-form-item label="媒体">
          <el-select
              v-model="searchInfo.m_id"
              filterable
              clearable
              :loading="mediaLoading"
              placeholder="请选择媒体"
              style="width: 180px"
          >
            <el-option
                v-for="item in mediaList"
                :key="item.id"
                :label="`${item.name}（${item.id}）`"
                :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="应用名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="操作系统">
          <el-tree-select
              v-model="searchInfo.os_type"
              :data="os_typeOptions"
              filterable
              clearable
              check-strictly
              style="width: 160px"
          />
        </el-form-item>

        <el-form-item label="接入方式">
          <el-tree-select
              v-model="searchInfo.access_type"
              :data="access_typeOptions"
              filterable
              clearable
              check-strictly
              style="width: 160px"
          />
        </el-form-item>

        <el-form-item label="是否启用">
          <el-tree-select
              v-model="searchInfo.enable"
              :data="enableOptions"
              filterable
              clearable
              check-strictly
              style="width: 160px"
          />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- ================= 表格 ================= -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button
            icon="delete"
            :disabled="!multipleSelection.length"
            @click="onDelete"
        >
          删除
        </el-button>
      </div>

      <el-table
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="应用名称" prop="name" width="140" />
        <el-table-column label="媒体名称" prop="media_name" width="150" />
        <el-table-column label="操作系统" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.os_type, os_typeOptions) }}
          </template>
        </el-table-column>
        <el-table-column label="接入方式" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.access_type, access_typeOptions) }}
          </template>
        </el-table-column>
        <el-table-column label="是否启用" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.enable, enableOptions) }}
          </template>
        </el-table-column>

        <el-table-column label="日期" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="scope">
            <el-button link type="primary" @click="updateSspAppFunc(scope.row)">编辑</el-button>
            <el-button link type="danger" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10,30,50,100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
      />
    </div>

    <!-- ================= 新增 / 编辑抽屉 ================= -->
    <el-drawer v-model="dialogFormVisible" destroy-on-close size="50%">
      <template #header>
        <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
      </template>

      <el-form ref="elFormRef" :model="formData" :rules="rules" label-width="80px">
        <el-form-item label="媒体" prop="m_id">
          <el-select
              v-model="formData.m_id"
              filterable
              clearable
              :loading="mediaLoading"
              :disabled="type === 'update'"
              placeholder="请选择媒体"
              style="width: 100%"
          >
            <el-option
                v-for="item in mediaList"
                :key="item.ID"
                :label="item.name + '(' + item.ID + ')'"
                :value="item.ID"
            />

          </el-select>
        </el-form-item>

        <el-form-item label="应用名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入应用名称"
          />
        </el-form-item>

        <el-form-item label="操作系统" prop="os_type">
          <el-tree-select
              v-model="formData.os_type"
              :data="os_typeOptions"
              check-strictly
              style="width:100%"
          />
        </el-form-item>

        <el-form-item label="接入方式" prop="access_type">
          <el-tree-select
              v-model="formData.access_type"
              :data="access_typeOptions"
              check-strictly
              style="width:100%"
          />
        </el-form-item>

        <el-form-item label="是否启用" prop="enable">
          <el-tree-select
              v-model="formData.enable"
              :data="enableOptions"
              check-strictly
              style="width:100%"
          />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" :loading="btnLoading" @click="enterDialog">
            确定
          </el-button>
          <el-button @click="closeDialog">取消</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getSspAppList,
  createSspApp,
  updateSspApp,
  deleteSspApp,
  deleteSspAppByIds,
  findSspApp
} from '@/api/ssp/sspApp'
import { getSspMediaList } from '@/api/ssp/sspMedia'
import { getDictFunc, formatDate, filterDict } from '@/utils/format'

/* ================= 基础状态 ================= */
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})

const mediaList = ref([])
const mediaLoading = ref(false)

const enableOptions = ref([])
const os_typeOptions = ref([])
const access_typeOptions = ref([])

/* ================= 表单 ================= */
const dialogFormVisible = ref(false)
const elFormRef = ref()
const btnLoading = ref(false)
const type = ref('create')

const formData = ref({
  m_id: undefined,
  name: '',
  os_type: '',
  access_type: '',
  enable: ''
})

const rules = {
  m_id: [{ required: true, message: '请选择媒体', trigger: 'change' }],
  name: [{ required: true, message: '请输入应用名称', trigger: 'blur' }],
  os_type: [{ required: true, message: '请选择操作系统', trigger: 'change' }],
  access_type: [{ required: true, message: '请选择接入方式', trigger: 'change' }],
  enable: [{ required: true, message: '请选择是否启用', trigger: 'change' }]
}

/* ================= 初始化 ================= */
onMounted(async () => {
  await loadMedia()
  await loadDict()
  getTableData()
})

const loadMedia = async () => {
  mediaLoading.value = true
  const res = await getSspMediaList({ enable: 1 })
  mediaList.value = res.data.list || []

  console.log("mediaList:",mediaList)
  mediaLoading.value = false
}

const loadDict = async () => {
  enableOptions.value = await getDictFunc('enable')
  os_typeOptions.value = await getDictFunc('os_type')
  access_typeOptions.value = await getDictFunc('access_type')
}

/* ================= 表格 ================= */
const getTableData = async () => {
  const res = await getSspAppList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

const handleSizeChange = val => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = val => {
  page.value = val
  getTableData()
}

/* ================= 新增 / 编辑 ================= */
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const updateSspAppFunc = async row => {
  const res = await findSspApp({ ID: row.ID })
  if (res.code === 0) {
    formData.value = { ...res.data, m_id: Number(res.data.m_id) }
    type.value = 'update'
    dialogFormVisible.value = true
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
}

const enterDialog = () => {
  elFormRef.value.validate(async valid => {
    if (!valid) return
    btnLoading.value = true
    const api = type.value === 'create' ? createSspApp : updateSspApp
    const res = await api(formData.value)
    btnLoading.value = false
    if (res.code === 0) {
      ElMessage.success('操作成功')
      closeDialog()
      getTableData()
    }
  })
}

/* ================= 删除 ================= */
const multipleSelection = ref([])

const handleSelectionChange = val => {
  multipleSelection.value = val
}

const deleteRow = row => {
  ElMessageBox.confirm('确认删除？', '提示').then(async () => {
    await deleteSspApp({ ID: row.ID })
    getTableData()
  })
}

const onDelete = () => {
  if (!multipleSelection.value.length) return
  const IDs = multipleSelection.value.map(i => i.ID)
  ElMessageBox.confirm('确认删除？', '提示').then(async () => {
    await deleteSspAppByIds({ IDs })
    getTableData()
  })
}
</script>
