// 時間割の個別データ型
export type Schedule = {
  schedule_id: number;
  user_id: number;
  day_of_week: number;
  time_slot: number;
  subject: string;
  location: string;
  schedule_url: string;
  created_at?: string;
  updated_at?: string;
};

// APIのレスポンス型
export type ApiResponse<T> = {
  status: string;
  message: string;
  error_detail: string;
  data: T;
};

// 時間割一覧取得時のデータ型
export type SchedulesResponse = {
  schedules: Schedule[];
  schedules_total: number;
};

// 時間割作成・更新時のデータ型
export type ScheduleResponse = Schedule;

// 削除時のレスポンス型
export type DeleteResponse = {
  status: string;
  message: string;
  error_detail: string;
};
