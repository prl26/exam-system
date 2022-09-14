<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="课程资源id:" prop="learn_resources_id">
          <el-input v-model.number="formData.learn_resources_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="章节id:" prop="chapter_id">
          <el-input v-model.number="formData.chapter_id" :clearable="true" placeholder="请输入" />
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
  name: 'Learn_resources_chapter_merge'
}
</script>

<script setup>
import {
  createLearn_resources_chapter_merge,
  updateLearn_resources_chapter_merge,
  findLearn_resources_chapter_merge
} from '@/api/bas_learn_resources_chapter_merge'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            learn_resources_id: 0,
            chapter_id: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findLearn_resources_chapter_merge({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.relearnResourcesChapterMerge
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
               res = await createLearn_resources_chapter_merge(formData.value)
               break
             case 'update':t
               res = await updateLearn_resources_chapter_merge(formData.value)
               break
             default:
               res = await createLearn_resources_chapter_merge(formData.value)
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
