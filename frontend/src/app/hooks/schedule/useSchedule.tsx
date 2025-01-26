import { useState } from "react";
import {
  fetchSchedules,
  createSchedule,
  updateSchedule,
  deleteSchedule,
} from "@/service/scheduleService";
import {
  Schedule,
  ApiResponse,
  SchedulesResponse,
  ScheduleResponse,
} from "@/types/schedule";

export const useSchedule = () => {
  const [schedules, setSchedules] = useState<Schedule[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Fetch Schedules
  const getSchedules = async (query_params?: Partial<Schedule>) => {
    setLoading(true);
    setError(null);
    try {
      const response: ApiResponse<SchedulesResponse> = await fetchSchedules(
        query_params
      );
      setSchedules(response.data.schedules);
    } catch (err) {
      console.error("Error fetching schedules:", err);
      setError("Failed to fetch schedules. Please try again later.");
    } finally {
      setLoading(false);
    }
  };

  // Create Schedule
  const addSchedule = async (
    user_id: number,
    day_of_week: number,
    time_slot: number,
    subject: string,
    location: string,
    schedule_url: string
  ) => {
    setLoading(true);
    setError(null);
    try {
      const response: ApiResponse<ScheduleResponse> = await createSchedule(
        user_id,
        day_of_week,
        time_slot,
        subject,
        location,
        schedule_url
      );
      setSchedules((prev) => [...prev, response.data]);
    } catch (err) {
      console.error("Error creating schedule:", err);
      setError("Failed to create schedule. Please try again later.");
    } finally {
      setLoading(false);
    }
  };

  // Update Schedule
  const modifySchedule = async (
    schedule_id: number,
    user_id: number,
    day_of_week: number,
    time_slot: number,
    subject: string,
    location: string,
    schedule_url: string
  ) => {
    setLoading(true);
    setError(null);
    try {
      const response: ApiResponse<ScheduleResponse> = await updateSchedule(
        schedule_id,
        user_id,
        day_of_week,
        time_slot,
        subject,
        location,
        schedule_url
      );
      setSchedules((prev) =>
        prev.map((schedule) =>
          schedule.schedule_id === schedule_id ? response.data : schedule
        )
      );
    } catch (err) {
      console.error("Error updating schedule:", err);
      setError("Failed to update schedule. Please try again later.");
    } finally {
      setLoading(false);
    }
  };

  // Delete Schedule
  const removeSchedule = async (schedule_id: number) => {
    setLoading(true);
    setError(null);
    try {
      await deleteSchedule(schedule_id);
      setSchedules((prev) =>
        prev.filter((schedule) => schedule.schedule_id !== schedule_id)
      );
    } catch (err) {
      console.error("Error deleting schedule:", err);
      setError("Failed to delete schedule. Please try again later.");
    } finally {
      setLoading(false);
    }
  };

  return {
    schedules,
    loading,
    error,
    getSchedules,
    addSchedule,
    modifySchedule,
    removeSchedule,
  };
};
