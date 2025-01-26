import {
  createScheduleApi,
  updateScheduleApi,
  getSchedulesApi,
  deleteScheduleApi,
} from "@/api/schedulesApi";
import {
  Schedule,
  ApiResponse,
  SchedulesResponse,
  ScheduleResponse,
} from "@/types/schedule";

// Create Schedule
export const createSchedule = async (
  user_id: number,
  day_of_week: number,
  time_slot: number,
  subject: string,
  location: string,
  schedule_url: string
): Promise<ApiResponse<ScheduleResponse>> => {
  try {
    return await createScheduleApi({
      user_id,
      day_of_week,
      time_slot,
      subject,
      location,
      schedule_url,
    });
  } catch (error) {
    console.error("Error creating schedule:", error);
    throw new Error("Failed to create schedule. Please try again later.");
  }
};

// Update Schedule
export const updateSchedule = async (
  schedule_id: number,
  user_id: number,
  day_of_week: number,
  time_slot: number,
  subject: string,
  location: string,
  schedule_url: string
): Promise<ApiResponse<ScheduleResponse>> => {
  try {
    return await updateScheduleApi({
      schedule_id,
      user_id,
      day_of_week,
      time_slot,
      subject,
      location,
      schedule_url,
    });
  } catch (error) {
    console.error("Error updating schedule:", error);
    throw new Error("Failed to update schedule. Please try again later.");
  }
};

// Fetch Schedules
export const fetchSchedules = async (
  query_params?: Partial<Schedule>
): Promise<ApiResponse<SchedulesResponse>> => {
  try {
    return await getSchedulesApi(query_params);
  } catch (error) {
    console.error("Error fetching schedules:", error);
    throw new Error("Failed to fetch schedules. Please try again later.");
  }
};

// Delete Schedule
export const deleteSchedule = async (
  schedule_id: number
): Promise<ApiResponse<{ message: string }>> => {
  try {
    return await deleteScheduleApi(schedule_id);
  } catch (error) {
    console.error("Error deleting schedule:", error);
    throw new Error("Failed to delete schedule. Please try again later.");
  }
};
