<template>
  <div class="profile">
    <!-- 未登录状态 -->
    <div v-if="!isLoggedIn" class="auth-container">
      <!-- 头部 -->
      <header class="header">
        <div class="user-icon">👤</div>
        <h1 class="title">我的账户</h1>
        <p class="subtitle">登录后享受更多功能</p>
      </header>

      <!-- 登录/注册切换 -->
      <div class="auth-tabs">
        <button
            class="tab-btn"
            :class="{ active: authMode === 'login' }"
            @click="authMode = 'login'"
        >
          登录
        </button>
        <button
            class="tab-btn"
            :class="{ active: authMode === 'register' }"
            @click="authMode = 'register'"
        >
          注册
        </button>
      </div>

      <!-- 登录表单 -->
      <div v-if="authMode === 'login'" class="auth-form">
        <h2 class="form-title">欢迎回来</h2>
        <form @submit.prevent="handleLogin">
          <div class="form-group">
            <label for="loginUsername">账号</label>
            <input
                id="loginUsername"
                v-model="loginForm.username"
                type="text"
                placeholder="请输入账号"
                class="form-input"
                required
            />
          </div>
          <div class="form-group">
            <label for="loginPassword">密码</label>
            <input
                id="loginPassword"
                v-model="loginForm.password"
                type="password"
                placeholder="请输入密码"
                class="form-input"
                required
            />
          </div>
          <button
              type="submit"
              class="auth-btn"
              :disabled="loading"
          >
            <span v-if="loading" class="loading-spinner"></span>
            {{ loading ? '登录中...' : '登录' }}
          </button>
        </form>
      </div>

      <!-- 注册表单 -->
      <div v-if="authMode === 'register'" class="auth-form">
        <h2 class="form-title">创建账户</h2>
        <form @submit.prevent="handleRegister">
          <div class="form-group">
            <label for="registerUsername">账号</label>
            <input
                id="registerUsername"
                v-model="registerForm.username"
                type="text"
                placeholder="请输入账号"
                class="form-input"
                required
            />
          </div>
          <div class="form-group">
            <label for="registerPassword">密码</label>
            <input
                id="registerPassword"
                v-model="registerForm.password"
                type="password"
                placeholder="请输入密码"
                class="form-input"
                required
            />
          </div>
          <div class="form-group">
            <label for="confirmPassword">确认密码</label>
            <input
                id="confirmPassword"
                v-model="registerForm.confirmPassword"
                type="password"
                placeholder="请再次输入密码"
                class="form-input"
                required
            />
          </div>
          <button
              type="submit"
              class="auth-btn"
              :disabled="loading"
          >
            <span v-if="loading" class="loading-spinner"></span>
            {{ loading ? '注册中...' : '注册' }}
          </button>
        </form>
      </div>
    </div>

    <!-- 已登录状态 -->
    <div v-else class="user-info-container">
      <!-- 用户信息卡片 -->
      <div class="user-card">
        <div class="user-avatar">
          <div class="avatar-icon">👤</div>
        </div>
        <div class="user-details">
          <h2 class="username">{{ userInfo.username }}</h2>
          <p class="user-level">普通用户</p>
          <p class="join-date">加入时间：{{ userInfo.joinDate }}</p>
        </div>
      </div>

      <!-- 功能菜单 -->
      <div class="menu-section">
        <div class="menu-item">
          <div class="menu-icon">📊</div>
          <div class="menu-content">
            <h3 class="menu-title">查询历史</h3>
            <p class="menu-desc">查看历史查询记录</p>
          </div>
          <div class="menu-arrow">›</div>
        </div>

        <div class="menu-item">
          <div class="menu-icon">⭐</div>
          <div class="menu-content">
            <h3 class="menu-title">我的收藏</h3>
            <p class="menu-desc">收藏的角色和记录</p>
          </div>
          <div class="menu-arrow">›</div>
        </div>

        <div class="menu-item">
          <div class="menu-icon">⚙️</div>
          <div class="menu-content">
            <h3 class="menu-title">设置</h3>
            <p class="menu-desc">个人设置和偏好</p>
          </div>
          <div class="menu-arrow">›</div>
        </div>

        <div class="menu-item">
          <div class="menu-icon">❓</div>
          <div class="menu-content">
            <h3 class="menu-title">帮助与反馈</h3>
            <p class="menu-desc">使用帮助和问题反馈</p>
          </div>
          <div class="menu-arrow">›</div>
        </div>
      </div>

      <!-- 退出登录 -->
      <div class="logout-section">
        <button class="logout-btn" @click="handleLogout">
          <span class="logout-icon">🚪</span>
          退出登录
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

// 定义组件名称
defineOptions({
  name: 'Profile'
})

// 响应式数据
const isLoggedIn = ref(false)
const authMode = ref('login')
const loading = ref(false)

const loginForm = reactive({
  username: '',
  password: ''
})

const registerForm = reactive({
  username: '',
  password: '',
  confirmPassword: ''
})

const userInfo = reactive({
  username: '',
  joinDate: ''
})

// 登录处理
const handleLogin = async () => {
  if (!loginForm.username || !loginForm.password) {
    ElMessage.error('请填写完整的登录信息')
    return
  }

  loading.value = true

  try {
    // 模拟登录API调用
    await new Promise(resolve => setTimeout(resolve, 1500))

    // 模拟登录成功
    userInfo.username = loginForm.username
    userInfo.joinDate = new Date().toLocaleDateString('zh-CN')
    isLoggedIn.value = true

    ElMessage.success('登录成功！')

    // 清空表单
    loginForm.username = ''
    loginForm.password = ''
  } catch (error) {
    ElMessage.error('登录失败，请检查账号密码')
  } finally {
    loading.value = false
  }
}

// 注册处理
const handleRegister = async () => {
  if (!registerForm.username || !registerForm.password || !registerForm.confirmPassword) {
    ElMessage.error('请填写完整的注册信息')
    return
  }

  if (registerForm.password !== registerForm.confirmPassword) {
    ElMessage.error('两次输入的密码不一致')
    return
  }

  if (registerForm.password.length < 6) {
    ElMessage.error('密码长度不能少于6位')
    return
  }

  loading.value = true

  try {
    // 模拟注册API调用
    await new Promise(resolve => setTimeout(resolve, 1500))

    ElMessage.success('注册成功！请登录')

    // 切换到登录模式并填入用户名
    authMode.value = 'login'
    loginForm.username = registerForm.username

    // 清空注册表单
    registerForm.username = ''
    registerForm.password = ''
    registerForm.confirmPassword = ''
  } catch (error) {
    ElMessage.error('注册失败，请稍后重试')
  } finally {
    loading.value = false
  }
}

// 退出登录
const handleLogout = () => {
  isLoggedIn.value = false
  userInfo.username = ''
  userInfo.joinDate = ''
  ElMessage.success('已退出登录')
}
</script>

<style scoped>
.profile {
  padding: 1rem;
  max-width: 600px;
  margin: 0 auto;
  padding-bottom: 5rem;
}

.auth-container {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.header {
  text-align: center;
  padding: 2rem 0;
}

.user-icon {
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

.auth-tabs {
  display: flex;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 16px;
  padding: 0.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.tab-btn {
  flex: 1;
  padding: 0.75rem 1rem;
  border: none;
  background: transparent;
  border-radius: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s;
  color: #6b7280;
}

.tab-btn.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
}

.auth-form {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 2rem;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.form-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 1.5rem;
  text-align: center;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
  margin-bottom: 0.5rem;
}

.form-input {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid rgba(102, 126, 234, 0.2);
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.3s;
  background: rgba(255, 255, 255, 0.8);
  box-sizing: border-box;
}

.form-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2);
  background: rgba(255, 255, 255, 1);
}

.auth-btn {
  width: 100%;
  padding: 0.875rem 1.5rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.auth-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.4);
}

.auth-btn:disabled {
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
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.user-info-container {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.user-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  padding: 2rem;
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.user-avatar {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.avatar-icon {
  font-size: 2.5rem;
  color: white;
}

.user-details {
  flex: 1;
}

.username {
  font-size: 1.5rem;
  font-weight: 600;
  color: #667eea;
  margin-bottom: 0.25rem;
}

.user-level {
  font-size: 0.9rem;
  color: #10b981;
  margin-bottom: 0.25rem;
  font-weight: 500;
}

.join-date {
  font-size: 0.8rem;
  color: #6b7280;
  margin: 0;
}

.menu-section {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 20px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  overflow: hidden;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem;
  cursor: pointer;
  transition: all 0.3s;
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
}

.menu-item:last-child {
  border-bottom: none;
}

.menu-item:hover {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.05) 0%, rgba(118, 75, 162, 0.05) 100%);
  transform: translateX(5px);
}

.menu-icon {
  font-size: 1.5rem;
  width: 40px;
  text-align: center;
}

.menu-content {
  flex: 1;
}

.menu-title {
  font-size: 1rem;
  font-weight: 500;
  color: #374151;
  margin-bottom: 0.25rem;
}

.menu-desc {
  font-size: 0.8rem;
  color: #6b7280;
  margin: 0;
}

.menu-arrow {
  font-size: 1.2rem;
  color: #6b7280;
}

.logout-section {
  text-align: center;
}

.logout-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 2rem;
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
  border: 1px solid rgba(239, 68, 68, 0.2);
  border-radius: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s;
}

.logout-btn:hover {
  background: rgba(239, 68, 68, 0.2);
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(239, 68, 68, 0.2);
}

.logout-icon {
  font-size: 1.1rem;
}

/* 移动端适配 */
@media (max-width: 640px) {
  .profile {
    padding: 0.5rem;
  }

  .user-card {
    flex-direction: column;
    text-align: center;
  }

  .menu-item {
    padding: 1rem;
  }
}
</style>