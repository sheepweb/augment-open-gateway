<template>
  <el-config-provider :locale="locale">
    <div id="app">
      <!-- 独立页面：不使用Layout组件，带过渡动画 -->
      <router-view v-if="isStandalonePage" v-slot="{ Component }">
        <transition name="page-fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
      <!-- 管理后台页面：使用Layout组件 -->
      <Layout v-else />
    </div>
  </el-config-provider>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import Layout from '@/components/Layout.vue'
import { systemAPI } from '@/api'

const locale = ref(zhCn)
const route = useRoute()

const loadTurnstileScript = () => {
  if (!window.turnstile && !document.querySelector('script[src*="turnstile"]')) {
    const script = document.createElement('script')
    script.src = 'https://challenges.cloudflare.com/turnstile/v0/api.js'
    script.async = true
    script.defer = true
    document.head.appendChild(script)
  }
}

// 仅在开启 Turnstile 时预加载脚本
onMounted(async () => {
  try {
    const config = await systemAPI.getFrontendConfig()
    window.__APP_FRONTEND_CONFIG__ = config
    if (config?.turnstile_enabled !== false) {
      loadTurnstileScript()
    }
  } catch (error) {
    console.error('加载前端公开配置失败，默认启用 Turnstile:', error)
    window.__APP_FRONTEND_CONFIG__ = {
      turnstile_enabled: true
    }
    loadTurnstileScript()
  }
})

// 计算当前页面是否为独立页面
const isStandalonePage = computed(() => {
  const currentPath = route.path
  // 独立页面：首页、用户登录/注册页、所有用户页面、关于页面
  return currentPath === '/' ||
         currentPath === '/register' ||
         currentPath === '/admin' ||
         currentPath === '/about' ||
         currentPath.startsWith('/user/')
})
</script>

<style>
#app {
  height: 100vh;
  margin: 0;
  padding: 0;
}

/* 页面过渡动画 - 淡入淡出 + 左右滑动 */
.page-fade-enter-active,
.page-fade-leave-active {
  transition: all 0.3s ease;
}

.page-fade-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.page-fade-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}

/* 为特定页面重置样式 */
#app:has(.user-dashboard) {
  height: auto;
  min-height: 100vh;
  background: transparent;
}

* {
  box-sizing: border-box;
}

body {
  margin: 0;
  padding: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
}

/* 为UserDashboard页面重置body样式 */
body:has(.user-dashboard) {
  background: transparent;
  overflow-x: hidden;
}
</style>
