<template>
  <div v-if="loading" class="flex flex-col items-center py-16 gap-4">
    <div class="w-10 h-10 border-3 border-gray-200 border-t-orange-500 rounded-full animate-spin"></div>
    <p class="text-sm text-gray-400">加载中...</p>
  </div>

  <div v-else-if="!pet" class="flex flex-col items-center py-16 text-center">
    <div class="w-20 h-20 bg-gray-100 rounded-full flex items-center justify-center mb-4">
      <PawPrint class="w-10 h-10 text-gray-300" />
    </div>
    <p class="text-gray-500">宠物不存在</p>
    <router-link to="/pets" class="mt-4 text-orange-500 hover:text-orange-600 font-medium text-sm">返回宠物列表</router-link>
  </div>

  <div v-else class="space-y-6">
    <!-- Back Link -->
    <router-link to="/pets" class="inline-flex items-center gap-1 text-sm text-gray-500 hover:text-gray-700 transition-colors">
      <ChevronLeft class="w-4 h-4" /> 返回
    </router-link>

    <!-- Pet Info Card -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <div class="h-52 bg-gradient-to-br from-orange-100 via-amber-50 to-orange-50 flex items-center justify-center relative overflow-hidden">
        <img v-if="pet.avatar_url" :src="pet.avatar_url" :alt="pet.name" class="w-full h-full object-cover" />
        <PawPrint v-else class="w-20 h-20 text-orange-200" />
        <span v-if="pet.status === 'memorial'" class="absolute top-4 right-4 flex items-center gap-1.5 text-xs bg-gray-800/70 backdrop-blur-sm text-white px-3 py-1.5 rounded-full">
          <Flower2 class="w-3.5 h-3.5" /> 纪念
        </span>
      </div>
      <div class="p-6">
        <div class="flex items-center justify-between mb-5">
          <h1 class="text-2xl font-bold text-gray-800">{{ pet.name }}</h1>
          <div class="flex gap-2">
            <router-link :to="`/pets/${pet.id}/edit`" class="inline-flex items-center gap-1.5 px-3 py-1.5 text-sm font-medium border border-gray-200 rounded-xl text-gray-600 hover:bg-gray-50 hover:border-gray-300 transition-all duration-200">
              <Pencil class="w-3.5 h-3.5" /> 编辑
            </router-link>
            <button @click="handleDelete" class="inline-flex items-center gap-1.5 px-3 py-1.5 text-sm font-medium border border-red-200 rounded-xl text-red-500 hover:bg-red-50 hover:border-red-300 transition-all duration-200">
              <Trash2 class="w-3.5 h-3.5" /> 删除
            </button>
          </div>
        </div>

        <!-- Pet Details Grid -->
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div class="bg-gray-50 rounded-xl p-3">
            <div class="flex items-center gap-1.5 text-gray-400 text-xs mb-1">
              <Layers class="w-3 h-3" /> 种类
            </div>
            <div class="font-medium text-gray-800">{{ pet.category?.name }}</div>
          </div>
          <div class="bg-gray-50 rounded-xl p-3">
            <div class="flex items-center gap-1.5 text-gray-400 text-xs mb-1">
              <Sparkles class="w-3 h-3" /> 品种
            </div>
            <div class="font-medium text-gray-800">{{ pet.breed?.name || '未知' }}</div>
          </div>
          <div class="bg-gray-50 rounded-xl p-3">
            <div class="flex items-center gap-1.5 text-gray-400 text-xs mb-1">
              <component :is="pet.gender === '公' ? Mars : Venus" class="w-3 h-3" /> 性别
            </div>
            <div class="font-medium text-gray-800">{{ pet.gender }}</div>
          </div>
          <div class="bg-gray-50 rounded-xl p-3">
            <div class="flex items-center gap-1.5 text-gray-400 text-xs mb-1">
              <Calendar class="w-3 h-3" /> 生日
            </div>
            <div class="font-medium text-gray-800">{{ pet.birthday }}</div>
          </div>
        </div>

        <p v-if="pet.bio" class="mt-4 pt-4 border-t border-gray-100 text-gray-600 text-sm leading-relaxed">{{ pet.bio }}</p>
      </div>
    </div>

    <!-- Tabs -->
    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <div class="flex border-b border-gray-100">
        <button v-for="tab in tabs" :key="tab.key" @click="activeTab = tab.key"
          class="flex-1 flex items-center justify-center gap-2 py-3.5 text-sm font-medium transition-all duration-200 relative"
          :class="activeTab === tab.key ? 'text-orange-500' : 'text-gray-400 hover:text-gray-600'">
          <component :is="tab.icon" class="w-4 h-4" />
          {{ tab.label }}
          <span v-if="activeTab === tab.key" class="absolute bottom-0 left-1/4 right-1/4 h-0.5 bg-orange-500 rounded-full"></span>
        </button>
      </div>

      <!-- Diary Tab -->
      <div v-if="activeTab === 'diary'" class="p-4 space-y-3 min-h-[200px]">
        <div v-if="loadingTab" class="flex items-center justify-center py-8">
          <div class="w-6 h-6 border-2 border-gray-200 border-t-orange-500 rounded-full animate-spin"></div>
        </div>
        <div v-else-if="diaries.length === 0" class="flex flex-col items-center py-8 text-center">
          <BookOpen class="w-10 h-10 text-gray-200 mb-3" />
          <p class="text-sm text-gray-400">暂无日记</p>
          <router-link to="/diaries/new" class="mt-2 text-xs text-orange-500 hover:text-orange-600">写一篇日记</router-link>
        </div>
        <router-link v-for="d in diaries" :key="d.id" :to="`/diaries/${d.id}`"
          class="block p-4 rounded-xl hover:bg-gray-50 transition-colors group">
          <div class="flex items-center justify-between mb-1.5">
            <span class="font-medium text-gray-700 group-hover:text-orange-500 transition-colors">{{ d.title }}</span>
            <span class="text-xs text-gray-400">{{ d.entry_date }}</span>
          </div>
          <p class="text-sm text-gray-500 line-clamp-2">{{ d.content }}</p>
          <div v-if="d.photos?.length" class="flex gap-2 mt-3">
            <img v-for="p in d.photos.slice(0, 3)" :key="p.id" :src="p.thumbnail_url || p.photo_url"
              class="w-10 h-10 rounded-lg object-cover" />
          </div>
        </router-link>
      </div>

      <!-- Health Tab -->
      <div v-if="activeTab === 'health'" class="p-4 space-y-3 min-h-[200px]">
        <div class="flex justify-between items-center">
          <span class="text-sm font-medium text-gray-600">健康记录</span>
          <router-link :to="`/health/${pet.id}`" class="inline-flex items-center gap-1 text-sm text-orange-500 hover:text-orange-600 font-medium transition-colors">
            查看全部 <ChevronRight class="w-3.5 h-3.5" />
          </router-link>
        </div>
        <div v-if="loadingTab" class="flex items-center justify-center py-8">
          <div class="w-6 h-6 border-2 border-gray-200 border-t-orange-500 rounded-full animate-spin"></div>
        </div>
        <div v-else-if="healthRecords.length === 0" class="flex flex-col items-center py-8 text-center">
          <Activity class="w-10 h-10 text-gray-200 mb-3" />
          <p class="text-sm text-gray-400">暂无健康记录</p>
        </div>
        <div v-else v-for="r in healthRecords.slice(0, 5)" :key="r.id"
          class="flex items-center justify-between p-3 bg-gray-50 rounded-xl">
          <div>
            <div class="font-medium text-sm text-gray-700">{{ r.title }}</div>
            <div class="text-xs text-gray-400 flex items-center gap-2">
              <span>{{ r.record_date }}</span>
              <span v-if="r.value" class="bg-white px-2 py-0.5 rounded-md">{{ r.value }} {{ r.unit }}</span>
            </div>
          </div>
          <span class="text-xs bg-white text-gray-600 px-2.5 py-1 rounded-lg border border-gray-200">{{ typeLabel(r.type) }}</span>
        </div>
      </div>

      <!-- Album Tab -->
      <div v-if="activeTab === 'album'" class="p-4 min-h-[200px]">
        <div v-if="loadingTab" class="flex items-center justify-center py-8">
          <div class="w-6 h-6 border-2 border-gray-200 border-t-orange-500 rounded-full animate-spin"></div>
        </div>
        <div v-else-if="photos.length === 0" class="flex flex-col items-center py-8 text-center">
          <Image class="w-10 h-10 text-gray-200 mb-3" />
          <p class="text-sm text-gray-400">暂无照片</p>
        </div>
        <div v-else class="grid grid-cols-3 md:grid-cols-4 gap-2">
          <img v-for="p in photos" :key="p.id" :src="p.thumbnail_url || p.photo_url"
            @click="lightboxUrl = p.photo_url"
            class="w-full aspect-square object-cover rounded-xl cursor-pointer hover:opacity-90 transition-opacity ring-2 ring-transparent hover:ring-orange-200" />
        </div>
      </div>
    </div>

    <!-- Lightbox -->
    <Teleport to="body">
      <Transition name="lightbox">
        <div v-if="lightboxUrl" class="fixed inset-0 bg-black/90 z-[100] flex items-center justify-center backdrop-blur-sm"
          @click="lightboxUrl = ''">
          <button @click="lightboxUrl = ''" class="absolute top-4 right-4 w-10 h-10 bg-white/10 rounded-full flex items-center justify-center hover:bg-white/20 transition-colors">
            <X class="w-5 h-5 text-white" />
          </button>
          <img :src="lightboxUrl" class="max-w-[90vw] max-h-[90vh] object-contain rounded-lg shadow-2xl" @click.stop />
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { usePetStore } from '../../stores/pet'
import { useDiaryStore } from '../../stores/diary'
import { healthApi } from '../../api/health'
import { petsApi } from '../../api/pets'
import {
  PawPrint, ChevronLeft, Flower2, Pencil, Trash2, Layers, Sparkles, Mars, Venus, Calendar,
  BookOpen, Activity, Image, ChevronRight, X
} from 'lucide-vue-next'

const tabs = [
  { key: 'diary', label: '日记', icon: BookOpen },
  { key: 'health', label: '健康', icon: Activity },
  { key: 'album', label: '相册', icon: Image },
]

const typeLabel = (type: string) => ({
  vaccine: '疫苗', deworming: '驱虫', weight: '体重', medical: '就医'
}[type] || type)

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

<style scoped>
.lightbox-enter-active { transition: all 0.3s ease-out; }
.lightbox-leave-active { transition: all 0.2s ease-in; }
.lightbox-enter-from { opacity: 0; }
.lightbox-leave-to { opacity: 0; }
</style>