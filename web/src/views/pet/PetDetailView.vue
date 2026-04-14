<template>
  <div v-if="loading" class="text-center py-12">
    <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto"></div>
  </div>
  <div v-else-if="!pet" class="text-center py-12 text-gray-500">宠物不存在</div>
  <div v-else class="space-y-6">
    <router-link to="/pets" class="text-gray-500 hover:text-gray-700">&larr; 返回</router-link>

    <!-- Pet Info Card -->
    <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
      <div class="h-48 bg-gradient-to-br from-orange-100 to-amber-50 flex items-center justify-center relative">
        <img v-if="pet.avatar_url" :src="pet.avatar_url" :alt="pet.name" class="w-full h-full object-cover" />
        <span v-else class="text-8xl">🐾</span>
        <span v-if="pet.status === 'memorial'" class="absolute top-3 right-3 text-xs bg-gray-800/60 text-white px-2 py-1 rounded-full">纪念</span>
      </div>
      <div class="p-6">
        <div class="flex items-center justify-between mb-4">
          <h1 class="text-2xl font-bold text-gray-800">{{ pet.name }}</h1>
          <div class="flex gap-2">
            <router-link :to="`/pets/${pet.id}/edit`" class="px-3 py-1.5 text-sm border border-gray-300 rounded-lg text-gray-600 hover:bg-gray-50">编辑</router-link>
            <button @click="handleDelete" class="px-3 py-1.5 text-sm border border-red-200 rounded-lg text-red-500 hover:bg-red-50">删除</button>
          </div>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-sm">
          <div><span class="text-gray-400">种类</span><div class="font-medium mt-1">{{ pet.category?.name }}</div></div>
          <div><span class="text-gray-400">品种</span><div class="font-medium mt-1">{{ pet.breed?.name || '未知' }}</div></div>
          <div><span class="text-gray-400">性别</span><div class="font-medium mt-1">{{ pet.gender }}</div></div>
          <div><span class="text-gray-400">生日</span><div class="font-medium mt-1">{{ pet.birthday }}</div></div>
        </div>
        <p v-if="pet.bio" class="mt-4 pt-4 border-t border-gray-100 text-gray-500 text-sm">{{ pet.bio }}</p>
      </div>
    </div>

    <!-- Tabs -->
    <div class="bg-white rounded-2xl shadow-sm">
      <div class="flex border-b border-gray-100">
        <button v-for="tab in tabs" :key="tab.key" @click="activeTab = tab.key"
          :class="['flex-1 py-3 text-sm font-medium text-center transition-colors', activeTab === tab.key ? 'text-orange-500 border-b-2 border-orange-500' : 'text-gray-400 hover:text-gray-600']">
          {{ tab.label }}
        </button>
      </div>

      <!-- Diary Tab -->
      <div v-if="activeTab === 'diary'" class="p-4 space-y-3">
        <div v-if="loadingTab" class="text-center py-8 text-gray-400">加载中...</div>
        <div v-else-if="diaries.length === 0" class="text-center py-8 text-gray-400">暂无日记</div>
        <router-link v-for="d in diaries" :key="d.id" :to="`/diaries/${d.id}`"
          class="block p-4 rounded-xl hover:bg-gray-50 transition-colors">
          <div class="flex items-center justify-between mb-1">
            <span class="font-medium text-gray-700">{{ d.title }}</span>
            <span class="text-xs text-gray-400">{{ d.entry_date }}</span>
          </div>
          <p class="text-sm text-gray-500 line-clamp-2">{{ d.content }}</p>
          <div v-if="d.photos?.length" class="flex gap-2 mt-2">
            <img v-for="p in d.photos.slice(0, 3)" :key="p.id" :src="p.thumbnail_url || p.photo_url" class="w-12 h-12 rounded-lg object-cover" />
          </div>
        </router-link>
      </div>

      <!-- Health Tab -->
      <div v-if="activeTab === 'health'" class="p-4 space-y-3">
        <div class="flex justify-between items-center">
          <span class="text-sm text-gray-500">健康记录</span>
          <router-link :to="`/health/${pet.id}`" class="text-sm text-orange-500 hover:underline">查看全部</router-link>
        </div>
        <div v-if="healthRecords.length === 0" class="text-center py-8 text-gray-400">暂无记录</div>
        <div v-for="r in healthRecords.slice(0, 5)" :key="r.id" class="flex items-center justify-between p-3 bg-gray-50 rounded-xl">
          <div>
            <div class="font-medium text-sm text-gray-700">{{ r.title }}</div>
            <div class="text-xs text-gray-400">{{ r.record_date }} {{ r.value ? `· ${r.value} ${r.unit}` : '' }}</div>
          </div>
          <span class="text-xs bg-gray-200 text-gray-600 px-2 py-0.5 rounded">{{ typeLabel(r.type) }}</span>
        </div>
      </div>

      <!-- Album Tab -->
      <div v-if="activeTab === 'album'" class="p-4">
        <div v-if="photos.length === 0" class="text-center py-8 text-gray-400">暂无照片</div>
        <div class="grid grid-cols-3 md:grid-cols-4 gap-2">
          <img v-for="p in photos" :key="p.id" :src="p.thumbnail_url || p.photo_url"
            @click="lightboxUrl = p.photo_url"
            class="w-full aspect-square object-cover rounded-xl cursor-pointer hover:opacity-80 transition-opacity" />
        </div>
      </div>
    </div>

    <!-- Lightbox -->
    <div v-if="lightboxUrl" class="fixed inset-0 bg-black/80 z-[100] flex items-center justify-center" @click="lightboxUrl = ''">
      <img :src="lightboxUrl" class="max-w-[90vw] max-h-[90vh] object-contain rounded-lg" @click.stop />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { usePetStore } from '../../stores/pet'
import { useDiaryStore } from '../../stores/diary'
import { healthApi } from '../../api/health'
import { petsApi } from '../../api/pets'

const tabs = [
  { key: 'diary', label: '日记' },
  { key: 'health', label: '健康' },
  { key: 'album', label: '相册' },
]
const typeLabel = (type: string) => ({ vaccine: '疫苗', deworming: '驱虫', weight: '体重', medical: '就医' }[type] || type)

const route = useRoute()
const router = useRouter()
const petStore = usePetStore()
const diaryStore = useDiaryStore()
const pet = ref<any>(null)
const loading = ref(true)
const loadingTab = ref(false)
const activeTab = ref('diary')
const diaries = ref<any[]>([])
const healthRecords = ref<any[]>([])
const photos = ref<any[]>([])
const lightboxUrl = ref('')

const petId = () => Number(route.params.id)

const loadTab = async () => {
  loadingTab.value = true
  try {
    if (activeTab.value === 'diary') {
      await diaryStore.fetchDiaries(petId())
      diaries.value = diaryStore.diaries
    } else if (activeTab.value === 'health') {
      const { data } = await healthApi.list(petId())
      healthRecords.value = data.data || data
    } else if (activeTab.value === 'album') {
      const { data } = await petsApi.getPhotos(petId())
      photos.value = data.data || data
    }
  } finally {
    loadingTab.value = false
  }
}

watch(activeTab, loadTab)

const handleDelete = async () => {
  if (!confirm('确定要删除这只宠物吗？')) return
  await petStore.deletePet(petId())
  router.push('/pets')
}

onMounted(async () => {
  try {
    await petStore.fetchPet(petId())
    pet.value = petStore.currentPet
    await loadTab()
  } finally {
    loading.value = false
  }
})
</script>