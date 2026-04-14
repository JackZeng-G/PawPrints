import http from './http'
import type { Pet } from '../types/pet'

export const petsApi = {
  list: (status?: string, page = 1, pageSize = 20) =>
    http.get<{ data: Pet[]; total: number }>('/pets', { params: { status, page, page_size: pageSize } }),

  get: (id: number) => http.get<Pet>(`/pets/${id}`),

  create: (pet: Partial<Pet>) => http.post<Pet>('/pets', pet),

  update: (id: number, pet: Partial<Pet>) => http.put<Pet>(`/pets/${id}`, pet),

  delete: (id: number) => http.delete(`/pets/${id}`),

  setMemorial: (id: number) => http.put(`/pets/${id}/memorial`),

  getPhotos: (id: number) => http.get(`/pets/${id}/photos`),
}