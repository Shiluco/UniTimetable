// components/layout/Layout.tsx
import styles from "./layout.module.css";
import { Logo } from "@/app/components/common/logo/logo";
import { NavMenu } from "@/app/components/navMenu/navMenu";
import { SearchBox } from "@/app/components/searchBox/searchBox";

interface LayoutProps {
  children: React.ReactNode;
}

export const Layout = ({ children }: LayoutProps) => (
  <div className={styles.gridContainer}>
    <div className={styles.gridItem}>
      <Logo className={styles.mb50} />
      <NavMenu />
    </div>
    <div className={`${styles.gridItem} ${styles.scrollable}`}>{children}</div>
    <div className={styles.gridItem}>
      <SearchBox departmentOptions={[]} majorOptions={[]} yearOptions={[]} />
    </div>
  </div>
);
