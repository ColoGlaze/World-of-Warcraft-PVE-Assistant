package repository

import (
	"WowServer/model"
	"fmt"

	"gorm.io/gorm"
)

// PlayerViolationRepository 数据库操作类
type PlayerViolationRepository struct {
	DB *gorm.DB
}

// NewPlayerViolationRepository 创建新实例
func NewPlayerViolationRepository(db *gorm.DB) *PlayerViolationRepository {
	return &PlayerViolationRepository{DB: db}
}

// GetPlayerViolations 查询玩家违规记录
func (repo *PlayerViolationRepository) GetPlayerViolations(playerName, serverName string) ([]model.PlayerViolation, error) {
	var results []model.PlayerViolation

	query := repo.DB.Model(&model.PlayerViolation{})

	if playerName != "" {
		query = query.Where("player_name LIKE ?", playerName)
	}
	if serverName != "" {
		query = query.Where("server_name = ?", serverName)
	}

	err := query.Order("violation_time DESC").Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	return results, nil
}

// GetServerNameProportions 查询每个服务器的违规记录占比
func (repo *PlayerViolationRepository) GetServerNameProportions(limit int) ([]map[string]interface{}, error) {
	// 统计每个服务器的记录数量
	var serverStats []struct {
		ServerName string `json:"server_name"`
		Count      int64  `json:"count"`
	}

	// 查询每个服务器的记录数
	query := repo.DB.Model(&model.PlayerViolation{}).
		Select("server_name, COUNT(*) as count").
		Group("server_name").
		Order("count DESC")

	// 如果设置了限制数量，则应用限制
	if limit > 0 {
		query = query.Limit(limit)
	}

	// 执行查询
	err := query.Scan(&serverStats).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query server stats: %w", err)
	}

	// 计算总记录数
	var totalCount int64
	err = repo.DB.Model(&model.PlayerViolation{}).Count(&totalCount).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query total count: %w", err)
	}

	// 计算占比并构建结果
	results := make([]map[string]interface{}, len(serverStats))
	for i, stat := range serverStats {
		proportion := float64(stat.Count) / float64(totalCount) * 100
		results[i] = map[string]interface{}{
			"server_name": stat.ServerName,
			"count":       stat.Count,
			"proportion":  proportion,
		}
	}

	return results, nil
}
