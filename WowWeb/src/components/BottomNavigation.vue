<template>
  <nav class="bottom-navigation">
    <router-link
        to="/"
        class="nav-item"
        :class="{ active: $route.name === 'Home' }"
    >
      <div class="nav-icon">🏠</div>
      <span class="nav-label">首页</span>
    </router-link>

    <router-link
        to="/profile"
        class="nav-item"
        :class="{ active: $route.name === 'Profile' }"
    >
      <div class="nav-icon">👤</div>
      <span class="nav-label">我的</span>
    </router-link>
  </nav>
</template>

<script setup>
// 底部导航组件
</script>

<style scoped>
.bottom-navigation {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-top: 1px solid rgba(102, 126, 234, 0.2);
  display: flex;
  padding: 0.5rem 0 calc(0.5rem + env(safe-area-inset-bottom));
  box-shadow:
      0 -8px 32px rgba(0, 0, 0, 0.1),
      0 0 0 1px rgba(255, 255, 255, 0.2);
  z-index: 1000;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-bottom: none;
}

.nav-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0.5rem 1rem;
  text-decoration: none;
  color: #667eea;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  border-radius: 12px;
  margin: 0 0.25rem;
}

.nav-item:hover {
  color: #764ba2;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  transform: translateY(-2px);
}

.nav-item.active {
  color: #764ba2;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.15) 0%, rgba(118, 75, 162, 0.15) 100%);
}

.nav-item.active::before {
  content: '';
  position: absolute;
  top: -1px;
  left: 50%;
  transform: translateX(-50%);
  width: 32px;
  height: 4px;
  background: linear-gradient(90deg, #667eea, #764ba2);
  border-radius: 2px;
  animation: activeGlow 2s ease-in-out infinite alternate;
}

@keyframes activeGlow {
  0% { box-shadow: 0 0 5px rgba(102, 126, 234, 0.5); }
  100% { box-shadow: 0 0 15px rgba(118, 75, 162, 0.8); }
}

.nav-icon {
  font-size: 1.5rem;
  margin-bottom: 0.25rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

.nav-item.active .nav-icon {
  transform: scale(1.15);
  animation: iconPulse 2s ease-in-out infinite;
}

.nav-item:hover .nav-icon {
  transform: scale(1.1) translateY(-1px);
}

@keyframes iconPulse {
  0%, 100% { filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1)); }
  50% { filter: drop-shadow(0 4px 8px rgba(102, 126, 234, 0.3)); }
}

.nav-label {
  font-size: 0.75rem;
  font-weight: 500;
  line-height: 1;
  transition: all 0.3s;
}

.nav-item.active .nav-label {
  font-weight: 600;
}

/* 移动端优化 */
@media (max-width: 640px) {
  .bottom-navigation {
    padding: 0.75rem 0 calc(0.75rem + env(safe-area-inset-bottom));
  }

  .nav-item {
    padding: 0.5rem 0.5rem;
    margin: 0 0.125rem;
  }

  .nav-icon {
    font-size: 1.25rem;
  }

  .nav-label {
    font-size: 0.6875rem;
  }
}

/* PC端优化 */
@media (min-width: 1024px) {
  .bottom-navigation {
    max-width: 800px;
    left: 50%;
    transform: translateX(-50%);
    border-radius: 20px 20px 0 0;
  }
}
</style>