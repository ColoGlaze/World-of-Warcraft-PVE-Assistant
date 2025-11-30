package model

import "time"

// SystemConfig 系统配置表模型
type SystemConfig struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ConfigKey   string    `json:"config_key" gorm:"type:varchar(100);not null;uniqueIndex:uk_config_key"`
	ConfigValue string    `json:"config_value" gorm:"type:text;not null"`
	Description string    `json:"description" gorm:"type:varchar(255);null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName 设置表名
func (SystemConfig) TableName() string {
	return "system_config"
}
