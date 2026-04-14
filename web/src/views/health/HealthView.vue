<template>
  <div class="space-y-6">
    <div class="flex items-center gap-4">
      <router-link to="/pets" class="text-gray-500 hover:text-gray-700">&larr; 返回宠物</router-link>
      <h1 class="text-2xl font-bold text-gray-800">健康记录</h1>
    </div>

    <div class="flex items-center justify-between">
      <div class="flex gap-2">
        <button
          v-for="t in types"
          :key="t.value"
          @click="filterType = t.value"
          :class="['px-4 py-2 rounded-lg', filterType === t.value ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-600']"
        >{{ t.label }}</button>
      </div>
      <button @click="showForm = true" class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600">
        + 添加记录
      </button>
    </div>

    <div v-if="loading" class="text-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
    </div>

    <div v-else class="space-y-3">
      <div v-for="record in filteredRecords" :key="record.id" class="bg-white rounded-lg shadow p-4 flex items-center justify-between">
        <div>
          <div class="flex items-center gap-2">
            <span class="font-medium text-gray-800">{{ record.title }}</span>
            <span class="text-xs bg-gray-100 px-2 py-1 rounded">{{ typeLabel(record.type) }}</span>
          </div>
          <div class="text-sm text-gray-500 mt-1">{{ record.record_date }} {{ record.value ? `· ${record.value} ${record.unit}` : '' }}</div>
          <div v-if="record.notes" class="text-sm text-gray-400 mt-1">{{ record.notes }}</div>
        </div>
        <div class="flex gap-2">
          <button @click="editRecord(record)" class="text-blue-500 text-sm hover:underline">编辑</button>
          <button @click="deleteRecord(record.id)" class="text-red-500 text-sm hover:underline">删除</button>
        </div>
      </div>
      <div v-if="filteredRecords.length === 0" class="text-center py-8 text-gray-500">暂无记录</div>
    </div>

    <div v-if="showForm" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-md space-y-4">
        <h2 class="text-lg font-bold">{{ editing ? '编辑记录' : '添加记录' }}</h2>
        <form @submit.prevent="handleSubmit" class="space-y-3">
          <select v-model="form.type" required class="w-full border border-gray-300 rounded-lg px-4 py-2">
            <option value="vaccine">疫苗</option>
            <option value="deworming">驱虫</option>
            <option value="weight">体重</option>
            <option value="medical">就医</option>
          </select>
          <input v-model="form.title" type="text" required placeholder="标题" class="w-full border border-gray-300 rounded-lg px-4 py-2" />
          <input v-model="form.record_date" type="date" required class="w-full border border-gray-300 rounded-lg px-4 py-2" />
          <div class="grid grid-cols-2 gap-3">
            <input v-model.number="form.value" type="number" step="0.1" placeholder="数值" class="border border-gray-300 rounded-lg px-4 py-2" />
            <input v-model="form.unit" type="text" placeholder="单位" class="border border-gray-300 rounded-lg px-4 py-2" />
          </div>
          <input v-model="form.vet_name" type="text" placeholder="兽医" class="w-full border border-gray-300 rounded-lg px-4 py-2" />
          <textarea v-model="form.notes" rows="2" placeholder="备注" class="w-full border border-gray-300 rounded-lg px-4 py-2"></textarea>
          <div class="flex justify-end gap-3">
            <button type="button" @click="closeForm" class="px-4 py-2 border border-gray-300 rounded-lg">取消</button>
            <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600">保存</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { healthApi } from '../../api/health'

const route = useRoute()
const petId = computed(() => Number(route.params.petId))
const records = ref<any[]>([])
const loading = ref(true)
const filterType = ref('')
const showForm = ref(false)
const editing = ref<number | null>(null)

const types = [
  { value: '', label: '全部' },
  { value: 'vaccine', label: '疫苗' },
  { value: 'deworming', label: '驱虫' },
  { value: 'weight', label: '体重' },
  { value: 'medical', label: '就医' },
]

const form = ref<any>({
  type: 'vaccine', title: '', record_date: new Date().toISOString().split('T')[0],
  value: null, unit: '', vet_name: '', notes: '',
})

const filteredRecords = computed(() => {
  if (!filterType.value) return records.value
  return records.value.filter((r: any) => r.type === filterType.value)
})

const typeLabel = (type: string) => {
  const map: Record<string, string> = { vaccine: '疫苗', deworming: '驱虫', weight: '体重', medical: '就医' }
  return map[type] || type
}

const editRecord = (record: any) => {
  editing.value = record.id
  form.value = { ...record }
  showForm.value = true
}

const closeForm = () => {
  showForm.value = false
  editing.value = null
  form.value = { type: 'vaccine', title: '', record_date: new Date().toISOString().split('T')[0], value: null, unit: '', vet_name: '', notes: '' }
}

const handleSubmit = async () => {
  const payload = { ...form.value, pet_id: petId.value }
  if (!payload.value) delete payload.value
  if (editing.value) {
    await healthApi.update(editing.value, payload)
  } else {
    await healthApi.create(payload)
  }
  closeForm()
  await loadRecords()
}

const deleteRecord = async (id: number) => {
  if (!confirm('确定删除？')) return
  await healthApi.delete(id)
  await loadRecords()
}

const loadRecords = async () => {
  loading.value = true
  try {
    const { data } = await healthApi.list(petId.value)
    records.value = data.data || data
  } finally {
    loading.value = false
  }
}

onMounted(loadRecords)
</script>