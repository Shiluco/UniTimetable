import { useState, useCallback } from "react";
import { login } from "@/api/authApi"; 

export const AuthService = () => {
  const [error, setError] = useState<string | null>(null);

  const loginService = useCallback(async (email: string, password: string) => {
    try {
      const response = await login(email, password); // 既存の login 関数を呼び出し
      setError(null); // エラーをリセット
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


  return { loginService,error };
};
