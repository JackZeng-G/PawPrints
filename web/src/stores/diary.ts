import { defineStore } from 'pinia'
import { diariesApi } from '../api/diaries'
import type { DiaryEntry, CreateDiaryInput } from '../types/diary'

export const useDiaryStore = defineStore('diary', {
  state: () => ({
    diaries: [] as DiaryEntry[],
    currentDiary: null as DiaryEntry | null,
    total: 0,
    loading: false,
  }),
  actions: {
    async fetchDiaries(petId?: number, keyword?: string) {
      this.loading = true
      try {
        const { data } = await diariesApi.list(petId, keyword)
        this.diaries = data.data
        this.total = data.total
      } finally {
        this.loading = false
      }
    },
    async fetchDiary(id: number) {
      const { data } = await diariesApi.get(id)
      this.currentDiary = data
    },
    async createDiary(input: CreateDiaryInput) {
      await diariesApi.create(input)
      await this.fetchDiaries()
    },
    async updateDiary(id: number, input: CreateDiaryInput) {
      await diariesApi.update(id, input)
      await this.fetchDiaries()
    },
    async deleteDiary(id: number) {
      await diariesApi.delete(id)
      await this.fetchDiaries()
    },
  },
})