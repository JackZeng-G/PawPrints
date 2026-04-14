<template>
  <div class="max-w-2xl mx-auto space-y-6">
    <div class="flex items-center gap-4">
      <router-link to="/diaries" class="text-gray-500 hover:text-gray-700">&larr; 返回</router-link>
      <h1 class="text-2xl font-bold text-gray-800">{{ isEdit ? '编辑日记' : '写日记' }}</h1>
    </div>

    <form @submit.prevent="handleSubmit" class="bg-white rounded-lg shadow p-6 space-y-4">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">标题 *</label>
        <input v-model="form.title" type="text" required class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">日期 *</label>
          <input v-model="form.entry_date" type="date" required class="w-full border border-gray-300 rounded-lg px-4 py-2" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">心情</label>
          <select v-model="form.mood" class="w-full border border-gray-300 rounded-lg px-4 py-2">
            <option value="">请选择</option>
            <option value="happy">😊 开心</option>
            <option value="calm">😌 平静</option>
            <option value="sad">😢 难过</option>
            <option value="excited">🤗 兴奋</option>
            <option value="worried">😟 担心</option>
          </select>
        </div>
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">关联宠物</label>
        <div class="flex flex-wrap gap-2">
          <label
            v-for="pet in pets"
            :key="pet.id"
            :class="['px-3 py-1 rounded-full cursor-pointer border', form.pet_ids.includes(pet.id) ? 'bg-orange-100 border-orange-400 text-orange-600' : 'bg-gray-50 border-gray-200 text-gray-600']"
          >
            <input type="checkbox" :value="pet.id" v-model="form.pet_ids" class="hidden" />
            {{ pet.name }}
          </label>
        </div>
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">内容 *</label>
        <textarea v-model="form.content" rows="8" required class="w-full border border-gray-300 rounded-lg px-4 py-2"></textarea>
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">照片</label>
        <ImageUploader :photos="form.photos" @remove="(i: number) => form.photos?.splice(i, 1)" @photos-added="addPhotos" />
      </div>

      <div class="flex justify-end gap-3 pt-4">
        <router-link to="/diaries" class="px-6 py-2 border border-gray-300 rounded-lg text-gray-600 hover:bg-gray-50">取消</router-link>
        <button type="submit" :disabled="submitting" class="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:opacity-50">
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