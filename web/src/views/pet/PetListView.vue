<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <h1 class="text-2xl font-bold text-gray-800">我的宠物</h1>
      <router-link to="/pets/new" class="bg-orange-500 text-white px-4 py-2 rounded-lg hover:bg-orange-600">
        + 添加宠物
      </router-link>
    </div>

    <div class="flex gap-2">
      <button
        @click="filterStatus = undefined"
        :class="['px-4 py-2 rounded-lg', filterStatus === undefined ? 'bg-orange-500 text-white' : 'bg-gray-100 text-gray-600']"
      >全部</button>
      <button
        @click="filterStatus = 'alive'"
        :class="['px-4 py-2 rounded-lg', filterStatus === 'alive' ? 'bg-orange-500 text-white' : 'bg-gray-100 text-gray-600']"
      >在世</button>
      <button
        @click="filterStatus = 'memorial'"
        :class="['px-4 py-2 rounded-lg', filterStatus === 'memorial' ? 'bg-orange-500 text-white' : 'bg-gray-100 text-gray-600']"
      >纪念</button>
    </div>

    <div v-if="loading" class="text-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto"></div>
    </div>

    <div v-else-if="pets.length === 0" class="text-center py-12 text-gray-500">
      还没有添加宠物，点击上方按钮添加第一个宠物吧！
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <router-link
        v-for="pet in pets"
        :key="pet.id"
        :to="`/pets/${pet.id}`"
        class="block bg-white rounded-lg shadow hover:shadow-md transition overflow-hidden"
      >
        <div class="h-32 bg-gray-100 flex items-center justify-center">
          <img v-if="pet.avatar_url" :src="pet.avatar_url" :alt="pet.name" class="w-full h-full object-cover" />
          <span v-else class="text-6xl">🐾</span>
        </div>
        <div class="p-4">
          <div class="flex items-center justify-between mb-2">
            <span class="text-lg font-bold text-gray-800">{{ pet.name }}</span>
            <span v-if="pet.status === 'memorial'" class="text-xs bg-gray-200 px-2 py-1 rounded">纪念</span>
          </div>
          <div class="text-sm text-gray-500">
            {{ pet.category?.name }} · {{ pet.breed?.name || '未知品种' }}
          </div>
          <div class="text-sm text-gray-500">{{ pet.gender }} · {{ pet.birthday }}</div>
        </div>
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { usePetStore } from '../../stores/pet'

const petStore = usePetStore()
const pets = ref<any[]>([])
const loading = ref(true)
const filterStatus = ref<string | undefined>(undefined)

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