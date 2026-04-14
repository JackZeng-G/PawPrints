<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center gap-4">
      <router-link to="/pets" class="inline-flex items-center gap-1 text-sm text-gray-500 hover:text-gray-700 transition-colors">
        <ChevronLeft class="w-4 h-4" /> 返回宠物
      </router-link>
      <h1 class="text-2xl font-bold text-gray-800">健康记录</h1>
    </div>

    <!-- Filter & Action -->
    <div class="flex items-center justify-between gap-4">
      <div class="flex gap-2 bg-white p-1.5 rounded-xl border border-gray-100 shadow-sm flex-1">
        <button v-for="t in types" :key="t.value" @click="filterType = t.value"
          class="flex-1 px-3 py-2 rounded-lg text-sm font-medium transition-all duration-200"
          :class="filterType === t.value ? 'bg-blue-500 text-white shadow-sm' : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'"
        >{{ t.label }}</button>
      </div>
      <button @click="showForm = true"
        class="inline-flex items-center gap-2 bg-blue-500 text-white px-4 py-2.5 rounded-xl hover:bg-blue-600 active:scale-[0.98] transition-all duration-200 shadow-sm shadow-blue-500/20 font-medium text-sm flex-shrink-0">
        <Plus class="w-4 h-4" /> 添加
      </button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center py-16 gap-4">
      <div class="w-10 h-10 border-3 border-gray-200 border-t-blue-500 rounded-full animate-spin"></div>
      <p class="text-sm text-gray-400">加载中...</p>
    </div>

    <template v-else>
      <!-- Weight Chart -->
      <div v-if="filterType === 'weight' && weightData.length > 1" class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5">
        <h3 class="text-sm font-medium text-gray-500 mb-4 flex items-center gap-2">
          <TrendingUp class="w-4 h-4 text-blue-500" /> 体重趋势
        </h3>
        <div ref="chartRef" style="width: 100%; height: 280px;"></div>
      </div>

      <!-- Records List -->
      <div v-if="filteredRecords.length === 0" class="flex flex-col items-center py-16 text-center">
        <Activity class="w-12 h-12 text-gray-200 mb-4" />
        <p class="text-gray-500 mb-1">暂无记录</p>
        <p class="text-sm text-gray-400">点击上方按钮添加健康记录</p>
      </div>

      <div v-else class="space-y-3">
        <div v-for="record in filteredRecords" :key="record.id"
          class="bg-white rounded-2xl border border-gray-100 p-4 flex items-center justify-between hover:shadow-sm transition-shadow duration-200">
          <div class="flex items-start gap-3">
            <div :class="['w-10 h-10 rounded-xl flex items-center justify-center flex-shrink-0', typeBgClass(record.type)]">
              <component :is="typeIcon(record.type)" class="w-5 h-5" :class="typeTextClass(record.type)" />
            </div>
            <div>
              <div class="flex items-center gap-2">
                <span class="font-medium text-gray-800">{{ record.title }}</span>
                <span class="text-xs bg-gray-100 text-gray-500 px-2 py-0.5 rounded-md">{{ typeLabel(record.type) }}</span>
              </div>
              <div class="text-sm text-gray-400 mt-1 flex items-center gap-2">
                <span>{{ record.record_date }}</span>
                <span v-if="record.value" class="bg-gray-50 px-2 py-0.5 rounded-md text-gray-600">{{ record.value }} {{ record.unit }}</span>
              </div>
              <div v-if="record.notes" class="text-sm text-gray-400 mt-1">{{ record.notes }}</div>
            </div>
          </div>
          <div class="flex gap-1">
            <button @click="editRecord(record)" class="p-2 text-gray-400 hover:text-blue-500 hover:bg-blue-50 rounded-lg transition-all duration-200">
              <Pencil class="w-4 h-4" />
            </button>
            <button @click="deleteRecord(record.id)" class="p-2 text-gray-400 hover:text-red-500 hover:bg-red-50 rounded-lg transition-all duration-200">
              <Trash2 class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    </template>

    <!-- Form Modal -->
    <Teleport to="body">
      <Transition name="modal">
        <div v-if="showForm" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4" @click.self="closeForm">
          <div class="bg-white rounded-2xl shadow-xl w-full max-w-md overflow-hidden">
            <div class="flex items-center justify-between px-6 py-4 border-b border-gray-100">
              <h2 class="text-lg font-bold text-gray-800">{{ editing ? '编辑记录' : '添加记录' }}</h2>
              <button @click="closeForm" class="p-1 text-gray-400 hover:text-gray-600 rounded-lg hover:bg-gray-100 transition-colors">
                <X class="w-5 h-5" />
              </button>
            </div>
            <form @submit.prevent="handleSubmit" class="p-6 space-y-4">
              <select v-model="form.type" required class="w-full border border-gray-200 rounded-xl px-4 py-2.5 bg-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all">
                <option value="vaccine">疫苗</option>
                <option value="deworming">驱虫</option>
                <option value="weight">体重</option>
                <option value="medical">就医</option>
              </select>
              <input v-model="form.title" type="text" required placeholder="标题" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all" />
              <input v-model="form.record_date" type="date" required class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all" />
              <div class="grid grid-cols-2 gap-3">
                <input v-model.number="form.value" type="number" step="0.1" placeholder="数值" class="border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all" />
                <input v-model="form.unit" type="text" placeholder="单位" class="border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all" />
              </div>
              <input v-model="form.vet_name" type="text" placeholder="兽医" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all" />
              <textarea v-model="form.notes" rows="2" placeholder="备注" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"></textarea>
              <div class="flex justify-end gap-3 pt-2">
                <button type="button" @click="closeForm" class="px-5 py-2.5 border border-gray-200 rounded-xl text-gray-600 hover:bg-gray-50 font-medium text-sm transition-all">取消</button>
                <button type="submit" class="px-5 py-2.5 bg-blue-500 text-white rounded-xl hover:bg-blue-600 font-medium text-sm shadow-sm shadow-blue-500/20 transition-all">保存</button>
              </div>
            </form>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { healthApi } from '../../api/health'
import {
  ChevronLeft, Plus, TrendingUp, Activity, Syringe, ShieldCheck, Scale, Stethoscope,
  Pencil, Trash2, X
} from 'lucide-vue-next'

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

const typeLabel = (type: string) => {
  const map: Record<string, string> = { vaccine: '疫苗', deworming: '驱虫', weight: '体重', medical: '就医' }
  return map[type] || type
}

const typeIcon = (type: string) => {
  const map: Record<string, any> = { vaccine: Syringe, deworming: ShieldCheck, weight: Scale, medical: Stethoscope }
  return map[type] || Activity
}

const typeBgClass = (type: string) => {
  const map: Record<string, string> = { vaccine: 'bg-green-50', deworming: 'bg-purple-50', weight: 'bg-blue-50', medical: 'bg-red-50' }
  return map[type] || 'bg-gray-50'
}

const typeTextClass = (type: string) => {
  const map: Record<string, string> = { vaccine: 'text-green-500', deworming: 'text-purple-500', weight: 'text-blue-500', medical: 'text-red-500' }
  return map[type] || 'text-gray-500'
}

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

<style scoped>
.modal-enter-active { transition: all 0.3s ease-out; }
.modal-leave-active { transition: all 0.2s ease-in; }
.modal-enter-from { opacity: 0; }
.modal-leave-to { opacity: 0; }
.modal-enter-from > div { transform: scale(0.95) translateY(10px); }
.modal-leave-to > div { transform: scale(0.95) translateY(10px); }
</style>