<template>
  <div class="mdt-planner">
    <!-- 返回按钮 -->
    <div class="back-button-container">
      <button @click="goBack" class="back-button">
        <span class="back-icon">←</span>
        返回首页
      </button>
    </div>

    <div class="coming-soon">
      <div class="icon">🗺️</div>
      <h2 class="title">大秘境MDT路线规划</h2>
      <p class="message">功能开发中，敬请期待</p>
      <div class="features">
        <div class="feature-item">
          <span class="feature-icon">🗺️</span>
          <span>地下城路线规划</span>
        </div>
        <div class="feature-item">
          <span class="feature-icon">👥</span>
          <span>小怪分组优化</span>
        </div>
        <div class="feature-item">
          <span class="feature-icon">⏱️</span>
          <span>时间计算分析</span>
        </div>
        <div class="feature-item">
          <span class="feature-icon">📊</span>
          <span>路线分享导出</span>
        </div>
      </div>
      <div class="progress-bar">
        <div class="progress-fill"></div>
      </div>
      <p class="progress-text">开发进度 25%</p>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

// 定义组件名称
defineOptions({
  name: 'MdtPlanner'
})

const router = useRouter()

// 返回首页
const goBack = () => {
  router.push('/')
}
</script>

<style scoped>
.mdt-planner {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: calc(100vh - 4rem);
  padding: 2rem 1rem 6rem;
}

.back-button-container {
  align-self: flex-start;
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

.coming-soon {
  text-align: center;
  max-width: 400px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 3rem 2rem;
  border-radius: 24px;
  box-shadow:
      0 20px 40px rgba(0, 0, 0, 0.1),
      0 0 0 1px rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  position: relative;
  overflow: hidden;
}

.coming-soon::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: conic-gradient(from 0deg, transparent, rgba(102, 126, 234, 0.1), transparent);
  animation: rotate 10s linear infinite;
  z-index: -1;
}

@keyframes rotate {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.icon {
  font-size: 4rem;
  margin-bottom: 1.5rem;
  animation: pulse 2s ease-in-out infinite alternate, float 3s ease-in-out infinite;
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.2));
}

@keyframes pulse {
  0% { transform: scale(1); }
  100% { transform: scale(1.1); }
}

@keyframes float {
  0%, 100% { transform: translateY(0px) scale(1); }
  50% { transform: translateY(-10px) scale(1.05); }
}

.title {
  font-size: 1.75rem;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 1rem;
  animation: titleShine 3s ease-in-out infinite;
}

@keyframes titleShine {
  0%, 100% { filter: brightness(1); }
  50% { filter: brightness(1.2); }
}

.message {
  font-size: 1.125rem;
  color: #667eea;
  margin-bottom: 2rem;
  font-weight: 500;
}

.features {
  margin-bottom: 2rem;
}

.feature-item {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  padding: 0.75rem;
  margin-bottom: 0.5rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 12px;
  color: #374151;
  font-weight: 500;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: 1px solid rgba(102, 126, 234, 0.2);
  position: relative;
  overflow: hidden;
}

.feature-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.feature-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.2);
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.15) 0%, rgba(118, 75, 162, 0.15) 100%);
}

.feature-item:hover::before {
  left: 100%;
}

.feature-icon {
  font-size: 1.25rem;
  animation: iconBounce 2s ease-in-out infinite;
}

@keyframes iconBounce {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}

.progress-bar {
  width: 100%;
  height: 8px;
  background: rgba(102, 126, 234, 0.2);
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 1rem;
  position: relative;
}

.progress-bar::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
  animation: progressShine 2s ease-in-out infinite;
}

@keyframes progressShine {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.progress-fill {
  width: 25%;
  height: 100%;
  background: linear-gradient(90deg, #667eea, #764ba2);
  border-radius: 8px;
  animation: shimmer 2s ease-in-out infinite alternate, progressGlow 3s ease-in-out infinite;
  position: relative;
}

@keyframes shimmer {
  0% { opacity: 0.8; }
  100% { opacity: 1; }
}

@keyframes progressGlow {
  0%, 100% { box-shadow: 0 0 10px rgba(102, 126, 234, 0.5); }
  50% { box-shadow: 0 0 20px rgba(118, 75, 162, 0.8); }
}

.progress-text {
  font-size: 0.875rem;
  color: #667eea;
  margin: 0;
  font-weight: 600;
}

/* 移动端适配 */
@media (max-width: 640px) {
  .mdt-planner {
    min-height: calc(100vh - 3rem);
    padding: 1rem 0.5rem 5rem;
  }

  .coming-soon {
    padding: 2rem 1.5rem;
  }

  .title {
    font-size: 1.5rem;
  }

  .icon {
    font-size: 3rem;
  }
}
</style>