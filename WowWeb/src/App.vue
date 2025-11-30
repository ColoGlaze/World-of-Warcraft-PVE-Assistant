<template>
  <div id="app">
    <AnnouncementModal />
    <main class="main-content">
      <router-view />
    </main>
    <BottomNavigation />
  </div>
</template>

<script setup>
import BottomNavigation from './components/BottomNavigation.vue'
import AnnouncementModal from "./components/AnnouncementModal.vue";
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: 'PingFang SC', 'Helvetica Neue', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
  position: relative;
}

#app::before {
  content: '';
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background:
      radial-gradient(circle at 20% 80%, rgba(120, 119, 198, 0.3) 0%, transparent 50%),
      radial-gradient(circle at 80% 20%, rgba(255, 119, 198, 0.3) 0%, transparent 50%),
      radial-gradient(circle at 40% 40%, rgba(120, 219, 255, 0.2) 0%, transparent 50%);
  pointer-events: none;
  z-index: 0;
}

.main-content {
  min-height: 100vh;
  padding-bottom: 4rem;
  position: relative;
  z-index: 1;
}

/* 全局样式重置 */
body {
  margin: 0;
  background: transparent;
}

/* 响应式容器 */
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

/* 移动端适配 */
@media (max-width: 640px) {
  .container {
    padding: 0 0.5rem;
  }

  .main-content {
    padding-bottom: 5rem;
  }
}

/* PC端适配 */
@media (min-width: 1024px) {
  .main-content {
    padding-bottom: 4rem;
  }
}

/* 滚动条样式 */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* 选择文本颜色 */
::selection {
  background: #3b82f6;
  color: white;
}

/* 焦点状态优化 */
*:focus {
  outline: 2px solid #3b82f6;
  outline-offset: 2px;
}

*:focus:not(:focus-visible) {
  outline: none;
}

</style>