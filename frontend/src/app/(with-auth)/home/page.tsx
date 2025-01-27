"use client";

import { useEffect } from "react";
import { Layout } from "@/app/components/layout/layout";
import { PostItem } from "@/app/components/postItem/postItem";
import styles from "./page.module.css";
import { useSchedule } from "@/app/hooks/schedule/useSchedule";

export default function HomePage() {
  const { schedules, loading, error, getSchedules } = useSchedule();

  useEffect(() => {
    const userString = localStorage.getItem("user");
    if (userString) {
      try {
        const userObj = JSON.parse(userString);
        if (userObj?.id) {
          // userId がある場合のみAPIコール
          getSchedules(userObj.id);
        }
      } catch (err) {
        console.error("Failed to parse user:", err);
      }
    }
  }, []);
  return (
    <Layout>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {loading ? <p>Loading schedules...</p> : <PostItem schedules={schedules} name="Sample Name" department="Sample Department" major="Sample Major" year={1} className={styles.timeline_postItem} comment="message" />}
    </Layout>
  );
}
