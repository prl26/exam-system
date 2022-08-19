<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="学生id">
          <el-input v-model="searchInfo.studentId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="课程Id">
          <el-input v-model="searchInfo.courseId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="课程名称">
          <el-input v-model="searchInfo.courseName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="教学班名称">
          <el-input v-model="searchInfo.teachClassName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="教学班id">
          <el-input v-model="searchInfo.teachClassId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="学生id" prop="studentId" width="120" />
        <el-table-column align="left" label="课程Id" prop="courseId" width="120" />
        <el-table-column align="left" label="课程名称" prop="courseName" width="120" />
        <el-table-column align="left" label="教学班名称" prop="teachClassName" width="120" />
        <el-table-column align="left" label="教学班id" prop="teachClassId" width="120" />
        <el-table-column align="left" label="考勤占比" prop="attendanceProportion" width="120" />
        <el-table-column align="left" label="考勤得分" prop="attendanceScore" width="120" />
        <el-table-column align="left" label="学习资源占比" prop="learnResourcesProportion" width="120" />
        <el-table-column align="left" label="学习资源得分" prop="learnResourcesScore" width="120" />
        <el-table-column align="left" label="过程化考核得分" prop="procedureScore" width="120" />
        <el-table-column align="left" label="过程化考核占比" prop="procedureProportion" width="120" />
        <el-table-column align="left" label="期末考试成绩" prop="examScrore" width="120" />
        <el-table-column align="left" label="期末考试占比" prop="examProporation" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateScoreFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="学生id:"  prop="studentId" >
          <el-input v-model.number="formData.studentId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课程Id:"  prop="courseId" >
          <el-input v-model.number="formData.courseId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="课程名称:"  prop="courseName" >
          <el-input v-model="formData.courseName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="教学班名称:"  prop="teachClassName" >
          <el-input v-model="formData.teachClassName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="教学班id:"  prop="teachClassId" >
          <el-input v-model.number="formData.teachClassId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="考勤占比:"  prop="attendanceProportion" >
          <el-input-number v-model="formData.attendanceProportion"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="考勤得分:"  prop="attendanceScore" >
          <el-input-number v-model="formData.attendanceScore"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="学习资源占比:"  prop="learnResourcesProportion" >
          <el-input-number v-model="formData.learnResourcesProportion"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="学习资源得分:"  prop="learnResourcesScore" >
          <el-input-number v-model="formData.learnResourcesScore"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="过程化考核得分:"  prop="procedureScore" >
          <el-input-number v-model="formData.procedureScore"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="过程化考核占比:"  prop="procedureProportion" >
          <el-input-number v-model="formData.procedureProportion"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="期末考试成绩:"  prop="examScrore" >
          <el-input-number v-model="formData.examScrore"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="期末考试占比:"  prop="examProporation" >
          <el-input-number v-model="formData.examProporation"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteScore,
  deleteScoreByIds,
  updateScore,
  findScore,
  getScoreList
} from '@/api/tea_score'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getScoreList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteScoreFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteScoreByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateScoreFunc = async(row) => {
    const res = await findScore({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rescore
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteScoreFunc = async (row) => {
    const res = await deleteScore({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
