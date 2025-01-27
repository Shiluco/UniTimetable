import { getDepartmentsApi, getMajorsApi } from "@/api/masterDataApi";
import { Department, Major } from "@/types/masterData";

// Fetch Departments
export const fetchDepartments = async (): Promise<Department[]> => {
  try {
    const departments = await getDepartmentsApi();
    console.log("Fetched departments:", departments);
    return departments;
  } catch (error) {
    console.error("Error fetching departments:", error);
    throw new Error("Failed to fetch departments. Please try again later.");
  }
};

// Fetch Majors
export const fetchMajors = async (): Promise<Major[]> => {
  try {
    return await getMajorsApi();
  } catch (error) {
    console.error("Error fetching majors:", error);
    throw new Error("Failed to fetch majors. Please try again later.");
  }
};