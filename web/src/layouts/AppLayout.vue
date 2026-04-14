<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Desktop Nav -->
    <nav class="hidden md:block bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        <router-link to="/" class="text-xl font-bold text-orange-500">🐾 PawPrints</router-link>
        <div class="flex gap-4">
          <router-link to="/pets" class="text-gray-600 hover:text-orange-500 transition-colors">宠物</router-link>
          <router-link to="/diaries" class="text-gray-600 hover:text-orange-500 transition-colors">日记</router-link>
          <router-link to="/gallery" class="text-gray-600 hover:text-orange-500 transition-colors">相册</router-link>
          <router-link to="/reminders" class="text-gray-600 hover:text-orange-500 transition-colors">提醒</router-link>
          <router-link to="/settings" class="text-gray-600 hover:text-orange-500 transition-colors">设置</router-link>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto px-4 py-6 pb-24 md:pb-6">
      <router-view />
    </main>

    <!-- Mobile Bottom Nav -->
    <nav class="md:hidden fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 z-50">
      <div class="flex items-center justify-around py-2">
        <router-link to="/" class="flex flex-col items-center gap-0.5 text-gray-500 hover:text-orange-500" :class="{ 'text-orange-500': $route.path === '/' }">
          <span class="text-xl">🏠</span>
          <span class="text-[10px]">首页</span>
        </router-link>
        <router-link to="/diaries" class="flex flex-col items-center gap-0.5 text-gray-500 hover:text-orange-500" :class="{ 'text-orange-500': $route.path.startsWith('/diaries') }">
          <span class="text-xl">📖</span>
          <span class="text-[10px]">日记</span>
        </router-link>
        <router-link to="/diaries/new" class="flex flex-col items-center -mt-5">
          <div class="w-12 h-12 bg-orange-500 rounded-full flex items-center justify-center shadow-lg hover:bg-orange-600 transition-colors">
            <span class="text-white text-2xl font-light">+</span>
          </div>
        </router-link>
        <router-link to="/pets" class="flex flex-col items-center gap-0.5 text-gray-500 hover:text-orange-500" :class="{ 'text-orange-500': $route.path.startsWith('/pets') }">
          <span class="text-xl">🐾</span>
          <span class="text-[10px]">宠物</span>
        </router-link>
        <router-link to="/settings" class="flex flex-col items-center gap-0.5 text-gray-500 hover:text-orange-500" :class="{ 'text-orange-500': $route.path === '/settings' }">
          <span class="text-xl">⚙️</span>
          <span class="text-[10px]">我的</span>
        </router-link>
      </div>
    </nav>

    <!-- Toast Notifications -->
    <ToastNotify />
  </div>
</template>

<script setup lang="ts">
import ToastNotify from '../components/ToastNotify.vue'
</script>
