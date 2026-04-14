<template>
  <div class="max-w-2xl mx-auto space-y-6">
    <!-- Header -->
    <div class="flex items-center gap-4">
      <router-link to="/diaries" class="inline-flex items-center gap-1 text-sm text-gray-500 hover:text-gray-700 transition-colors">
        <ChevronLeft class="w-4 h-4" /> 返回
      </router-link>
      <h1 class="text-2xl font-bold text-gray-800">{{ isEdit ? '编辑日记' : '写日记' }}</h1>
    </div>

    <form @submit.prevent="handleSubmit" class="bg-white rounded-2xl border border-gray-100 shadow-sm p-6 space-y-5">
      <!-- Title -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1.5">标题 <span class="text-red-400">*</span></label>
        <input v-model="form.title" type="text" required placeholder="给日记起个名字"
          class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all" />
      </div>

      <!-- Date & Mood -->
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">日期 <span class="text-red-400">*</span></label>
          <input v-model="form.entry_date" type="date" required
            class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">心情</label>
          <select v-model="form.mood" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all">
            <option value="">请选择</option>
            <option value="happy">😊 开心</option>
            <option value="calm">😌 平静</option>
            <option value="sad">😢 难过</option>
            <option value="excited">🤗 兴奋</option>
            <option value="worried">😟 担心</option>
          </select>
        </div>
      </div>

      <!-- Pets -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1.5">关联宠物</label>
        <div class="flex flex-wrap gap-2">
          <label v-for="pet in pets" :key="pet.id"
            :class="['inline-flex items-center gap-1.5 px-3 py-1.5 rounded-xl cursor-pointer border font-medium text-sm transition-all duration-200',
              form.pet_ids.includes(pet.id) ? 'bg-orange-50 border-orange-300 text-orange-600' : 'bg-white border-gray-200 text-gray-500 hover:border-gray-300']"
          >
            <input type="checkbox" :value="pet.id" v-model="form.pet_ids" class="hidden" />
            {{ pet.name }}
          </label>
        </div>
      </div>

      <!-- Content -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1.5">内容 <span class="text-red-400">*</span></label>
        <textarea v-model="form.content" rows="8" required placeholder="记录今天的故事..."
          class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all resize-y"></textarea>
      </div>

      <!-- Photos -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1.5">照片</label>
        <ImageUploader :photos="form.photos" @remove="(i: number) => form.photos?.splice(i, 1)" @photos-added="addPhotos" />
      </div>

      <!-- Actions -->
      <div class="flex justify-end gap-3 pt-2">
        <router-link to="/diaries"
          class="px-5 py-2.5 border border-gray-200 rounded-xl text-gray-600 hover:bg-gray-50 font-medium text-sm transition-all">
          取消
        </router-link>
        <button type="submit" :disabled="submitting"
          class="inline-flex items-center gap-2 px-5 py-2.5 bg-blue-500 text-white rounded-xl hover:bg-blue-600 disabled:opacity-50 font-medium text-sm shadow-sm shadow-blue-500/20 transition-all">
          <Loader2 v-if="submitting" class="w-4 h-4 animate-spin" />
          {{ submitting ? '保存中...' : '保存' }}
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useDiaryStore } from '../../stores/diary'
import { usePetStore } from '../../stores/pet'
import ImageUploader from '../../components/ImageUploader.vue'
import { ChevronLeft, Loader2 } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const diaryStore = useDiaryStore()
const petStore = usePetStore()
const isEdit = computed(() => !!route.params.id)
const submitting = ref(false)

const pets = ref<any[]>([])
const form = ref<any>({
  title: '',
  content: '',
  entry_date: new Date().toISOString().split('T')[0],
  mood: '',
  pet_ids: [] as number[],
  photos: [] as { photo_url: string; thumbnail_url: string }[],
})

const addPhotos = (newPhotos: { photo_url: string; thumbnail_url: string }[]) => {
  form.value.photos.push(...newPhotos)
}

const handleSubmit = async () => {
  submitting.value = true
  try {
    const payload = { ...form.value }
    if (!payload.mood) delete payload.mood
    if (!payload.photos?.length) delete payload.photos
    if (isEdit.value) {
      await diaryStore.updateDiary(Number(route.params.id), payload)
    } else {
      await diaryStore.createDiary(payload)
    }
    router.push('/diaries')
  } finally {
    submitting.value = false
  }
}

onMounted(async () => {
  await petStore.fetchPets()
  pets.value = petStore.pets

  if (isEdit.value) {
    await diaryStore.fetchDiary(Number(route.params.id))
    const diary = diaryStore.currentDiary
    if (diary) {
      form.value = {
        title: diary.title,
        content: diary.content,
        entry_date: diary.entry_date,
        mood: diary.mood || '',
        pet_ids: diary.pets?.map((p: any) => p.pet_id) || [],
        photos: diary.photos?.map((p: any) => ({ photo_url: p.photo_url, thumbnail_url: p.thumbnail_url })) || [],
      }
    }
  }
})
</script>