<template>
  <div class="space-y-8">
    <!-- Header -->
    <div class="text-center py-6">
      <div class="flex items-center justify-center gap-2 mb-2">
        <PawPrint class="w-8 h-8 text-orange-500" />
        <h1 class="text-3xl font-bold text-gray-800">萌宠档案馆</h1>
      </div>
      <p class="text-gray-500 text-sm">记录宠物的每一个美好瞬间</p>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center py-16 gap-4">
      <div class="w-10 h-10 border-3 border-gray-200 border-t-orange-500 rounded-full animate-spin"></div>
      <p class="text-sm text-gray-400">加载中...</p>
    </div>

    <template v-else>
      <!-- Stats Cards -->
      <div class="grid grid-cols-3 gap-4">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 hover:shadow-md transition-shadow duration-200">
          <div class="w-10 h-10 bg-orange-50 rounded-xl flex items-center justify-center mb-3">
            <Heart class="w-5 h-5 text-orange-500" />
          </div>
          <div class="text-2xl font-bold text-gray-800">{{ stats.total_pets }}</div>
          <div class="text-xs text-gray-400 mt-0.5">宠物总数</div>
        </div>
        <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 hover:shadow-md transition-shadow duration-200">
          <div class="w-10 h-10 bg-blue-50 rounded-xl flex items-center justify-center mb-3">
            <BookOpen class="w-5 h-5 text-blue-500" />
          </div>
          <div class="text-2xl font-bold text-gray-800">{{ stats.total_diaries }}</div>
          <div class="text-xs text-gray-400 mt-0.5">日记总数</div>
        </div>
        <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 hover:shadow-md transition-shadow duration-200">
          <div class="w-10 h-10 bg-amber-50 rounded-xl flex items-center justify-center mb-3">
            <Bell class="w-5 h-5 text-amber-500" />
          </div>
          <div class="text-2xl font-bold text-gray-800">{{ stats.active_reminders }}</div>
          <div class="text-xs text-gray-400 mt-0.5">待办提醒</div>
        </div>
      </div>

      <!-- My Pets -->
      <div v-if="pets.length" class="bg-white rounded-2xl shadow-sm border border-gray-100 p-6">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-bold text-gray-800">我的宠物</h2>
          <router-link to="/pets" class="text-sm text-orange-500 hover:text-orange-600 font-medium flex items-center gap-1 transition-colors">
            查看全部 <ChevronRight class="w-4 h-4" />
          </router-link>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <router-link
            v-for="pet in pets"
            :key="pet.id"
            :to="`/pets/${pet.id}`"
            class="flex flex-col items-center p-4 rounded-xl hover:bg-orange-50 transition-all duration-200 group"
          >
            <div v-if="pet.avatar_url" class="w-16 h-16 rounded-full mb-2 overflow-hidden ring-2 ring-transparent group-hover:ring-orange-200 transition-all">
              <img :src="pet.avatar_url" :alt="pet.name" class="w-full h-full object-cover" />
            </div>
            <div v-else class="w-16 h-16 rounded-full mx-auto mb-2 bg-orange-50 flex items-center justify-center">
              <PawPrint class="w-7 h-7 text-orange-400" />
            </div>
            <div class="text-center font-medium text-gray-800 text-sm">{{ pet.name }}</div>
            <div class="text-center text-xs text-gray-400">{{ pet.category?.name }}</div>
          </router-link>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="flex justify-center gap-3">
        <router-link to="/pets/new" class="inline-flex items-center gap-2 bg-orange-500 text-white px-5 py-2.5 rounded-xl hover:bg-orange-600 active:scale-[0.98] transition-all duration-200 shadow-sm shadow-orange-500/20 font-medium text-sm">
          <Plus class="w-4 h-4" /> 添加宠物
        </router-link>
        <router-link to="/diaries/new" class="inline-flex items-center gap-2 bg-white text-gray-700 px-5 py-2.5 rounded-xl hover:bg-gray-50 active:scale-[0.98] transition-all duration-200 border border-gray-200 font-medium text-sm">
          <BookOpen class="w-4 h-4" /> 写日记
        </router-link>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { usePetStore } from '../../stores/pet'
import { statsApi } from '../../api/common'
import { PawPrint, Heart, BookOpen, Bell, Plus, ChevronRight } from 'lucide-vue-next'

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
