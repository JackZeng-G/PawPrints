export interface HealthRecord {
  id: number
  pet_id: number
  type: 'vaccine' | 'deworming' | 'weight' | 'medical'
  title: string
  value: number
  unit: string
  notes: string
  record_date: string
  next_due_date?: string
  vet_name: string
  clinic_name: string
  medication: string
  diagnosis: string
  created_at: string
  updated_at: string
}