<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold text-gray-800">相册</h1>

    <div v-if="loading" class="text-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto"></div>
    </div>

    <div v-else-if="groups.length === 0" class="text-center py-16 text-gray-400">
      <span class="text-5xl block mb-4">📷</span>
      还没有照片，写日记时上传图片吧
    </div>

    <div v-for="group in groups" :key="group.pet.id" class="space-y-3">
      <div class="flex items-center gap-3">
        <div class="w-8 h-8 rounded-full bg-orange-100 flex items-center justify-center text-sm">
          {{ group.pet.name.charAt(0) }}
        </div>
        <h2 class="font-semibold text-gray-700">{{ group.pet.name }}</h2>
        <span class="text-sm text-gray-400">{{ group.photos.length }} 张</span>
      </div>
      <div class="grid grid-cols-3 md:grid-cols-5 lg:grid-cols-6 gap-2">
        <div v-for="(p, i) in group.photos" :key="p.id" class="relative group cursor-pointer" @click="openLightbox(group, i)">
          <img :src="p.thumbnail_url || p.photo_url" class="w-full aspect-square object-cover rounded-xl" />
          <div v-if="p.entry_date" class="absolute bottom-0 inset-x-0 bg-gradient-to-t from-black/50 to-transparent rounded-b-xl px-2 py-1">
            <span class="text-[10px] text-white">{{ p.entry_date }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Lightbox -->
    <div v-if="lightbox.open" class="fixed inset-0 bg-black/90 z-[100] flex items-center justify-center" @click="lightbox.open = false">
      <button @click.stop="prev" class="absolute left-4 text-white/60 hover:text-white text-3xl">&lsaquo;</button>
      <img :src="lightbox.url" class="max-w-[85vw] max-h-[85vh] object-contain rounded-lg" @click.stop />
      <button @click.stop="next" class="absolute right-4 text-white/60 hover:text-white text-3xl">&rsaquo;</button>
      <div class="absolute bottom-6 text-white/70 text-sm">{{ lightbox.date }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { usePetStore } from '../../stores/pet'
import { petsApi } from '../../api/pets'

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
    } catch {}
  }
  loading.value = false
})
</script>