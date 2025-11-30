<template>
  <div class="character-search">
    <!-- 返回按钮 -->
    <div class="back-button-container">
      <button @click="goBack" class="back-button">
        <span class="back-icon">←</span>
        返回首页
      </button>
    </div>

    <!-- 头部 -->
    <header class="header">
      <div class="search-icon">🔍</div>
      <h1 class="title">角色搜索结果</h1>
      <p class="subtitle">搜索关键词：{{ searchQuery }}</p>
    </header>

    <!-- 搜索表单 -->
    <div class="search-form">
      <div class="search-input-container">
        <input
          v-model="newSearchQuery"
          type="text"
          placeholder="输入角色名称..."
          class="search-input"
          @keyup.enter="handleSearch"
        />
        <button 
          class="search-btn"
          @click="handleSearch"
          :disabled="loading || !newSearchQuery.trim()"
        >
          <span v-if="loading" class="loading-spinner"></span>
          <span v-else class="search-icon">🔍</span>
        </button>
      </div>
    </div>

    <!-- 搜索结果 -->
    <div v-if="searchResults.length > 0" class="search-results">
      <div class="results-header">
        <h3 class="results-title">找到 {{ searchResults.length }} 个角色(如用户关闭英雄榜角色信息访问权限则不予显示)</h3>
      </div>
      
      <div class="results-list">
        <div
          v-for="character in searchResults"
          :key="character.id"
          class="character-item"
          @click="goToCharacterDetail(character)"
        >
          <div class="character-avatar">
            <img 
              :src="character.avatar_url" 
              alt="角色头像" 
              class="character-avatar-img" 
              :class="character.faction"
            />
            <div class="level-badge">{{ character.level }}</div>
          </div>
          
          <div class="character-info">
            <div class="character-name-section">
              <h4 class="character-name">{{ character.name }}</h4>
              <div class="character-tags">
                <span class="faction-tag" :class="character.faction">
                  {{ character.faction === 'alliance' ? '联盟' : '部落' }}
                </span>
                <span class="class-tag" :class="character.class">
                  {{ character.className }}
                </span>
              </div>
            </div>
            
            <div class="character-details">
              <div class="detail-item">
                <span class="detail-label">服务器：</span>
                <span class="detail-value">{{ character.server }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">种族：</span>
                <span class="detail-value">{{ character.race }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">等级：</span>
                <span class="detail-value">{{ character.level }}</span>
              </div>
            </div>
          </div>
          
          <!-- <div class="character-stats">
            <div class="stat-item">
              <div class="stat-value">{{ character.itemLevel }}</div>
              <div class="stat-label">装等</div>
            </div>
            <div class="stat-item">
              <div class="stat-value">{{ character.mythicScore }}</div>
              <div class="stat-label">大秘境分数</div>
            </div>
          </div> -->
          
          <div class="view-arrow">›</div>
        </div>
      </div>
    </div>

    <!-- 无结果提示 -->
    <div v-else-if="hasSearched && !loading" class="no-results">
      <div class="no-results-icon">😔</div>
      <h3 class="no-results-title">未找到角色</h3>
      <p class="no-results-text">没有找到名为 "{{ searchQuery }}" 的角色，请检查角色名称是否正确。</p>
      <div class="search-tips">
        <h4>搜索提示：</h4>
        <ul>
          <li>确保角色名称拼写正确</li>
          <li>角色名称区分大小写</li>
          <li>可以尝试搜索角色名称的一部分</li>
        </ul>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
      <div class="loading-spinner-large"></div>
      <p class="loading-text">正在搜索角色...</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { penaltyAPI } from '../api/penalty'

defineOptions({
  name: 'CharacterSearch'
})

const route = useRoute()
const router = useRouter()

const searchQuery = ref('')
const newSearchQuery = ref('')
const searchResults = ref([])
const loading = ref(false)
const hasSearched = ref(false)

onMounted(() => {
  searchQuery.value = route.query.name || ''
  newSearchQuery.value = searchQuery.value
  if (searchQuery.value) {
    performSearch(searchQuery.value)
  }
})

const handleSearch = () => {
  if (newSearchQuery.value.trim()) {
    searchQuery.value = newSearchQuery.value.trim()
    performSearch(searchQuery.value)
    // 更新URL
    router.replace({
      query: { name: searchQuery.value }
    })
  }
}

const performSearch = async (query) => {
  loading.value = true
  hasSearched.value = true
  
  try {
    // 调用真实API接口
    const response = await penaltyAPI.getPlayerList({ keyword: query })
    
    // 处理返回的数据格式
    if (response.status === 'success' && response.data?.data?.roles) {
      // 将API返回的角色数据转换为组件需要的格式
      searchResults.value = response.data.data.roles.map((role, index) => ({
        id: `${role.server_id}-${role.role_name}`, // 创建唯一ID
        name: role.role_name,
        server: role.server_name,
        faction: role.faction === '联盟' ? 'alliance' : 'horde',
        race: role.sub_text?.split(' ')[1] || '未知', // 从sub_text中提取种族信息
        class: `class-${role.class_id}`,
        className: role.class_name,
        level: role.level,
        // itemLevel: 0, // API中没有提供装等信息
        // mythicScore: 0, // API中没有提供大秘境分数
        avatar_url: role.avatar_url || '' 
      }))
    } else {
      searchResults.value = []
    }
  } catch (error) {
    console.error('搜索失败:', error)
    searchResults.value = []
  } finally {
    loading.value = false
  }
}

const goToCharacterDetail = (character) => {
  // router.push({
  //   path: '/character-detail',
  //   query: {
  //     id: character.id,
  //     name: character.name,
  //     server: character.server
  //   }
  // })
}

const goBack = () => {
  router.push('/')
}
</script>

<style scoped>
.character-search {
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

.search-icon {
  font-size: 4rem;
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

.search-form {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 1.5rem;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  margin-bottom: 2rem;
  width: 100%;
  box-sizing: border-box;
}

.search-input-container {
  display: flex;
  gap: 0.5rem;
  align-items: center;
  width: 100%;
}

.search-input {
  flex: 1;
  padding: 0.75rem;
  height: 48px;
  border: 2px solid rgba(102, 126, 234, 0.2);
  border-radius: 8px;
  font-size: 1rem;
  background: rgba(255, 255, 255, 0.8);
  transition: all 0.3s;
  box-sizing: border-box;
  width: 100%;
  min-width: 0;
}

.search-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2);
  background: rgba(255, 255, 255, 1);
}

.search-btn {
  padding: 0.75rem 1.5rem;
  height: 48px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  min-width: 120px;
  box-sizing: border-box;
}

.search-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.search-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.search-btn .search-icon {
  font-size: 1.5rem;
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
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.search-results {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  overflow: hidden;
}

.results-header {
  padding: 1.5rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.results-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #667eea;
  margin: 0;
}

.results-list {
  display: flex;
  flex-direction: column;
}

.character-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem;
  cursor: pointer;
  transition: all 0.3s;
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
  position: relative;
  overflow: hidden;
}

.character-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.05), transparent);
  transition: left 0.5s;
}

.character-item:hover {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
  transform: translateX(5px);
}

.character-item:hover::before {
  left: 100%;
}

.character-item:last-child {
  border-bottom: none;
}

.character-avatar {
  position: relative;
  flex-shrink: 0;
}

/* 修改：添加头像图片样式 */
.character-avatar-img {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  position: relative;
}

.character-avatar-img.alliance {
  border: 2px solid #0066cc;
  box-shadow: 0 0 15px rgba(0, 102, 204, 0.3);
}

.character-avatar-img.horde {
  border: 2px solid #cc0000;
  box-shadow: 0 0 15px rgba(204, 0, 0, 0.3);
}

.level-badge {
  position: absolute;
  bottom: -5px;
  right: -5px;
  background: linear-gradient(135deg, #fbbf24, #f59e0b);
  color: white;
  font-size: 0.75rem;
  font-weight: 600;
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.character-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.character-name-section {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.character-name {
  font-size: 1.25rem;
  font-weight: 600;
  color: #374151;
  margin: 0;
}

.character-tags {
  display: flex;
  gap: 0.5rem;
}

.faction-tag, .class-tag {
  padding: 0.25rem 0.75rem;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: 500;
}

.faction-tag.alliance {
  background: rgba(0, 102, 204, 0.1);
  color: #0066cc;
  border: 1px solid rgba(0, 102, 204, 0.2);
}

.faction-tag.horde {
  background: rgba(204, 0, 0, 0.1);
  color: #cc0000;
  border: 1px solid rgba(204, 0, 0, 0.2);
}

.class-tag {
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
  border: 1px solid rgba(102, 126, 234, 0.2);
}

.character-details {
  display: flex;
  gap: 1rem;
}

.detail-item {
  font-size: 0.875rem;
}

.detail-label {
  color: #6b7280;
}

.detail-value {
  color: #374151;
  font-weight: 500;
}

.character-stats {
  display: flex;
  gap: 1rem;
  flex-shrink: 0;
}

.stat-item {
  text-align: center;
}

.stat-value {
  font-size: 1.125rem;
  font-weight: 600;
  color: #667eea;
}

.stat-label {
  font-size: 0.75rem;
  color: #6b7280;
}

.view-arrow {
  font-size: 1.5rem;
  color: #6b7280;
  flex-shrink: 0;
}

.no-results {
  text-align: center;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.no-results-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
}

.no-results-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #374151;
  margin-bottom: 0.5rem;
}

.no-results-text {
  color: #6b7280;
  margin-bottom: 2rem;
}

.search-tips {
  text-align: left;
  max-width: 400px;
  margin: 0 auto;
}

.search-tips h4 {
  color: #667eea;
  margin-bottom: 0.5rem;
}

.search-tips ul {
  color: #6b7280;
  font-size: 0.875rem;
}

.loading-state {
  text-align: center;
  padding: 3rem 2rem;
}

.loading-spinner-large {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(102, 126, 234, 0.2);
  border-top: 4px solid #667eea;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1rem;
}

.loading-text {
  color: #667eea;
  font-size: 1.125rem;
  font-weight: 500;
}

/* 移动端适配 */
@media (max-width: 640px) {
  .character-search {
    padding: 0.5rem;
  }
  
  .character-item {
    flex-direction: column;
    text-align: center;
    gap: 1rem;
  }
  
  .character-name-section {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .character-details {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .character-stats {
    justify-content: center;
  }
}
</style>