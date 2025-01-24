"use strict";

import { LoginForm } from "@/app/components/loginForm/loginForm";
import { Text } from "@/app/components/common/text/text";
import Image from "next/image";
import styles from "./page.module.css";

export default function LoginPage() {
  return (
    <div className={styles.wrapper}>
      <Text variant="h2">みんなの時間割を共有する履修登録特化型SNS</Text>
      <Image src="/assets/univTimeTableLogo.svg" alt="logo" width={500} height={300} />
      <LoginForm />
    </div>
  );
}
