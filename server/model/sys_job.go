package model

//SysJob 定时任务调度实体
type SysJob struct {
    
        Id int `gorm:"primarykey;comment:'主键ID'"` // 主键ID
    
        Name string `gorm:"comment:'调度名称'"` // 调度名称
    
        InvokeTarget string `gorm:"comment:'调用目标字符串'"` // 调用目标字符串
    
        CronExpression string `gorm:"comment:'cron执行表达式'"` // cron执行表达式
    
        Concurrent string `gorm:"comment:'是否并发执行（0允许 1禁止）'"` // 是否并发执行（0允许 1禁止）
    
        Status string `gorm:"comment:'状态（0正常 1暂停）'"` // 状态（0正常 1暂停）
    
        IsDelete int `gorm:"comment:'是否删除: 0=否, 1=是'"` // 是否删除: 0=否, 1=是
    
        CreateTime int64 `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
    
    
        CreateBy string `gorm:"comment:'创建者'"` // 创建者
    
        UpdateTime int64 `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
    
    
        UpdateBy string `gorm:"comment:'更新者'"` // 更新者
    
        DeleteTime int64 `gorm:"comment:'删除时间'"` // 删除时间
    
    
        Remark string `gorm:"comment:'备注信息'"` // 备注信息
}
