<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold text-gray-800">提醒</h1>

    <div class="flex gap-2">
      <button
        @click="tab = 'active'"
        :class="['px-4 py-2 rounded-lg', tab === 'active' ? 'bg-orange-500 text-white' : 'bg-gray-100 text-gray-600']"
      >进行中</button>
      <button
        @click="tab = 'all'"
        :class="['px-4 py-2 rounded-lg', tab === 'all' ? 'bg-orange-500 text-white' : 'bg-gray-100 text-gray-600']"
      >全部</button>
    </div>

    <button @click="showForm = true" class="bg-orange-500 text-white px-4 py-2 rounded-lg hover:bg-orange-600">
      + 新建提醒
    </button>

    <div v-if="loading" class="text-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-orange-500 mx-auto"></div>
    </div>

    <div v-else class="space-y-3">
      <div v-for="reminder in reminders" :key="reminder.id" class="bg-white rounded-lg shadow p-4 flex items-center justify-between">
        <div>
          <div class="flex items-center gap-2">
            <span :class="['font-medium', reminder.completed ? 'text-gray-400 line-through' : 'text-gray-800']">
              {{ reminder.title }}
            </span>
            <span v-if="reminder.completed" class="text-xs bg-green-100 text-green-600 px-2 py-1 rounded">已完成</span>
          </div>
          <div class="text-sm text-gray-500 mt-1">
            {{ reminder.remind_at }}
            <span v-if="reminder.description"> · {{ reminder.description }}</span>
          </div>
        </div>
        <div v-if="!reminder.completed" class="flex gap-2">
          <button @click="completeReminder(reminder.id)" class="text-green-500 text-sm hover:underline">完成</button>
          <button @click="snoozeReminder(reminder.id)" class="text-blue-500 text-sm hover:underline">稍后</button>
          <button @click="deleteReminder(reminder.id)" class="text-red-500 text-sm hover:underline">删除</button>
        </div>
      </div>
      <div v-if="reminders.length === 0" class="text-center py-8 text-gray-500">暂无提醒</div>
    </div>

    <div v-if="showForm" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div class="bg-white rounded-lg shadow-xl p-6 w-full max-w-md space-y-4">
        <h2 class="text-lg font-bold">新建提醒</h2>
        <form @submit.prevent="handleSubmit" class="space-y-3">
          <input v-model="form.title" type="text" required placeholder="标题" class="w-full border border-gray-300 rounded-lg px-4 py-2" />
          <textarea v-model="form.description" rows="2" placeholder="描述" class="w-full border border-gray-300 rounded-lg px-4 py-2"></textarea>
          <input v-model="form.remind_at" type="datetime-local" required class="w-full border border-gray-300 rounded-lg px-4 py-2" />
          <select v-model="form.type" class="w-full border border-gray-300 rounded-lg px-4 py-2">
            <option value="custom">自定义</option>
            <option value="vaccine_due">疫苗到期</option>
            <option value="deworming_due">驱虫到期</option>
          </select>
          <div class="flex justify-end gap-3">
            <button type="button" @click="showForm = false" class="px-4 py-2 border border-gray-300 rounded-lg">取消</button>
            <button type="submit" class="px-4 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600">保存</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { remindersApi } from '../../api/reminders'

const reminders = ref<any[]>([])
const loading = ref(true)
const tab = ref('active')
const showForm = ref(false)
const form = ref<any>({
  title: '', description: '', remind_at: '', type: 'custom', pet_id: 0,
})

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