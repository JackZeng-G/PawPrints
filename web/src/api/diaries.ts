import http from './http'
import type { DiaryEntry, CreateDiaryInput } from '../types/diary'

export const diariesApi = {
  list: (petId?: number, keyword?: string, page = 1, pageSize = 20) =>
    http.get<{ data: DiaryEntry[]; total: number }>('/diaries', { params: { pet_id: petId, keyword, page, page_size: pageSize } }),

  get: (id: number) => http.get<DiaryEntry>(`/diaries/${id}`),

  create: (input: CreateDiaryInput) => http.post<DiaryEntry>('/diaries', input),

  update: (id: number, input: CreateDiaryInput) => http.put<DiaryEntry>(`/diaries/${id}`, input),

  delete: (id: number) => http.delete(`/diaries/${id}`),
}