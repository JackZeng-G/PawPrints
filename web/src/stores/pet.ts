import { defineStore } from 'pinia'
import { petsApi } from '../api/pets'
import type { Pet } from '../types/pet'

export const usePetStore = defineStore('pet', {
  state: () => ({
    pets: [] as Pet[],
    currentPet: null as Pet | null,
    total: 0,
    loading: false,
  }),
  actions: {
    async fetchPets(status?: string) {
      this.loading = true
      try {
        const { data } = await petsApi.list(status)
        this.pets = data.data
        this.total = data.total
      } finally {
        this.loading = false
      }
    },
    async fetchPet(id: number) {
      const { data } = await petsApi.get(id)
      this.currentPet = data
    },
    async createPet(pet: Partial<Pet>) {
      await petsApi.create(pet)
      await this.fetchPets()
    },
    async updatePet(id: number, pet: Partial<Pet>) {
      await petsApi.update(id, pet)
      await this.fetchPets()
    },
    async deletePet(id: number) {
      await petsApi.delete(id)
      await this.fetchPets()
    },
  },
})
