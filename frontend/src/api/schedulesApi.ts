import { Fetcher } from "@/util/fetcher";
import {
  Schedule,
  ApiResponse,
  SchedulesResponse,
  ScheduleResponse,
  DeleteResponse,
} from "@/types/schedule";


const  url = "/schedules";

// 時間割の取得
export const getSchedulesApi = async (
  query_params?: Partial<Schedule>
): Promise<ApiResponse<SchedulesResponse>> => {
  const queryString = query_params
    ? "?" +
      Object.entries(query_params)
        .filter(([, value]) => value !== undefined)
        .map(
          ([key, value]) =>
            `${encodeURIComponent(key)}=${encodeURIComponent(
              value as string | number
            )}`
        )
        .join("&")
    : "";

  const response = await Fetcher<ApiResponse<SchedulesResponse>>(
    `${url}${queryString}`,
    {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    }
  );

  return response;
};

// 時間割の作成
export const createScheduleApi = async (
  schedule: Omit<Schedule, "schedule_id" | "created_at" | "updated_at">
): Promise<ApiResponse<ScheduleResponse>> => {
  const response = await Fetcher<ApiResponse<ScheduleResponse>>(
    `${url}`,
    {
      method: "POST",
      body: JSON.stringify(schedule),
      headers: { "Content-Type": "application/json" },
    }
  );

  return response;
};

// 時間割の更新
export const updateScheduleApi = async (
  schedule: Omit<Schedule, "created_at">
): Promise<ApiResponse<ScheduleResponse>> => {
  const response = await Fetcher<ApiResponse<ScheduleResponse>>(`${url}`, {
    method: "PUT",
    body: JSON.stringify(schedule),
    headers: { "Content-Type": "application/json" },
  });

  return response;
};

// 時間割の削除
export const deleteScheduleApi = async (
  schedule_id: number
): Promise<ApiResponse<DeleteResponse>> => {
  const response = await Fetcher<ApiResponse<DeleteResponse>>(
    `${url}?schedule_id=${schedule_id}`,
    {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
    }
  );

  return response;
};
