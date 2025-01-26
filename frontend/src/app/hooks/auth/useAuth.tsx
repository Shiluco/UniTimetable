import { useCallback, useContext, useState } from "react";
import { AuthService } from "@/service/authService"; // AuthService をインポート
import { LoginResponse } from "@/types/auth";
import UserContext from "../user/userInfoContext";

export const useAuth = () => {
  const { loginService, error } = AuthService();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [loginResponse, setLoginResponse] = useState<LoginResponse>();
  const { setUserInfo } = useContext(UserContext);

  const handleLogin = async (email: string, password: string) => {
    setIsLoading(true); // ローディング状態を開始
    const response = await loginService(email, password); // AuthService の handleLogin を呼び出し
    localStorage.setItem("authToken", response.data.accessToken); // ローカルストレージにトークンを保存（例としてトークン）
    setIsLoading(false); // ローディング状態を終了
    setLoginResponse(response); // ログインレスポンスをセット
    setUserInfo(response.data.user); // ユーザー情報をセット
  };

  const handleLogout = useCallback(() => {
    // ローカルストレージのトークンを削除
    localStorage.removeItem("authToken");
    window.location.href = "/login";
  }, []);

  return { handleLogin, handleLogout, isLoading, error, loginResponse }; // 必要なデータや関数を返す
};
