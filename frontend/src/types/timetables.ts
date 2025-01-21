export interface Schedule {
  schedule_id: number;
  user_id: number;
  day_of_week: number;
  time_slot: number;
  subject: string;
  location: string;
  schedule_url?: string;
  created_at: string;
  updated_at: string;
}

export interface TimetableProps {
  schedules: Schedule[];
  className?: string;
}
