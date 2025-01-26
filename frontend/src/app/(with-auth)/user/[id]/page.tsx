"use client";

import styles from "./page.module.css";
import { PostItem } from "@/app/components/postItem/postItem";
import { ProfileCard } from "@/app/components/profileCard/profileCard";
import { Layout } from "@/app/components/layout/layout";

const sampleSchedules = [
  {
    schedule_id: 1,
    user_id: 1,
    day_of_week: 1,
    time_slot: 1,
    subject: "人工知能概論",
    location: "情13",
    schedule_url: "/schedules/1",
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
  {
    schedule_id: 2,
    user_id: 1,
    day_of_week: 3,
    time_slot: 2,
    subject: "応用プログラミングC",
    location: "オンライン",
    schedule_url: "/schedules/2",
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
];

export default function UserPage() {
  return (
    <Layout>
      <ProfileCard name="Sample Name" department="Sample Department" major="Sample Major" year={1} description="Sample Description" className={styles.mb50} />
      <PostItem schedules={sampleSchedules} name="Sample Name" department="Sample Department" major="Sample Major" year={1} className={styles.timeline_postItem} comment="message" />
    </Layout>
  );
}
