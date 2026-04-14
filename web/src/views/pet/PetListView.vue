<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <h1 class="text-2xl font-bold text-gray-800">我的宠物</h1>
      <router-link to="/pets/new" class="inline-flex items-center gap-2 bg-orange-500 text-white px-4 py-2 rounded-xl hover:bg-orange-600 active:scale-[0.98] transition-all duration-200 shadow-sm shadow-orange-500/20 font-medium text-sm">
        <Plus class="w-4 h-4" /> 添加宠物
      </router-link>
    </div>

    <!-- Filter Tabs -->
    <div class="flex gap-2 bg-white p-1.5 rounded-xl border border-gray-100 shadow-sm">
      <button v-for="tab in tabs" :key="tab.value"
        @click="filterStatus = tab.value"
        class="flex-1 px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200"
        :class="filterStatus === tab.value ? 'bg-orange-500 text-white shadow-sm' : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'"
      >{{ tab.label }}</button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center py-16 gap-4">
      <div class="w-10 h-10 border-3 border-gray-200 border-t-orange-500 rounded-full animate-spin"></div>
      <p class="text-sm text-gray-400">加载中...</p>
    </div>

    <!-- Empty State -->
    <div v-else-if="pets.length === 0" class="flex flex-col items-center py-16 text-center">
      <div class="w-20 h-20 bg-orange-50 rounded-full flex items-center justify-center mb-4">
        <PawPrint class="w-10 h-10 text-orange-300" />
      </div>
      <p class="text-gray-500 mb-1">还没有添加宠物</p>
      <p class="text-sm text-gray-400 mb-4">点击上方按钮添加第一个宠物吧！</p>
      <router-link to="/pets/new" class="inline-flex items-center gap-2 text-orange-500 hover:text-orange-600 font-medium text-sm">
        <Plus class="w-4 h-4" /> 立即添加
      </router-link>
    </div>

    <!-- Pet Grid -->
    <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <router-link
        v-for="pet in pets"
        :key="pet.id"
        :to="`/pets/${pet.id}`"
        class="group bg-white rounded-2xl border border-gray-100 overflow-hidden hover:shadow-lg hover:-translate-y-0.5 transition-all duration-200"
      >
        <div class="h-40 bg-gradient-to-br from-orange-50 to-amber-50 flex items-center justify-center relative overflow-hidden">
          <img v-if="pet.avatar_url" :src="pet.avatar_url" :alt="pet.name" class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300" />
          <PawPrint v-else class="w-16 h-16 text-orange-200" />
          <span v-if="pet.status === 'memorial'" class="absolute top-3 right-3 flex items-center gap-1 text-xs bg-gray-800/60 backdrop-blur-sm text-white px-2.5 py-1 rounded-full">
            <Flower2 class="w-3 h-3" /> 纪念
          </span>
        </div>
        <div class="p-4">
          <div class="flex items-center justify-between mb-2">
            <span class="text-lg font-bold text-gray-800">{{ pet.name }}</span>
          </div>
          <div class="flex items-center gap-2 text-sm text-gray-500">
            <span class="bg-gray-50 px-2 py-0.5 rounded-md text-xs">{{ pet.category?.name }}</span>
            <span class="text-gray-300">·</span>
            <span>{{ pet.breed?.name || '未知品种' }}</span>
          </div>
          <div class="flex items-center gap-3 mt-3 text-xs text-gray-400">
            <span class="flex items-center gap-1">
              <component :is="pet.gender === '公' ? Mars : Venus" class="w-3 h-3" />
              {{ pet.gender }}
            </span>
            <span class="flex items-center gap-1">
              <Calendar class="w-3 h-3" />
              {{ pet.birthday || '未知生日' }}
            </span>
          </div>
        </div>
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { usePetStore } from '../../stores/pet'
import { Plus, PawPrint, Flower2, Mars, Venus, Calendar } from 'lucide-vue-next'

const petStore = usePetStore()
const pets = ref<any[]>([])
const loading = ref(true)
const filterStatus = ref<string | undefined>(undefined)

const tabs = [
  { value: undefined, label: '全部' },
  { value: 'alive', label: '在世' },
  { value: 'memorial', label: '纪念' },
]

watch(filterStatus, async () => {
  loading.value = true
  await petStore.fetchPets(filterStatus.value)
  pets.value = petStore.pets
  loading.value = false
})

onMounted(async () => {
  await petStore.fetchPets()
  pets.value = petStore.pets
  loading.value = false
})
</script>
