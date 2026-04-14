<template>
  <div class="space-y-8">
    <!-- Header -->
    <h1 class="text-2xl font-bold text-gray-800">相册</h1>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center py-16 gap-4">
      <div class="w-10 h-10 border-3 border-gray-200 border-t-orange-500 rounded-full animate-spin"></div>
      <p class="text-sm text-gray-400">加载中...</p>
    </div>

    <!-- Empty State -->
    <div v-else-if="groups.length === 0" class="flex flex-col items-center py-16 text-center">
      <div class="w-20 h-20 bg-gray-100 rounded-full flex items-center justify-center mb-4">
        <Image class="w-10 h-10 text-gray-300" />
      </div>
      <p class="text-gray-500 mb-1">还没有照片</p>
      <p class="text-sm text-gray-400 mb-4">写日记时上传图片吧</p>
      <router-link to="/diaries/new" class="inline-flex items-center gap-2 text-orange-500 hover:text-orange-600 font-medium text-sm">
        <Plus class="w-4 h-4" /> 写一篇日记
      </router-link>
    </div>

    <!-- Photo Groups -->
    <div v-for="group in groups" :key="group.pet.id" class="space-y-3">
      <div class="flex items-center gap-3">
        <router-link :to="`/pets/${group.pet.id}`"
          class="w-10 h-10 rounded-full bg-orange-50 flex items-center justify-center overflow-hidden ring-2 ring-transparent hover:ring-orange-200 transition-all"
        >
          <img v-if="group.pet.avatar_url" :src="group.pet.avatar_url" class="w-full h-full object-cover" />
          <span v-else class="text-orange-400 font-medium text-sm">{{ group.pet.name.charAt(0) }}</span>
        </router-link>
        <router-link :to="`/pets/${group.pet.id}`" class="font-semibold text-gray-700 hover:text-orange-500 transition-colors">
          {{ group.pet.name }}
        </router-link>
        <span class="text-xs text-gray-400 bg-gray-50 px-2 py-0.5 rounded-md">{{ group.photos.length }} 张</span>
      </div>
      <div class="grid grid-cols-3 md:grid-cols-5 lg:grid-cols-6 gap-2">
        <div v-for="(p, i) in group.photos" :key="p.id"
          class="relative group cursor-pointer"
          @click="openLightbox(group, i)"
        >
          <img :src="p.thumbnail_url || p.photo_url"
            class="w-full aspect-square object-cover rounded-xl ring-2 ring-transparent group-hover:ring-orange-200 transition-all duration-200"
          />
          <div v-if="p.entry_date"
            class="absolute bottom-0 inset-x-0 bg-gradient-to-t from-black/60 to-transparent rounded-b-xl px-2 py-1.5 opacity-0 group-hover:opacity-100 transition-opacity"
          >
            <span class="text-[10px] text-white font-medium">{{ p.entry_date }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Lightbox -->
    <Teleport to="body">
      <Transition name="lightbox">
        <div v-if="lightbox.open" class="fixed inset-0 bg-black/90 z-[100] flex items-center justify-center backdrop-blur-sm"
          @click="lightbox.open = false"
        >
          <button @click.stop="prev"
            class="absolute left-4 w-12 h-12 bg-white/10 rounded-full flex items-center justify-center hover:bg-white/20 transition-colors"
          >
            <ChevronLeft class="w-6 h-6 text-white" />
          </button>

          <img :src="lightbox.url" class="max-w-[85vw] max-h-[85vh] object-contain rounded-lg shadow-2xl" @click.stop />

          <button @click.stop="next"
            class="absolute right-4 w-12 h-12 bg-white/10 rounded-full flex items-center justify-center hover:bg-white/20 transition-colors"
          >
            <ChevronRight class="w-6 h-6 text-white" />
          </button>

          <button @click="lightbox.open = false"
            class="absolute top-4 right-4 w-10 h-10 bg-white/10 rounded-full flex items-center justify-center hover:bg-white/20 transition-colors"
          >
            <X class="w-5 h-5 text-white" />
          </button>

          <div v-if="lightbox.date" class="absolute bottom-6 text-white/70 text-sm font-medium">
            {{ lightbox.date }}
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { usePetStore } from '../../stores/pet'
import { petsApi } from '../../api/pets'
import { Image, Plus, ChevronLeft, ChevronRight, X } from 'lucide-vue-next'

const petStore = usePetStore()
const loading = ref(true)
const groups = ref<{ pet: any; photos: any[] }[]>([])

const lightbox = reactive({ open: false, url: '', date: '', groupIndex: 0, photoIndex: 0 })

const openLightbox = (group: any, index: number) => {
  const photo = group.photos[index]
  lightbox.groupIndex = groups.value.indexOf(group)
  lightbox.photoIndex = index
  lightbox.url = photo.photo_url
  lightbox.date = photo.entry_date || ''
  lightbox.open = true
}

const prev = () => {
  const g = lightbox.groupIndex
  const p = lightbox.photoIndex - 1
  if (p >= 0) {
    lightbox.photoIndex = p
  } else if (g > 0) {
    lightbox.groupIndex = g - 1
    lightbox.photoIndex = groups.value[g - 1].photos.length - 1
  }
  const photo = groups.value[lightbox.groupIndex]?.photos[lightbox.photoIndex]
  if (photo) { lightbox.url = photo.photo_url; lightbox.date = photo.entry_date || '' }
}

const next = () => {
  const g = lightbox.groupIndex
  const p = lightbox.photoIndex + 1
  if (p < groups.value[g].photos.length) {
    lightbox.photoIndex = p
  } else if (g < groups.value.length - 1) {
    lightbox.groupIndex = g + 1
    lightbox.photoIndex = 0
  }
  const photo = groups.value[lightbox.groupIndex]?.photos[lightbox.photoIndex]
  if (photo) { lightbox.url = photo.photo_url; lightbox.date = photo.entry_date || '' }
}

onMounted(async () => {
  await petStore.fetchPets()
  for (const pet of petStore.pets) {
    try {
      const { data } = await petsApi.getPhotos(pet.id)
      const photos = data.data || data || []
      if (photos.length) groups.value.push({ pet, photos })
    } catch { }
  }
  loading.value = false
})
</script>

<style scoped>
.lightbox-enter-active { transition: all 0.3s ease-out; }
.lightbox-leave-active { transition: all 0.2s ease-in; }
.lightbox-enter-from { opacity: 0; }
.lightbox-leave-to { opacity: 0; }
</style>