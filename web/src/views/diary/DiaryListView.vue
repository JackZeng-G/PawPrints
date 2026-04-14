<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <h1 class="text-2xl font-bold text-gray-800">宠物日记</h1>
      <router-link to="/diaries/new" class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600">
        + 写日记
      </router-link>
    </div>

    <div class="flex gap-4">
      <input
        v-model="keyword"
        @input="searchDebounced"
        type="text"
        placeholder="搜索日记..."
        class="flex-1 border border-gray-300 rounded-lg px-4 py-2"
      />
    </div>

    <div v-if="loading" class="text-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
    </div>

    <div v-else-if="diaries.length === 0" class="text-center py-12 text-gray-500">
      还没有日记，开始记录吧！
    </div>

    <div v-else class="space-y-4">
      <router-link
        v-for="diary in diaries"
        :key="diary.id"
        :to="`/diaries/${diary.id}`"
        class="block bg-white rounded-lg shadow hover:shadow-md transition p-6"
      >
        <div class="flex items-center justify-between mb-2">
          <h2 class="text-lg font-bold text-gray-800">{{ diary.title }}</h2>
          <span class="text-sm text-gray-400">{{ diary.entry_date }}</span>
        </div>
        <div v-if="diary.pets?.length" class="flex gap-2 mb-2">
          <span v-for="p in diary.pets" :key="p.pet_id" class="text-xs bg-orange-100 text-orange-600 px-2 py-1 rounded">
            {{ p.pet_name }}
          </span>
        </div>
        <p class="text-gray-600 line-clamp-2">{{ diary.content }}</p>
        <div v-if="diary.photos?.length" class="flex gap-2 mt-3">
          <img
            v-for="photo in diary.photos.slice(0, 4)"
            :key="photo.id"
            :src="photo.thumbnail_url || photo.photo_url"
            class="w-16 h-16 rounded object-cover"
          />
        </div>
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useDiaryStore } from '../../stores/diary'

const diaryStore = useDiaryStore()
const diaries = ref<any[]>([])
const loading = ref(true)
const keyword = ref('')
let timer: ReturnType<typeof setTimeout>

const searchDebounced = () => {
  clearTimeout(timer)
  timer = setTimeout(() => loadData(), 300)
}

const loadData = async () => {
  loading.value = true
  await diaryStore.fetchDiaries(undefined, keyword.value || undefined)
  diaries.value = diaryStore.diaries
  loading.value = false
}

onMounted(loadData)
</script>