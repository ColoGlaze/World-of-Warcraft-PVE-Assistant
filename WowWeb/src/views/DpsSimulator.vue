<template>
  <div class="dps-simulator">
    <!-- 返回按钮 -->
    <div class="back-button-container">
      <button @click="goBack" class="back-button">
        <span class="back-icon">←</span>
        返回首页
      </button>
    </div>

    <!-- 头部 -->
    <header class="header">
      <div class="dps-icon">⚔️</div>
      <h1 class="title">DPS模拟器</h1>
      <p class="subtitle">专业的伤害计算与装备优化工具</p>
    </header>

    <!-- 功能模块 -->
    <div class="features-grid">
      <div class="feature-card">
        <div class="feature-icon">🎯</div>
        <div class="feature-content">
          <h3 class="feature-title">最佳装备</h3>
          <p class="feature-desc">为你的角色找到最佳装备和天赋</p>
        </div>
        <div class="feature-actions">
          <button class="feature-btn">开始优化</button>
        </div>
      </div>
      
      <div class="feature-card">
        <div class="feature-icon">⚙️</div>
        <div class="feature-content">
          <h3 class="feature-title">装备优化</h3>
          <p class="feature-desc">找到你最好的掉落物</p>
        </div>
        <div class="feature-actions">
          <button class="feature-btn">查找掉落</button>
        </div>
      </div>
      
      <div class="feature-card">
        <div class="feature-icon">⚡</div>
        <div class="feature-content">
          <h3 class="feature-title">快速模拟</h3>
          <p class="feature-desc">运行快速模拟以获取你的DPS和详细统计数据</p>
        </div>
        <div class="feature-actions">
          <button class="feature-btn">开始模拟</button>
        </div>
      </div>
      
      <div class="feature-card">
        <div class="feature-icon">🔧</div>
        <div class="feature-content">
          <h3 class="feature-title">高级设置</h3>
          <p class="feature-desc">运行(几乎)任何原始的SimulationCraft脚本</p>
        </div>
        <div class="feature-actions">
          <button class="feature-btn">高级模式</button>
        </div>
      </div>
      
      <div class="feature-card">
        <div class="feature-icon">📊</div>
        <div class="feature-content">
          <h3 class="feature-title">属性权重</h3>
          <p class="feature-desc">计算各属性对DPS提升的权重值</p>
        </div>
        <div class="feature-actions">
          <button class="feature-btn">计算权重</button>
        </div>
      </div>
      
      <div class="feature-card">
        <div class="feature-icon">⚖️</div>
        <div class="feature-content">
          <h3 class="feature-title">装备比较</h3>
          <p class="feature-desc">比较不同装备组合的DPS差异</p>
        </div>
        <div class="feature-actions">
          <button class="feature-btn">开始比较</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

// DPS模拟器页面，暂时显示开发中状态

// 定义组件名称，用于keep-alive
defineOptions({
  name: 'DpsSimulator'
})

const router = useRouter()

// 返回首页
const goBack = () => {
  router.push('/')
}
</script>

<style scoped>
.dps-simulator {
  display: flex;
  flex-direction: column;
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

.header {
  text-align: center;
  margin-bottom: 3rem;
  padding: 2rem 0;
}

.dps-icon {
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
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: repeat(3, 1fr);
  gap: 2rem;
}

.feature-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 1.5rem;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  display: flex;
  align-items: center;
  gap: 1.5rem;
  text-align: left;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

.feature-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.1), transparent);
  transition: left 0.5s;
}

.feature-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(102, 126, 234, 0.2);
}

.feature-card:hover::before {
  left: 100%;
}

.feature-icon {
  font-size: 2.5rem;
  flex-shrink: 0;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 16px;
  border: 1px solid rgba(102, 126, 234, 0.2);
  animation: featureIconFloat 3s ease-in-out infinite;
}

@keyframes featureIconFloat {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-5px); }
}

.feature-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.feature-title {
  font-size: 1.125rem;
  font-weight: 600;
  color: #667eea;
  margin: 0;
}

.feature-desc {
  color: #6b7280;
  line-height: 1.6;
  margin: 0;
  font-size: 0.875rem;
}

.feature-actions {
  flex-shrink: 0;
  display: flex;
  align-items: center;
}

.feature-btn {
  padding: 0.625rem 1.5rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 10px;
  font-weight: 500;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s;
  position: relative;
  overflow: hidden;
  white-space: nowrap;
}

.feature-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
  left: -100%;
}

.feature-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.feature-btn:hover::before {
  left: 100%;
}

/* 移动端适配 */
@media (max-width: 640px) {
  .dps-simulator {
    padding: 0.5rem;
  }
  
  .title {
    font-size: 2rem;
  }
  
  .features-grid {
    grid-template-columns: 1fr;
    grid-template-rows: auto;
    gap: 1rem;
  }
  
  .dps-icon {
    font-size: 3rem;
  }
  
  .feature-card {
    padding: 1.5rem;
    gap: 1.5rem;
    flex-direction: column;
    text-align: center;
    align-items: center;
  }
  
  .feature-icon {
    font-size: 2.5rem;
    width: 60px;
    height: 60px;
  }
  
  .feature-content {
    text-align: center;
    align-items: center;
  }
  
  .feature-title {
    font-size: 1.125rem;
  }
  
  .feature-desc {
    font-size: 0.875rem;
    text-align: center;
  }
  
  .feature-actions {
    width: 100%;
    justify-content: center;
  }
  
  .feature-btn {
    padding: 0.75rem 2rem;
    font-size: 0.875rem;
    width: 100%;
    max-width: 200px;
  }
}
</style>