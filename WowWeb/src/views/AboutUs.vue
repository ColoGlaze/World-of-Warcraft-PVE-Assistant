<template>
  <div class="about-us">
    <!-- 返回按钮 -->
    <div class="back-button-container">
      <button @click="goBack" class="back-button">
        <span class="back-icon">←</span>
        返回
      </button>
    </div>

    <!-- 头部 -->
    <header class="header">
      <div class="wow-icon">🛡️</div>
      <h1 class="title">关于我们</h1>
      <p class="subtitle">了解魔兽世界PVE处罚查询网站</p>
    </header>

    <!-- 内容区域 -->
    <div class="content">
      <!-- 网站介绍 -->
      <section class="section">
        <div class="section-header">
          <div class="section-icon">🌟</div>
          <h2 class="section-title">网站介绍</h2>
        </div>
        <div class="section-content">
          <p>魔兽世界PVE处罚查询网站是一个专门公益为魔兽世界玩家提供PVE活动处罚记录查询服务的平台。我们致力于帮助玩家快速了解自己角色的PVE处罚记录。</p>
        </div>
      </section>

      <!-- 主要功能 -->
      <section class="section">
        <div class="section-header">
          <div class="section-icon">⚔️</div>
          <h2 class="section-title">主要功能</h2>
        </div>
        <div class="section-content">
          <div class="feature-grid">
            <div class="feature-card">
              <div class="feature-icon">🔍</div>
              <h3>处罚查询</h3>
              <p>快速查询角色的PVE处罚记录，包括装备回收、战网冻结等不同类型的处罚信息。</p>
            </div>
            <div class="feature-card">
              <div class="feature-icon">📊</div>
              <h3>DPS模拟</h3>
              <p>即将推出的DPS模拟器功能，帮助玩家优化装备搭配和低保选取。</p>
            </div>
            <div class="feature-card">
              <div class="feature-icon">📈</div>
              <h3>数据分析</h3>
              <p>提供详细的处罚统计和趋势分析，帮助了解游戏环境变化。</p>
            </div>
            <div class="feature-card">
              <div class="feature-icon">🛡️</div>
              <h3>安全保障</h3>
              <p>严格保护用户隐私，所有查询数据均经过加密处理。</p>
            </div>
          </div>
        </div>
      </section>

      <!-- 使用说明 -->
      <section class="section">
        <div class="section-header">
          <div class="section-icon">📖</div>
          <h2 class="section-title">使用说明</h2>
        </div>
        <div class="section-content">
          <div class="steps">
            <div class="step">
              <div class="step-number">1</div>
              <div class="step-content">
                <h4>输入角色信息</h4>
                <p>在搜索页面输入要查询的角色名称，可选择性地输入服务器名称以获得更精确的结果。</p>
              </div>
            </div>
            <div class="step">
              <div class="step-number">2</div>
              <div class="step-content">
                <h4>执行查询</h4>
                <p>点击"查询处罚"按钮，系统将在数据库中搜索相关的处罚记录。</p>
              </div>
            </div>
            <div class="step">
              <div class="step-number">3</div>
              <div class="step-content">
                <h4>查看结果</h4>
                <p>查询结果将显示处罚类型、时间、原因等详细信息，帮助您了解具体情况。</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 联系我们 -->
      <section class="section">
        <div class="section-header">
          <div class="section-icon">📧</div>
          <h2 class="section-title">联系我们</h2>
        </div>
        <div class="section-content">
          <div class="contact-info">
            <div class="contact-item">
              <div class="contact-icon">✉️</div>
              <div class="contact-details">
                <h4>邮箱联系</h4>
                <p>2814421796@qq.com</p>
              </div>
            </div>
            <div class="contact-item">
              <div class="contact-icon">💬</div>
              <div class="contact-details">
                <h4>绿玩交流QQ群</h4>
                <p>535745593</p>
              </div>
            </div>
            <div class="contact-item">
              <div class="contact-icon">🐛</div>
              <div class="contact-details">
                <h4>问题反馈</h4>
                <p>发现bug或有建议请及时联系我们</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- 免责声明 -->
      <section class="section">
        <div class="section-header">
          <div class="section-icon">⚠️</div>
          <h2 class="section-title">免责声明</h2>
        </div>
        <div class="section-content">
          <div class="disclaimer">
            <p>本网站提供的处罚查询服务仅供参考，具体的处罚决定以暴雪娱乐官方为准。我们不对查询结果的准确性承担法律责任，也不参与任何处罚决定的制定过程。</p>
            <p>使用本网站即表示您同意我们的服务条款和隐私政策。我们承诺保护用户隐私，不会将个人信息用于商业用途。</p>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

// 定义组件名称
defineOptions({
  name: 'AboutUs'
})

const router = useRouter()

// 返回上一页
const goBack = () => {
  router.go(-1)
}
</script>

<style scoped>
.about-us {
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
  margin-bottom: 3rem;
  padding: 2rem 0;
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
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.content {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.section {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  padding: 2rem;
  box-shadow:
      0 8px 32px rgba(0, 0, 0, 0.1),
      0 0 0 1px rgba(255, 255, 255, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.3);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

.section::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.05), transparent);
  transition: left 0.5s;
}

.section:hover {
  transform: translateY(-5px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.section:hover::before {
  left: 100%;
}

.section-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.section-icon {
  font-size: 2rem;
  animation: iconPulse 2s ease-in-out infinite alternate;
}

@keyframes iconPulse {
  0% { transform: scale(1); }
  100% { transform: scale(1.1); }
}

.section-title {
  font-size: 1.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0;
}

.section-content {
  color: #374151;
  line-height: 1.6;
}

.section-content p {
  margin-bottom: 1rem;
}

.feature-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-top: 1rem;
}

.feature-card {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  padding: 1.5rem;
  border-radius: 16px;
  border: 1px solid rgba(102, 126, 234, 0.2);
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
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 15px 30px rgba(102, 126, 234, 0.2);
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.15) 0%, rgba(118, 75, 162, 0.15) 100%);
}

.feature-card:hover::before {
  left: 100%;
}

.feature-card .feature-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  display: block;
  animation: featureIconFloat 3s ease-in-out infinite;
}

@keyframes featureIconFloat {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-5px); }
}

.feature-card h3 {
  font-size: 1.25rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 0.5rem;
}

.feature-card p {
  font-size: 0.9rem;
  color: #6b7280;
  margin: 0;
}

.steps {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.step {
  display: flex;
  gap: 1rem;
  align-items: flex-start;
}

.step-number {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 1.1rem;
  flex-shrink: 0;
  animation: stepPulse 2s ease-in-out infinite alternate;
}

@keyframes stepPulse {
  0% { box-shadow: 0 0 10px rgba(102, 126, 234, 0.3); }
  100% { box-shadow: 0 0 20px rgba(118, 75, 162, 0.5); }
}

.step-content h4 {
  font-size: 1.1rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 0.5rem;
}

.step-content p {
  margin: 0;
  color: #6b7280;
}

.contact-info {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}

.contact-item {
  display: flex;
  gap: 1rem;
  align-items: flex-start;
  padding: 1.5rem;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
  border-radius: 12px;
  border: 1px solid rgba(102, 126, 234, 0.1);
  transition: all 0.3s;
}

.contact-item:hover {
  transform: translateY(-3px);
  box-shadow: 0 10px 25px rgba(102, 126, 234, 0.15);
}

.contact-icon {
  font-size: 1.5rem;
  animation: contactIconBounce 2s ease-in-out infinite;
}

@keyframes contactIconBounce {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}

.contact-details h4 {
  font-size: 1rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 0.25rem;
}

.contact-details p {
  margin: 0;
  color: #6b7280;
  font-size: 0.9rem;
}

.disclaimer {
  background: linear-gradient(135deg, rgba(251, 191, 36, 0.1) 0%, rgba(245, 158, 11, 0.1) 100%);
  padding: 1.5rem;
  border-radius: 12px;
  border-left: 4px solid #f59e0b;
}

.disclaimer p {
  margin-bottom: 1rem;
  color: #92400e;
  font-size: 0.9rem;
}

.disclaimer p:last-child {
  margin-bottom: 0;
}

/* 移动端适配 */
@media (max-width: 640px) {
  .about-us {
    padding: 0.5rem;
  }

  .title {
    font-size: 2rem;
  }

  .section {
    padding: 1.5rem;
  }

  .feature-grid {
    grid-template-columns: 1fr;
  }

  .contact-info {
    grid-template-columns: 1fr;
  }

  .step {
    flex-direction: column;
    text-align: center;
  }

  .step-number {
    align-self: center;
  }
}
</style>