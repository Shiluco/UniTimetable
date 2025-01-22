import { Fetcher } from "@/util/fetcher";
import { User } from "@/types/user";

export const createUserApi = async (
  user_name: string,
  email: string,
  password: string,
  department_id: number,
  major_id: number,
  grade: number
): Promise<User> => {
  const response = await Fetcher<User>("/users", {
    method: "POST",
    body: JSON.stringify({
      user_name,
      email,
      password,
      department_id,
      major_id,
      grade,
    }),
    headers: { "Content-Type": "application/json" },
  });
  return response;
};


export const updateUserApi = async (
  user_id: number,
  username: string,
  email: string,
  comment: string,
  icon_url: string,
  department_id: number,
  major_id: number,
  grade: number
): Promise<User> => {
  const response = await Fetcher<User>("/users", {
    method: "PUT",
    body: JSON.stringify({
      user_id,
      username,
      email,
      comment,
      icon_url,
      department_id,
      major_id,
      grade,
    }),
    headers: { "Content-Type": "application/json" },
  });
  return response;
};

export const getUsersApi = async (query_params?: {
  username?: string;
  email?: string;
  department_id?: number;
  major_id?: number;
  grade?: number;
}): Promise<User[]> => {
  // クエリパラメータをURLエンコード
  const queryString = query_params
    ? "?" +
      Object.entries(query_params)
        .filter(([value]) => value !== undefined) // undefinedを除外
        .map(
          ([key, value]) =>
            `${encodeURIComponent(key)}=${encodeURIComponent(
              value as string | number
            )}`
        )
        .join("&")
    : "";

  // Fetchリクエストを送信
  const response = await Fetcher<User[]>(`/users${queryString}`, {
    method: "GET",
    headers: { "Content-Type": "application/json" }, // 必要に応じて追加
  });

  if (!response) {
    throw new Error("Failed to fetch users. Please try again later.");
  }

  return response;
};


export const deleteUserApi = async (user_id: number): Promise<void> => {
  // URLに`user_id`を含める
  await Fetcher<void>(`/users/${user_id}`, {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
  });
};
