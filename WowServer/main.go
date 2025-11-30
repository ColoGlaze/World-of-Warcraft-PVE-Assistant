package main

import (
	"WowServer/repository"
	"WowServer/service"
	"WowServer/utils"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// ApiResponse 统一的 API 响应结构
type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 创建统一的成功响应
func successResponse(data interface{}, message string) ApiResponse {
	return ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// 创建统一的失败响应
func errorResponse(message string) ApiResponse {
	return ApiResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
}

// 初始化数据库连接
func initDB() {
	dsn := "ranking:ranking@tcp(localhost:3306)/ranking?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
}

// getViolationsHandler 处理违规记录查询请求
func getViolationsHandler(c *gin.Context, service *service.PlayerViolationService) {
	playerName := c.Query("player_name")
	serverName := c.Query("server_name")

	// 处理玩家名称
	processedName := utils.ProcessPlayerName(playerName)

	// 获取违规记录
	results, err := service.GetViolations(processedName, serverName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("查询失败"))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, successResponse(results, "查询成功"))
}

// getServerStatusHandler 处理服务器状态查询请求
func getServerStatusHandler(c *gin.Context) {
	// 获取查询参数
	serverName := c.Query("name")
	sortOrder := c.DefaultQuery("sort", "desc") // 默认降序

	// 调用Blizzard API获取服务器状态
	url := "https://webapi.blizzard.cn/wow-armory-server/api/server_status?server_type=wow_mainline"
	response, err := utils.HTTPGet(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("获取服务器状态失败: "+err.Error()))
		return
	}

	// 提取data部分
	data, exists := response["data"]
	if !exists {
		c.JSON(http.StatusInternalServerError, errorResponse("响应数据格式错误"))
		return
	}

	// 转换为map以便访问List字段
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, errorResponse("数据格式转换失败"))
		return
	}

	// 获取服务器列表
	list, listExists := dataMap["List"]
	if !listExists {
		c.JSON(http.StatusInternalServerError, errorResponse("服务器列表不存在"))
		return
	}

	// 将列表转换为[]interface{}
	serverList, ok := list.([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, errorResponse("服务器列表格式错误"))
		return
	}

	// 过滤和处理服务器数据
	var filteredAndSorted []interface{}
	for _, server := range serverList {
		serverMap, ok := server.(map[string]interface{})
		if !ok {
			continue
		}

		// 按服务器名称过滤
		if serverName != "" {
			name, nameExists := serverMap["name"].(string)
			if !nameExists || !strings.Contains(name, serverName) {
				continue
			}
		}

		filteredAndSorted = append(filteredAndSorted, serverMap)
	}

	// 按population_type排序
	// 定义population_type的优先级映射
	populationPriority := map[string]int{
		"FULL":   4,
		"HIGH":   3,
		"MEDIUM": 2,
		"LOW":    1,
	}

	// 排序函数
	sort.Slice(filteredAndSorted, func(i, j int) bool {
		serverI := filteredAndSorted[i].(map[string]interface{})
		serverJ := filteredAndSorted[j].(map[string]interface{})

		// 获取population_type
		popTypeI, popTypeIExists := serverI["population_type"].(string)
		popTypeJ, popTypeJExists := serverJ["population_type"].(string)

		// 处理不存在的情况
		if !popTypeIExists && !popTypeJExists {
			return false
		} else if !popTypeIExists {
			return sortOrder == "asc"
		} else if !popTypeJExists {
			return sortOrder == "desc"
		}

		// 获取优先级
		priorityI := populationPriority[popTypeI]
		priorityJ := populationPriority[popTypeJ]

		// 根据排序顺序返回结果
		if sortOrder == "asc" {
			return priorityI < priorityJ
		} else {
			return priorityI > priorityJ
		}
	})

	// 按category分组
	groupedResult := make(map[string][]interface{})
	// 预定义分类顺序
	categories := []string{"一区", "三区", "五区", "十区", "推荐服务器"}

	// 初始化所有分类
	for _, category := range categories {
		groupedResult[category] = []interface{}{}
	}

	// 其他分类
	groupedResult["其他"] = []interface{}{}

	// 将服务器分配到相应的分类
	for _, server := range filteredAndSorted {
		serverMap := server.(map[string]interface{})
		category, categoryExists := serverMap["category"].(string)
		if !categoryExists {
			groupedResult["其他"] = append(groupedResult["其他"], serverMap)
			continue
		}

		// 检查是否是预定义分类
		found := false
		for _, predefined := range categories {
			if category == predefined {
				groupedResult[category] = append(groupedResult[category], serverMap)
				found = true
				break
			}
		}

		// 如果不是预定义分类，放到"其他"类别
		if !found {
			groupedResult["其他"] = append(groupedResult["其他"], serverMap)
		}
	}

	// 返回分组后的结果
	c.JSON(http.StatusOK, successResponse(groupedResult, "获取服务器状态成功"))
}

// getServerNameProportionsHandler 处理服务器名称占比统计请求
func getServerNameProportionsHandler(c *gin.Context, service *service.PlayerViolationService) {
	// 从请求参数中获取查询个数，默认为0（不限制）
	limitStr := c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse("无效的查询个数参数"))
		return
	}

	// 获取服务器名称占比统计
	results, err := service.GetServerNameProportions(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("查询服务器占比失败: "+err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, successResponse(results, "获取服务器占比成功"))
}

// getMythicPlusAffixesHandler 处理大米词缀信息查询请求
func getMythicPlusAffixesHandler(c *gin.Context) {
	// Raider.io词缀API URL
	url := "https://raider.io/api/v1/mythic-plus/affixes?region=cn&locale=cn"

	// 调用HTTPGet获取数据
	response, err := utils.HTTPGet(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("获取词缀信息失败: "+err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, successResponse(response, "获取词缀信息成功"))
}

// getSearchHandler 处理搜索请求
func getSearchHandler(c *gin.Context, systemConfigService *service.SystemConfigService) {
	// 获取前端传来的keyword参数
	keyword := c.Query("keyword")

	// 构建Blizzard搜索API的URL
	url := "https://webapi.blizzard.cn/wow-armory-server/api/search?keyword=" + keyword

	// 从数据库获取cookie值
	cookieValue, err := systemConfigService.GetConfigByKey("blizzard_api_cookie")
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("获取cookie配置失败: "+err.Error()))
		return
	}

	// 设置cookie
	headers := map[string]string{
		"Cookie": cookieValue,
	}

	// 调用HTTPGetWithHeaders获取数据
	response, err := utils.HTTPGetWithHeaders(url, headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse("获取搜索结果失败: "+err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, successResponse(response, "获取搜索结果成功"))
}

func main() {
	initDB()

	// 初始化仓库和服务
	playerViolationRepo := repository.NewPlayerViolationRepository(db)
	playerViolationService := service.NewPlayerViolationService(playerViolationRepo)

	// 初始化系统配置服务
	systemConfigRepo := repository.NewSystemConfigRepository(db)
	systemConfigService := service.NewSystemConfigService(systemConfigRepo)

	// 设置 Gin 为 Release 模式
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 创建 API 路由组
	api := r.Group("/player")
	{
		// 违规记录相关接口
		violations := api.Group("/violations")
		{
			// 查询违规记录
			violations.GET("/info", func(c *gin.Context) {
				getViolationsHandler(c, playerViolationService)
			})
			// 查询服务器名称占比统计
			violations.GET("/server-proportions", func(c *gin.Context) {
				getServerNameProportionsHandler(c, playerViolationService)
			})
		}
	}

	// 服务器状态相关接口
	server := r.Group("/server")
	{
		// 查询服务器状态
		server.GET("/status", getServerStatusHandler)
	}

	// 大米词缀相关接口
	mythic := r.Group("/mythic")
	{
		// 查询大米词缀信息
		mythic.GET("/affixes", getMythicPlusAffixesHandler)
	}

	// 添加搜索相关接口
	search := r.Group("/search")
	{
		// 搜索功能
		search.GET("/playerList", func(c *gin.Context) {
			getSearchHandler(c, systemConfigService)
		})
	}

	// 启动服务，监听 28144 端口
	err := r.Run(":28144")
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
