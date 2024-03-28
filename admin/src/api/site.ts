import request from '@/utils/request'

// 站点列表
export function site_list(params?: Record<string, any>) {
    return request.get({ url: '/site/list', params })
}
// 站点列表-所有
export function site_list_all(params?: Record<string, any>) {
    return request.get({ url: '/site/listAll', params })
}

// 站点详情
export function site_detail(params: Record<string, any>) {
    return request.get({ url: '/site/detail', params })
}

// 站点新增
export function site_add(params: Record<string, any>) {
    return request.post({ url: '/site/add', params })
}

// 站点编辑
export function site_edit(params: Record<string, any>) {
    return request.post({ url: '/site/edit', params })
}

// 站点删除
export function site_delete(params: Record<string, any>) {
    return request.post({ url: '/site/del', params })
}
