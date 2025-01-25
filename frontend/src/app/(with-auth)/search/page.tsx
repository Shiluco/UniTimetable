"use client";

import { Layout } from "@/app/components/layout/layout";
import styles from "./page.module.css";
import { PostItem } from "@/app/components/postItem/postItem";
import { SearchBox } from "@/app/components/searchBox/searchBox";

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

export default function SearchPage() {
  return (
    <Layout>
      <SearchBox
        departmentOptions={[
          { label: "Department 1", value: "Department 1" },
          { label: "Department 2", value: "Department 2" },
        ]}
        majorOptions={[
          { label: "Major 1", value: "Major 1" },
          { label: "Major 2", value: "Major 2" },
        ]}
        yearOptions={[
          { label: "1", value: "1" },
          { label: "2", value: "2" },
          { label: "3", value: "3" },
          { label: "4", value: "4" },
        ]}
        className={styles.mb50}
      />
      <PostItem schedules={sampleSchedules} name="Sample Name" department="Sample Department" major="Sample Major" year={1} className={styles.timeline_postItem} comment="message" />
    </Layout>
  );
}
