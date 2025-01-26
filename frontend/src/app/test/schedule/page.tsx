"use client";

import React, { useEffect } from "react";
import { useSchedule } from "@/app/hooks/schedule/useSchedule";

const ScheduleComponent = () => {
  const {
    schedules,
    loading,
    error,
    getSchedules,
    addSchedule,
    modifySchedule,
    removeSchedule,
  } = useSchedule();

  useEffect(() => {
    // 初期ロードでスケジュールを取得
    getSchedules();
  },[]);

  const handleAddSchedule = () => {
    addSchedule(1, 1, 1, "Math", "Room A", "http://example.com");
  };

  const handleUpdateSchedule = () => {
    modifySchedule(1, 1, 2, 1, "Science", "Room B", "http://example.com");
  };

  const handleDeleteSchedule = () => {
    removeSchedule(1);
  };

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <div>
      <h1>Schedules</h1>
      <ul>
        {schedules.map((schedule) => (
          <li key={schedule.schedule_id}>
            {schedule.subject} - {schedule.location}
          </li>
        ))}
      </ul>
      <button onClick={handleAddSchedule}>Add Schedule</button>
      <button onClick={handleUpdateSchedule}>Update Schedule</button>
      <button onClick={handleDeleteSchedule}>Delete Schedule</button>
    </div>
  );
};

export default ScheduleComponent;
