export interface Pet {
  id: number
  name: string
  category_id: number
  breed_id: number
  gender: string
  birthday: string
  passing_date?: string
  avatar_url: string
  bio: string
  status: 'alive' | 'memorial'
  created_at: string
  updated_at: string
  category: PetCategory
  breed?: PetBreed
}

export interface PetCategory {
  id: number
  name: string
  icon: string
  sort_order: number
}

export interface PetBreed {
  id: number
  category_id: number
  name: string
  expected_lifespan: string
  size: string
}
