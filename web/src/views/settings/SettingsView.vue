<template>
  <div class="max-w-2xl mx-auto space-y-6">
    <h1 class="text-2xl font-bold text-gray-800">设置</h1>

    <div class="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden">
      <!-- Data Management -->
      <div class="p-6">
        <h2 class="text-lg font-bold text-gray-800 mb-4 flex items-center gap-2">
          <Database class="w-5 h-5 text-gray-400" /> 数据管理
        </h2>
        <div class="space-y-2">
          <button @click="handleExport"
            class="w-full flex items-center justify-between px-4 py-3.5 border border-gray-100 rounded-xl hover:bg-gray-50 hover:border-gray-200 transition-all duration-200 group"
          >
            <div class="flex items-center gap-3">
              <div class="w-9 h-9 bg-blue-50 rounded-lg flex items-center justify-center">
                <Download class="w-4 h-4 text-blue-500" />
              </div>
              <div class="text-left">
                <div class="text-sm font-medium text-gray-700">导出数据</div>
                <div class="text-xs text-gray-400">JSON 格式</div>
              </div>
            </div>
            <ChevronRight class="w-4 h-4 text-gray-300 group-hover:text-gray-400 transition-colors" />
          </button>

          <button @click="handleBackup"
            class="w-full flex items-center justify-between px-4 py-3.5 border border-gray-100 rounded-xl hover:bg-gray-50 hover:border-gray-200 transition-all duration-200 group"
          >
            <div class="flex items-center gap-3">
              <div class="w-9 h-9 bg-green-50 rounded-lg flex items-center justify-center">
                <HardDrive class="w-4 h-4 text-green-500" />
              </div>
              <div class="text-left">
                <div class="text-sm font-medium text-gray-700">备份数据库</div>
                <div class="text-xs text-gray-400">SQLite 格式</div>
              </div>
            </div>
            <ChevronRight class="w-4 h-4 text-gray-300 group-hover:text-gray-400 transition-colors" />
          </button>

          <label class="w-full flex items-center justify-between px-4 py-3.5 border border-gray-100 rounded-xl hover:bg-gray-50 hover:border-gray-200 transition-all duration-200 group cursor-pointer"
          >
            <div class="flex items-center gap-3">
              <div class="w-9 h-9 bg-amber-50 rounded-lg flex items-center justify-center">
                <Upload class="w-4 h-4 text-amber-500" />
              </div>
              <div class="text-left">
                <div class="text-sm font-medium text-gray-700">导入数据</div>
                <div class="text-xs text-gray-400">从 JSON 文件恢复</div>
              </div>
            </div>
            <ChevronRight class="w-4 h-4 text-gray-300 group-hover:text-gray-400 transition-colors" />
            <input type="file" accept=".json" @change="handleImport" class="hidden" />
          </label>
        </div>
      </div>

      <!-- About -->
      <div class="border-t border-gray-100 p-6">
        <h2 class="text-lg font-bold text-gray-800 mb-3 flex items-center gap-2">
          <Info class="w-5 h-5 text-gray-400" /> 关于
        </h2>
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 bg-orange-50 rounded-xl flex items-center justify-center">
            <PawPrint class="w-5 h-5 text-orange-500" />
          </div>
          <div>
            <p class="font-medium text-gray-800">萌宠档案馆</p>
            <p class="text-xs text-gray-400">v1.0 · 记录宠物的每一个美好瞬间</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { dataApi } from '../../api/common'
import { Database, Download, HardDrive, Upload, ChevronRight, Info, PawPrint } from 'lucide-vue-next'

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