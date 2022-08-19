<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="填空题描述:" prop="describe">
          <el-input v-model="formData.describe" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否要求有序:" prop="is_order">
          <el-input v-model.number="formData.is_order" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="可填项:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
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
  name: 'QuestionBankSupplyBlank'
}
</script>

<script setup>
import {
  createQuestionBankSupplyBlank,
  updateQuestionBankSupplyBlank,
  findQuestionBankSupplyBlank
} from '@/api/questionBank_supply_blank'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            describe: '',
            is_order: 0,
            num: 0,
        })
// 验证规则
const rule = reactive({
               describe : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               is_order : [{
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
      const res = await findQuestionBankSupplyBlank({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.requestionBank_supply_blank
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
               res = await createQuestionBankSupplyBlank(formData.value)
               break
             case 'update':
               res = await updateQuestionBankSupplyBlank(formData.value)
               break
             default:
               res = await createQuestionBankSupplyBlank(formData.value)
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
