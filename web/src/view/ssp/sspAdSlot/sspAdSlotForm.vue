
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="媒体Id:" prop="media_id">
    <el-input v-model.number="formData.media_id" :clearable="true" placeholder="请输入媒体Id" />
</el-form-item>
        <el-form-item label="应用id:" prop="app_id">
    <el-input v-model.number="formData.app_id" :clearable="true" placeholder="请输入应用id" />
</el-form-item>
        <el-form-item label="广告位名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入广告位名称" />
</el-form-item>
        <el-form-item label="内部广告位名称:" prop="name_alise">
    <el-input v-model="formData.name_alise" :clearable="true" placeholder="请输入内部广告位名称" />
</el-form-item>
        <el-form-item label="广告类型:" prop="scene_id">
    <el-input v-model.number="formData.scene_id" :clearable="true" placeholder="请输入广告类型" />
</el-form-item>
        <el-form-item label="结算方式:" prop="ssp_pay_type">
    <el-tree-select v-model="formData.ssp_pay_type" placeholder="请选择结算方式" :data="pay_typeOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
</el-form-item>
        <el-form-item label="分成系数:" prop="ssp_deal_ratio">
    <el-input v-model.number="formData.ssp_deal_ratio" :clearable="true" placeholder="请输入分成系数" />
</el-form-item>
        <el-form-item label="交互类型:" prop="interaction_type">
    <el-input v-model.number="formData.interaction_type" :clearable="true" placeholder="请输入交互类型" />
</el-form-item>
        <el-form-item label="广告位高:" prop="height">
    <el-input v-model.number="formData.height" :clearable="true" placeholder="请输入广告位高" />
</el-form-item>
        <el-form-item label="广告位宽:" prop="width">
    <el-input v-model.number="formData.width" :clearable="true" placeholder="请输入广告位宽" />
</el-form-item>
        <el-form-item label="广告位图片:" prop="ad_image">
    <el-input v-model="formData.ad_image" :clearable="true" placeholder="请输入广告位图片" />
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
  createSsp_ad_slot,
  updateSsp_ad_slot,
  findSsp_ad_slot
} from '@/api/ssp/sspAdSlot'

defineOptions({
    name: 'Ssp_ad_slotForm'
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
const pay_typeOptions = ref([])
const enableOptions = ref([])
const formData = ref({
            media_id: undefined,
            app_id: undefined,
            name: '',
            name_alise: '',
            scene_id: undefined,
            ssp_pay_type: '',
            ssp_deal_ratio: undefined,
            interaction_type: undefined,
            height: undefined,
            width: undefined,
            ad_image: '',
            enable: '',
        })
// 验证规则
const rule = reactive({
               media_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               app_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               scene_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               ssp_pay_type : [{
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
      const res = await findSsp_ad_slot({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    pay_typeOptions.value = await getDictFunc('pay_type')
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
               res = await createSsp_ad_slot(formData.value)
               break
             case 'update':
               res = await updateSsp_ad_slot(formData.value)
               break
             default:
               res = await createSsp_ad_slot(formData.value)
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
