<template>
  <div class="max-w-2xl mx-auto space-y-6">
    <h1 class="text-2xl font-bold text-gray-800">设置</h1>

    <div class="bg-white rounded-lg shadow p-6 space-y-6">
      <div>
        <h2 class="text-lg font-bold text-gray-800 mb-3">数据管理</h2>
        <div class="space-y-3">
          <button @click="handleExport" class="w-full text-left px-4 py-3 border border-gray-200 rounded-lg hover:bg-gray-50 flex items-center justify-between">
            <span>导出数据 (JSON)</span>
            <span class="text-gray-400">&rarr;</span>
          </button>
          <button @click="handleBackup" class="w-full text-left px-4 py-3 border border-gray-200 rounded-lg hover:bg-gray-50 flex items-center justify-between">
            <span>备份数据库 (SQLite)</span>
            <span class="text-gray-400">&rarr;</span>
          </button>
          <div>
            <label class="w-full text-left px-4 py-3 border border-gray-200 rounded-lg hover:bg-gray-50 flex items-center justify-between cursor-pointer">
              <span>导入数据</span>
              <span class="text-gray-400">&rarr;</span>
              <input type="file" accept=".json" @change="handleImport" class="hidden" />
            </label>
          </div>
        </div>
      </div>

      <div class="border-t pt-6">
        <h2 class="text-lg font-bold text-gray-800 mb-3">关于</h2>
        <p class="text-gray-600">萌宠档案馆 v1.0</p>
        <p class="text-sm text-gray-400 mt-1">记录宠物的每一个美好瞬间</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { dataApi } from '../../api/common'

const handleExport = async () => {
  const { data } = await dataApi.export()
  const blob = new Blob([JSON.stringify(data, null, 2)], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `pawprints-export-${new Date().toISOString().split('T')[0]}.json`
  a.click()
  URL.revokeObjectURL(url)
}

const handleBackup = async () => {
  const { data } = await dataApi.backup()
  const url = URL.createObjectURL(data)
  const a = document.createElement('a')
  a.href = url
  a.download = `pawprints-backup-${new Date().toISOString().split('T')[0]}.db`
  a.click()
  URL.revokeObjectURL(url)
}

const handleImport = async (e: Event) => {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (!file) return
  if (!confirm('导入将覆盖现有数据，确定继续吗？')) return
  await dataApi.import(file)
  alert('导入成功！')
}
</script>