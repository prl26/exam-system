<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="试卷id:" prop="paperId">
          <el-input v-model.number="formData.paperId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="题目id:" prop="questionId">
          <el-input v-model.number="formData.questionId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所占分值:" prop="score">
          <el-input v-model.number="formData.score" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="题目类型:" prop="questionType">
          <el-select v-model="formData.questionType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in ProblemTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  name: 'PaperQuestionMerge'
}
</script>

<script setup>
import {
  createPaperQuestionMerge,
  updatePaperQuestionMerge,
  findPaperQuestionMerge
} from '@/api/exam_paper_question_merge'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const ProblemTypeOptions = ref([])
const formData = ref({
            paperId: 0,
            questionId: 0,
            score: 0,
            questionType: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPaperQuestionMerge({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repaperQuestionMerge
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ProblemTypeOptions.value = await getDictFunc('ProblemType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPaperQuestionMerge(formData.value)
               break
             case 'update':
               res = await updatePaperQuestionMerge(formData.value)
               break
             default:
               res = await createPaperQuestionMerge(formData.value)
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
