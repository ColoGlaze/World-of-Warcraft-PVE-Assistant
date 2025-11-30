package service

import (
	"WowServer/model"
	"WowServer/repository"
)

// PlayerViolationService 服务层，处理业务逻辑
type PlayerViolationService struct {
	Repo *repository.PlayerViolationRepository
}

// NewPlayerViolationService 创建服务实例
func NewPlayerViolationService(repo *repository.PlayerViolationRepository) *PlayerViolationService {
	return &PlayerViolationService{Repo: repo}
}

// GetViolations 获取玩家违规记录
func (s *PlayerViolationService) GetViolations(playerName, serverName string) ([]model.PlayerViolation, error) {
	return s.Repo.GetPlayerViolations(playerName, serverName)
}

// GetServerNameProportions 获取服务器名称占比统计
func (s *PlayerViolationService) GetServerNameProportions(limit int) ([]map[string]interface{}, error) {
	return s.Repo.GetServerNameProportions(limit)
}
