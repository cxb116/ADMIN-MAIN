
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="公司名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入公司名称" />
</el-form-item>
        <el-form-item label="预算隐射:" prop="dsp_code">
    <el-input v-model.number="formData.dsp_code" :clearable="true" placeholder="请输入预算隐射" />
</el-form-item>
        <el-form-item label="请求地址:" prop="url">
    <el-input v-model="formData.url" :clearable="true" placeholder="请输入请求地址" />
</el-form-item>
        <el-form-item label="请求方法:" prop="method">
    <el-tree-select v-model="formData.method" placeholder="请选择请求方法" :data="method_typeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="超时时间:" prop="timeout">
    <el-input v-model.number="formData.timeout" :clearable="true" placeholder="请输入超时时间" />
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
  createDspCompany,
  updateDspCompany,
  findDspCompany
} from '@/api/dsp/dspCompany'

defineOptions({
    name: 'DspCompanyForm'
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
const method_typeOptions = ref([])
const formData = ref({
            name: '',
            dsp_code: undefined,
            url: '',
            method: '',
            timeout: undefined,
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dsp_code : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               url : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               method : [{
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
      const res = await findDspCompany({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    method_typeOptions.value = await getDictFunc('method_type')
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
               res = await createDspCompany(formData.value)
               break
             case 'update':
               res = await updateDspCompany(formData.value)
               break
             default:
               res = await createDspCompany(formData.value)
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
