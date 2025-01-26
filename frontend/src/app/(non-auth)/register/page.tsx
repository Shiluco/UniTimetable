"use client";

import { LoginLayout } from "@/app/components/layout/loginLayout";
import { RegisterForm } from "@/app/components/registerForm/registerForm";

export default function LoginPage() {
  return (
    <LoginLayout>
      <RegisterForm departmentOptions={[]} majorOptions={[]} yearOptions={[]} />
    </LoginLayout>
  );
}
