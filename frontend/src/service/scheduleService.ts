import { createScheduleApi, updateScheduleApi, getSchedulesApi, deleteScheduleApi } from "@/api/schedulesApi";
import { Schedule, ApiResponse, SchedulesResponse, ScheduleResponse, DeleteResponse } from "@/types/schedule";

interface CreateScheduleParams {
  user_id: number;
  day_of_week: number;
  time_slot: number;
  subject: string;
  location: string;
  schedule_url: string;
}

interface UpdateScheduleParams extends CreateScheduleParams {
  schedule_id: number;
}

export const createSchedule = async (params: CreateScheduleParams): Promise<ApiResponse<ScheduleResponse>> => {
  try {
    return await createScheduleApi(params);
  } catch (error) {
    return {
      status: "error",
      message: "Failed to create schedule",
      error_detail: error instanceof Error ? error.message : "Unknown error",
      data: {} as ScheduleResponse,
    };
  }
};

export const updateSchedule = async (params: UpdateScheduleParams): Promise<ApiResponse<ScheduleResponse>> => {
  try {
    return await updateScheduleApi(params);
  } catch (error) {
    return {
      status: "error",
      message: "Failed to update schedule",
      error_detail: error instanceof Error ? error.message : "Unknown error",
      data: {} as ScheduleResponse,
    };
  }
};

export const fetchSchedules = async (query_params?: Partial<Schedule>): Promise<ApiResponse<SchedulesResponse>> => {
  try {
    return await getSchedulesApi(query_params as number);
  } catch (error) {
    return {
      status: "error",
      message: "Failed to fetch schedules",
      error_detail: error instanceof Error ? error.message : "Unknown error",
      data: { schedules: [], schedules_total: 0 },
    };
  }
};

export const deleteSchedule = async (schedule_id: number): Promise<DeleteResponse> => {
  try {
    return await deleteScheduleApi(schedule_id);
  } catch (error) {
    return {
      status: "error",
      message: "Failed to delete schedule",
      error_detail: error instanceof Error ? error.message : "Unknown error",
    };
  }
};
