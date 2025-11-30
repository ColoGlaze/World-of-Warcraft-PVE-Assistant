<template>
  <div class="search-penalty">
    <!-- 返回首页按钮 -->
    <div class="back-button-container">
      <button @click="goToHome" class="back-button">
        <span class="back-icon">←</span>
        返回首页
      </button>
    </div>

    <!-- 头部标题 -->
    <header class="header">
      <div class="wow-icon">⚔️</div>
      <h1 class="title">魔兽世界PVE处罚查询</h1>
      <p class="subtitle">查询角色的PVE处罚记录</p>
    </header>

    <!-- 搜索表单 -->
    <div class="search-form">
      <div class="form-group">
        <label for="characterName">角色名称</label>
        <input id="characterName" v-model="searchForm.characterName" type="text" placeholder="请输入角色名称"
          class="form-input" @keyup.enter="handleSearch" />
      </div>

      <div class="form-group">
        <label for="server">服务器</label>
        <input id="server" v-model="searchForm.server" type="text" placeholder="请输入服务器名称(可选)" class="form-input" />
      </div>

      <button class="search-btn" :disabled="loading || !searchForm.characterName" @click="handleSearch">
        <span v-if="loading" class="loading-spinner"></span>
        {{ loading ? '查询中...' : '查询处罚' }}
      </button>
    </div>

    <!-- 搜索结果 -->
    <div v-if="searchResults.length > 0" class="search-results">
      <h3 class="results-title">查询结果</h3>
      <div class="results-list">
        <div v-for="result in searchResults" :key="result.id" class="result-item">
          <div class="character-info">
            <h4 class="character-name">{{ result.player_name }}</h4>
            <span class="server-name">{{ result.server_name }}</span>
          </div>
          <div class="penalty-info">
            <span class="penalty-type" :class="getPenaltyTypeClass(result.violation_count)">
              {{ result.violation_count }}
            </span>
            <span class="penalty-date">{{ formatDate(result.violation_time) }}</span>
          </div>
          <div class="penalty-reason">
            处罚原因：{{ result.violation_reason }}
          </div>
        </div>
      </div>
    </div>

    <!-- 无结果提示 -->
    <div v-else-if="hasSearched && !loading" class="no-results">
      <div class="green-player-cert">
        <div class="cert-background">
          <div class="cert-border">
            <div class="cert-content">
              <div class="cert-header">
                <div class="wow-emblem">🛡️</div>
                <h3 class="cert-title">绿色玩家认证</h3>
                <div class="cert-subtitle">GREEN PLAYER CERTIFICATE</div>
              </div>

              <div class="cert-body">
                <div class="player-info">
                  <div class="player-name">{{ searchForm.characterName }}</div>
                  <div class="server-info">{{ searchForm.server || '' }}</div>
                </div>

                <div class="cert-text">
                  <p>该角色无任何PVE处罚记录</p>
                  <p>纯绿玩家！</p>
                </div>

                <div class="cert-stamp">
                  <div class="stamp-circle">
                    <div class="stamp-inner">
                      <div class="stamp-text-top">纯绿玩家</div>
                      <div class="stamp-icon">✓</div>
                      <div class="stamp-text-bottom">VERIFIED</div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="cert-footer">
                <div class="cert-date">认证日期：{{ new Date().toLocaleDateString('zh-CN') }}</div>
                <div class="cert-id">证书编号：WOW-{{ Date.now().toString().slice(-8) }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 错误提示 -->
    <div v-if="errorMessage" class="error-message">
      {{ errorMessage }}
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { penaltyAPI } from '../api/penalty'
import { ElMessage } from "element-plus";

// 响应式数据
const route = useRoute()
const router = useRouter()
const searchForm = ref({
  characterName: '',
  server: ''
})

const searchResults = ref([])
const loading = ref(false)
const hasSearched = ref(false)
const errorMessage = ref('')

// 搜索处理
const handleSearch = async () => {
  if (!searchForm.value.characterName.trim()) {
    errorMessage.value = '请输入角色名称'
    return
  }

  loading.value = true
  errorMessage.value = ''
  hasSearched.value = true

  try {
    const params = {
      player_name: searchForm.value.characterName.trim(),
      server_name: searchForm.value.server
    }

    const response = await penaltyAPI.searchPenalty(params)
    searchResults.value = response.data || []
    loading.value = false

    if (searchResults.value.length > 0) {
      ElMessage({
        message: `查询成功！找到 ${searchResults.value.length} 条处罚记录`,
        type: 'success',
        duration: 3000,
        showClose: true
      })
    }
  } catch (error) {
    errorMessage.value = error.message || '搜索失败，请稍后重试'
    searchResults.value = []
    loading.value = false

    ElMessage({
      message: error.message || '搜索失败，请稍后重试',
      type: 'error',
      duration: 3000,
      showClose: true
    })
  }
}

// 获取处罚类型样式类
const getPenaltyTypeClass = (type) => {
  const typeMap = {
    '首次违规': 'warning',
    '暂停': 'suspend',
    '多次违规': 'ban'
  }
  return typeMap[type] || 'default'
}

// 格式化日期
const formatDate = (dateString) => {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}年${month}月${day}日上榜`;
}

// 返回首页
const goToHome = () => {
  router.push('/')
}
</script>

<style scoped>
.search-penalty {
  padding: 1rem;
  max-width: 800px;
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

.wow-icon {
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
  font-size: 2.2rem;
  font-weight: 800;
  background: linear-gradient(135deg, #667eea 0%, white 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 0.5rem;
  text-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
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
  font-size: 1rem;
  color: rgba(255, 255, 255, 0.9);
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

@keyframes aboutIconPulse {
  0% {
    transform: scale(1);
  }

  100% {
    transform: scale(1.1);
  }
}

.search-form {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 1.5rem;
  border-radius: 20px;
  box-shadow:
    0 8px 32px rgba(0, 0, 0, 0.1),
    0 0 0 1px rgba(255, 255, 255, 0.2);
  margin-bottom: 2rem;
  border: 1px solid rgba(255, 255, 255, 0.3);
  animation: formFloat 6s ease-in-out infinite;
}

@keyframes formFloat {

  0%,
  100% {
    transform: translateY(0px);
  }

  50% {
    transform: translateY(-5px);
  }
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group:last-of-type {
  margin-bottom: 2rem;
}

.form-group label {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
  margin-bottom: 0.5rem;
}

.form-input,
.form-input {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid rgba(102, 126, 234, 0.2);
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
  background: rgba(255, 255, 255, 0.8);
}

.form-input:focus,
.form-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow:
    0 0 0 3px rgba(102, 126, 234, 0.2),
    0 4px 12px rgba(102, 126, 234, 0.15);
  background: rgba(255, 255, 255, 1);
}

.search-btn {
  width: 100%;
  padding: 0.875rem 1.5rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  position: relative;
  overflow: hidden;
}

.search-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.search-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
  transform: translateY(-2px) scale(1.02);
  box-shadow:
    0 8px 25px rgba(102, 126, 234, 0.4),
    0 0 0 1px rgba(255, 255, 255, 0.1);
}

.search-btn:hover:not(:disabled)::before {
  left: 100%;
}

.search-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
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

.search-results {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  box-shadow:
    0 8px 32px rgba(0, 0, 0, 0.1),
    0 0 0 1px rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  overflow: hidden;
}

.results-title {
  padding: 1.5rem;
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: #1f2937;
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
  position: relative;
}

.results-list {}

.result-item {
  padding: 1.5rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

.result-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.05), transparent);
  transition: left 0.5s;
}

.results-list>.result-item:not(:first-child) {
  border-top: 1px solid rgba(102, 126, 234, 0.1);
}

.result-item:hover {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.08) 0%, rgba(118, 75, 162, 0.08) 100%);
  transform: translateX(5px);
  box-shadow: inset 4px 0 0 rgba(102, 126, 234, 0.3);
}

.result-item:hover::before {
  left: 100%;
}

.character-info {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 0.75rem;
}

.character-name {
  font-size: 1.125rem;
  font-weight: 600;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0;
}

.server-name {
  padding: 0.25rem 0.75rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  color: #667eea;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 500;
  border: 1px solid rgba(102, 126, 234, 0.2);
}

.penalty-info {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 0.75rem;
}

.penalty-type {
  padding: 0.3rem 0.8rem;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  position: relative;
  overflow: hidden;
  animation: pulseGlow 2s ease-in-out infinite alternate;
}

.penalty-type.warning {
  background: linear-gradient(135deg, #fbbf24, #f59e0b);
  color: white;
  box-shadow: 0 4px 15px rgba(251, 191, 36, 0.4);
}

.penalty-type.suspend {
  background: linear-gradient(135deg, #f97316, #ea580c);
  color: white;
  box-shadow: 0 4px 15px rgba(249, 115, 22, 0.4);
}

.penalty-type.ban {
  background: linear-gradient(135deg, #ef4444, #dc2626);
  color: white;
  box-shadow: 0 4px 15px rgba(239, 68, 68, 0.4);
}

@keyframes pulseGlow {
  0% {
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
  }

  100% {
    box-shadow: 0 6px 20px rgba(0, 0, 0, 0.4);
  }
}

.penalty-date {
  color: #667eea;
  font-size: 0.875rem;
  font-weight: 500;
}

.penalty-reason {
  color: #374151;
  font-size: 0.875rem;
  line-height: 1.5;
  padding: 0.75rem;
  background: rgba(102, 126, 234, 0.05);
  border-radius: 8px;
  border-left: 3px solid rgba(102, 126, 234, 0.3);
}

.no-results {
  display: flex;
  justify-content: center;
  padding: 2rem 1rem;
}

.green-player-cert {
  max-width: 500px;
  width: 100%;
}

.cert-background {
  background: linear-gradient(135deg, #10b981 0%, #059669 100%);
  padding: 3px;
  border-radius: 20px;
  position: relative;
  overflow: hidden;
  animation: certGlow 3s ease-in-out infinite alternate;
}

@keyframes certGlow {
  0% {
    box-shadow: 0 0 20px rgba(16, 185, 129, 0.5);
  }

  100% {
    box-shadow: 0 0 40px rgba(5, 150, 105, 0.8);
  }
}

.cert-background::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(from 0deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  animation: certRotate 8s linear infinite;
  z-index: 0;
}

@keyframes certRotate {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

.cert-border {
  background: rgba(255, 255, 255, 0.98);
  border-radius: 18px;
  padding: 2px;
  position: relative;
  z-index: 1;
}

.cert-content {
  background: linear-gradient(135deg, #f0fdf4 0%, #ecfdf5 100%);
  border-radius: 16px;
  padding: 2rem;
  position: relative;
  overflow: hidden;
}

.cert-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background:
    radial-gradient(circle at 20% 20%, rgba(16, 185, 129, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 80% 80%, rgba(5, 150, 105, 0.1) 0%, transparent 50%);
  pointer-events: none;
}

.cert-header {
  text-align: center;
  margin-bottom: 1.5rem;
  position: relative;
  z-index: 1;
}

.wow-emblem {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  animation: emblemFloat 3s ease-in-out infinite;
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.2));
}

@keyframes emblemFloat {

  0%,
  100% {
    transform: translateY(0px) rotate(0deg);
  }

  50% {
    transform: translateY(-5px) rotate(5deg);
  }
}

.cert-title {
  font-size: 1.5rem;
  font-weight: 800;
  color: #059669;
  margin: 0 0 0.25rem 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.cert-subtitle {
  font-size: 0.75rem;
  color: #6b7280;
  font-weight: 500;
  letter-spacing: 2px;
  text-transform: uppercase;
}

.cert-body {
  position: relative;
  z-index: 1;
}

.player-info {
  text-align: center;
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: rgba(16, 185, 129, 0.1);
  border-radius: 12px;
  border: 2px solid rgba(16, 185, 129, 0.2);
}

.player-name {
  font-size: 1.25rem;
  font-weight: 700;
  color: #059669;
  margin-bottom: 0.25rem;
}

.server-info {
  font-size: 0.875rem;
  color: #6b7280;
  font-weight: 500;
}

.cert-text {
  text-align: center;
  margin-bottom: 1.5rem;
}

.cert-text p {
  color: #374151;
  font-size: 0.9rem;
  line-height: 1.5;
  margin-bottom: 0.5rem;
  font-weight: 500;
}

.cert-stamp {
  display: flex;
  justify-content: center;
  margin-bottom: 1.5rem;
}

.stamp-circle {
  width: 100px;
  height: 100px;
  border: 4px solid #10b981;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(16, 185, 129, 0.1) 0%, rgba(5, 150, 105, 0.2) 100%);
  position: relative;
  animation: stampPulse 2s ease-in-out infinite alternate;
}

@keyframes stampPulse {
  0% {
    transform: scale(1);
    box-shadow: 0 0 20px rgba(16, 185, 129, 0.3);
  }

  100% {
    transform: scale(1.05);
    box-shadow: 0 0 30px rgba(5, 150, 105, 0.5);
  }
}

.stamp-circle::before {
  content: '';
  position: absolute;
  top: -4px;
  left: -4px;
  right: -4px;
  bottom: -4px;
  border: 2px dashed #10b981;
  border-radius: 50%;
  animation: stampRotate 10s linear infinite;
}

@keyframes stampRotate {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

.stamp-inner {
  text-align: center;
  position: relative;
  z-index: 1;
}

.stamp-text-top {
  font-size: 0.7rem;
  font-weight: 700;
  color: #059669;
  margin-bottom: 0.25rem;
  letter-spacing: 1px;
}

.stamp-icon {
  font-size: 1.5rem;
  color: #10b981;
  font-weight: 900;
  margin: 0.25rem 0;
  animation: checkPulse 1.5s ease-in-out infinite;
}

@keyframes checkPulse {

  0%,
  100% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.2);
  }
}

.stamp-text-bottom {
  font-size: 0.6rem;
  font-weight: 600;
  color: #6b7280;
  margin-top: 0.25rem;
  letter-spacing: 1px;
}

.cert-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
  border-top: 2px solid rgba(16, 185, 129, 0.2);
  font-size: 0.75rem;
  color: #6b7280;
  position: relative;
  z-index: 1;
}

.cert-date,
.cert-id {
  font-weight: 500;
}

.error-message {
  background: linear-gradient(135deg, rgba(239, 68, 68, 0.1) 0%, rgba(220, 38, 38, 0.1) 100%);
  color: #ef4444;
  padding: 1rem;
  border-radius: 12px;
  margin-top: 1rem;
  text-align: center;
  border: 1px solid rgba(239, 68, 68, 0.2);
  backdrop-filter: blur(10px);
}

/* 移动端适配 */
@media (max-width: 640px) {
  .search-penalty {
    padding: 0.5rem;
  }

  .title {
    font-size: 1.8rem;
  }

  .character-info {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }

  .penalty-info {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
}
</style>