package service

import (
	"WowServer/repository"
)

// SystemConfigService 系统配置服务

type SystemConfigService struct {
	Repo *repository.SystemConfigRepository
}

// NewSystemConfigService 创建系统配置服务实例
func NewSystemConfigService(repo *repository.SystemConfigRepository) *SystemConfigService {
	return &SystemConfigService{Repo: repo}
}

// GetConfigByKey 根据配置键获取配置值
func (s *SystemConfigService) GetConfigByKey(key string) (string, error) {
	return s.Repo.GetConfigByKey(key)
}
