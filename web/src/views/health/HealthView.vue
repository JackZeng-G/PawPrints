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

    <!-- Weight Chart -->
    <div v-if="filterType === 'weight' && weightData.length > 1" class="bg-white rounded-2xl shadow-sm p-4">
      <h3 class="text-sm font-medium text-gray-500 mb-3">体重趋势</h3>
      <div ref="chartRef" style="width: 100%; height: 280px;"></div>
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
import { ref, computed, onMounted, watch, nextTick } from 'vue'
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

const chartRef = ref<HTMLDivElement | null>(null)
const weightData = ref<{ date: string; value: number }[]>([])
let chart: any = null

const initChart = async () => {
  if (!chartRef.value || weightData.value.length < 2) return
  const echarts = await import('echarts')
  if (!chart) chart = echarts.init(chartRef.value)
  const option = {
    grid: { left: 40, right: 20, top: 20, bottom: 30 },
    xAxis: { type: 'category', data: weightData.value.map(d => d.date), axisLine: { lineStyle: { color: '#e5e7eb' } }, axisLabel: { color: '#9ca3af' } },
    yAxis: { type: 'value', name: 'kg', splitLine: { lineStyle: { color: '#f3f4f6' } }, axisLabel: { color: '#9ca3af' } },
    series: [{
      type: 'line',
      data: weightData.value.map(d => d.value),
      smooth: true,
      lineStyle: { color: '#3b82f6', width: 3 },
      itemStyle: { color: '#3b82f6' },
      areaStyle: { color: { type: 'linear', x: 0, y: 0, x2: 0, y2: 1, colorStops: [{ offset: 0, color: 'rgba(59, 130, 246, 0.3)' }, { offset: 1, color: 'rgba(59, 130, 246, 0.05)' }] } },
    }],
    tooltip: { trigger: 'axis', formatter: (p: any) => `${p[0].name}<br/>体重: ${p[0].value} kg` },
  }
  chart.setOption(option)
}

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

watch(filterType, async () => {
  if (filterType.value === 'weight' && weightData.value.length > 1) {
    await nextTick()
    initChart()
  }
})

const loadRecords = async () => {
  loading.value = true
  try {
    const { data } = await healthApi.list(petId.value)
    records.value = data.data || data
    // Extract weight data for chart
    weightData.value = records.value
      .filter((r: any) => r.type === 'weight' && r.value)
      .map((r: any) => ({ date: r.record_date, value: r.value }))
      .sort((a: any, b: any) => String(a.date).localeCompare(String(b.date)))
    if (filterType.value === 'weight' && weightData.value.length > 1) {
      await nextTick()
      initChart()
    }
  } finally {
    loading.value = false
  }
}

onMounted(loadRecords)
</script>