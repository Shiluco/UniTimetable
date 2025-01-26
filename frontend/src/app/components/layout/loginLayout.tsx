"use client";

import React from "react";
import { Text } from "@/app/components/common/text/text";
import Image from "next/image";
import styles from "./loginLayout.module.css";

interface LoginLayoutProps {
  children: React.ReactNode;
}

export const LoginLayout = ({ children }: LoginLayoutProps) => {
  return (
    <div className={styles.wrapper}>
      <Text variant="h2">みんなの時間割を共有する履修登録特化型SNS</Text>
      <Image src="/assets/univTimeTableLogo.svg" alt="logo" width={500} height={300} />
      {children}
    </div>
  );
};
