"use client";

import { LoginLayout } from "@/app/components/layout/loginLayout";
import { LoginForm } from "@/app/components/loginForm/loginForm";

export default function LoginPage() {
  return (
    <LoginLayout>
      <LoginForm />
    </LoginLayout>
  );
}
