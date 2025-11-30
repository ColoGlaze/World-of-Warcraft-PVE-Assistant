package model

import "time"

// PlayerViolation 数据表结构映射
type PlayerViolation struct {
	ID              int64     `json:"id"`
	PlayerName      string    `json:"player_name"`
	ServerName      string    `json:"server_name"`
	ViolationReason string    `json:"violation_reason"`
	ViolationCount  string    `json:"violation_count"`
	ViolationTime   time.Time `json:"violation_time"`
}
