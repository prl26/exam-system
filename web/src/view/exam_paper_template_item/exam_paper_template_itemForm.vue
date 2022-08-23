<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="章节:" prop="chapter">
          <el-input v-model="formData.chapter" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="题目类型:" prop="problemType">
          <el-select v-model="formData.problemType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in ProblemTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="难度:" prop="difficulty">
          <el-select v-model="formData.difficulty" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in DifficultyOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="数量:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="分数:" prop="score">
          <el-input v-model.number="formData.score" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="试卷模板id:" prop="templateId">
          <el-input v-model.number="formData.templateId" :clearable="true" placeholder="请输入" />
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
  name: 'PaperTemplateItem'
}
</script>

<script setup>
import {
  createPaperTemplateItem,
  updatePaperTemplateItem,
  findPaperTemplateItem
} from '@/api/exam_paper_template_item'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const ProblemTypeOptions = ref([])
const DifficultyOptions = ref([])
const formData = ref({
            chapter: '',
            problemType: undefined,
            difficulty: undefined,
            num: 0,
            score: 0,
            templateId: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPaperTemplateItem({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repaperTemplateItem
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ProblemTypeOptions.value = await getDictFunc('ProblemType')
    DifficultyOptions.value = await getDictFunc('Difficulty')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPaperTemplateItem(formData.value)
               break
             case 'update':
               res = await updatePaperTemplateItem(formData.value)
               break
             default:
               res = await createPaperTemplateItem(formData.value)
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
