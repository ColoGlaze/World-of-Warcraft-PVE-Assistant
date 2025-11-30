package repository

import (
	"WowServer/model"
	"fmt"

	"gorm.io/gorm"
)

// SystemConfigRepository 系统配置仓库

type SystemConfigRepository struct {
	DB *gorm.DB
}

// NewSystemConfigRepository 创建系统配置仓库实例
func NewSystemConfigRepository(db *gorm.DB) *SystemConfigRepository {
	return &SystemConfigRepository{DB: db}
}

// GetConfigByKey 根据配置键获取配置值
func (repo *SystemConfigRepository) GetConfigByKey(key string) (string, error) {
	var config model.SystemConfig
	err := repo.DB.Where("config_key = ?", key).First(&config).Error
	if err != nil {
		return "", fmt.Errorf("failed to get config by key %s: %w", key, err)
	}
	return config.ConfigValue, nil
}
