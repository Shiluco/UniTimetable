import { getDepartmentsApi, getMajorsApi } from "@/api/masterDataApi";
import { Department, Major } from "@/types/masterData";

// Fetch Departments
export const fetchDepartments = async (query_params?: {
  name?: string;
}): Promise<Department[]> => {
  try {
    return await getDepartmentsApi(query_params);
  } catch (error) {
    console.error("Error fetching departments:", error);
    throw new Error("Failed to fetch departments. Please try again later.");
  }
};

// Fetch Majors
export const fetchMajors = async (query_params?: {
  department_id?: number;
  name?: string;
}): Promise<Major[]> => {
  try {
    return await getMajorsApi(query_params);
  } catch (error) {
    console.error("Error fetching majors:", error);
    throw new Error("Failed to fetch majors. Please try again later.");
  }
};
