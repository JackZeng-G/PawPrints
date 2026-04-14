<template>
  <div v-if="loading" class="text-center py-12">
    <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto"></div>
  </div>
  <div v-else-if="!pet" class="text-center py-12 text-gray-500">宠物不存在</div>
  <div v-else class="space-y-6">
    <div class="flex items-center gap-4">
      <router-link to="/pets" class="text-gray-500 hover:text-gray-700">&larr; 返回</router-link>
    </div>

    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="h-48 bg-gray-100 flex items-center justify-center">
        <img v-if="pet.avatar_url" :src="pet.avatar_url" :alt="pet.name" class="w-full h-full object-cover" />
        <span v-else class="text-8xl">🐾</span>
      </div>
      <div class="p-6">
        <div class="flex items-center justify-between mb-4">
          <h1 class="text-2xl font-bold text-gray-800">{{ pet.name }}</h1>
          <div class="flex gap-2">
            <router-link :to="`/pets/${pet.id}/edit`" class="px-4 py-2 border border-gray-300 rounded-lg text-gray-600 hover:bg-gray-50">
              编辑
            </router-link>
            <router-link :to="`/health/${pet.id}`" class="px-4 py-2 border border-blue-300 rounded-lg text-blue-600 hover:bg-blue-50">
              健康记录
            </router-link>
            <button @click="handleDelete" class="px-4 py-2 border border-red-300 rounded-lg text-red-600 hover:bg-red-50">
              删除
            </button>
          </div>
        </div>

        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-sm">
          <div>
            <span class="text-gray-500">种类</span>
            <div class="font-medium">{{ pet.category?.name }}</div>
          </div>
          <div>
            <span class="text-gray-500">品种</span>
            <div class="font-medium">{{ pet.breed?.name || '未知' }}</div>
          </div>
          <div>
            <span class="text-gray-500">性别</span>
            <div class="font-medium">{{ pet.gender }}</div>
          </div>
          <div>
            <span class="text-gray-500">生日</span>
            <div class="font-medium">{{ pet.birthday }}</div>
          </div>
        </div>

        <div v-if="pet.bio" class="mt-4 pt-4 border-t border-gray-100">
          <p class="text-gray-600">{{ pet.bio }}</p>
        </div>

        <div v-if="pet.status === 'memorial' && pet.passing_date" class="mt-4 pt-4 border-t border-gray-100">
          <span class="text-sm text-gray-500">离别日期: {{ pet.passing_date }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { usePetStore } from '../../stores/pet'

const route = useRoute()
const router = useRouter()
const petStore = usePetStore()
const pet = ref<any>(null)
const loading = ref(true)

const handleDelete = async () => {
  if (!confirm('确定要删除这只宠物吗？')) return
  await petStore.deletePet(Number(route.params.id))
  router.push('/pets')
}

onMounted(async () => {
  try {
    await petStore.fetchPet(Number(route.params.id))
    pet.value = petStore.currentPet
  } finally {
    loading.value = false
  }
})
</script>