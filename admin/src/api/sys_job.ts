import request from '@/utils/request'

// 定时任务调度列表
export function sys_job_list(params?: Record<string, any>) {
    return request.get({ url: '/sys_job/list', params })
}
// 定时任务调度列表-所有
export function sys_job_list_all(params?: Record<string, any>) {
    return request.get({ url: '/sys_job/listAll', params })
}

// 定时任务调度详情
export function sys_job_detail(params: Record<string, any>) {
    return request.get({ url: '/sys_job/detail', params })
}

// 定时任务调度新增
export function sys_job_add(params: Record<string, any>) {
    return request.post({ url: '/sys_job/add', params })
}

// 定时任务调度编辑
export function sys_job_edit(params: Record<string, any>) {
    return request.post({ url: '/sys_job/edit', params })
}

// 定时任务调度删除
export function sys_job_delete(params: Record<string, any>) {
    return request.post({ url: '/sys_job/del', params })
}
