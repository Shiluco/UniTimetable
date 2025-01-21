import { Fetcher } from "@/util/fetcher";
import { LoginResponse } from "@/types/auth";

export const login = async (
  email: string,
  password: string
): Promise<LoginResponse> => {
  const response = await Fetcher<LoginResponse>("auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
    headers: { "Content-Type": "application/json" },
    authRequired: false,
  });

  // アクセストークンをローカルストレージに保存
  localStorage.setItem("authToken", response.data.accessToken);

  return response;
};

export const logout = async () => {
  // ローカルストレージのトークンを削除
  localStorage.removeItem("authToken");
};
