<template>
  <div class="space-y-8">
    <div class="text-center py-8">
      <h1 class="text-3xl font-bold text-gray-800 mb-2">🐾 萌宠档案馆</h1>
      <p class="text-gray-500">记录宠物的每一个美好瞬间</p>
    </div>

    <div v-if="loading" class="text-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto"></div>
      <p class="mt-4 text-gray-500">加载中...</p>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-4xl mb-2">🐾</div>
        <div class="text-2xl font-bold text-gray-800">{{ stats.total_pets }}</div>
        <div class="text-gray-500">宠物总数</div>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-4xl mb-2">📖</div>
        <div class="text-2xl font-bold text-gray-800">{{ stats.total_diaries }}</div>
        <div class="text-gray-500">日记总数</div>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <div class="text-4xl mb-2">🔔</div>
        <div class="text-2xl font-bold text-gray-800">{{ stats.active_reminders }}</div>
        <div class="text-gray-500">待办提醒</div>
      </div>
    </div>

    <div v-if="pets.length" class="bg-white rounded-lg shadow p-6">
      <h2 class="text-xl font-bold text-gray-800 mb-4">我的宠物</h2>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <router-link
          v-for="pet in pets"
          :key="pet.id"
          :to="`/pets/${pet.id}`"
          class="block p-4 rounded-lg border border-gray-200 hover:border-orange-300 hover:shadow transition"
        >
          <div v-if="pet.avatar_url" class="w-16 h-16 rounded-full mx-auto mb-2 overflow-hidden">
            <img :src="pet.avatar_url" :alt="pet.name" class="w-full h-full object-cover" />
          </div>
          <div v-else class="w-16 h-16 rounded-full mx-auto mb-2 bg-gray-100 flex items-center justify-center text-2xl">
            🐾
          </div>
          <div class="text-center font-medium text-gray-800">{{ pet.name }}</div>
          <div class="text-center text-sm text-gray-500">{{ pet.category?.name }}</div>
        </router-link>
      </div>
    </div>

    <div class="flex justify-center gap-4">
      <router-link to="/pets/new" class="bg-orange-500 text-white px-6 py-3 rounded-lg hover:bg-orange-600 transition">
        + 添加宠物
      </router-link>
      <router-link to="/diaries/new" class="bg-blue-500 text-white px-6 py-3 rounded-lg hover:bg-blue-600 transition">
        + 写日记
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { usePetStore } from '../../stores/pet'
import { statsApi } from '../../api/common'

const petStore = usePetStore()
const pets = ref<any[]>([])
const stats = ref({ total_pets: 0, total_diaries: 0, active_reminders: 0 })
const loading = ref(true)

onMounted(async () => {
  try {
    await petStore.fetchPets()
    pets.value = petStore.pets.slice(0, 8)
    const { data: statsData } = await statsApi.overview()
    stats.value = statsData
  } finally {
    loading.value = false
  }
})
</script>