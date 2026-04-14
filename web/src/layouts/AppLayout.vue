<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Desktop Nav -->
    <nav class="hidden md:block bg-white shadow-sm sticky top-0 z-40">
      <div class="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        <router-link to="/" class="flex items-center gap-2 text-xl font-bold text-orange-500 hover:text-orange-600 transition-colors">
          <PawPrint class="w-6 h-6" />
          <span>PawPrints</span>
        </router-link>
        <div class="flex items-center gap-1">
          <router-link v-for="item in navItems" :key="item.path" :to="item.path"
            class="flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200"
            :class="$route.path === item.path || $route.path.startsWith(item.path + '/') ? 'text-orange-600 bg-orange-50' : 'text-gray-600 hover:text-gray-900 hover:bg-gray-100'">
            <component :is="item.icon" class="w-4 h-4" />
            {{ item.label }}
          </router-link>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto px-4 py-6 pb-24 md:pb-6">
      <router-view />
    </main>

    <!-- Mobile Bottom Nav -->
    <nav class="md:hidden fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 z-50 safe-area-pb">
      <div class="flex items-center justify-around py-2">
        <router-link v-for="item in mobileNavItems" :key="item.path" :to="item.path"
          class="flex flex-col items-center gap-0.5 py-1 px-3 rounded-lg transition-all duration-200 min-w-[64px]"
          :class="isActiveRoute(item.path) ? 'text-orange-500' : 'text-gray-400 hover:text-gray-600'">
          <component :is="item.icon" class="w-5 h-5" :class="{ 'scale-110': isActiveRoute(item.path) }" />
          <span class="text-[10px] font-medium">{{ item.label }}</span>
        </router-link>

        <!-- Floating Add Button -->
        <router-link to="/diaries/new" class="flex flex-col items-center -mt-5">
          <div class="w-12 h-12 bg-orange-500 rounded-full flex items-center justify-center shadow-lg shadow-orange-500/30 hover:bg-orange-600 hover:scale-105 active:scale-95 transition-all duration-200">
            <Plus class="w-6 h-6 text-white" stroke-width="2.5" />
          </div>
        </router-link>

        <router-link v-for="item in mobileNavItems2" :key="item.path" :to="item.path"
          class="flex flex-col items-center gap-0.5 py-1 px-3 rounded-lg transition-all duration-200 min-w-[64px]"
          :class="isActiveRoute(item.path) ? 'text-orange-500' : 'text-gray-400 hover:text-gray-600'">
          <component :is="item.icon" class="w-5 h-5" :class="{ 'scale-110': isActiveRoute(item.path) }" />
          <span class="text-[10px] font-medium">{{ item.label }}</span>
        </router-link>
      </div>
    </nav>

    <!-- Toast Notifications -->
    <ToastNotify />
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { PawPrint, Home, BookOpen, Image, Bell, Settings, Plus, Heart, User } from 'lucide-vue-next'
import ToastNotify from '../components/ToastNotify.vue'

const route = useRoute()

const navItems = [
  { path: '/pets', label: '宠物', icon: Heart },
  { path: '/diaries', label: '日记', icon: BookOpen },
  { path: '/gallery', label: '相册', icon: Image },
  { path: '/reminders', label: '提醒', icon: Bell },
  { path: '/settings', label: '设置', icon: Settings },
]

const mobileNavItems = [
  { path: '/', label: '首页', icon: Home },
  { path: '/diaries', label: '日记', icon: BookOpen },
]

const mobileNavItems2 = [
  { path: '/pets', label: '宠物', icon: Heart },
  { path: '/settings', label: '我的', icon: User },
]

const isActiveRoute = (path: string) => {
  if (path === '/') return route.path === '/'
  return route.path === path || route.path.startsWith(path + '/')
}
</script>

<style scoped>
.safe-area-pb {
  padding-bottom: env(safe-area-inset-bottom, 0);
}
</style>
