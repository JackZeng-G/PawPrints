export interface DiaryEntry {
  id: number
  title: string
  content: string
  entry_date: string
  mood: string
  created_at: string
  updated_at: string
  pets: DiaryPet[]
  photos: DiaryPhoto[]
}

export interface DiaryPet {
  id: number
  pet_id: number
  pet_name: string
}

export interface DiaryPhoto {
  id: number
  photo_url: string
  thumbnail_url: string
  sort_order: number
}

export interface CreateDiaryInput {
  title: string
  content: string
  entry_date: string
  mood?: string
  pet_ids: number[]
  photos?: { photo_url: string; thumbnail_url: string }[]
}