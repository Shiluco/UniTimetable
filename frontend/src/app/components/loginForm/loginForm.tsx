"use client";

import { TextBox } from "@/app/components/common/textBox/textBox";
import "@/style/loginForm.scss";
import { Button } from "@/app/components/common/button/button";
import { useState } from "react";
import { useAuthService } from "@/service/useAuthService";

export const LoginForm = () => {
  const [email, setEmail] = useState("");
  const [emailError, setEmailError] = useState("");
  const [password, setPassword] = useState("");
  const [passwordError, setPasswordError] = useState("");
  const { loginService } = useAuthService();

  const handleEmailChange = (value: string) => {
    setEmail(value);
    if (!value) {
      setEmailError("メールアドレスを入力してください");
    } else {
      setEmailError("");
    }
  };

  const handlePasswordChange = (value: string) => {
    setPassword(value);
    if (!value) {
      setPasswordError("パスワードを入力してください");
    } else {
      setPasswordError("");
    }
  };

  const handleSubmit = () => {
    if (!email) {
      setEmailError("メールアドレスを入力してください");
    }
    if (!password) {
      setPasswordError("パスワードを入力してください");
    }

    loginService(email, password);
  };

  return (
    <div className="loginForm">
      <TextBox placeholder="静大メアド" type="search" className="mb-20" onChange={handleEmailChange} />
      {emailError && <div className="loginFormError">{emailError}</div>}
      <TextBox placeholder="パスワード" type="password" className="mb-20" onChange={handlePasswordChange} />
      {passwordError && <div className="loginFormError">{passwordError}</div>}
      <div className="loginForm__button flexWrap">
        <Button label="新規登録" type="normal" reverse={true} onClick={() => (window.location.href = "/register")} />
        <Button label="ログイン" type="normal" onClick={handleSubmit} />
      </div>
    </div>
  );
};
