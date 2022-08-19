<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="学生id:" prop="studentId">
          <el-input v-model.number="formData.studentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课程Id:" prop="courseId">
          <el-input v-model.number="formData.courseId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课程名称:" prop="courseName">
          <el-input v-model="formData.courseName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="教学班名称:" prop="teachClassName">
          <el-input v-model="formData.teachClassName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="教学班id:" prop="teachClassId">
          <el-input v-model.number="formData.teachClassId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="考勤占比:" prop="attendanceProportion">
          <el-input-number v-model="formData.attendanceProportion" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="考勤得分:" prop="attendanceScore">
          <el-input-number v-model="formData.attendanceScore" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="学习资源占比:" prop="learnResourcesProportion">
          <el-input-number v-model="formData.learnResourcesProportion" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="学习资源得分:" prop="learnResourcesScore">
          <el-input-number v-model="formData.learnResourcesScore" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="过程化考核得分:" prop="procedureScore">
          <el-input-number v-model="formData.procedureScore" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="过程化考核占比:" prop="procedureProportion">
          <el-input-number v-model="formData.procedureProportion" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="期末考试成绩:" prop="examScrore">
          <el-input-number v-model="formData.examScrore" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="期末考试占比:" prop="examProporation">
          <el-input-number v-model="formData.examProporation" :precision="2" :clearable="true"></el-input-number>
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
  name: 'Score'
}
</script>

<script setup>
import {
  createScore,
  updateScore,
  findScore
} from '@/api/tea_score'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            studentId: 0,
            courseId: 0,
            courseName: '',
            teachClassName: '',
            teachClassId: 0,
            attendanceProportion: 0,
            attendanceScore: 0,
            learnResourcesProportion: 0,
            learnResourcesScore: 0,
            procedureScore: 0,
            procedureProportion: 0,
            examScrore: 0,
            examProporation: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findScore({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rescore
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createScore(formData.value)
               break
             case 'update':
               res = await updateScore(formData.value)
               break
             default:
               res = await createScore(formData.value)
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
