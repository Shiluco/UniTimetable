"use client";

import { useEffect, useState } from "react";
import { Layout } from "@/app/components/layout/layout";
import { PostItem } from "@/app/components/postItem/postItem";
import styles from "./page.module.css";
import { usePost } from "@/app/hooks/post/usePost";
import { useSchedule } from "@/app/hooks/schedule/useSchedule";
import { Schedule } from "@/types/schedule";

export default function HomePage() {
  const { postsInfo, loading, error, getPosts } = usePost();
  const { schedules, getSchedules } = useSchedule();
  const [mySchedules, setMySchedules] = useState<Schedule[]>([]);

  useEffect(() => {
    const userString = localStorage.getItem("user");
    if (userString) {
      try {
        const userObj = JSON.parse(userString);
        if (userObj?.id) {
          // userId がある場合のみAPIコール
          getPosts(userObj.id);
        }
      } catch (err) {
        console.error("Failed to parse user:", err);
      }
    }

    getSchedules();
  }, []);

  useEffect(() => {
    const enrolledScheduleIds = postsInfo.map((post) => post.schedule_id);
    setMySchedules(schedules.filter((schedule) => enrolledScheduleIds.includes(schedule.schedule_id)));
  }, [postsInfo, schedules]);
  return (
    <Layout>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {loading ? <p>Loading schedules...</p> : <PostItem schedules={mySchedules} name="Sample Name" department="Sample Department" major="Sample Major" year={1} className={styles.timeline_postItem} comment="message" />}
    </Layout>
  );
}
