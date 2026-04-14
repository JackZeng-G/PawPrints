export interface Reminder {
  id: number
  title: string
  description: string
  remind_at: string
  type: 'vaccine_due' | 'deworming_due' | 'custom'
  pet_id: number
  health_record_id: number
  completed: boolean
  completed_at?: string
  snoozed_until?: string
  created_at: string
  updated_at: string
}
