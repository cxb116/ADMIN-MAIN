
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="媒体名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入媒体名称" />
</el-form-item>
        <el-form-item label="媒体账号:" prop="account">
    <el-input v-model="formData.account" :clearable="true" placeholder="请输入媒体账号" />
</el-form-item>
        <el-form-item label="密码:" prop="password">
    <el-input v-model="formData.password" :clearable="true" placeholder="请输入密码" />
</el-form-item>
        <el-form-item label="公司名称:" prop="media_company_name">
    <el-input v-model="formData.media_company_name" :clearable="true" placeholder="请输入公司名称" />
</el-form-item>
        <el-form-item label="公司简称:" prop="media_company_short">
    <el-input v-model="formData.media_company_short" :clearable="true" placeholder="请输入公司简称" />
</el-form-item>
        <el-form-item label="社会信用代码:" prop="media_company_code">
    <el-input v-model="formData.media_company_code" :clearable="true" placeholder="请输入社会信用代码" />
</el-form-item>
        <el-form-item label="营业执照:" prop="media_company_license">
    <el-input v-model="formData.media_company_license" :clearable="true" placeholder="请输入营业执照" />
</el-form-item>
        <el-form-item label="公司地址:" prop="media_company_address">
    <el-input v-model="formData.media_company_address" :clearable="true" placeholder="请输入公司地址" />
</el-form-item>
        <el-form-item label="法人姓名:" prop="media_owner_name">
    <el-input v-model="formData.media_owner_name" :clearable="true" placeholder="请输入法人姓名" />
</el-form-item>
        <el-form-item label="联系人:" prop="contact_name">
    <el-input v-model="formData.contact_name" :clearable="true" placeholder="请输入联系人" />
</el-form-item>
        <el-form-item label="联系电话:" prop="contact_phone">
    <el-input v-model="formData.contact_phone" :clearable="true" placeholder="请输入联系电话" />
</el-form-item>
        <el-form-item label="联系邮箱:" prop="contact_email">
    <el-input v-model="formData.contact_email" :clearable="true" placeholder="请输入联系邮箱" />
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
  createSspMedia,
  updateSspMedia,
  findSspMedia
} from '@/api/ssp/sspMedia'

defineOptions({
    name: 'SspMediaForm'
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
const formData = ref({
            name: '',
            account: '',
            password: '',
            media_company_name: '',
            media_company_short: '',
            media_company_code: '',
            media_company_license: '',
            media_company_address: '',
            media_owner_name: '',
            contact_name: '',
            contact_phone: '',
            contact_email: '',
            enable: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               account : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               media_company_name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               media_owner_name : [{
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
      const res = await findSspMedia({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    enableOptions.value = await getDictFunc('enable')
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
               res = await createSspMedia(formData.value)
               break
             case 'update':
               res = await updateSspMedia(formData.value)
               break
             default:
               res = await createSspMedia(formData.value)
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
