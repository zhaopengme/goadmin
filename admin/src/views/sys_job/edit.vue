<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
                        <el-form-item label="调度名称" prop="name">
                            <el-input v-model="formData.name" placeholder="请输入调度名称" />
                        </el-form-item>
                        <el-form-item label="调用目标字符串" prop="invokeTarget">
                            <el-input
                                v-model="formData.invokeTarget"
                                placeholder="请输入调用目标字符串"
                                type="textarea"
                                :autosize="{ minRows: 4, maxRows: 6 }"
                            />
                        </el-form-item>
                        <el-form-item label="cron执行表达式" prop="cronExpression">
                            <el-input v-model="formData.cronExpression" placeholder="请输入cron执行表达式" />
                        </el-form-item>
                        <el-form-item label="是否并发执行（0允许 1禁止）" prop="concurrent">
                            <el-input v-model="formData.concurrent" placeholder="请输入是否并发执行（0允许 1禁止）" />
                        </el-form-item>
                        <el-form-item label="状态（0正常 1暂停）" prop="status">
                            <el-radio-group v-model="formData.status" placeholder="请选择状态（0正常 1暂停）">
                                <el-radio
                                    v-for="(item, index) in dictData.status"
                                    :key="index"
                                    :label="item.value"
                                    :disabled="!item.status"
                                >
                                    {{ item.name }}
                                </el-radio>
                            </el-radio-group>
                        </el-form-item>
                        <el-form-item label="备注信息" prop="remark">
                            <el-input
                                v-model="formData.remark"
                                placeholder="请输入备注信息"
                                type="textarea"
                                :autosize="{ minRows: 4, maxRows: 6 }"
                            />
                        </el-form-item>
            
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {  sys_job_edit, sys_job_add, sys_job_detail } from '@/api/sys_job'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import type { PropType } from 'vue'
defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    }
})
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑定时任务调度' : '新增定时任务调度'
})

const formData = reactive({
    id: '',
    name: '',
    invokeTarget: '',
    cronExpression: '',
    concurrent: '',
    status: '',
    remark: '',
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入主键ID',
            trigger: ['blur']
        }
    ],
    name: [
        {
            required: true,
            message: '请输入调度名称',
            trigger: ['blur']
        }
    ],
    invokeTarget: [
        {
            required: true,
            message: '请输入调用目标字符串',
            trigger: ['blur']
        }
    ],
    cronExpression: [
        {
            required: true,
            message: '请输入cron执行表达式',
            trigger: ['blur']
        }
    ],
    concurrent: [
        {
            required: true,
            message: '请输入是否并发执行（0允许 1禁止）',
            trigger: ['blur']
        }
    ],
    status: [
        {
            required: true,
            message: '请选择状态（0正常 1暂停）',
            trigger: ['blur']
        }
    ],
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    mode.value == 'edit' ? await sys_job_edit(data) : await sys_job_add(data)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
        }
    }
}

const getDetail = async (row: Record<string, any>) => {
    const data = await sys_job_detail({
        id: row.id
    })
    setFormData(data)
}

const handleClose = () => {
    emit('close')
}

defineExpose({
    open,
    setFormData,
    getDetail
})
</script>
