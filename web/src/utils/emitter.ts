import { reactive, readonly } from 'vue'

type ToastEvent = { type: 'success' | 'error' | 'info'; message: string }

const state = reactive<{ toasts: { id: number; message: string; type: 'success' | 'error' | 'info'; leaving: boolean }[] }>({
  toasts: [],
})
let nextId = 0

export const emitter = {
  emit(_type: 'toast', data: ToastEvent) {
    const id = nextId++
    state.toasts.push({ id, message: data.message, type: data.type, leaving: false })
    setTimeout(() => {
      const t = state.toasts.find(t => t.id === id)
      if (t) t.leaving = true
      setTimeout(() => {
        state.toasts = state.toasts.filter(t => t.id !== id)
      }, 300)
    }, 3000)
  },
  get toasts() { return readonly(state).toasts },
}
