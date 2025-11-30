<template>
  <div class="penalty-ranking">
    <!-- 返回按钮 -->
    <div class="back-button-container">
      <button @click="goBack" class="back-button">
        <span class="back-icon">←</span>
        返回首页
      </button>
    </div>

    <!-- 头部 -->
    <header class="header">
      <div class="ranking-icon">📊</div>
      <h1 class="title">处罚角色服务器占比排行</h1>
      <p class="subtitle">各服务器处罚角色数量统计</p>
    </header>

    <!-- 统计概览 -->
    <div class="stats-overview">
      <div class="stat-card">
        <div class="stat-icon">👥</div>
        <div class="stat-info">
          <div class="stat-number">{{ totalPenalties }}位</div>
          <div class="stat-label">总处罚角色</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">🌐</div>
        <div class="stat-info">
          <div class="stat-number">{{ totalServers }}个</div>
          <div class="stat-label">涉及服务器</div>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">📈</div>
        <div class="stat-info">
          <div class="stat-number">{{ updateTime }}</div>
          <div class="stat-label">最后更新</div>
        </div>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
      <div class="loading-spinner"></div>
      <p>正在加载数据...</p>
    </div>

    <!-- 错误提示 -->
    <div v-else-if="error" class="error-message">
      {{ error }}
    </div>

    <!-- 完整排行榜 -->
    <div v-else class="ranking-section">
      <div class="section-header">
        <h2 class="section-title">完整排行榜</h2>
        <div class="filter-controls">
          <select v-model="sortOrder" class="sort-select">
            <option value="desc">占比从高到低</option>
            <option value="asc">占比从低到高</option>
          </select>
        </div>
      </div>

      <div class="ranking-table">
        <div class="table-header">
          <div class="header-rank">排名</div>
          <div class="header-server">服务器</div>
          <div class="header-count">处罚角色数量</div>
          <div class="header-percentage">占比</div>
          <div class="header-chart">占比图</div>
        </div>

       <div class="table-body">
          <div
              v-for="(server, index) in sortedServers"
              :key="server.id"
              class="table-row"
              :class="{ 'top-three': index < 3 }"
          >
            <div class="cell-rank">
              <div class="rank-badge" :class="getRankClass(index)">
                {{ index + 1 }}
              </div>
            </div>
            <div class="cell-server">
              <div class="server-info">
                <div class="server-name">{{ server.name }}</div>
                <div class="server-zone" v-if="server.zone">{{ server.zone }}</div>
              </div>
            </div>
            <div class="cell-count">
              <div class="count-number">{{ server.count }}</div>
            </div>
            <div class="cell-percentage">
              <div class="percentage-number">{{ server.percentage }}%</div>
            </div>
            <div class="cell-chart">
              <div class="chart-container">
                <div
                    class="chart-bar"
                    :class="getRankClass(index)"
                    :style="{ width: (server.percentage / maxPercentage * 100) + '%' }"
                ></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 数据说明 -->
    <div v-if="!loading" class="data-note">
      <div class="note-icon">ℹ️</div>
      <div class="note-content">
        <h3>数据说明</h3>
        <p>• 统计时间范围：最近30天</p>
        <p>• 数据来源：PVE处罚查询系统</p>
        <p>• 更新频率：每日凌晨自动更新</p>
        <p>• 占比计算：各服务器处罚数量 / 总处罚数量 × 100%</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { penaltyAPI } from '../api/penalty'

// 定义组件名称
defineOptions({
  name: 'PenaltyRanking'
})

const router = useRouter()
const sortOrder = ref('desc')
const loading = ref(false)
const error = ref('')

// 完整服务器数据
const serverData = ref([])

// 获取处罚服务器占比排行
const fetchPenaltyRanking = async () => {
loading.value = true
  error.value = ''
  
  try {
    const response = await penaltyAPI.getServerPenaltyList(0)
    if (response.status === 'success' && response.data) {
      // 转换数据格式以适应组件需求
      serverData.value = response.data.map((server, index) => ({
        id: index + 1,
        name: server.server_name,
        zone: server.zone || '',
        count: server.count,
        percentage: server.proportion.toFixed(3)
      }))
    }
  } catch (err) {
    ElMessage.error('获取数据失败，请稍后重试:'+err)
  } finally {
    loading.value = false
  }
}

// 计算属性
const totalPenalties = computed(() => {
  return serverData.value.reduce((sum, server) => sum + server.count, 0)
})

const totalServers = computed(() => {
  return serverData.value.length
})

const updateTime = computed(() => {
  const now = new Date();
  const year = now.getFullYear();
  const month = now.getMonth() + 1;
  const date = now.getDate();
  return `${year}年${month}月${date}日`;
})

const maxPercentage = computed(() => {
  return Math.max(...serverData.value.map(server => server.percentage))
})

const sortedServers = computed(() => {
  const sorted = [...serverData.value]
  if (sortOrder.value === 'desc') {
    return sorted.sort((a, b) => b.percentage - a.percentage)
  } else {
    return sorted.sort((a, b) => a.percentage - b.percentage)
  }
})

// 返回首页
const goBack = () => {
  router.push('/')
}

// 获取排名样式类
const getRankClass = (index) => {
  if (index === 0) return 'rank-first'
  if (index === 1) return 'rank-second'
  if (index === 2) return 'rank-third'
  return 'rank-normal'
}

onMounted(() => {
  fetchPenaltyRanking()
})
</script>

<style scoped>
.penalty-ranking {
  padding: 1rem;
  max-width: 1000px;
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

.ranking-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
  animation: iconFloat 3s ease-in-out infinite;
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.3));
}

@keyframes iconFloat {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-8px); }
}

.title {
  font-size: 2.2rem;
  font-weight: 800;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 0.5rem;
}

.subtitle {
  font-size: 1rem;
  color: rgba(255, 255, 255, 0.9);
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.stats-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
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
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
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

.ranking-section {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  padding: 1.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  margin-bottom: 2rem;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.section-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #667eea;
  margin: 0;
}

.filter-controls {
  display: flex;
  gap: 1rem;
}

.sort-select {
  padding: 0.5rem 1rem;
  border: 2px solid rgba(102, 126, 234, 0.2);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.8);
  color: #374151;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s;
}

.sort-select:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2);
}

.ranking-table {
  width: 100%;
}

.table-header {
  display: grid;
  grid-template-columns: 60px 1fr 100px 80px 120px;
  gap: 1rem;
  padding: 1rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 12px;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: #667eea;
  font-size: 0.875rem;
}

.table-body {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.table-row {
  display: grid;
  grid-template-columns: 60px 1fr 100px 80px 120px;
  gap: 1rem;
  padding: 1rem;
  background: rgba(255, 255, 255, 0.8);
  border-radius: 12px;
  transition: all 0.3s;
  align-items: center;
  position: relative;
  overflow: hidden;
}

.table-row::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.1), transparent);
  transition: left 0.5s;
}

.table-row:hover {
  background: rgba(102, 126, 234, 0.05);
  transform: translateX(5px);
}

.table-row:hover::before {
  left: 100%;
}

.table-row.top-three {
  background: linear-gradient(135deg, rgba(255, 215, 0, 0.1) 0%, rgba(255, 237, 78, 0.05) 100%);
  border: 1px solid rgba(255, 215, 0, 0.2);
}

.rank-badge {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 0.875rem;
  color: white;
  animation: rankPulse 2s ease-in-out infinite alternate;
}

.rank-first {
  background: linear-gradient(135deg, #ffd700, #ffed4e);
  box-shadow: 0 0 15px rgba(255, 215, 0, 0.5);
}

.rank-second {
  background: linear-gradient(135deg, #c0c0c0, #e5e7eb);
  box-shadow: 0 0 15px rgba(192, 192, 192, 0.5);
}

.rank-third {
  background: linear-gradient(135deg, #cd7f32, #d97706);
  box-shadow: 0 0 15px rgba(205, 127, 50, 0.5);
}

.rank-normal {
  background: linear-gradient(135deg, #667eea, #764ba2);
  box-shadow: 0 0 10px rgba(102, 126, 234, 0.3);
}

@keyframes rankPulse {
  0% { transform: scale(1); }
  100% { transform: scale(1.05); }
}

.server-info {
  display: flex;
  flex-direction: column;
}

.server-name {
  font-weight: 500;
  color: #374151;
  margin-bottom: 0.25rem;
}

.server-zone {
  font-size: 0.75rem;
  color: #6b7280;
  padding: 0.125rem 0.5rem;
  background: rgba(102, 126, 234, 0.1);
  border-radius: 12px;
  display: inline-block;
  width: fit-content;
}

.count-number {
  font-weight: 600;
  color: #374151;
  text-align: center;
}

.percentage-number {
  font-weight: 600;
  color: #667eea;
}

.chart-container {
  width: 100%;
  height: 8px;
  background: rgba(102, 126, 234, 0.2);
  border-radius: 4px;
  overflow: hidden;
  position: relative;
}

.chart-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: chartShine 2s ease-in-out infinite;
}

@keyframes chartShine {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.chart-bar {
  height: 100%;
  border-radius: 4px;
  transition: width 0.8s ease-out;
  animation: chartGlow 2s ease-in-out infinite alternate;
}

.chart-bar.rank-first {
  background: linear-gradient(90deg, #ffd700, #ffed4e);
}

.chart-bar.rank-second {
  background: linear-gradient(90deg, #c0c0c0, #e5e7eb);
}

.chart-bar.rank-third {
  background: linear-gradient(90deg, #cd7f32, #d97706);
}

.chart-bar.rank-normal {
  background: linear-gradient(90deg, #667eea, #764ba2);
}

@keyframes chartGlow {
  0% { opacity: 0.8; }
  100% { opacity: 1; }
}

.data-note {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 16px;
  padding: 1.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  display: flex;
  gap: 1rem;
}

.note-icon {
  font-size: 1.5rem;
  flex-shrink: 0;
}

.note-content h3 {
  font-size: 1rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 0.5rem;
}

.note-content p {
  font-size: 0.875rem;
  color: #6b7280;
  margin-bottom: 0.25rem;
  line-height: 1.4;
}

.note-content p:last-child {
  margin-bottom: 0;
}

.loading-state {
  text-align: center;
  padding: 3rem;
  color: #667eea;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(102, 126, 234, 0.2);
  border-top: 4px solid #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1rem;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.error-message {
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
  padding: 1.5rem;
  border-radius: 16px;
  text-align: center;
  margin-bottom: 2rem;
  border: 1px solid rgba(239, 68, 68, 0.2);
}

/* 移动端适配 */
@media (max-width: 768px) {
  .penalty-ranking {
    padding: 0.5rem;
  }

  .title {
    font-size: 1.8rem;
  }

  .stats-overview {
    grid-template-columns: 1fr;
  }

  .section-header {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }

  .table-header,
  .table-row {
    grid-template-columns: 50px 1fr 80px 60px 100px;
    gap: 0.5rem;
    padding: 0.75rem;
    font-size: 0.8rem;
  }

  .rank-badge {
    width: 28px;
    height: 28px;
    font-size: 0.75rem;
  }

  .server-name {
    font-size: 0.875rem;
  }

  .server-zone {
    font-size: 0.6875rem;
  }

  .data-note {
    flex-direction: column;
    text-align: center;
  }
}
</style>