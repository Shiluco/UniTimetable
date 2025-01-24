"use client";

import { TextBox } from "@/app/components/common/textBox/textBox";
import "@/style/loginForm.scss";
import { Button } from "@/app/components/common/button/button";

export const LoginForm = () => {
  return (
    <div className="loginForm">
      <TextBox placeholder="静大メアド" type="search" className="mb-20" />
      <TextBox placeholder="パスワード" type="password" />
      <div className="loginForm__button flexWrap">
        <Button label="新規登録" type="normal" reverse={true} />
        <Button label="ログイン" type="normal" />
      </div>
    </div>
  );
};
