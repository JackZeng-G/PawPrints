<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <h1 class="text-2xl font-bold text-gray-800">提醒</h1>
      <button @click="showForm = true"
        class="inline-flex items-center gap-2 bg-orange-500 text-white px-4 py-2.5 rounded-xl hover:bg-orange-600 active:scale-[0.98] transition-all duration-200 shadow-sm shadow-orange-500/20 font-medium text-sm"
      >
        <Plus class="w-4 h-4" /> 新建
      </button>
    </div>

    <!-- Tabs -->
    <div class="flex gap-2 bg-white p-1.5 rounded-xl border border-gray-100 shadow-sm">
      <button v-for="t in tabOptions" :key="t.value" @click="tab = t.value"
        class="flex-1 px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200"
        :class="tab === t.value ? 'bg-orange-500 text-white shadow-sm' : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'"
      >{{ t.label }}</button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center py-16 gap-4">
      <div class="w-10 h-10 border-3 border-gray-200 border-t-orange-500 rounded-full animate-spin"></div>
      <p class="text-sm text-gray-400">加载中...</p>
    </div>

    <template v-else>
      <!-- Empty State -->
      <div v-if="reminders.length === 0" class="flex flex-col items-center py-16 text-center">
        <div class="w-20 h-20 bg-orange-50 rounded-full flex items-center justify-center mb-4">
          <Bell class="w-10 h-10 text-orange-300" />
        </div>
        <p class="text-gray-500 mb-1">暂无提醒</p>
        <p class="text-sm text-gray-400">点击上方按钮创建提醒</p>
      </div>

      <!-- Reminder List -->
      <div v-else class="space-y-3">
        <div v-for="reminder in reminders" :key="reminder.id"
          class="bg-white rounded-2xl border border-gray-100 p-4 hover:shadow-sm transition-shadow duration-200"
        >
          <div class="flex items-start justify-between">
            <div class="flex items-start gap-3">
              <div :class="['w-10 h-10 rounded-xl flex items-center justify-center flex-shrink-0 mt-0.5',
                reminder.completed ? 'bg-green-50' : typeBgClass(reminder.type)]">
                <component :is="reminder.completed ? CheckCircle2 : Bell"
                  :class="['w-5 h-5', reminder.completed ? 'text-green-500' : typeTextClass(reminder.type)]" />
              </div>
              <div>
                <div class="flex items-center gap-2">
                  <span :class="['font-medium text-sm', reminder.completed ? 'text-gray-400 line-through' : 'text-gray-800']">
                    {{ reminder.title }}
                  </span>
                  <span v-if="reminder.completed" class="text-xs bg-green-50 text-green-600 px-2 py-0.5 rounded-md font-medium">已完成</span>
                  <span v-else :class="['text-xs px-2 py-0.5 rounded-md font-medium', typeBadgeClass(reminder.type)]">
                    {{ typeLabel(reminder.type) }}
                  </span>
                </div>
                <div class="text-xs text-gray-400 mt-1 flex items-center gap-1.5">
                  <Clock class="w-3 h-3" /> {{ reminder.remind_at }}
                </div>
                <div v-if="reminder.description" class="text-xs text-gray-400 mt-1">{{ reminder.description }}</div>
              </div>
            </div>
            <div v-if="!reminder.completed" class="flex gap-1">
              <button @click="completeReminder(reminder.id)"
                class="p-2 text-gray-400 hover:text-green-500 hover:bg-green-50 rounded-lg transition-all duration-200"
                title="完成">
                <Check class="w-4 h-4" />
              </button>
              <button @click="snoozeReminder(reminder.id)"
                class="p-2 text-gray-400 hover:text-blue-500 hover:bg-blue-50 rounded-lg transition-all duration-200"
                title="稍后提醒">
                <Clock class="w-4 h-4" />
              </button>
              <button @click="deleteReminder(reminder.id)"
                class="p-2 text-gray-400 hover:text-red-500 hover:bg-red-50 rounded-lg transition-all duration-200"
                title="删除">
                <Trash2 class="w-4 h-4" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </template>

    <!-- Form Modal -->
    <Teleport to="body">
      <Transition name="modal">
        <div v-if="showForm" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4" @click.self="showForm = false">
          <div class="bg-white rounded-2xl shadow-xl w-full max-w-md overflow-hidden">
            <div class="flex items-center justify-between px-6 py-4 border-b border-gray-100">
              <h2 class="text-lg font-bold text-gray-800">新建提醒</h2>
              <button @click="showForm = false" class="p-1 text-gray-400 hover:text-gray-600 rounded-lg hover:bg-gray-100 transition-colors">
                <X class="w-5 h-5" />
              </button>
            </div>
            <form @submit.prevent="handleSubmit" class="p-6 space-y-4">
              <input v-model="form.title" type="text" required placeholder="标题"
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all" />
              <textarea v-model="form.description" rows="2" placeholder="描述"
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all resize-y"></textarea>
              <input v-model="form.remind_at" type="datetime-local" required
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all" />
              <select v-model="form.type"
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 bg-white focus:ring-2 focus:ring-orange-500 focus:border-orange-500 transition-all"
              >
                <option value="custom">自定义</option>
                <option value="vaccine_due">疫苗到期</option>
                <option value="deworming_due">驱虫到期</option>
              </select>
              <div class="flex justify-end gap-3 pt-2">
                <button type="button" @click="showForm = false"
                  class="px-5 py-2.5 border border-gray-200 rounded-xl text-gray-600 hover:bg-gray-50 font-medium text-sm transition-all">
                  取消
                </button>
                <button type="submit"
                  class="px-5 py-2.5 bg-orange-500 text-white rounded-xl hover:bg-orange-600 font-medium text-sm shadow-sm shadow-orange-500/20 transition-all">
                  保存
                </button>
              </div>
            </form>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { remindersApi } from '../../api/reminders'
import { Plus, Bell, Check, CheckCircle2, Clock, Trash2, X } from 'lucide-vue-next'

const reminders = ref<any[]>([])
const loading = ref(true)
const tab = ref('active')
const showForm = ref(false)

const tabOptions = [
  { value: 'active', label: '进行中' },
  { value: 'all', label: '全部' },
]

const form = ref<any>({
  title: '', description: '', remind_at: '', type: 'custom', pet_id: 0,
})

const typeLabel = (type: string) => ({
  custom: '自定义', vaccine_due: '疫苗', deworming_due: '驱虫'
}[type] || type)

const typeBgClass = (type: string) => ({
  vaccine_due: 'bg-green-50', deworming_due: 'bg-purple-50', custom: 'bg-amber-50'
}[type] || 'bg-gray-50')

const typeTextClass = (type: string) => ({
  vaccine_due: 'text-green-500', deworming_due: 'text-purple-500', custom: 'text-amber-500'
}[type] || 'text-gray-500')

const typeBadgeClass = (type: string) => ({
  vaccine_due: 'bg-green-50 text-green-600', deworming_due: 'bg-purple-50 text-purple-600', custom: 'bg-amber-50 text-amber-600'
}[type] || 'bg-gray-50 text-gray-600')

const loadReminders = async () => {
  loading.value = true
  try {
    if (tab.value === 'active') {
      const { data } = await remindersApi.active()
      reminders.value = Array.isArray(data) ? data : (data as any).data || []
    } else {
      const { data } = await remindersApi.list()
      reminders.value = data.data || data
    }
  } finally {
    loading.value = false
  }
}

const completeReminder = async (id: number) => {
  await remindersApi.complete(id)
  await loadReminders()
}

const snoozeReminder = async (id: number) => {
  await remindersApi.snooze(id, 60)
  await loadReminders()
}

const deleteReminder = async (id: number) => {
  if (!confirm('确定删除？')) return
  await remindersApi.delete(id)
  await loadReminders()
}

const handleSubmit = async () => {
  await remindersApi.create(form.value)
  showForm.value = false
  form.value = { title: '', description: '', remind_at: '', type: 'custom', pet_id: 0 }
  await loadReminders()
}

watch(tab, loadReminders)
onMounted(loadReminders)
</script>

<style scoped>
.modal-enter-active { transition: all 0.3s ease-out; }
.modal-leave-active { transition: all 0.2s ease-in; }
.modal-enter-from { opacity: 0; }
.modal-leave-to { opacity: 0; }
.modal-enter-from > div { transform: scale(0.95) translateY(10px); }
.modal-leave-to > div { transform: scale(0.95) translateY(10px); }
</style>