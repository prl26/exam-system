<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="课程id:" prop="course_id">
          <el-input v-model.number="formData.course_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学期id:" prop="term_id">
          <el-input v-model.number="formData.term_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否还有其他专业的学生:" prop="has_more">
          <el-select v-model="formData.has_more" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in has_moreOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属班级id:" prop="belong_class_id">
          <el-input v-model.number="formData.belong_class_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="教学班名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="所属老师id:" prop="teacher_id">
          <el-input v-model.number="formData.teacher_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="考勤得分占比:" prop="attendance_proportion">
          <el-input-number v-model="formData.attendance_proportion" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="学习资源得分占比:" prop="learn_resource_proportion">
          <el-input-number v-model="formData.learn_resource_proportion" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="期末考得分占比:" prop="final_exam_proportion">
          <el-input-number v-model="formData.final_exam_proportion" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="过程化考核得分:" prop="procedure_exam_proportion">
          <el-input-number v-model="formData.procedure_exam_proportion" :precision="2" :clearable="true"></el-input-number>
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
  name: 'TeachClass'
}
</script>

<script setup>
import {
  createTeachClass,
  updateTeachClass,
  findTeachClass
} from '@/api/teachClass'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const has_moreOptions = ref([])
const formData = ref({
            course_id: 0,
            term_id: 0,
            has_more: undefined,
            belong_class_id: 0,
            name: '',
            teacher_id: 0,
            attendance_proportion: 0,
            learn_resource_proportion: 0,
            final_exam_proportion: 0,
            procedure_exam_proportion: 0,
        })
// 验证规则
const rule = reactive({
               course_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               term_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               belong_class_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               attendance_proportion : [{
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
      const res = await findTeachClass({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reteachClass
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    has_moreOptions.value = await getDictFunc('has_more')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createTeachClass(formData.value)
               break
             case 'update':
               res = await updateTeachClass(formData.value)
               break
             default:
               res = await createTeachClass(formData.value)
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
