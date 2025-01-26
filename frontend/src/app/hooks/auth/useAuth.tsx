import { useCallback, useContext, useState } from "react";
import { useAuthService } from "@/service/useAuthService"; // AuthService をインポート
import { LoginResponse } from "@/types/auth";
import UserContext from "../user/userInfoContext";

export const useAuth = () => {
  const { loginService, error } = useAuthService();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [loginResponse, setLoginResponse] = useState<LoginResponse>();
  const { setUserInfo } = useContext(UserContext);

  const handleLogin = async (email: string, password: string) => {
    setIsLoading(true); // ローディング状態を開始
    const response = await loginService(email, password); // AuthService の handleLogin を呼び出し
    localStorage.setItem("accessToken", response.data.accessToken);
    localStorage.setItem("refreshToken", response.data.refreshToken);
    localStorage.setItem("user", JSON.stringify(response.data.user));
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
