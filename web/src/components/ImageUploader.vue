<template>
  <div>
    <div
      @dragover.prevent="dragging = true"
      @dragleave="dragging = false"
      @drop.prevent="handleDrop"
      @click="inputRef?.click()"
      :class="['border-2 border-dashed rounded-xl p-6 text-center cursor-pointer transition-colors', dragging ? 'border-orange-400 bg-orange-50' : 'border-gray-300 hover:border-gray-400']"
    >
      <div class="text-3xl mb-2">{{ dragging ? '📥' : '📁' }}</div>
      <p class="text-sm text-gray-500">{{ dragging ? '松开上传' : '拖拽图片到这里，或点击选择' }}</p>
      <p class="text-xs text-gray-400 mt-1">支持 JPG/PNG/WebP</p>
    </div>
    <input ref="inputRef" type="file" accept="image/*" multiple class="hidden" @change="handleInput" />

    <div v-if="photos.length" class="flex flex-wrap gap-2 mt-3">
      <div v-for="(photo, i) in photos" :key="i" class="relative group w-20 h-20">
        <img :src="photo.photo_url" class="w-full h-full rounded-lg object-cover" />
        <button type="button" @click="emit('remove', i)"
          class="absolute -top-1 -right-1 w-5 h-5 bg-red-500 text-white rounded-full text-xs flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
          &times;
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { uploadApi } from '../api/common'

const props = defineProps<{ photos: { photo_url: string; thumbnail_url: string }[] }>()
const emit = defineEmits<{
  remove: [index: number]
  photosAdded: [photos: { photo_url: string; thumbnail_url: string }[]]
}>()

const inputRef = ref<HTMLInputElement | null>(null)
const dragging = ref(false)

const uploadAndEmit = async (files: FileList) => {
  const newPhotos: { photo_url: string; thumbnail_url: string }[] = []
  for (const file of Array.from(files)) {
    if (!file.type.startsWith('image/')) continue
    const result = await uploadApi.upload(file)
    newPhotos.push({ photo_url: result.url, thumbnail_url: result.thumbnail_url })
  }
  if (newPhotos.length) emit('photosAdded', newPhotos)
}

const handleDrop = async (e: DragEvent) => {
  dragging.value = false
  if (e.dataTransfer?.files) await uploadAndEmit(e.dataTransfer.files)
}

const handleInput = async (e: Event) => {
  const files = (e.target as HTMLInputElement).files
  if (files) await uploadAndEmit(files)
  ;(e.target as HTMLInputElement).value = ''
}
</script>
