
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
  createDspLaunch,
  updateDspLaunch,
  findDspLaunch
} from '@/api/dsp/dspLaunch'

defineOptions({
    name: 'DspLaunchForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 数组控制组件
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDspLaunch({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    launch_strategyOptions.value = await getDictFunc('launch_strategy')
    direction_typeOptions.value = await getDictFunc('direction_type')
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
