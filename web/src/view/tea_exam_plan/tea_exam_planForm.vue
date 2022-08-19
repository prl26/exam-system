<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="考试名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="教学班id:" prop="teachClassId">
          <el-input v-model.number="formData.teachClassId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="考试时长:" prop="time">
          <el-date-picker v-model="formData.time" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="考试时间:" prop="startTime">
          <el-date-picker v-model="formData.startTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="结束时间:" prop="endTime">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="课程Id:" prop="courseId">
          <el-input v-model.number="formData.courseId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="考试模板Id:" prop="templateId">
          <el-input v-model.number="formData.templateId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="state">
          <el-select v-model="formData.state" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="是否审核:" prop="audit">
          <el-select v-model="formData.audit" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in auditOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="考试类型:" prop="type">
          <el-select v-model="formData.type" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in ExamTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="通过分数:" prop="passScore">
          <el-input-number v-model="formData.passScore" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="权重:" prop="weight">
          <el-input-number v-model="formData.weight" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ExamPlan'
}
</script>

<script setup>
import {
  createExamPlan,
  updateExamPlan,
  findExamPlan
} from '@/api/tea_exam_plan'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const ExamTypeOptions = ref([])
const statusOptions = ref([])
const auditOptions = ref([])
const formData = ref({
            name: '',
            teachClassId: 0,
            time: new Date(),
            startTime: new Date(),
            endTime: new Date(),
            courseId: 0,
            templateId: 0,
            state: undefined,
            audit: undefined,
            type: undefined,
            passScore: 0,
            weight: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findExamPlan({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reexamPlan
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ExamTypeOptions.value = await getDictFunc('ExamType')
    statusOptions.value = await getDictFunc('status')
    auditOptions.value = await getDictFunc('audit')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createExamPlan(formData.value)
               break
             case 'update':
               res = await updateExamPlan(formData.value)
               break
             default:
               res = await createExamPlan(formData.value)
               break
           }
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
