<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="调度名称" prop="name">
                    <el-input  v-model="queryParams.name" />
                </el-form-item>
                <el-form-item label="状态（0正常 1暂停）" prop="status">
                    <el-select
                        v-model="queryParams.status"
                        clearable
                    >
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.status"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button v-perms="['admin:sys_job:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <el-table
                class="mt-4"
                size="large"
                v-loading="pager.loading"
                :data="pager.lists"
            >
                <el-table-column label="调度名称" prop="name" min-width="130" />
                <el-table-column label="调用目标字符串" prop="invokeTarget" min-width="130" />
                <el-table-column label="cron执行表达式" prop="cronExpression" min-width="130" />
                <el-table-column label="是否并发执行（0允许 1禁止）" prop="concurrent" min-width="130" />
                <el-table-column label="状态（0正常 1暂停）" prop="status" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.status" :value="row.status" />
                    </template>
                </el-table-column>
                <el-table-column label="是否删除: 0=否, 1=是" prop="isDelete" min-width="130" />
                <el-table-column label="创建时间" prop="createTime" min-width="130" />
                <el-table-column label="创建者" prop="createBy" min-width="130" />
                <el-table-column label="更新时间" prop="updateTime" min-width="130" />
                <el-table-column label="更新者" prop="updateBy" min-width="130" />
                <el-table-column label="备注信息" prop="remark" min-width="130" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:sys_job:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:sys_job:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
                        >
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <edit-popup
            v-if="showEdit"
            ref="editRef"
            :dict-data="dictData"
            @success="getLists"
            @close="showEdit = false"
        />
    </div>
</template>
<script lang="ts" setup>
import { sys_job_delete, sys_job_list } from '@/api/sys_job'
import { useDictData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name:"sys_job"
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    name: '',
    status: '',
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: sys_job_list,
    params: queryParams
})
const { dictData } = useDictData<{
    status: any[]
}>(['status'])


const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await sys_job_delete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

getLists()
</script>
