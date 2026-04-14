import { ref } from 'vue'

const toasts = ref<{ id: number; message: string; type: 'success' | 'error' | 'info'; leaving: boolean }[]>([])
let nextId = 0

export function useToast() {
  const show = (message: string, type: 'success' | 'error' | 'info' = 'info') => {
    const id = nextId++
    toasts.value.push({ id, message, type, leaving: false })
    setTimeout(() => {
      const t = toasts.value.find(t => t.id === id)
      if (t) t.leaving = true
      setTimeout(() => {
        toasts.value = toasts.value.filter(t => t.id !== id)
      }, 300)
    }, 3000)
  }

  return {
    toasts,
    success: (msg: string) => show(msg, 'success'),
    error: (msg: string) => show(msg, 'error'),
    info: (msg: string) => show(msg, 'info'),
  }
}
