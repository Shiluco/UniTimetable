import { Fetcher } from "@/util/fetcher";
import { Department, Major } from "@/types/masterData";
// 学部一覧取得API
export const getDepartmentsApi = async (): Promise<Department[]> => {
  const response = await Fetcher<Department[]>("/departments", {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });

  return response;
};

// 学科一覧取得API
export const getMajorsApi = async (): Promise<Major[]> => {
  const response = await Fetcher<Major[]>("/majors", {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });

  return response;
};