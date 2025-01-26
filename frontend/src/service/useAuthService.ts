import { useState, useCallback } from "react";
import { login } from "@/api/authApi";
import { register } from "@/api/authApi";

export const useAuthService = () => {
  const [error, setError] = useState<string | null>(null);

  const loginService = useCallback(async (email: string, password: string) => {
    try {
      const response = await login(email, password); // 既存の login 関数を呼び出し
      setError(null); // エラーをリセット
      window.location.href = "/home";
      return response; // ログインレスポンスを返す
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message); // `Error`型のmessageを使用
      } else {
        setError("不明なエラーが発生しました。");
      }
      throw err;
    }
  }, []);

  const registerService = useCallback(async (name: string, email: string, password: string) => {
    try {
      const response = await register(name, email, password); // 既存の Register 関数を呼び出し
      setError(null); // エラーをリセット
      window.location.href = "/login";
      return response; // ログインレスポンスを返す
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message); // `Error`型のmessageを使用
      } else {
        setError("不明なエラーが発生しました。");
      }
      throw err;
    }
  }, []);

  return { loginService, registerService, error };
};
