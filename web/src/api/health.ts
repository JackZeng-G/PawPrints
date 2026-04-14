import http from './http'
import type { HealthRecord } from '../types/health'

export const healthApi = {
  list: (petId?: number, type?: string, page = 1, pageSize = 20) =>
    http.get<{ data: HealthRecord[]; total: number }>('/health', { params: { pet_id: petId, type, page, page_size: pageSize } }),

  create: (record: Partial<HealthRecord>) => http.post<HealthRecord>('/health', record),

  update: (id: number, record: Partial<HealthRecord>) => http.put<HealthRecord>(`/health/${id}`, record),

  delete: (id: number) => http.delete(`/health/${id}`),

  weightChart: (petId: number) => http.get<HealthRecord[]>(`/health/${petId}/weight-chart`),

  upcoming: (petId: number) => http.get<HealthRecord[]>(`/health/${petId}/upcoming`),
}