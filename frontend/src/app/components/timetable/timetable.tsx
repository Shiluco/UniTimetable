// timetable.tsx
import React from "react";
import { TimetableBlock } from "@/app/components/common/timetableBlock/timetableBlock";
import { TimetableProps } from "@/types/timetables";
import "@/style/timetable.scss";

export const Timetable = ({ schedules, className = "" }: TimetableProps) => {
  const days = ["Mon.", "Tue.", "Web.", "Thu.", "Fri."];
  const periods = [1, 2, 3, 4, 5];

  const getSchedule = (day: number, period: number) => {
    return schedules.find((schedule) => schedule.day_of_week === day && schedule.time_slot === period);
  };

  return (
    <div className={`timetable ${className}`}>
      <div className="timetable-grid">
        <div className="timetable-header">
          <div className="empty-cell" />
          {days.map((day) => (
            <div key={day} className="day-cell">
              {day}
            </div>
          ))}
        </div>

        {periods.map((period) => (
          <div key={period} className="timetable-row">
            <div className="period-cell">{period}</div>
            {days.map((day, dayIndex) => {
              const schedule = getSchedule(dayIndex + 1, period);
              return (
                <div key={`${period}-${day}`} className="class-cell">
                  <TimetableBlock text={schedule?.subject} />
                </div>
              );
            })}
          </div>
        ))}
      </div>
    </div>
  );
};
