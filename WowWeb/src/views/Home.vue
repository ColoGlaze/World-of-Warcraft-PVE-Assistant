<template>
  <div class="home">
    <!-- 头部搜索区 -->
    <header class="header">
      <div class="search-section">
        <div class="wow-icon">⚔️</div>
        <h1 class="title">魔兽世界助手</h1>
        <p class="subtitle">您的专业PVE伙伴</p>

        <!-- 快速搜索栏 -->
        <div class="quick-search">
          <input
              v-model="quickSearchText"
              type="text"
              placeholder="搜索角色信息..."
              class="search-input"
              @keyup.enter="handleQuickSearch"
          />
          <button
              class="search-btn"
              @click="handleQuickSearch"
              :disabled="!quickSearchText.trim()"
          >
            <span class="search-icon">🔍</span>
          </button>
        </div>
      </div>
    </header>

    <!-- 金刚区 -->
    <section class="diamond-section">
      <div class="diamond-grid">
        <router-link to="/search" class="diamond-item">
          <div class="diamond-icon">🔍</div>
          <h3 class="diamond-title">处罚查询</h3>
          <p class="diamond-desc">查询角色PVE处罚记录</p>
        </router-link>

        <router-link to="/dps-simulator" class="diamond-item">
          <div class="diamond-icon">⚔️</div>
          <h3 class="diamond-title">DPS模拟</h3>
          <p class="diamond-desc">伤害计算与优化</p>
        </router-link>

        <router-link to="/about" class="diamond-item">
          <div class="diamond-icon">ℹ️</div>
          <h3 class="diamond-title">关于我们</h3>
          <p class="diamond-desc">了解网站信息</p>
        </router-link>

        <router-link to="/server-status" class="diamond-item">
          <div class="diamond-icon">🌐</div>
          <h3 class="diamond-title">服务器状态</h3>
          <p class="diamond-desc">实时服务器监控</p>
        </router-link>

        <router-link to="/mdt" class="diamond-item">
          <div class="diamond-icon">🗺️</div>
          <h3 class="diamond-title">大秘境MDT</h3>
          <p class="diamond-desc">路线规划工具</p>
        </router-link>

        <div class="diamond-item coming-soon">
          <div class="diamond-icon">🛠️</div>
          <h3 class="diamond-title">更多工具</h3>
          <p class="diamond-desc">敬请期待</p>
        </div>
      </div>
    </section>

    <!-- 信息展示区 -->
    <section class="info-section">
      <!-- 公告信息 -->
      <div class="info-card ranking-card" @click="goToRanking">
        <div class="card-header">
          <div class="card-icon">📊</div>
          <h3 class="card-title">处罚服务器占比排行</h3>
          <div class="view-more">查看详情 ›</div>
        </div>
        <div class="card-content">
          <div class="ranking-list">
            <div
                v-for="(server, index) in topServers"
                :key="server.id"
                class="ranking-item"
            >
              <div class="rank-info">
                <div class="rank-number" :class="getRankClass(index)">{{ index + 1 }}</div>
                <div class="server-name">{{ server.name }}</div>
              </div>
              <div class="rank-data">
                <div class="percentage">{{ server.percentage }}%</div>
                <div class="bar-container">
                  <div
                      class="bar-fill"
                      :class="getRankClass(index)"
                      :style="{ width: server.percentage + '%' }"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 统计信息 -->
      <div class="info-card">
        <div class="card-header">
          <div class="card-icon">💰</div>
          <h3 class="card-title">今日金价比例</h3>
        </div>
        <div class="card-content">
          <div class="gold-price-display">
            <div class="price-main">
              <div class="price-ratio">1.5:10,000</div>
              <div class="price-label">人民币 : 游戏金币</div>
            </div>
          </div>
          <div class="price-note">
            <p>💡 比例说明：1.5元人民币 = 10000游戏金币</p>
            <p>📈 数据更新时间：{{ getPriceUpdateTime() }}</p>
          </div>
        </div>
      </div>

      <!-- 史诗钥石词缀 -->
      <div class="info-card">
        <div class="card-header">
          <div class="card-icon">🔑</div>
          <h3 class="card-title">本周史诗钥石词缀</h3>
        </div>
        <div class="card-content">
          <div class="affix-week">
            <div class="week-info">
              <div class="week-number">第 {{ getCurrentWeek() }} 周</div>
              <div class="week-date">{{ getWeekDateRange() }}</div>
            </div>
            <div class="affix-list">
              <div 
                v-for="(affix, index) in mythicAffixes"
                :key="affix.id"
                class="affix-item"
                :class="getAffixLevelClass(index)"
              >
                <div class="affix-icon">
                  <img 
                    :src="affix.icon_url?.trim()"
                    :alt="affix.name"
                    v-if="affix.icon_url?.trim()"
                    @error="(e) => {e.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAiIGhlaWdodD0iMjAiIHZpZXdCb3g9IjAgMCAyMCAyMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48Y2lyY2xlIGN4PSIxMCIgY3k9IjEwIiByPSI4IiBmaWxsPSIjZWY0NDRkIi8+PHBhdGggZD0iTTYgMTZjMC0yLjIgMS44LTQgNC00czQgMS44IDQgNC0xLjggNC00IDQtNCAxLjgtNCA0eiIgZmlsbD0icmdiYSgyNDUsMTc3LDE3NywwLjIpIi8+PHRleHQgeD0iMTAiIHk9IjE4IiBmb250LWZhbWlseT0iQXJpYWwiIGZvbnQtc2l6ZT0iNiIgdGV4dC1hbmNob3I9Im1pZGRsZSI+4oCU5b2T5qilPC90ZXh0Pjwvc3ZnPg==';}"
                  />
                  <!-- 默认图标，如果没有图片URL -->
                  <span v-else>💎</span>
                </div>
                <div class="affix-info">
                  <div class="affix-name">{{ affix.name }}</div>
                  <div class="affix-desc">{{ affix.description }}</div>
                </div>
              </div>
              <!-- 如果没有数据，显示加载中 -->
              <div v-if="mythicAffixes.length === 0" class="loading-affixes">
                加载词缀数据中...
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 使用指南 -->
      <div class="info-card">
        <div class="card-header">
          <div class="card-icon">📖</div>
          <h3 class="card-title">使用指南</h3>
        </div>
        <div class="card-content">
          <div class="guide-steps">
            <div class="guide-step">
              <div class="step-number">1</div>
              <div class="step-text">点击"处罚查询"进入查询页面</div>
            </div>
            <div class="guide-step">
              <div class="step-number">2</div>
              <div class="step-text">输入角色名称和服务器信息</div>
            </div>
            <div class="guide-step">
              <div class="step-number">3</div>
              <div class="step-text">查看详细的处罚记录或绿玩认证</div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { penaltyAPI } from '../api/penalty'

// 定义组件名称
defineOptions({
  name: 'Home'
})

const router = useRouter()
const quickSearchText = ref('')

// 服务器处罚排行数据
const topServers = ref([])
// 史诗钥石词缀数据
const mythicAffixes = ref([])

// 获取服务器处罚排行
const fetchTopServers = async () => {
try {
    const response = await penaltyAPI.getServerPenaltyList(5) // 获取前5名服务器
    if (response.status === 'success' && response.data) {
      // 转换数据格式以适应组件需求
      topServers.value = response.data.map((server, index) => ({
        id: index + 1,
        name: server.server_name,
        percentage: server.proportion.toFixed(1) // 保留1位小数
      }))
    }
  } catch (error) {
    el.message.error('获取服务器排行数据失败:' + error)
  }
}

// 获取史诗钥石词缀数据
const fetchMythicAffixes = async () => {
  try {
    const response = await penaltyAPI.getMethicAffix()
    if (response.status === 'success' && response.data?.affix_details) {
      mythicAffixes.value = response.data.affix_details
    }
  } catch (error) {
    console.error('获取史诗钥石词缀数据失败:', error)
  }
}

const getAffixLevelClass = (index) => {
  // 根据索引设置不同等级的样式
  const levels = ['affix-level-4', 'affix-level-7', 'affix-level-10', 'affix-level-12']
  return levels[index] || 'affix-level-4'
}

// 获取当前周数
const getCurrentWeek = () => {
  const startDate = new Date('2025-08-14')
  const currentDate = new Date()
  if (currentDate < startDate) {
    return 0
  }
  const diffTime = Math.abs(currentDate - startDate)
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  const weekNumber = Math.floor(diffDays / 7) + 1
  return weekNumber
}

// 获取本周日期范围
const getWeekDateRange = () => {
  const now = new Date();
  const today = now.getDay(); // 0-6，0是周日，4是周四
  
  // 计算上个周四的日期
  const lastThursday = new Date(now);
  const daysSinceThursday = today < 4 ? today + 3 : today - 4;
  lastThursday.setDate(now.getDate() - daysSinceThursday);
  
  // 计算下个周四的日期
  const nextThursday = new Date(lastThursday);
  nextThursday.setDate(lastThursday.getDate() + 7);
  
  // 格式化日期函数
  const formatDate = (date) => {
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    return `${month.toString().padStart(2, '0')}月${day.toString().padStart(2, '0')}日`;
  };
  
  return `${formatDate(lastThursday)}—${formatDate(nextThursday)}`;
}

// 获取金价更新时间
const getPriceUpdateTime = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = now.getMonth() + 1;
  const date = now.getDate();
  return `${year}年${month}月${date}日`;
}

// 快速搜索处理
const handleQuickSearch = () => {
  if (quickSearchText.value.trim()) {
    router.push({
      path: '/character-search',
      query: { name: quickSearchText.value.trim() }
    })
  }
}

// 跳转到排行详情页面
const goToRanking = () => {
  router.push('/penalty-ranking')
}

// 获取排名样式类
const getRankClass = (index) => {
  if (index === 0) return 'rank-first'
  if (index === 1) return 'rank-second'
  if (index === 2) return 'rank-third'
  return 'rank-normal'
}

onMounted(() => {
  fetchTopServers()
  fetchMythicAffixes()
})
</script>

<style scoped>
.home {
  padding: 1rem;
  max-width: 1200px;
  margin: 0 auto;
  padding-bottom: 5rem;
}

.header {
  text-align: center;
  margin-bottom: 2rem;
  padding: 2rem 0;
}

.search-section {
  max-width: 600px;
  margin: 0 auto;
}

.wow-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
  animation: iconFloat 3s ease-in-out infinite, iconGlow 2s ease-in-out infinite alternate;
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.3));
}

@keyframes iconFloat {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-10px) rotate(5deg); }
}

@keyframes iconGlow {
  0% { filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.3)) drop-shadow(0 0 20px rgba(102, 126, 234, 0.3)); }
  100% { filter: drop-shadow(0 6px 16px rgba(0, 0, 0, 0.4)) drop-shadow(0 0 30px rgba(118, 75, 162, 0.5)); }
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
  0% { filter: drop-shadow(0 0 5px rgba(102, 126, 234, 0.3)); }
  100% { filter: drop-shadow(0 0 20px rgba(118, 75, 162, 0.5)); }
}

.subtitle {
  font-size: 1.1rem;
  color: rgba(255, 255, 255, 0.9);
  margin: 0 0 2rem 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.quick-search {
  display: flex;
  gap: 0.5rem;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 0.5rem;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.search-input {
  flex: 1;
  padding: 0.75rem 1rem;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  background: rgba(255, 255, 255, 0.8);
  transition: all 0.3s;
}

.search-input:focus {
  outline: none;
  background: rgba(255, 255, 255, 1);
  box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.3);
}

.search-btn {
  padding: 0.75rem 1.5rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.search-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.search-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.search-icon {
  font-size: 1.2rem;
}

.diamond-section {
  margin-bottom: 3rem;
}

.diamond-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 1.5rem;
}

.diamond-item {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 2rem;
  border-radius: 20px;
  text-decoration: none;
  color: inherit;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  text-align: center;
  position: relative;
  overflow: hidden;
}

.diamond-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.1), transparent);
  transition: left 0.5s;
}

.diamond-item:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(102, 126, 234, 0.2);
}

.diamond-item:hover::before {
  left: 100%;
}

.diamond-item.coming-soon {
  opacity: 0.7;
  cursor: default;
}

.diamond-item.coming-soon:hover {
  transform: none;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.diamond-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
  animation: diamondIconFloat 3s ease-in-out infinite;
}

@keyframes diamondIconFloat {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-5px); }
}

.diamond-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 0.5rem;
}

.diamond-desc {
  font-size: 0.9rem;
  color: #6b7280;
  margin: 0;
}

.info-section {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1.5rem;
}

.info-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  padding: 1.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  transition: all 0.3s;
}

.info-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.15);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.card-icon {
  font-size: 1.5rem;
  animation: cardIconPulse 2s ease-in-out infinite alternate;
}

@keyframes cardIconPulse {
  0% { transform: scale(1); }
  100% { transform: scale(1.1); }
}

.card-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: #667eea;
  margin: 0;
}

.card-content {
  color: #374151;
}

.ranking-card {
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.ranking-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(102, 126, 234, 0.2);
}

.card-header {
  position: relative;
}

.view-more {
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
  font-size: 0.875rem;
  color: #667eea;
  font-weight: 500;
  transition: all 0.3s;
}

.ranking-card:hover .view-more {
  color: #764ba2;
  transform: translateY(-50%) translateX(3px);
}

.ranking-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.ranking-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  background: rgba(102, 126, 234, 0.05);
  border-radius: 8px;
  transition: all 0.3s;
  position: relative;
  overflow: hidden;
}

.ranking-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.1), transparent);
  transition: left 0.5s;
}

.ranking-item:hover {
  background: rgba(102, 126, 234, 0.1);
  transform: translateX(3px);
}

.ranking-item:hover::before {
  left: 100%;
}

.rank-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.rank-number {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 0.8rem;
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
  100% { transform: scale(1.1); }
}

.server-name {
  font-weight: 500;
  color: #374151;
}

.rank-data {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  min-width: 120px;
}

.percentage {
  font-weight: 600;
  color: #667eea;
  font-size: 0.9rem;
  min-width: 40px;
  text-align: right;
}

.bar-container {
  width: 60px;
  height: 6px;
  background: rgba(102, 126, 234, 0.2);
  border-radius: 3px;
  overflow: hidden;
  position: relative;
}

.bar-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: barShine 2s ease-in-out infinite;
}

@keyframes barShine {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.bar-fill {
  height: 100%;
  border-radius: 3px;
  transition: width 0.8s ease-out;
  animation: barGlow 2s ease-in-out infinite alternate;
}

.bar-fill.rank-first {
  background: linear-gradient(90deg, #ffd700, #ffed4e);
}

.bar-fill.rank-second {
  background: linear-gradient(90deg, #c0c0c0, #e5e7eb);
}

.bar-fill.rank-third {
  background: linear-gradient(90deg, #cd7f32, #d97706);
}

.bar-fill.rank-normal {
  background: linear-gradient(90deg, #667eea, #764ba2);
}

@keyframes barGlow {
  0% { opacity: 0.8; }
  100% { opacity: 1; }
}

.announcement-item {
  display: flex;
  gap: 1rem;
  margin-bottom: 0.75rem;
  padding: 0.75rem;
  background: rgba(102, 126, 234, 0.05);
  border-radius: 8px;
  transition: all 0.3s;
}

.announcement-item:hover {
  background: rgba(102, 126, 234, 0.1);
  transform: translateX(5px);
}

.announcement-date {
  font-size: 0.8rem;
  color: #667eea;
  font-weight: 500;
  white-space: nowrap;
}

.announcement-text {
  font-size: 0.9rem;
  line-height: 1.4;
}

.gold-price-display {
  display: flex;
  justify-content: center;
  margin-bottom: 1rem;
}

.price-main {
  text-align: center;
  padding: 2rem;
  background: linear-gradient(135deg, rgba(251, 191, 36, 0.1) 0%, rgba(245, 158, 11, 0.1) 100%);
  border-radius: 16px;
  border: 1px solid rgba(251, 191, 36, 0.2);
  transition: all 0.3s;
  min-width: 200px;
}

.price-main:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 25px rgba(251, 191, 36, 0.2);
}

.price-ratio {
  font-size: 2.5rem;
  font-weight: 700;
  color: #f59e0b;
  margin-bottom: 0.5rem;
  animation: priceGlow 2s ease-in-out infinite alternate;
}

@keyframes priceGlow {
  0% { text-shadow: 0 0 10px rgba(245, 158, 11, 0.3); }
  100% { text-shadow: 0 0 20px rgba(245, 158, 11, 0.6); }
}

.price-label {
  font-size: 0.9rem;
  color: #6b7280;
  font-weight: 500;
}

.price-note {
  margin-top: 1rem;
  padding: 0.75rem;
  background: rgba(251, 191, 36, 0.05);
  border-radius: 8px;
  border-left: 3px solid rgba(251, 191, 36, 0.3);
}

.price-note p {
  font-size: 0.8rem;
  color: #6b7280;
  margin-bottom: 0.25rem;
}

.price-note p:last-child {
  margin-bottom: 0;
}
.guide-steps {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.guide-step {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.step-number {
  width: 32px;
  height: 32px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 0.9rem;
  flex-shrink: 0;
}

.step-text {
  font-size: 0.9rem;
  line-height: 1.4;
}

/* 移动端适配 */
@media (max-width: 640px) {
  .home {
    padding: 0.5rem;
  }

  .title {
    font-size: 2rem;
  }

  .diamond-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 1rem;
  }

  .diamond-item {
    padding: 1.5rem 1rem;
  }

  .diamond-icon {
    font-size: 2.5rem;
  }

  .info-section {
    grid-template-columns: 1fr;
  }

  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 0.75rem;
  }

  .price-main {
    padding: 1.5rem;
    min-width: auto;
  }

  .price-ratio {
    font-size: 2rem;
  }

  .guide-step {
    flex-direction: column;
    text-align: center;
  }
}

/* 史诗钥石词缀样式 */
.affix-week {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.week-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 8px;
  border: 1px solid rgba(102, 126, 234, 0.2);
}

.week-number {
  font-weight: 600;
  color: #667eea;
  font-size: 0.9rem;
}

.week-date {
  font-size: 0.8rem;
  color: #6b7280;
}

.affix-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.affix-item {
  display: flex;
  gap: 0.75rem;
  padding: 0.75rem;
  border-radius: 8px;
  border: 1px solid;
  transition: all 0.3s;
  position: relative;
  overflow: hidden;
}

.affix-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  transition: left 0.5s;
}

.affix-item:hover::before {
  left: 100%;
}

.affix-level-4 {
  background: linear-gradient(135deg, rgba(34, 197, 94, 0.1) 0%, rgba(21, 128, 61, 0.1) 100%);
  border-color: rgba(34, 197, 94, 0.3);
}

.affix-level-7 {
  background: linear-gradient(135deg, rgba(251, 191, 36, 0.1) 0%, rgba(245, 158, 11, 0.1) 100%);
  border-color: rgba(251, 191, 36, 0.3);
}

.affix-level-10 {
  background: linear-gradient(135deg, rgba(239, 68, 68, 0.1) 0%, rgba(220, 38, 38, 0.1) 100%);
  border-color: rgba(239, 68, 68, 0.3);
}

.affix-icon {
  font-size: 1.5rem;
  flex-shrink: 0;
  animation: affixIconPulse 2s ease-in-out infinite alternate;
}

.affix-info {
  flex: 1;
}

.affix-name {
  font-weight: 600;
  margin-bottom: 0.25rem;
}

.affix-level-4 .affix-name {
  color: #22c55e;
}

.affix-level-7 .affix-name {
  color: #fbbf24;
}

.affix-level-10 .affix-name {
  color: #ef4444;
}

.affix-desc {
  font-size: 0.8rem;
  color: #6b7280;
  line-height: 1.4;
}

@media (max-width: 640px) {
  .week-info {
    flex-direction: column;
    gap: 0.25rem;
    text-align: center;
  }

  .affix-item {
    flex-direction: column;
    text-align: center;
  }
}
</style>