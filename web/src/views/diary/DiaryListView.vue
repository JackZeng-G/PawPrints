<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <h1 class="text-2xl font-bold text-gray-800">宠物日记</h1>
      <router-link to="/diaries/new"
        class="inline-flex items-center gap-2 bg-blue-500 text-white px-4 py-2.5 rounded-xl hover:bg-blue-600 active:scale-[0.98] transition-all duration-200 shadow-sm shadow-blue-500/20 font-medium text-sm">
        <Plus class="w-4 h-4" /> 写日记
      </router-link>
    </div>

    <!-- Search -->
    <div class="relative">
      <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-gray-400" />
      <input v-model="keyword" @input="searchDebounced" type="text" placeholder="搜索日记..."
        class="w-full border border-gray-200 rounded-xl pl-10 pr-4 py-2.5 bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all" />
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center py-16 gap-4">
      <div class="w-10 h-10 border-3 border-gray-200 border-t-blue-500 rounded-full animate-spin"></div>
      <p class="text-sm text-gray-400">加载中...</p>
    </div>

    <!-- Empty State -->
    <div v-else-if="diaries.length === 0" class="flex flex-col items-center py-16 text-center">
      <div class="w-20 h-20 bg-blue-50 rounded-full flex items-center justify-center mb-4">
        <BookOpen class="w-10 h-10 text-blue-300" />
      </div>
      <p class="text-gray-500 mb-1">还没有日记</p>
      <p class="text-sm text-gray-400 mb-4">开始记录宠物的美好瞬间吧</p>
      <router-link to="/diaries/new" class="inline-flex items-center gap-2 text-blue-500 hover:text-blue-600 font-medium text-sm">
        <Plus class="w-4 h-4" /> 写第一篇日记
      </router-link>
    </div>

    <!-- Diary List -->
    <div v-else class="space-y-4">
      <router-link v-for="diary in diaries" :key="diary.id" :to="`/diaries/${diary.id}`"
        class="block bg-white rounded-2xl border border-gray-100 p-5 hover:shadow-md hover:-translate-y-0.5 transition-all duration-200 group">
        <div class="flex items-center justify-between mb-2">
          <h2 class="text-lg font-bold text-gray-800 group-hover:text-blue-500 transition-colors">{{ diary.title }}</h2>
          <span class="text-xs text-gray-400 flex-shrink-0">{{ diary.entry_date }}</span>
        </div>
        <div v-if="diary.pets?.length" class="flex gap-1.5 mb-2">
          <span v-for="p in diary.pets" :key="p.pet_id"
            class="text-xs bg-orange-50 text-orange-600 px-2 py-0.5 rounded-md font-medium">
            {{ p.pet_name }}
          </span>
        </div>
        <p class="text-sm text-gray-500 line-clamp-2">{{ diary.content }}</p>
        <div v-if="diary.photos?.length" class="flex gap-2 mt-3">
          <img v-for="photo in diary.photos.slice(0, 4)" :key="photo.id" :src="photo.thumbnail_url || photo.photo_url"
            class="w-14 h-14 rounded-xl object-cover" />
          <div v-if="diary.photos.length > 4"
            class="w-14 h-14 rounded-xl bg-gray-100 flex items-center justify-center text-xs text-gray-500 font-medium">
            +{{ diary.photos.length - 4 }}
          </div>
        </div>
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useDiaryStore } from '../../stores/diary'
import { Plus, Search, BookOpen } from 'lucide-vue-next'

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