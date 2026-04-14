<template>
  <div class="min-h-screen bg-gray-50">
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        <router-link to="/" class="text-xl font-bold text-orange-500">🐾 PawPrints</router-link>
        <div class="flex gap-4">
          <router-link to="/pets" class="text-gray-600 hover:text-orange-500">宠物</router-link>
          <router-link to="/diaries" class="text-gray-600 hover:text-orange-500">日记</router-link>
          <router-link to="/reminders" class="text-gray-600 hover:text-orange-500">提醒</router-link>
          <router-link to="/settings" class="text-gray-600 hover:text-orange-500">设置</router-link>
        </div>
      </div>
    </nav>
    <main class="max-w-7xl mx-auto px-4 py-6">
      <router-view />
    </main>
  </div>
</template>
