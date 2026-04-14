<template>
  <Teleport to="body">
    <div class="fixed top-4 right-4 z-[9999] flex flex-col gap-2 pointer-events-none">
      <TransitionGroup name="toast">
        <div v-for="toast in toasts" :key="toast.id"
          :class="[
            'pointer-events-auto px-4 py-3 rounded-xl shadow-lg shadow-gray-200/50 text-sm font-medium flex items-center gap-3 min-w-[280px] max-w-[400px] transition-all duration-300',
            toast.leaving && 'opacity-0 translate-x-4',
            toast.type === 'success' && 'bg-white text-emerald-600 border border-emerald-100',
            toast.type === 'error' && 'bg-white text-red-600 border border-red-100',
            toast.type === 'info' && 'bg-white text-blue-600 border border-blue-100',
          ]"
        >
          <div :class="['w-8 h-8 rounded-lg flex items-center justify-center flex-shrink-0',
            toast.type === 'success' && 'bg-emerald-50',
            toast.type === 'error' && 'bg-red-50',
            toast.type === 'info' && 'bg-blue-50'
          ]"
          >
            <Check v-if="toast.type === 'success'" class="w-4 h-4" />
            <XCircle v-else-if="toast.type === 'error'" class="w-4 h-4" />
            <Info v-else class="w-4 h-4" />
          </div>
          <span class="flex-1">{{ toast.message }}</span>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { emitter } from '../utils/emitter'
import { Check, XCircle, Info } from 'lucide-vue-next'

const toasts = computed(() => emitter.toasts)
</script>

<style scoped>
.toast-enter-active { transition: all 0.3s ease-out; }
.toast-leave-active { transition: all 0.2s ease-in; }
.toast-enter-from { opacity: 0; transform: translateX(40px) scale(0.95); }
.toast-leave-to { opacity: 0; transform: translateX(40px) scale(0.95); }
</style>