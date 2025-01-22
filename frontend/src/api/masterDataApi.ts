import { Fetcher } from "@/util/fetcher";
import { Department, Major } from "@/types/masterData";

// 学部一覧取得API
export const getDepartmentsApi = async (query_params?: {
  name?: string;
}): Promise<Department[]> => {
  const queryString = query_params
    ? "?" +
      Object.entries(query_params)
        .filter(([value]) => value !== undefined)
        .map(
          ([key, value]) =>
            `${encodeURIComponent(key)}=${encodeURIComponent(value as string)}`
        )
        .join("&")
    : "";

  const response = await Fetcher<Department[]>(`/departments${queryString}`, {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });

  return response;
};

// 学科一覧取得API
export const getMajorsApi = async (query_params?: {
  department_id?: number;
  name?: string;
}): Promise<Major[]> => {
  const queryString = query_params
    ? "?" +
      Object.entries(query_params)
        .filter(([value]) => value !== undefined)
        .map(
          ([key, value]) =>
            `${encodeURIComponent(key)}=${encodeURIComponent(
              value as string | number
            )}`
        )
        .join("&")
    : "";

  const response = await Fetcher<Major[]>(`/majors${queryString}`, {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });

  return response;
};
