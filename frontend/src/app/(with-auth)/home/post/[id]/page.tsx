"use client";

import { Comment } from "@/app/components/comment/comment";
import styles from "./page.module.css";
import { Logo } from "@/app/components/common/logo/logo";
import { NavMenu } from "@/app/components/navMenu/navMenu";
import { PostItem } from "@/app/components/postItem/postItem";
import { SearchBox } from "@/app/components/searchBox/searchBox";
import { EditReply } from "@/app/components/editReply/editReply";

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

export default function HomePage() {
  return (
    <div className={styles.gridContainer}>
      <div className={styles.gridItem}>
        <Logo className={styles.mb50} />
        <NavMenu />
      </div>
      <div className={`${styles.gridItem} ${styles.scrollable}`}>
        <PostItem schedules={sampleSchedules} name="Sample Name" department="Sample Department" major="Sample Major" year={1} className={styles.timeline_postItem} comment="message" />
        <Comment name="Sample Name" department="Sample Department" major="Sample Major" year={2} comment="Sample Comment" className={styles.mb30} />
        <EditReply name="Sample Name" department="Sample Department" major="Sample Major" year={1} />
      </div>
      <div className={styles.gridItem}>
        <SearchBox departmentOptions={[]} majorOptions={[]} yearOptions={[]} />
      </div>
    </div>
  );
}
