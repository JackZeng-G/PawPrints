<template>
  <div v-if="loading" class="text-center py-12">
    <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
  </div>
  <div v-else-if="!diary" class="text-center py-12 text-gray-500">日记不存在</div>
  <div v-else class="max-w-3xl mx-auto space-y-6">
    <div class="flex items-center gap-4">
      <router-link to="/diaries" class="text-gray-500 hover:text-gray-700">&larr; 返回</router-link>
    </div>

    <div class="bg-white rounded-lg shadow p-6">
      <div class="flex items-center justify-between mb-4">
        <h1 class="text-2xl font-bold text-gray-800">{{ diary.title }}</h1>
        <div class="flex gap-2">
          <router-link :to="`/diaries/${diary.id}/edit`" class="px-4 py-2 border border-gray-300 rounded-lg text-gray-600 hover:bg-gray-50">
            编辑
          </router-link>
          <button @click="handleDelete" class="px-4 py-2 border border-red-300 rounded-lg text-red-600 hover:bg-red-50">
            删除
          </button>
        </div>
      </div>

      <div class="flex items-center gap-4 text-sm text-gray-500 mb-4">
        <span>{{ diary.entry_date }}</span>
        <span v-if="diary.mood">心情: {{ diary.mood }}</span>
      </div>

      <div v-if="diary.pets?.length" class="flex gap-2 mb-4">
        <span v-for="p in diary.pets" :key="p.pet_id" class="text-sm bg-orange-100 text-orange-600 px-3 py-1 rounded-full">
          {{ p.pet_name }}
        </span>
      </div>

      <div class="prose max-w-none text-gray-700 whitespace-pre-wrap">{{ diary.content }}</div>

      <div v-if="diary.photos?.length" class="grid grid-cols-3 gap-4 mt-6">
        <img
          v-for="photo in diary.photos"
          :key="photo.id"
          :src="photo.photo_url"
          class="w-full rounded-lg object-cover"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useDiaryStore } from '../../stores/diary'

const route = useRoute()
const router = useRouter()
const diaryStore = useDiaryStore()
const diary = ref<any>(null)
const loading = ref(true)

const handleDelete = async () => {
  if (!confirm('确定要删除这篇日记吗？')) return
  await diaryStore.deleteDiary(Number(route.params.id))
  router.push('/diaries')
}

onMounted(async () => {
  try {
    await diaryStore.fetchDiary(Number(route.params.id))
    diary.value = diaryStore.currentDiary
  } finally {
    loading.value = false
  }
})
</script>