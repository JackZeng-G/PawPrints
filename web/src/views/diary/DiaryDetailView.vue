<template>
  <div v-if="loading" class="flex flex-col items-center py-16 gap-4">
    <div class="w-10 h-10 border-3 border-gray-200 border-t-blue-500 rounded-full animate-spin"></div>
    <p class="text-sm text-gray-400">加载中...</p>
  </div>

  <div v-else-if="!diary" class="flex flex-col items-center py-16 text-center">
    <div class="w-20 h-20 bg-gray-100 rounded-full flex items-center justify-center mb-4">
      <BookOpen class="w-10 h-10 text-gray-300" />
    </div>
    <p class="text-gray-500">日记不存在</p>
    <router-link to="/diaries" class="mt-4 text-blue-500 hover:text-blue-600 font-medium text-sm">返回日记列表</router-link>
  </div>

  <div v-else class="max-w-3xl mx-auto space-y-6">
    <!-- Back Link -->
    <router-link to="/diaries" class="inline-flex items-center gap-1 text-sm text-gray-500 hover:text-gray-700 transition-colors">
      <ChevronLeft class="w-4 h-4" /> 返回
    </router-link>

    <!-- Diary Card -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 p-6">
      <!-- Header -->
      <div class="flex items-center justify-between mb-4">
        <h1 class="text-xl font-bold text-gray-800">{{ diary.title }}</h1>
        <div class="flex gap-2">
          <router-link :to="`/diaries/${diary.id}/edit`"
            class="inline-flex items-center gap-1.5 px-3 py-1.5 text-sm font-medium border border-gray-200 rounded-xl text-gray-600 hover:bg-gray-50 hover:border-gray-300 transition-all"
          >
            <Pencil class="w-3.5 h-3.5" /> 编辑
          </router-link>
          <button @click="handleDelete"
            class="inline-flex items-center gap-1.5 px-3 py-1.5 text-sm font-medium border border-red-200 rounded-xl text-red-500 hover:bg-red-50 hover:border-red-300 transition-all"
          >
            <Trash2 class="w-3.5 h-3.5" /> 删除
          </button>
        </div>
      </div>

      <!-- Meta Info -->
      <div class="flex items-center gap-4 text-sm text-gray-400 mb-4">
        <span class="flex items-center gap-1.5">
          <Calendar class="w-3.5 h-3.5" /> {{ diary.entry_date }}
        </span>
        <span v-if="diary.mood" class="flex items-center gap-1.5 bg-gray-50 px-2 py-0.5 rounded-md">
          {{ moodIcon(diary.mood) }} {{ moodLabel(diary.mood) }}
        </span>
      </div>

      <!-- Pets -->
      <div v-if="diary.pets?.length" class="flex gap-2 mb-4">
        <router-link v-for="p in diary.pets" :key="p.pet_id" :to="`/pets/${p.pet_id}`"
          class="text-sm bg-orange-50 text-orange-600 px-3 py-1 rounded-full hover:bg-orange-100 transition-colors"
        >
          {{ p.pet_name }}
        </router-link>
      </div>

      <!-- Content -->
      <div class="prose max-w-none text-gray-700 whitespace-pre-wrap leading-relaxed">{{ diary.content }}</div>

      <!-- Photos -->
      <div v-if="diary.photos?.length" class="grid grid-cols-3 gap-3 mt-6">
        <img v-for="photo in diary.photos" :key="photo.id" :src="photo.photo_url"
          class="w-full aspect-square object-cover rounded-xl cursor-pointer hover:opacity-90 transition-opacity ring-2 ring-transparent hover:ring-blue-200"
          @click="lightboxUrl = photo.photo_url" />
      </div>
    </div>

    <!-- Lightbox -->
    <Teleport to="body">
      <Transition name="lightbox">
        <div v-if="lightboxUrl" class="fixed inset-0 bg-black/90 z-[100] flex items-center justify-center backdrop-blur-sm"
          @click="lightboxUrl = ''">
          <button @click="lightboxUrl = ''" class="absolute top-4 right-4 w-10 h-10 bg-white/10 rounded-full flex items-center justify-center hover:bg-white/20 transition-colors"
          >
            <X class="w-5 h-5 text-white" />
          </button>
          <img :src="lightboxUrl" class="max-w-[90vw] max-h-[90vh] object-contain rounded-lg shadow-2xl" @click.stop />
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useDiaryStore } from '../../stores/diary'
import { BookOpen, ChevronLeft, Pencil, Trash2, Calendar, X } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const diaryStore = useDiaryStore()
const diary = ref<any>(null)
const loading = ref(true)
const lightboxUrl = ref('')

const moodIcon = (mood: string) => ({
  happy: '😊', calm: '😌', sad: '😢', excited: '🤗', worried: '😟'
}[mood] || '')

const moodLabel = (mood: string) => ({
  happy: '开心', calm: '平静', sad: '难过', excited: '兴奋', worried: '担心'
}[mood] || mood)

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

<style scoped>
.lightbox-enter-active { transition: all 0.3s ease-out; }
.lightbox-leave-active { transition: all 0.2s ease-in; }
.lightbox-enter-from { opacity: 0; }
.lightbox-leave-to { opacity: 0; }
</style>