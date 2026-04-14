<template>
  <Teleport to="body">
    <div class="fixed top-4 right-4 z-[9999] flex flex-col gap-2 pointer-events-none">
      <TransitionGroup name="toast">
        <div
          v-for="toast in toasts"
          :key="toast.id"
          :class="[
            'pointer-events-auto px-5 py-3 rounded-xl shadow-lg text-sm font-medium flex items-center gap-2 min-w-[240px] max-w-[400px] backdrop-blur-sm transition-all duration-300',
            toast.leaving && 'opacity-0 translate-x-4',
            toast.type === 'success' && 'bg-emerald-50/90 text-emerald-700 border border-emerald-200',
            toast.type === 'error' && 'bg-red-50/90 text-red-700 border border-red-200',
            toast.type === 'info' && 'bg-blue-50/90 text-blue-700 border border-blue-200',
          ]"
        >
          <span v-if="toast.type === 'success'">&#10003;</span>
          <span v-else-if="toast.type === 'error'">&#10007;</span>
          <span v-else>&#8505;</span>
          <span>{{ toast.message }}</span>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { emitter } from '../utils/emitter'
const toasts = computed(() => emitter.toasts)
</script>

<style scoped>
.toast-enter-active { transition: all 0.3s ease-out; }
.toast-leave-active { transition: all 0.3s ease-in; }
.toast-enter-from { opacity: 0; transform: translateX(40px); }
.toast-leave-to { opacity: 0; transform: translateX(40px); }
</style>
