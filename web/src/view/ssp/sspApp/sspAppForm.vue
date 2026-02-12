
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="媒体id:" prop="m_id">
    <el-input v-model.number="formData.m_id" :clearable="true" placeholder="请输入媒体id" />
</el-form-item>
        <el-form-item label="应用名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入应用名称" />
</el-form-item>
        <el-form-item label="操作系统:" prop="os_type">
    <el-tree-select v-model="formData.os_type" placeholder="请选择操作系统" :data="os_typeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="接入方式:" prop="access_type">
    <el-tree-select v-model="formData.access_type" placeholder="请选择接入方式" :data="access_typeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="包名:" prop="pkg">
    <el-input v-model="formData.pkg" :clearable="true" placeholder="请输入包名" />
</el-form-item>
        <el-form-item label="下载地址:" prop="download_url">
    <el-input v-model="formData.download_url" :clearable="true" placeholder="请输入下载地址" />
</el-form-item>
        <el-form-item label="是否启用:" prop="enable">
    <el-tree-select v-model="formData.enable" placeholder="请选择是否启用" :data="enableOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSspApp,
  updateSspApp,
  findSspApp
} from '@/api/ssp/sspApp'

defineOptions({
    name: 'SspAppForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const enableOptions = ref([])
const os_typeOptions = ref([])
const access_typeOptions = ref([])
const formData = ref({
            m_id: undefined,
            name: '',
            os_type: '',
            access_type: '',
            pkg: '',
            download_url: '',
            enable: '',
        })
// 验证规则
const rule = reactive({
               m_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               os_type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               access_type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               enable : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSspApp({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    enableOptions.value = await getDictFunc('enable')
    os_typeOptions.value = await getDictFunc('os_type')
    access_typeOptions.value = await getDictFunc('access_type')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createSspApp(formData.value)
               break
             case 'update':
               res = await updateSspApp(formData.value)
               break
             default:
               res = await createSspApp(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
