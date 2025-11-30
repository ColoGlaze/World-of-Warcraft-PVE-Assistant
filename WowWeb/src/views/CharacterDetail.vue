<template>
  <div class="character-detail">
    <!-- 返回按钮 -->
    <div class="back-button-container">
      <button @click="goBack" class="back-button">
        <span class="back-icon">←</span>
        返回搜索结果
      </button>
    </div>

    <!-- 角色基本信息 -->
    <div class="character-header">
      <div class="character-avatar-section">
        <div class="character-avatar" :class="characterData.faction">
          <div class="avatar-placeholder">
            {{ characterData.race.charAt(0) }}
          </div>
          <div class="level-badge">{{ characterData.level }}</div>
        </div>
        <div class="character-basic-info">
          <h1 class="character-name">{{ characterData.name }}</h1>
          <div class="character-tags">
            <span class="faction-tag" :class="characterData.faction">
              {{ characterData.faction === 'alliance' ? '联盟' : '部落' }}
            </span>
            <span class="class-tag">{{ characterData.className }}</span>
            <span class="server-tag">{{ characterData.server }}</span>
          </div>
          <div class="character-meta">
            <span>{{ characterData.race }} {{ characterData.className }}</span>
            <span>等级 {{ characterData.level }}</span>
          </div>
        </div>
      </div>
      
      <div class="character-stats-overview">
        <div class="stat-card">
          <div class="stat-value">{{ characterData.itemLevel }}</div>
          <div class="stat-label">装备等级</div>
        </div>
        <div class="stat-card">
          <div class="stat-value">{{ characterData.mythicScore }}</div>
          <div class="stat-label">大秘境分数</div>
        </div>
        <div class="stat-card">
          <div class="stat-value">{{ characterData.raidProgress }}</div>
          <div class="stat-label">团本进度</div>
        </div>
      </div>
    </div>

    <!-- 详细信息标签页 -->
    <div class="detail-tabs">
      <div class="tab-buttons">
        <button 
          v-for="tab in tabs" 
          :key="tab.id"
          class="tab-button"
          :class="{ active: activeTab === tab.id }"
          @click="activeTab = tab.id"
        >
          <span class="tab-icon">{{ tab.icon }}</span>
          <span class="tab-text">{{ tab.name }}</span>
        </button>
      </div>

      <!-- 装备信息 -->
      <div v-if="activeTab === 'equipment'" class="tab-content">
        <div class="equipment-section">
          <h3 class="section-title">当前装备</h3>
          <div class="equipment-grid">
            <div 
              v-for="item in characterData.equipment" 
              :key="item.slot"
              class="equipment-item"
              :class="getItemQualityClass(item.quality)"
            >
              <div class="item-icon">{{ item.icon }}</div>
              <div class="item-info">
                <div class="item-name">{{ item.name }}</div>
                <div class="item-level">物品等级 {{ item.itemLevel }}</div>
                <div class="item-slot">{{ item.slotName }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 天赋加点 -->
      <div v-if="activeTab === 'talents'" class="tab-content">
        <div class="talents-section">
          <h3 class="section-title">天赋配置</h3>
          <div class="talent-builds">
            <div 
              v-for="build in characterData.talents" 
              :key="build.id"
              class="talent-build"
            >
              <div class="build-header">
                <h4 class="build-name">{{ build.name }}</h4>
                <span class="build-type">{{ build.type }}</span>
              </div>
              <div class="talent-tree">
                <div 
                  v-for="talent in build.talents" 
                  :key="talent.id"
                  class="talent-node"
                  :class="{ active: talent.selected }"
                >
                  <div class="talent-icon">{{ talent.icon }}</div>
                  <div class="talent-name">{{ talent.name }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 属性面板 -->
      <div v-if="activeTab === 'attributes'" class="tab-content">
        <div class="attributes-section">
          <h3 class="section-title">角色属性</h3>
          <div class="attributes-grid">
            <div class="attribute-category">
              <h4 class="category-title">基础属性</h4>
              <div class="attribute-list">
                <div 
                  v-for="attr in characterData.attributes.basic" 
                  :key="attr.name"
                  class="attribute-item"
                >
                  <span class="attr-name">{{ attr.name }}</span>
                  <span class="attr-value">{{ attr.value }}</span>
                </div>
              </div>
            </div>
            
            <div class="attribute-category">
              <h4 class="category-title">战斗属性</h4>
              <div class="attribute-list">
                <div 
                  v-for="attr in characterData.attributes.combat" 
                  :key="attr.name"
                  class="attribute-item"
                >
                  <span class="attr-name">{{ attr.name }}</span>
                  <span class="attr-value">{{ attr.value }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 大秘境记录 -->
      <div v-if="activeTab === 'mythic'" class="tab-content">
        <div class="mythic-section">
          <h3 class="section-title">史诗钥石地下城记录</h3>
          <div class="mythic-records">
            <div 
              v-for="record in characterData.mythicRecords" 
              :key="record.id"
              class="mythic-record"
              :class="getTimingClass(record.timing)"
            >
              <div class="dungeon-info">
                <div class="dungeon-icon">🏰</div>
                <div class="dungeon-details">
                  <div class="dungeon-name">{{ record.dungeonName }}</div>
                  <div class="dungeon-level">+{{ record.level }}</div>
                </div>
              </div>
              
              <div class="record-stats">
                <div class="completion-time">
                  <span class="time-value">{{ record.completionTime }}</span>
                  <span class="time-label">完成时间</span>
                </div>
                <div class="timing-status" :class="getTimingClass(record.timing)">
                  {{ getTimingText(record.timing) }}
                </div>
              </div>
              
              <div class="record-date">{{ formatDate(record.date) }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 团本进度 -->
      <div v-if="activeTab === 'raid'" class="tab-content">
        <div class="raid-section">
          <h3 class="section-title">团队副本进度</h3>
          <div class="raid-progress">
            <div 
              v-for="raid in characterData.raidProgress" 
              :key="raid.id"
              class="raid-instance"
            >
              <div class="raid-header">
                <div class="raid-icon">🏛️</div>
                <div class="raid-info">
                  <h4 class="raid-name">{{ raid.name }}</h4>
                  <div class="raid-difficulty">{{ raid.difficulty }}</div>
                </div>
                <div class="raid-progress-summary">
                  {{ raid.clearedBosses }}/{{ raid.totalBosses }}
                </div>
              </div>
              
              <div class="boss-list">
                <div 
                  v-for="boss in raid.bosses" 
                  :key="boss.id"
                  class="boss-item"
                  :class="{ killed: boss.killed }"
                >
                  <div class="boss-icon">{{ boss.killed ? '✅' : '❌' }}</div>
                  <div class="boss-name">{{ boss.name }}</div>
                  <div class="boss-date" v-if="boss.killed">
                    {{ formatDate(boss.killDate) }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

defineOptions({
  name: 'CharacterDetail'
})

const route = useRoute()
const router = useRouter()

const activeTab = ref('equipment')
const characterData = ref({})

const tabs = [
  { id: 'equipment', name: '装备', icon: '⚔️' },
  { id: 'talents', name: '天赋', icon: '🌟' },
  { id: 'attributes', name: '属性', icon: '📊' },
  { id: 'mythic', name: '大秘境', icon: '🗝️' },
  { id: 'raid', name: '团本', icon: '🏛️' }
]

// 模拟角色详细数据
const mockCharacterData = {
  id: 1,
  name: '德玛西亚',
  server: '暴风城',
  faction: 'alliance',
  race: '人类',
  class: 'warrior',
  className: '战士',
  level: 80,
  itemLevel: 485,
  mythicScore: 2850,
  raidProgress: '8/8 英雄',
  equipment: [
    { slot: 'head', slotName: '头部', name: '泰坦锻造头盔', itemLevel: 489, quality: 'epic', icon: '⛑️' },
    { slot: 'neck', slotName: '颈部', name: '龙鳞项链', itemLevel: 486, quality: 'epic', icon: '📿' },
    { slot: 'shoulder', slotName: '肩部', name: '守护者护肩', itemLevel: 485, quality: 'epic', icon: '🛡️' },
    { slot: 'chest', slotName: '胸部', name: '钢铁胸甲', itemLevel: 490, quality: 'legendary', icon: '🦺' },
    { slot: 'waist', slotName: '腰部', name: '力量腰带', itemLevel: 483, quality: 'epic', icon: '🔗' },
    { slot: 'legs', slotName: '腿部', name: '战士护腿', itemLevel: 487, quality: 'epic', icon: '🦵' },
    { slot: 'feet', slotName: '脚部', name: '钢铁战靴', itemLevel: 484, quality: 'epic', icon: '🥾' },
    { slot: 'mainhand', slotName: '主手', name: '屠龙剑', itemLevel: 492, quality: 'legendary', icon: '⚔️' }
  ],
  talents: [
    {
      id: 1,
      name: '防护专精',
      type: '防御',
      talents: [
        { id: 1, name: '盾牌猛击', icon: '🛡️', selected: true },
        { id: 2, name: '复仇', icon: '⚡', selected: true },
        { id: 3, name: '盾墙', icon: '🏰', selected: false },
        { id: 4, name: '嘲讽', icon: '😤', selected: true }
      ]
    }
  ],
  attributes: {
    basic: [
      { name: '力量', value: '2,450' },
      { name: '敏捷', value: '1,200' },
      { name: '智力', value: '800' },
      { name: '耐力', value: '3,200' }
    ],
    combat: [
      { name: '攻击强度', value: '4,900' },
      { name: '暴击等级', value: '1,850' },
      { name: '急速等级', value: '1,650' },
      { name: '精通等级', value: '1,420' },
      { name: '护甲值', value: '28,500' }
    ]
  },
  mythicRecords: [
    {
      id: 1,
      dungeonName: '奈萨里奥的巢穴',
      level: 15,
      completionTime: '32:45',
      timing: 'intime',
      date: '2024-01-15'
    },
    {
      id: 2,
      dungeonName: '蕨皮山谷',
      level: 14,
      completionTime: '28:30',
      timing: 'intime',
      date: '2024-01-14'
    },
    {
      id: 3,
      dungeonName: '注能大厅',
      level: 16,
      completionTime: '35:20',
      timing: 'overtime',
      date: '2024-01-13'
    }
  ],
  raidProgress: [
    {
      id: 1,
      name: '阿贝鲁斯，焚影熔炉',
      difficulty: '英雄难度',
      clearedBosses: 8,
      totalBosses: 9,
      bosses: [
        { id: 1, name: '卡扎拉', killed: true, killDate: '2024-01-10' },
        { id: 2, name: '融合室', killed: true, killDate: '2024-01-10' },
        { id: 3, name: '被遗忘的实验', killed: true, killDate: '2024-01-11' },
        { id: 4, name: '突击队长', killed: true, killDate: '2024-01-11' },
        { id: 5, name: '拉什南', killed: true, killDate: '2024-01-12' },
        { id: 6, name: '守护者', killed: true, killDate: '2024-01-12' },
        { id: 7, name: '回响', killed: true, killDate: '2024-01-13' },
        { id: 8, name: '萨卡雷斯', killed: true, killDate: '2024-01-13' },
        { id: 9, name: '萨格拉斯', killed: false, killDate: null }
      ]
    }
  ]
}

onMounted(() => {
  // 根据路由参数加载角色数据
  const characterId = route.query.id
  // 这里应该根据ID从API获取数据，现在使用模拟数据
  characterData.value = mockCharacterData
})

const getItemQualityClass = (quality) => {
  const qualityMap = {
    'common': 'quality-common',
    'uncommon': 'quality-uncommon', 
    'rare': 'quality-rare',
    'epic': 'quality-epic',
    'legendary': 'quality-legendary'
  }
  return qualityMap[quality] || 'quality-common'
}

const getTimingClass = (timing) => {
  return timing === 'intime' ? 'timing-success' : 'timing-overtime'
}

const getTimingText = (timing) => {
  return timing === 'intime' ? '限时完成' : '超时完成'
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('zh-CN')
}

const goBack = () => {
  router.go(-1)
}
</script>

<style scoped>
.character-detail {
  padding: 1rem;
  max-width: 1200px;
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

.character-header {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  padding: 2rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  margin-bottom: 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.character-avatar-section {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.character-avatar {
  position: relative;
}

.avatar-placeholder {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  font-weight: 700;
  color: white;
}

.character-avatar.alliance .avatar-placeholder {
  background: linear-gradient(135deg, #0066cc, #004499);
  box-shadow: 0 0 20px rgba(0, 102, 204, 0.4);
}

.character-avatar.horde .avatar-placeholder {
  background: linear-gradient(135deg, #cc0000, #990000);
  box-shadow: 0 0 20px rgba(204, 0, 0, 0.4);
}

.level-badge {
  position: absolute;
  bottom: -5px;
  right: -5px;
  background: linear-gradient(135deg, #fbbf24, #f59e0b);
  color: white;
  font-size: 0.875rem;
  font-weight: 600;
  padding: 0.375rem 0.75rem;
  border-radius: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.character-basic-info {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.character-name {
  font-size: 2rem;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0;
}

.character-tags {
  display: flex;
  gap: 0.75rem;
}

.faction-tag, .class-tag, .server-tag {
  padding: 0.375rem 0.875rem;
  border-radius: 16px;
  font-size: 0.875rem;
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

.server-tag {
  background: rgba(118, 75, 162, 0.1);
  color: #764ba2;
  border: 1px solid rgba(118, 75, 162, 0.2);
}

.character-meta {
  color: #6b7280;
  font-size: 0.875rem;
  display: flex;
  gap: 1rem;
}

.character-stats-overview {
  display: flex;
  gap: 1.5rem;
}

.stat-card {
  text-align: center;
  padding: 1rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 12px;
  border: 1px solid rgba(102, 126, 234, 0.2);
  min-width: 100px;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: 700;
  color: #667eea;
  margin-bottom: 0.25rem;
}

.stat-label {
  font-size: 0.75rem;
  color: #6b7280;
  font-weight: 500;
}

.detail-tabs {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  overflow: hidden;
}

.tab-buttons {
  display: flex;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.tab-button {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 1rem;
  border: none;
  background: transparent;
  cursor: pointer;
  transition: all 0.3s;
  color: #6b7280;
  font-weight: 500;
}

.tab-button:hover {
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
}

.tab-button.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.tab-icon {
  font-size: 1.125rem;
}

.tab-content {
  padding: 2rem;
}

.section-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 1.5rem;
}

.equipment-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1rem;
}

.equipment-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  border-radius: 12px;
  border: 2px solid;
  transition: all 0.3s;
}

.equipment-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
}

.quality-common {
  border-color: #9ca3af;
  background: rgba(156, 163, 175, 0.1);
}

.quality-uncommon {
  border-color: #10b981;
  background: rgba(16, 185, 129, 0.1);
}

.quality-rare {
  border-color: #3b82f6;
  background: rgba(59, 130, 246, 0.1);
}

.quality-epic {
  border-color: #8b5cf6;
  background: rgba(139, 92, 246, 0.1);
}

.quality-legendary {
  border-color: #f59e0b;
  background: rgba(245, 158, 11, 0.1);
}

.item-icon {
  font-size: 2rem;
  flex-shrink: 0;
}

.item-info {
  flex: 1;
}

.item-name {
  font-weight: 600;
  color: #374151;
  margin-bottom: 0.25rem;
}

.item-level {
  font-size: 0.875rem;
  color: #667eea;
  margin-bottom: 0.25rem;
}

.item-slot {
  font-size: 0.75rem;
  color: #6b7280;
}

.talent-builds {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.talent-build {
  border: 1px solid rgba(102, 126, 234, 0.2);
  border-radius: 12px;
  padding: 1.5rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
}

.build-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.build-name {
  font-size: 1.125rem;
  font-weight: 600;
  color: #374151;
  margin: 0;
}

.build-type {
  padding: 0.25rem 0.75rem;
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: 500;
}

.talent-tree {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 1rem;
}

.talent-node {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  padding: 1rem;
  border-radius: 8px;
  border: 2px solid rgba(156, 163, 175, 0.3);
  background: rgba(156, 163, 175, 0.1);
  transition: all 0.3s;
}

.talent-node.active {
  border-color: #667eea;
  background: rgba(102, 126, 234, 0.1);
}

.talent-icon {
  font-size: 1.5rem;
}

.talent-name {
  font-size: 0.875rem;
  font-weight: 500;
  text-align: center;
  color: #374151;
}

.attributes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.attribute-category {
  border: 1px solid rgba(102, 126, 234, 0.2);
  border-radius: 12px;
  padding: 1.5rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
}

.category-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 1rem;
}

.attribute-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.attribute-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem;
  background: rgba(255, 255, 255, 0.5);
  border-radius: 8px;
}

.attr-name {
  color: #374151;
  font-weight: 500;
}

.attr-value {
  color: #667eea;
  font-weight: 600;
}

.mythic-records {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.mythic-record {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem;
  border-radius: 12px;
  border: 2px solid;
  transition: all 0.3s;
}

.mythic-record:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
}

.timing-success {
  border-color: #10b981;
  background: rgba(16, 185, 129, 0.1);
}

.timing-overtime {
  border-color: #f59e0b;
  background: rgba(245, 158, 11, 0.1);
}

.dungeon-info {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 1;
}

.dungeon-icon {
  font-size: 2rem;
}

.dungeon-details {
  display: flex;
  flex-direction: column;
}

.dungeon-name {
  font-weight: 600;
  color: #374151;
}

.dungeon-level {
  font-size: 0.875rem;
  color: #667eea;
  font-weight: 500;
}

.record-stats {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.completion-time {
  text-align: center;
}

.time-value {
  display: block;
  font-size: 1.125rem;
  font-weight: 600;
  color: #374151;
}

.time-label {
  font-size: 0.75rem;
  color: #6b7280;
}

.timing-status {
  padding: 0.375rem 0.875rem;
  border-radius: 16px;
  font-size: 0.875rem;
  font-weight: 500;
}

.timing-status.timing-success {
  background: rgba(16, 185, 129, 0.2);
  color: #10b981;
}

.timing-status.timing-overtime {
  background: rgba(245, 158, 11, 0.2);
  color: #f59e0b;
}

.record-date {
  font-size: 0.875rem;
  color: #6b7280;
}

.raid-progress {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.raid-instance {
  border: 1px solid rgba(102, 126, 234, 0.2);
  border-radius: 12px;
  padding: 1.5rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
}

.raid-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.raid-icon {
  font-size: 2rem;
}

.raid-info {
  flex: 1;
}

.raid-name {
  font-size: 1.125rem;
  font-weight: 600;
  color: #374151;
  margin: 0 0 0.25rem 0;
}

.raid-difficulty {
  font-size: 0.875rem;
  color: #667eea;
}

.raid-progress-summary {
  font-size: 1.25rem;
  font-weight: 600;
  color: #667eea;
}

.boss-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 1rem;
}

.boss-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1rem;
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.5);
  transition: all 0.3s;
}

.boss-item.killed {
  background: rgba(16, 185, 129, 0.1);
  border: 1px solid rgba(16, 185, 129, 0.2);
}

.boss-item:not(.killed) {
  background: rgba(156, 163, 175, 0.1);
  border: 1px solid rgba(156, 163, 175, 0.2);
}

.boss-icon {
  font-size: 1.25rem;
}

.boss-name {
  flex: 1;
  font-weight: 500;
  color: #374151;
}

.boss-date {
  font-size: 0.75rem;
  color: #6b7280;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .character-detail {
    padding: 0.5rem;
  }
  
  .character-header {
    flex-direction: column;
    gap: 1.5rem;
    text-align: center;
  }
  
  .character-stats-overview {
    justify-content: center;
  }
  
  .tab-buttons {
    flex-wrap: wrap;
  }
  
  .tab-button {
    flex: 1 1 auto;
    min-width: 120px;
  }
  
  .equipment-grid {
    grid-template-columns: 1fr;
  }
  
  .attributes-grid {
    grid-template-columns: 1fr;
  }
  
  .mythic-record {
    flex-direction: column;
    text-align: center;
  }
  
  .boss-list {
    grid-template-columns: 1fr;
  }
}
</style>