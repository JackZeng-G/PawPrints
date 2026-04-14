<template>
  <div class="max-w-2xl mx-auto space-y-6">
    <!-- Header -->
    <div class="flex items-center gap-4">
      <router-link to="/pets" class="inline-flex items-center gap-1 text-sm text-gray-500 hover:text-gray-700 transition-colors">
        <ChevronLeft class="w-4 h-4" /> 返回
      </router-link>
      <h1 class="text-2xl font-bold text-gray-800">{{ isEdit ? '编辑宠物' : '添加宠物' }}</h1>
    </div>

    <form @submit.prevent="handleSubmit" class="bg-white rounded-2xl border border-gray-100 shadow-sm p-6 space-y-5">
      <!-- Avatar Preview -->
      <div v-if="form.avatar_url" class="flex justify-center mb-2">
        <div class="w-24 h-24 rounded-full overflow-hidden ring-4 ring-orange-100">
          <img :src="form.avatar_url" alt="头像" class="w-full h-full object-cover" />
        </div>
      </div>

      <!-- Name -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1.5">名字 <span class="text-red-400">*</span></label>
        <input v-model="form.name" type="text" required placeholder="宠物的名字"
          class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all" />
      </div>

      <!-- Avatar Upload -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1.5">头像</label>
        <label class="inline-flex items-center gap-2 px-4 py-2 border border-gray-200 rounded-xl text-sm text-gray-600 hover:bg-gray-50 hover:border-gray-300 transition-all cursor-pointer">
          <Upload class="w-4 h-4" /> 选择图片
          <input type="file" accept="image/*" @change="handleAvatar" class="hidden" />
        </label>
      </div>

      <!-- Category & Breed -->
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">种类 <span class="text-red-400">*</span></label>
          <select v-model="form.category_id" required @change="loadBreeds"
            class="w-full border border-gray-200 rounded-xl px-4 py-2.5 bg-white focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all"
          >
            <option value="">请选择</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">品种</label>
          <select v-model="form.breed_id"
            class="w-full border border-gray-200 rounded-xl px-4 py-2.5 bg-white focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all"
          >
            <option value="">请选择</option>
            <option v-for="breed in breeds" :key="breed.id" :value="breed.id">{{ breed.name }}</option>
          </select>
        </div>
      </div>

      <!-- Gender & Birthday -->
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">性别 <span class="text-red-400">*</span></label>
          <select v-model="form.gender" required
            class="w-full border border-gray-200 rounded-xl px-4 py-2.5 bg-white focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all"
          >
            <option value="公">公</option>
            <option value="母">母</option>
            <option value="未知">未知</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">生日 <span class="text-red-400">*</span></label>
          <input v-model="form.birthday" type="date" required
            class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all" />
        </div>
      </div>

      <!-- Bio -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1.5">简介</label>
        <textarea v-model="form.bio" rows="3" placeholder="介绍一下你的宠物..."
          class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all resize-y"></textarea>
      </div>

      <!-- Status (Edit only) -->
      <div v-if="isEdit" class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">状态</label>
          <select v-model="form.status"
            class="w-full border border-gray-200 rounded-xl px-4 py-2.5 bg-white focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all"
          >
            <option value="alive">在世</option>
            <option value="memorial">纪念</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1.5">离别日期</label>
          <input v-model="form.passing_date" type="date"
            class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all" />
        </div>
      </div>

      <!-- Actions -->
      <div class="flex justify-end gap-3 pt-2">
        <router-link to="/pets"
          class="px-5 py-2.5 border border-gray-200 rounded-xl text-gray-600 hover:bg-gray-50 font-medium text-sm transition-all"
        >
          取消
        </router-link>
        <button type="submit" :disabled="submitting"
          class="inline-flex items-center gap-2 px-5 py-2.5 bg-orange-500 text-white rounded-xl hover:bg-orange-600 disabled:opacity-50 font-medium text-sm shadow-sm shadow-orange-500/20 transition-all"
        >
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
import { usePetStore } from '../../stores/pet'
import { categoriesApi, uploadApi } from '../../api/common'
import { ChevronLeft, Upload, Loader2 } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const petStore = usePetStore()
const isEdit = computed(() => !!route.params.id)
const submitting = ref(false)

const form = ref<any>({
  name: '',
  category_id: '',
  breed_id: '',
  gender: '未知',
  birthday: '',
  bio: '',
  status: 'alive',
  passing_date: '',
  avatar_url: '',
})

const categories = ref<any[]>([])
const breeds = ref<any[]>([])

const loadBreeds = async () => {
  if (!form.value.category_id) { breeds.value = []; return }
  const { data } = await categoriesApi.breeds(form.value.category_id)
  breeds.value = data.data || data
}

const handleAvatar = async (e: Event) => {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (!file) return
  const result = await uploadApi.upload(file)
  form.value.avatar_url = result.url
}

const handleSubmit = async () => {
  submitting.value = true
  try {
    const payload = { ...form.value }
    if (!payload.breed_id) delete payload.breed_id
    if (!payload.passing_date) delete payload.passing_date
    if (isEdit.value) {
      await petStore.updatePet(Number(route.params.id), payload)
    } else {
      await petStore.createPet(payload)
    }
    router.push('/pets')
  } finally {
    submitting.value = false
  }
}

onMounted(async () => {
  const { data } = await categoriesApi.list()
  categories.value = data.data || data

  if (isEdit.value) {
    await petStore.fetchPet(Number(route.params.id))
    const pet = petStore.currentPet
    if (pet) {
      form.value = {
        name: pet.name,
        category_id: pet.category_id,
        breed_id: pet.breed_id,
        gender: pet.gender,
        birthday: pet.birthday,
        bio: pet.bio,
        status: pet.status,
        passing_date: pet.passing_date || '',
        avatar_url: pet.avatar_url,
      }
      await loadBreeds()
    }
  }
})
</script>