
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="预算位名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入预算位名称" />
</el-form-item>
        <el-form-item label="广告类型:" prop="scene_id">
    <el-input v-model.number="formData.scene_id" :clearable="true" placeholder="请输入广告类型" />
</el-form-item>
        <el-form-item label="预算方广告位:" prop="dsp_slot_code">
    <el-input v-model="formData.dsp_slot_code" :clearable="true" placeholder="请输入预算方广告位" />
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
        <el-form-item label="公司产品:" prop="cascaderValue">
    <el-cascader
      v-model="cascaderValue"
      :options="cascaderOptions"
      :props="{ expandTrigger: 'hover' }"
      placeholder="请选择公司产品"
      style="width: 100%"
      clearable
      filterable
    />
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <RichEdit v-model="formData.remark"/>
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
  createDspSlotInfo,
  updateDspSlotInfo,
  findDspSlotInfo
} from '@/api/dsp/dspSlotInfo'
import {
  Cascader
} from '@/api/dsp/dspProduct'

defineOptions({
    name: 'DspSlotInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const pay_typeOptions = ref([])
const cascaderOptions = ref([])
const cascaderValue = ref([])
const formData = ref({
            name: '',
            scene_id: undefined,
            dsp_slot_code: '',
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
            remark: '',
        })
// 验证规则
const rule = reactive({
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
               dsp_slot_code : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cascaderValue : [{
                   required: true,
                   message: '请选择公司产品',
                   trigger: ['change','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    // 获取级联选择器数据
    const cascaderRes = await Cascader()
    if (cascaderRes.code === 0) {
      cascaderOptions.value = cascaderRes.data
    }

    if (route.query.id) {
      const res = await findDspSlotInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        // 编辑模式：将后端返回的 company_id 和 product_id 转换为级联选择器格式
        if (res.data.dsp_company_id && res.data.dsp_product_id) {
          cascaderValue.value = [
            String(res.data.dsp_company_id),
            String(res.data.dsp_product_id)
          ]
        }
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    pay_typeOptions.value = await getDictFunc('pay_type')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false

         // 将级联选择器的值转换为后端需要的格式
         if (cascaderValue.value && cascaderValue.value.length === 2) {
           formData.value.dsp_company_id = Number(cascaderValue.value[0])
           formData.value.dsp_product_id = Number(cascaderValue.value[1])
         }

            let res
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
 
