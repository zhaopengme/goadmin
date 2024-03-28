package model

//Site 站点实体
type Site struct {
    
        Id int `gorm:"primarykey;comment:'主键ID'"` // 主键ID
    
        Name string `gorm:"comment:'名称'"` // 名称
    
        Url string `gorm:"comment:'站点url'"` // 站点url
    
        IsDelete int `gorm:"comment:'是否删除: [0=否, 1=是]'"` // 是否删除: [0=否, 1=是]
    
        CreateTime int64 `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
    
    
        UpdateTime int64 `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
    
    
        DeleteTime int64 `gorm:"comment:'删除时间'"` // 删除时间
    
}
