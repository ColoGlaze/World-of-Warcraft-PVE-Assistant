<template>
  <div class="server-status">
    <!-- 返回按钮 -->
    <div class="back-button-container">
      <button @click="goBack" class="back-button">
        <span class="back-icon">←</span>
        返回首页
      </button>
    </div>

    <!-- 头部 -->
    <header class="header">
      <div class="server-icon">🌐</div>
      <h1 class="title">服务器状态</h1>
      <p class="subtitle">实时监控魔兽世界服务器状态</p>
    </header>

    <!-- 刷新按钮 -->
    <div class="refresh-section">
      <!-- 搜索框 -->
      <div class="search-container">
        <input v-model="searchQuery" type="text" placeholder="搜索服务器名称..." class="search-input"
          @input="handleSearchInput" />
        <div class="search-icon">🔍</div>
      </div>

      <!-- 排序按钮 -->
      <div class="sort-container">
        <button @click="toggleSort" class="sort-btn">
          <span class="sort-icon">{{ getSortIcon() }}</span>
          <span class="sort-text">{{ getSortText() }}</span>
        </button>
      </div>

      <button @click="refreshStatus" class="refresh-btn" :disabled="loading">
        <span v-if="loading" class="loading-spinner"></span>
        <span v-else class="refresh-icon">🔄</span>
        {{ loading ? '刷新中...' : '刷新状态' }}
      </button>
      <span class="last-update">最后更新：{{ lastUpdateTime }}</span>
    </div>

    <!-- 统计信息 -->
    <div class="stats-section">
      <div class="stat-card">
        <div class="stat-icon">✅</div>
        <div class="stat-info">
          <div class="stat-number">{{ onlineServers }}</div>
          <div class="stat-label">在线服务器</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">⚠️</div>
        <div class="stat-info">
          <div class="stat-number">{{ maintenanceServers }}</div>
          <div class="stat-label">维护中</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">❌</div>
        <div class="stat-info">
          <div class="stat-number">{{ offlineServers }}</div>
          <div class="stat-label">离线服务器</div>
        </div>
      </div>
    </div>

    <!-- 服务器列表 -->
    <div class="server-list">
      <div v-if="loading" class="loading-state">
        <div class="loading-spinner"></div>
        <p>正在加载服务器列表...</p>
      </div>

      <div v-else-if="allServers.length === 0" class="empty-state">
        <p>暂无服务器数据</p>
      </div>

      <div v-else>
        <div v-for="server in allServers" :key="server.id" class="server-item"
          :class="getServerStatusClass(server.status_type)">
          <div class="server-main-info">
            <div class="server-status-indicator">
              <div class="status-dot" :class="server.status_type.toLowerCase()"></div>
              <span class="status-text">{{ server.status_name }}</span>
            </div>
            <div class="server-details">
              <h3 class="server-name">{{ server.name }}</h3>
              <div class="server-meta">
                <span class="server-zone">{{ server.category }}</span>
                <span class="server-type">{{ server.type_name }}</span>
              </div>
            </div>
          </div>
          <div class="server-load">
            <div class="load-info">
              <span class="load-label">负载</span>
              <span class="load-value" :class="getLoadClass(server.population_name)">
                {{ server.population_name }}
              </span>
            </div>
            <div class="load-bar">
              <div class="load-fill" :class="getLoadClass(server.population_name)"
                :style="{ width: getLoadPercentage(server.population_name) + '%' }"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { penaltyAPI } from '../api/penalty'
import { ElMessage } from 'element-plus'

// 定义组件名称
defineOptions({
  name: 'ServerStatus'
})

const router = useRouter()
const loading = ref(false)
const lastUpdateTime = ref('')
const searchQuery = ref('')
const sortOrder = ref('desc')

// 服务器列表数据
const serverList = ref([])
// 防抖计时器
let searchTimer = null

// 防抖函数
const debounce = (func, wait) => {
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(searchTimer)
      func(...args)
    }
    clearTimeout(searchTimer)
    searchTimer = setTimeout(later, wait)
  }
}

// 处理搜索输入
const handleSearchInput = debounce(() => {
  searchServers()
}, 500) // 500毫秒防抖时间

// 计算统计数据
const onlineServers = computed(() => {
  let count = 0
  Object.values(serverList.value).forEach(servers => {
    count += servers.filter(server => server.status_type === 'UP').length
  })
  return count
})

const maintenanceServers = computed(() => {
  let count = 0
  Object.values(serverList.value).forEach(servers => {
    count += servers.filter(server => server.status_type === 'MAINTENANCE').length
  })
  return count
})

const offlineServers = computed(() => {
  let count = 0
  Object.values(serverList.value).forEach(servers => {
    count += servers.filter(server => server.status_type === 'DOWN').length
  })
  return count
})

const allServers = computed(() => {
  const servers = []
  Object.values(serverList.value).forEach(categoryServers => {
    servers.push(...categoryServers)
  })
  return servers
})

// 获取服务器列表
const getServersStatusList = async (params = {}) => {
  loading.value = true
  try {
    const response = await penaltyAPI.getServerList(params)
    if (response.status === 'success' && response.data) {
      serverList.value = response.data
    }
  } catch (error) {
    console.error('获取服务器列表失败:', error)
    ElMessage.error('获取服务器列表失败，请稍后重试')
  } finally {
    loading.value = false
  }
}

// 搜索服务器
const searchServers = async () => {
  const params = {
    name: searchQuery.value.trim(),
    sort: sortOrder.value
  }
  await getServersStatusList(params)
}

// 返回首页
const goBack = () => {
  router.push('/')
}

// 刷新状态
const refreshStatus = async () => {
  const params = {
    name: searchQuery.value.trim(),
    sort: sortOrder.value
  }

  try {
    await getServersStatusList(params)
    updateLastUpdateTime()
  } catch (error) {
    console.error('刷新失败:', error)
  }
}

// 更新最后更新时间
const updateLastUpdateTime = () => {
  const now = new Date()
  lastUpdateTime.value = now.toLocaleTimeString('zh-CN')
}

// 切换排序状态
const toggleSort = async () => {
  sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  await searchServers()
}

// 获取排序图标
const getSortIcon = () => {
  return sortOrder.value === 'asc' ? '↑' : '↓'
}

// 获取排序文本
const getSortText = () => {
  return sortOrder.value === 'asc' ? '状态升序' : '状态降序'
}

// 获取服务器状态样式类
const getServerStatusClass = (status) => {
  switch (status) {
    case 'UP': return 'server-online'
    case 'MAINTENANCE': return 'server-maintenance'
    case 'DOWN': return 'server-offline'
    default: return ''
  }
}

// 获取负载样式类
const getLoadClass = (load) => {
  switch (load) {
    case '低': return 'load-low'
    case '中': return 'load-medium'
    case '高':
    case '满': return 'load-high'
    default: return 'load-low'
  }
}

// 获取负载百分比
const getLoadPercentage = (load) => {
  switch (load) {
    case '低': return 30
    case '中': return 60
    case '高': return 80
    case '满': return 95
    default: return 30
  }
}

// 组件挂载时初始化
onMounted(() => {
  updateLastUpdateTime()
  getServersStatusList()
})
</script>

<style scoped>
.server-status {
  padding: 1rem;
  max-width: 900px;
  margin: 0 auto;
  padding-bottom: 5rem;
}

.back-button-container {
  margin-bottom: 1rem;
}

.back-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(102, 126, 234, 0.2);
  border-radius: 12px;
  color: #667eea;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.back-button:hover {
  background: rgba(102, 126, 234, 0.1);
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.2);
}

.back-icon {
  font-size: 1.2rem;
  transition: transform 0.3s;
}

.back-button:hover .back-icon {
  transform: translateX(-3px);
}

.header {
  text-align: center;
  margin-bottom: 2rem;
  padding: 2rem 0;
}

.server-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
  animation: iconFloat 3s ease-in-out infinite, iconGlow 2s ease-in-out infinite alternate;
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.3));
}

@keyframes iconFloat {

  0%,
  100% {
    transform: translateY(0px) rotate(0deg);
  }

  50% {
    transform: translateY(-10px) rotate(5deg);
  }
}

@keyframes iconGlow {
  0% {
    filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 20px rgba(102, 126, 234, 0.3));
  }

  100% {
    filter: drop-shadow(0 6px 16px rgba(0, 0, 0, 0.4)) drop-shadow(0 0 30px rgba(118, 75, 162, 0.5));
  }
}

.title {
  font-size: 2.5rem;
  font-weight: 800;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 0.5rem;
  animation: titleGlow 3s ease-in-out infinite alternate;
}

@keyframes titleGlow {
  0% {
    filter: drop-shadow(0 0 5px rgba(102, 126, 234, 0.3));
  }

  100% {
    filter: drop-shadow(0 0 20px rgba(118, 75, 162, 0.5));
  }
}

.subtitle {
  font-size: 1.1rem;
  color: rgba(255, 255, 255, 0.9);
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.refresh-section {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 1rem;
  margin-bottom: 2rem;
  padding: 1rem;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  flex-wrap: wrap;
}

.search-container {
  position: relative;
  flex: 1;
  min-width: 200px;
  max-width: 300px;
}

.search-input {
  width: 100%;
  padding: 0.75rem 2.5rem 0.75rem 1rem;
  border: 2px solid rgba(102, 126, 234, 0.2);
  border-radius: 8px;
  font-size: 0.9rem;
  background: rgba(255, 255, 255, 0.8);
  transition: all 0.3s;
  box-sizing: border-box;
}

.search-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2);
  background: rgba(255, 255, 255, 1);
}

.search-icon {
  position: absolute;
  right: 0.75rem;
  top: 50%;
  transform: translateY(-50%);
  color: #6b7280;
  font-size: 1rem;
  pointer-events: none;
}

.sort-container {
  flex-shrink: 0;
}

.sort-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
  border: 1px solid rgba(102, 126, 234, 0.2);
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s;
}

.sort-btn:hover {
  background: rgba(102, 126, 234, 0.2);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.2);
}

.sort-icon {
  font-size: 1rem;
  transition: transform 0.3s;
}

.sort-btn:hover .sort-icon {
  transform: scale(1.1);
}

.sort-text {
  white-space: nowrap;
}

.refresh-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s;
}

.refresh-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.refresh-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.loading-spinner {
  width: 16px;
  height: 16px;
  border: 2px solid #ffffff30;
  border-top: 2px solid #ffffff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

.refresh-icon {
  font-size: 1.1rem;
  animation: refreshRotate 2s ease-in-out infinite;
}

@keyframes refreshRotate {

  0%,
  100% {
    transform: rotate(0deg);
  }

  50% {
    transform: rotate(180deg);
  }
}

.last-update {
  font-size: 0.875rem;
  color: #6b7280;
}

.server-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 2rem;
}

.loading-state,
.empty-state {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 3rem;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  text-align: center;
  color: #6b7280;
}

.server-item {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 1.5rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: relative;
  overflow: hidden;
}

.server-online {
  border-left: 4px solid #10b981;
}

.server-maintenance {
  border-left: 4px solid #f59e0b;
}

.server-offline {
  border-left: 4px solid #ef4444;
}

.server-main-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.server-status-indicator {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
}

.status-dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  animation: statusPulse 2s ease-in-out infinite;
}

.status-dot.up {
  background: #10b981;
  box-shadow: 0 0 10px rgba(16, 185, 129, 0.5);
}

.status-dot.maintenance {
  background: #f59e0b;
  box-shadow: 0 0 10px rgba(245, 158, 11, 0.5);
}

.status-dot.down {
  background: #ef4444;
  box-shadow: 0 0 10px rgba(239, 68, 68, 0.5);
}

@keyframes statusPulse {

  0%,
  100% {
    opacity: 1;
  }

  50% {
    opacity: 0.5;
  }
}

.status-text {
  font-size: 0.75rem;
  font-weight: 500;
}

.server-online .status-text {
  color: #10b981;
}

.server-maintenance .status-text {
  color: #f59e0b;
}

.server-offline .status-text {
  color: #ef4444;
}

.server-details {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.server-name {
  font-size: 1.125rem;
  font-weight: 600;
  color: #667eea;
  margin: 0;
}

.server-meta {
  display: flex;
  gap: 1rem;
}

.server-zone,
.server-type {
  font-size: 0.8rem;
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
  font-weight: 500;
}

.server-zone {
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
}

.server-type {
  background: rgba(118, 75, 162, 0.1);
  color: #764ba2;
}

.server-load {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 0.5rem;
  min-width: 100px;
}

.load-info {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.load-label {
  font-size: 0.8rem;
  color: #6b7280;
}

.load-value {
  font-size: 0.875rem;
  font-weight: 600;
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
}

.load-low {
  background: rgba(34, 197, 94, 0.1);
  color: #22c55e;
}

.load-medium {
  background: rgba(251, 191, 36, 0.1);
  color: #fbbf24;
}

.load-high {
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
}

.load-bar {
  width: 80px;
  height: 4px;
  background: rgba(0, 0, 0, 0.1);
  border-radius: 2px;
  overflow: hidden;
}

.load-fill {
  height: 100%;
  border-radius: 2px;
  transition: width 0.3s ease;
  animation: loadPulse 2s ease-in-out infinite alternate;
}

@keyframes loadPulse {
  0% {
    opacity: 0.8;
  }

  100% {
    opacity: 1;
  }
}

.stats-section {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem
}

.stat-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 1.5rem;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  display: flex;
  align-items: center;
  gap: 1rem;
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 12px 35px rgba(0, 0, 0, 0.15);
}

.stat-icon {
  font-size: 2rem;
  animation: statIconBounce 2s ease-in-out infinite;
}

@keyframes statIconBounce {

  0%,
  100% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.1);
  }
}

.stat-info {
  display: flex;
  flex-direction: column;
}

.stat-number {
  font-size: 1.5rem;
  font-weight: 700;
  color: #667eea;
}

.stat-label {
  font-size: 0.875rem;
  color: #6b7280;
}

/* 移动端适配 */
@media (max-width: 640px) {
  .server-status {
    padding: 0.5rem;
  }

  .title {
    font-size: 2rem;
  }

  .refresh-section {
    flex-direction: column-reverse;
    gap: 1rem;
    align-items: stretch;
  }

  .search-container {
    max-width: none;
  }

  .sort-btn {
    justify-content: center;
  }

  .server-item {
    flex-direction: row;
    gap: 0.5rem;
    text-align: left;
    padding: 1rem;
  }

  .server-main-info {
    flex-direction: row;
    flex: 1;
  }

  .server-load {
    align-items: flex-end;
    min-width: 80px;
  }

  .stats-section {
    grid-template-columns: 1fr;
  }
}
</style>