import {
  createUserApi,
  updateUserApi,
  getUsersApi,
  deleteUserApi,
} from "@/api/userApi";
import { User } from "@/types/user";

// Fetch Users
export const fetchUsers = async (query_params?: {
  username?: string;
  email?: string;
  department_id?: number;
  major_id?: number;
  grade?: number;
}): Promise<User[]> => {
  try {
    return await getUsersApi(query_params);
  } catch (error) {
    console.error("Error fetching users:", error);
    throw new Error("Failed to fetch users. Please try again later.");
  }
};

// Add User
export const addUser = async (
  user_name: string,
  email: string,
  password: string,
  department_id: number,
  major_id: number,
  grade: number
): Promise<User> => {
  try {
    return await createUserApi(
      user_name,
      email,
      password,
      department_id,
      major_id,
      grade
    );
  } catch (error) {
    console.error("Error adding user:", error);
    throw new Error(
      "Failed to add user. Please check the details and try again."
    );
  }
};

// Edit User
export const editUser = async (
  user_id: number,
  username: string,
  email: string,
  comment: string,
  icon_url: string,
  department_id: number,
  major_id: number,
  grade: number
): Promise<User> => {
  try {
    return await updateUserApi(
      user_id,
      username,
      email,
      comment,
      icon_url,
      department_id,
      major_id,
      grade
    );
  } catch (error) {
    console.error("Error editing user:", error);
    throw new Error("Failed to update user details. Please try again.");
  }
};

// Remove User
export const removeUser = async (user_id: number): Promise<void> => {
  try {
    return await deleteUserApi(user_id);
  } catch (error) {
    console.error("Error removing user:", error);
    throw new Error("Failed to remove user. Please try again later.");
  }
};
