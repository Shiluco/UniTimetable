import "@/style/registerForm.scss";
import { Button } from "@/app/components/common/button/button";
import { SelectBox } from "@/app/components/common/selectBox/selectBox";
import { TextBox } from "@/app/components/common/textBox/textBox";
import { Option } from "@/types/selectBox";
import { useState } from "react";
interface RegisterFormProps {
  departmentOptions: Option[];
  majorOptions: Option[];
  yearOptions: Option[];
  className?: string;
}

export const RegisterForm = ({ departmentOptions, majorOptions, yearOptions, className }: RegisterFormProps) => {
  const [email, setEmail] = useState("");
  const [emailError, setEmailError] = useState("");
  const [password, setPassword] = useState("");
  const [passwordError, setPasswordError] = useState("");

  const validateEmail = (email: string) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      setEmailError("有効なメールアドレスを入力してください");
      return false;
    }

    if (!email.endsWith("@shizuoka.ac.jp")) {
      setEmailError("静岡大学のメールアドレスを入力してください");
      return false;
    }

    setEmailError("");
    return true;
  };

  const handleEmailChange = (value: string) => {
    setEmail(value);
    if (value) {
      validateEmail(value);
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
      return;
    }

    if (!password) {
      setPasswordError("パスワードを入力してください");
      return;
    }

    if (!validateEmail(email)) {
      return;
    }

    console.log("Valid email:", email);
  };

  return (
    <div className={`registerForm ${className}`}>
      <SelectBox className="w-100 mb-10" options={departmentOptions} placeholder="学部" />
      <SelectBox className="w-100 mb-10" options={majorOptions} placeholder="学科" />
      <SelectBox className="w-100 mb-10" options={yearOptions} placeholder="学年" />
      <TextBox placeholder="ニックネーム" className="w-100 mb-10" type="search" />
      <TextBox placeholder="静大メアド" type="search" className="mb-10" value={email} onChange={handleEmailChange} />
      {emailError && <div className="registerFormError">{emailError}</div>}
      <TextBox placeholder="パスワード" type="password" className="mb-20" value={password} onChange={handlePasswordChange} />
      {passwordError && <div className="registerFormError">{passwordError}</div>}
      <div className="registerForm__button flexWrap">
        <Button label="前画面に戻る" type="normal" reverse={true} onClick={() => (window.location.href = "/login")} />
        <Button label="登録する" type="normal" onClick={handleSubmit} />
      </div>
    </div>
  );
};
