import { Fetcher } from "@/util/fetcher";
import { Schedule, ApiResponse, SchedulesResponse, ScheduleResponse, DeleteResponse } from "@/types/schedule";

const url = "schedules";

// 時間割の取得（例: ユーザーIDで取得）
export const getSchedulesApi = async (userId: number): Promise<ApiResponse<SchedulesResponse>> => {
  // GET /schedules/:id
  const response = await Fetcher<ApiResponse<SchedulesResponse>>(`${url}/${userId}`, {
    method: "GET",
    authRequired: true, // ★ここで明示的に指定
    headers: { "Content-Type": "application/json" },
  });
  return response;
};

// 時間割の作成
export const createScheduleApi = async (schedule: Omit<Schedule, "schedule_id" | "created_at" | "updated_at">): Promise<ApiResponse<ScheduleResponse>> => {
  // POST /schedules
  const response = await Fetcher<ApiResponse<ScheduleResponse>>(url, {
    method: "POST",
    body: JSON.stringify(schedule),
    headers: { "Content-Type": "application/json" },
  });

  return response;
};

// 時間割の更新
export const updateScheduleApi = async (schedule: Omit<Schedule, "created_at">): Promise<ApiResponse<ScheduleResponse>> => {
  // schedule_id をパスから渡す形にする
  const { schedule_id, ...updateData } = schedule;
  if (!schedule_id) {
    throw new Error("schedule_id is required for update.");
  }

  // PUT /schedules/:schedule_id
  const response = await Fetcher<ApiResponse<ScheduleResponse>>(`${url}/${schedule_id}`, {
    method: "PUT",
    body: JSON.stringify(updateData),
    headers: { "Content-Type": "application/json" },
  });

  return response;
};

// 時間割の削除
export const deleteScheduleApi = async (schedule_id: number): Promise<ApiResponse<DeleteResponse>> => {
  // DELETE /schedules/:schedule_id
  const response = await Fetcher<ApiResponse<DeleteResponse>>(`${url}/${schedule_id}`, {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
  });

  return response;
};
