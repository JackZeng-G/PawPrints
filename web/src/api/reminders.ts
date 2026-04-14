import http from './http'
import type { Reminder } from '../types/reminder'

export const remindersApi = {
  list: (completed = false, page = 1, pageSize = 20) =>
    http.get<{ data: Reminder[]; total: number }>('/reminders', { params: { completed, page, page_size: pageSize } }),

  active: () => http.get<Reminder[]>('/reminders/active'),

  create: (reminder: Partial<Reminder>) => http.post<Reminder>('/reminders', reminder),

  update: (id: number, reminder: Partial<Reminder>) => http.put<Reminder>(`/reminders/${id}`, reminder),

  delete: (id: number) => http.delete(`/reminders/${id}`),

  complete: (id: number) => http.put(`/reminders/${id}/complete`),

  snooze: (id: number, minutes = 30) => http.put(`/reminders/${id}/snooze`, { minutes }),
}