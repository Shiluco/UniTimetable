"use client";

import styles from "./page.module.css";
import { Logo } from "@/app/components/common/logo/logo";
import { NavMenu } from "@/app/components/navMenu/navMenu";
import { PostItem } from "@/app/components/postItem/postItem";
import { SearchBox } from "@/app/components/searchBox/searchBox";

export default function HomePage() {
  return (
    <div className={styles.gridContainer}>
      <div className={styles.gridItem}>
        <Logo className={styles.mb50} />
        <NavMenu />
      </div>
      <div className={`${styles.gridItem} ${styles.scrollable}`}>
        <PostItem schedules={[]} name="Sample Name" department="Sample Department" major="Sample Major" year={1} className={styles.timeline_postItem} comment="message" />
      </div>
      <div className={styles.gridItem}>
        <SearchBox departmentOptions={[]} majorOptions={[]} yearOptions={[]} />
      </div>
    </div>
  );
}
