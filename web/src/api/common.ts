import http from './http'

export const uploadApi = {
  upload: async (file: File) => {
    const formData = new FormData()
    formData.append('file', file)
    const res = await http.post<{ url: string; thumbnail_url: string }>('/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    return res.data
  },
}

export const categoriesApi = {
  list: () => http.get('/categories'),
  breeds: (categoryId: number) => http.get(`/categories/${categoryId}/breeds`),
}

export const statsApi = {
  overview: () => http.get('/stats/overview'),
  timeline: (days = 30) => http.get('/stats/timeline', { params: { days } }),
}

export const dataApi = {
  export: () => http.get('/data/export'),
  backup: () => http.get('/data/backup', { responseType: 'blob' }),
  import: (file: File) => {
    const formData = new FormData()
    formData.append('file', file)
    return http.post('/data/import', formData, { headers: { 'Content-Type': 'multipart/form-data' } })
  },
}
