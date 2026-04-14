<template>
  <div class="max-w-2xl mx-auto space-y-6">
    <div class="flex items-center gap-4">
      <router-link to="/pets" class="text-gray-500 hover:text-gray-700">&larr; 返回</router-link>
      <h1 class="text-2xl font-bold text-gray-800">{{ isEdit ? '编辑宠物' : '添加宠物' }}</h1>
    </div>

    <form @submit.prevent="handleSubmit" class="bg-white rounded-lg shadow p-6 space-y-4">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">名字 *</label>
        <input v-model="form.name" type="text" required class="w-full border border-gray-300 rounded-lg px-4 py-2 focus:ring-2 focus:ring-orange-500 focus:border-orange-500" />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">头像</label>
        <input type="file" accept="image/*" @change="handleAvatar" class="w-full text-sm text-gray-500" />
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">种类 *</label>
          <select v-model="form.category_id" required @change="loadBreeds" class="w-full border border-gray-300 rounded-lg px-4 py-2">
            <option value="">请选择</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">品种</label>
          <select v-model="form.breed_id" class="w-full border border-gray-300 rounded-lg px-4 py-2">
            <option value="">请选择</option>
            <option v-for="breed in breeds" :key="breed.id" :value="breed.id">{{ breed.name }}</option>
          </select>
        </div>
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">性别 *</label>
          <select v-model="form.gender" required class="w-full border border-gray-300 rounded-lg px-4 py-2">
            <option value="公">公</option>
            <option value="母">母</option>
            <option value="未知">未知</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">生日 *</label>
          <input v-model="form.birthday" type="date" required class="w-full border border-gray-300 rounded-lg px-4 py-2" />
        </div>
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">简介</label>
        <textarea v-model="form.bio" rows="3" class="w-full border border-gray-300 rounded-lg px-4 py-2"></textarea>
      </div>

      <div v-if="isEdit" class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
          <select v-model="form.status" class="w-full border border-gray-300 rounded-lg px-4 py-2">
            <option value="alive">在世</option>
            <option value="memorial">纪念</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">离别日期</label>
          <input v-model="form.passing_date" type="date" class="w-full border border-gray-300 rounded-lg px-4 py-2" />
        </div>
      </div>

      <div class="flex justify-end gap-3 pt-4">
        <router-link to="/pets" class="px-6 py-2 border border-gray-300 rounded-lg text-gray-600 hover:bg-gray-50">取消</router-link>
        <button type="submit" :disabled="submitting" class="px-6 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600 disabled:opacity-50">
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