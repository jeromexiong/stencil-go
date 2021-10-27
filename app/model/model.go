package model

// Admin [...]
type Admin struct {
	ID          uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"id"`                 // 管理员id
	Nickname    string `gorm:"column:nickname;type:varchar(32);not null;default:''" json:"nickname"`          // 姓名
	Username    string `gorm:"unique;column:username;type:varchar(64);not null;default:''" json:"username"`   // 用户名
	Telephone   string `gorm:"unique;column:telephone;type:varchar(20);not null;default:''" json:"telephone"` // 联系方式
	Pwd         string `gorm:"column:pwd;type:varchar(32);not null;default:''" json:"pwd"`                    // 密码
	Salt        string `gorm:"column:salt;type:varchar(32);not null;default:''" json:"salt"`                  // 密码加盐
	LastIP      string `gorm:"column:last_ip;type:varchar(64);not null;default:''" json:"last_ip"`            // 最后登录id
	CreatedTime int64  `gorm:"column:created_time;type:bigint(20);not null;default:0" json:"created_time"`
	UpdatedTime int64  `gorm:"column:updated_time;type:bigint(20);not null;default:0" json:"updated_time"`
}

// CasbinRule [...]
type CasbinRule struct {
	ID    int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	Ptype string `gorm:"column:ptype;type:varchar(255)" json:"ptype"`
	V0    string `gorm:"column:v0;type:varchar(255)" json:"v0"`
	V1    string `gorm:"column:v1;type:varchar(255)" json:"v1"`
	V2    string `gorm:"column:v2;type:varchar(255)" json:"v2"`
	V3    string `gorm:"column:v3;type:varchar(255)" json:"v3"`
	V4    string `gorm:"column:v4;type:varchar(255)" json:"v4"`
	V5    string `gorm:"column:v5;type:varchar(255)" json:"v5"`
	V6    string `gorm:"column:v6;type:varchar(255)" json:"v6"`
}

// File [...]
type File struct {
	ID          uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"id"`               // 文件ID
	Name        string `gorm:"column:name;type:varchar(200);not null;default:''" json:"name"`               // 原始文件名
	Savename    string `gorm:"column:savename;type:char(100);not null;default:''" json:"savename"`          // 保存名称
	Savepath    string `gorm:"column:savepath;type:varchar(100);not null;default:''" json:"savepath"`       // 文件保存路径
	Savepathp   string `gorm:"column:savepathp;type:varchar(100);not null;default:''" json:"savepathp"`     // 转码后路径
	Ext         string `gorm:"column:ext;type:char(10);not null;default:''" json:"ext"`                     // 文件后缀
	Mime        string `gorm:"column:mime;type:char(200);not null;default:''" json:"mime"`                  // 文件mime类型
	Size        uint32 `gorm:"column:size;type:int(10) unsigned;not null;default:0" json:"size"`            // 文件大小 单位 B
	Width       uint32 `gorm:"column:width;type:int(10) unsigned;not null;default:0" json:"width"`          // 尺寸 宽度 图片/视频
	Height      uint32 `gorm:"column:height;type:int(10) unsigned;not null;default:0" json:"height"`        // 尺寸 高度 图片/视频
	Duration    uint32 `gorm:"column:duration;type:int(10) unsigned;not null;default:0" json:"duration"`    // 时长 音频/视频
	Md5         string `gorm:"column:md5;type:char(32);not null;default:''" json:"md5"`                     // 文件md5
	Category    string `gorm:"column:category;type:char(32);not null;default:''" json:"category"`           // 所属分类
	Location    uint8  `gorm:"column:location;type:tinyint(3) unsigned;not null;default:0" json:"location"` // 文件保存位置 0 阿里云
	URL         string `gorm:"column:url;type:varchar(255);not null;default:''" json:"url"`                 // 远程地址
	CreatedTime int64  `gorm:"column:created_time;type:bigint(20);not null;default:0" json:"created_time"`
	UpdatedTime int64  `gorm:"column:updated_time;type:bigint(20);not null;default:0" json:"updated_time"`
}
