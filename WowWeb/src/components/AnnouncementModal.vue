<template>
  <div v-if="showModal" class="announcement-overlay" @click="closeModal">
    <div class="announcement-modal" @click.stop>
      <!-- 关闭按钮 -->
      <button class="close-btn" @click="closeModal">×</button>

      <!-- 公告标题 -->
      <div class="announcement-header">
        <div class="wow-logo">🛡️</div>
        <h2 class="announcement-title">重要公告</h2>
      </div>

      <!-- 公告内容 -->
      <div class="announcement-content">
        <div class="notice-item">
          <div class="notice-icon">📢</div>
          <div class="notice-text">
            <h3>欢迎使用魔兽世界PVE处罚查询系统</h3>
            <p>本网站提供公正、透明的PVE处罚记录查询服务。QQ交流群：<a href="https://qun.qq.com/universal-share/share?ac=1&authKey=mvWKoHHUaMLCuyiD16ID7dqvAdD9WxWfk4DloraAbG9WQoU2L3DiYGPFHox9Rqfi&busi_data=eyJncm91cENvZGUiOiI1MzU3NDU1OTMiLCJ0b2tlbiI6IkRoV0U3OXB0S1RLRVdick9NRXAyV0lRa1VNc29WaFdGeUl2TlgweTVZdS8zS1RGWVVwdEtkWUpwZ01RaTNKVmMiLCJ1aW4iOiIyODE0NDIxNzk2In0%3D&data=0zvqacnSXxIYN_zazhW84DjBlkUwFbD5Z0hJxAJBdRoNtjYUVieeOM0dTGHSJiC6sHIZNvIR0gQ1WRYoC77CpA&svctype=4&tempid=h5_group_info" target="_blank">535745593</a></p>
          </div>
        </div>

        <div class="notice-item">
          <div class="notice-icon">⚠️</div>
          <div class="notice-text">
            <h3>免责声明</h3>
            <p>查询结果仅供参考，具体处罚决定以暴雪娱乐官方为准。我们不参与任何处罚决定的制定过程。</p>
          </div>
        </div>

        <div class="notice-item">
          <div class="notice-icon">🔒</div>
          <div class="notice-text">
            <h3>隐私保护</h3>
            <p>我们严格保护用户隐私，所有查询数据均经过加密处理，不会用于商业用途。</p>
          </div>
        </div>

        <div class="notice-item">
          <div class="notice-icon">🎮</div>
          <div class="notice-text">
            <h3>更新内容</h3>
            <p>数据内容已更新至{{ formattedCurrentDate }}最新处罚公告，后期网站将更新MDT大秘境工具和DPS模拟等相关功能。</p>
          </div>
        </div>
      </div>

      <!-- 底部按钮 -->
      <div class="announcement-footer">
        <div class="checkbox-container">
          <label class="checkbox-label">
            <input
                type="checkbox"
                v-model="dontShowToday"
                class="checkbox-input"
            />
            <span class="checkbox-custom"></span>
            <span class="checkbox-text">今日不再提示</span>
          </label>
        </div>
        <button class="confirm-btn" @click="closeModal">
          <span class="btn-icon">✨</span>
          我不管，我纯绿玩
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted, computed} from 'vue'

const showModal = ref(false)
const dontShowToday = ref(false)

// 获取今天的日期字符串
const getTodayString = () => {
  return new Date().toDateString()
}

const formattedCurrentDate = computed(() => {
  const now = new Date()
  const year = now.getFullYear()
  const month = String(now.getMonth() + 1).padStart(2, '0')
  const day = String(now.getDate()).padStart(2, '0')
  return `${year}年${month}月${day}日`
})

// 检查今天是否已经显示过公告
const hasSeenTodayAnnouncement = () => {
  const lastSeenDate = localStorage.getItem('lastAnnouncementDate')
  const today = getTodayString()
  return lastSeenDate === today
}

// 组件挂载时显示公告
onMounted(() => {
  // 检查今天是否已经显示过公告
  if (!hasSeenTodayAnnouncement()) {
    setTimeout(() => {
      showModal.value = true
    }, 500) // 延迟500ms显示，让页面先加载完成
  }
})

// 关闭公告
const closeModal = () => {
  showModal.value = false

  // 如果勾选了"今日不再提示"，则记录今天的日期
  if (dontShowToday.value) {
    const today = getTodayString()
    localStorage.setItem('lastAnnouncementDate', today)
  }
}
</script>

<style scoped>
.announcement-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
  padding: 1rem;
  animation: fadeIn 0.3s ease-out;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.announcement-modal {
  background: rgba(255, 255, 255, 0.98);
  backdrop-filter: blur(20px);
  border-radius: 24px;
  max-width: 500px;
  width: 100%;
  max-height: 80vh;
  overflow: hidden;
  position: relative;
  box-shadow:
      0 20px 60px rgba(0, 0, 0, 0.3),
      0 0 0 1px rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  animation: slideIn 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  flex-direction: column;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(-30px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.close-btn {
  position: absolute;
  top: 1rem;
  right: 1rem;
  width: 32px;
  height: 32px;
  border: none;
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
  border-radius: 50%;
  font-size: 1.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s;
  z-index: 1;
}

.close-btn:hover {
  background: rgba(102, 126, 234, 0.2);
  transform: scale(1.1);
}

.announcement-header {
  text-align: center;
  padding: 2rem 2rem 1rem;
}

.wow-logo {
  font-size: 3rem;
  margin-bottom: 1rem;
  animation: logoFloat 3s ease-in-out infinite;
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.2));
}

@keyframes logoFloat {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-8px) rotate(3deg); }
}

.announcement-title {
  font-size: 1.75rem;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0;
  animation: titleGlow 2s ease-in-out infinite alternate;
}

@keyframes titleGlow {
  0% { filter: drop-shadow(0 0 5px rgba(102, 126, 234, 0.3)); }
  100% { filter: drop-shadow(0 0 15px rgba(118, 75, 162, 0.5)); }
}

.announcement-content {
  padding: 0 2rem;
  max-height: 400px;
  overflow-y: auto;
  flex: 1;
}

.notice-item {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
  border-radius: 12px;
  border: 1px solid rgba(102, 126, 234, 0.1);
  transition: all 0.3s;
  position: relative;
  overflow: hidden;
}

.notice-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.1), transparent);
  transition: left 0.5s;
}

.notice-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.15);
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.08) 0%, rgba(118, 75, 162, 0.08) 100%);
}

.notice-item:hover::before {
  left: 100%;
}

.notice-icon {
  font-size: 1.5rem;
  flex-shrink: 0;
  animation: iconPulse 2s ease-in-out infinite alternate;
}

@keyframes iconPulse {
  0% { transform: scale(1); }
  100% { transform: scale(1.1); }
}

.notice-text h3 {
  font-size: 1rem;
  font-weight: 600;
  color: #667eea;
  margin: 0 0 0.5rem 0;
}

.notice-text p {
  font-size: 0.875rem;
  color: #374151;
  margin: 0;
  line-height: 1.5;
}

.announcement-footer {
  padding: 1.5rem 2rem 2rem;
  text-align: center;
  flex-shrink: 0;
}

.checkbox-container {
  margin-bottom: 1.5rem;
  display: flex;
  justify-content: center;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  cursor: pointer;
  font-size: 0.9rem;
  color: #374151;
  transition: all 0.3s;
  padding: 0.5rem;
  border-radius: 8px;
}

.checkbox-label:hover {
  background: rgba(102, 126, 234, 0.05);
  color: #667eea;
}

.checkbox-input {
  display: none;
}

.checkbox-custom {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(102, 126, 234, 0.3);
  border-radius: 4px;
  position: relative;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  background: rgba(255, 255, 255, 0.8);
}

.checkbox-custom::after {
  content: '';
  position: absolute;
  top: 1px;
  left: 5px;
  width: 6px;
  height: 10px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg) scale(0);
  transition: transform 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

.checkbox-input:checked + .checkbox-custom {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-color: #667eea;
  animation: checkboxPulse 0.3s ease-out;
}

.checkbox-input:checked + .checkbox-custom::after {
  transform: rotate(45deg) scale(1);
}

@keyframes checkboxPulse {
  0% { transform: scale(1); }
  50% { transform: scale(1.1); }
  100% { transform: scale(1); }
}

.checkbox-text {
  font-weight: 500;
  user-select: none;
}

.confirm-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.875rem 2rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
}

.confirm-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.confirm-btn:hover {
  background: linear-gradient(135deg, #5a67d8 0%, #6b46c1 100%);
  transform: translateY(-2px) scale(1.05);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.confirm-btn:hover::before {
  left: 100%;
}

.btn-icon {
  font-size: 1.1rem;
  animation: sparkle 1.5s ease-in-out infinite;
}

@keyframes sparkle {
  0%, 100% { transform: scale(1) rotate(0deg); }
  50% { transform: scale(1.2) rotate(180deg); }
}

/* 移动端适配 */
@media (max-width: 640px) {
  .announcement-overlay {
    padding: 0.5rem;
  }

  .announcement-modal {
    max-height: 60vh;
    height: 60vh;
  }

  .announcement-header {
    padding: 1.5rem 1.5rem 1rem;
    flex-shrink: 0;
  }

  .announcement-content {
    padding: 0 1.5rem;
    flex: 1;
    overflow-y: auto;
  }

  .announcement-footer {
    padding: 1.5rem;
    flex-shrink: 0;
  }

  .announcement-title {
    font-size: 1.5rem;
  }

  .wow-logo {
    font-size: 2.5rem;
  }

  .notice-item {
    flex-direction: column;
    text-align: center;
  }

  .confirm-btn {
    width: 100%;
    justify-content: center;
  }
}

/* 滚动条样式 */
.announcement-content::-webkit-scrollbar {
  width: 4px;
}

.announcement-content::-webkit-scrollbar-track {
  background: rgba(102, 126, 234, 0.1);
  border-radius: 2px;
}

.announcement-content::-webkit-scrollbar-thumb {
  background: rgba(102, 126, 234, 0.3);
  border-radius: 2px;
}

.announcement-content::-webkit-scrollbar-thumb:hover {
  background: rgba(102, 126, 234, 0.5);
}
</style>