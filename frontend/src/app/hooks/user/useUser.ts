import { useCallback, useContext } from "react";
import UserInfoContext from "@/app/hooks/user/userInfoContext"; // コンテキストのインポート
import {
  fetchUsers,
  addUser,
  editUser,
  removeUser,
} from "@/service/userService";

export const useUser = () => {
  const { usersInfo, setUsersInfo } = useContext(UserInfoContext);

  const handleError = (error: unknown): string => {
    if (error instanceof Error) {
      return error.message;
    }
    return "An unexpected error occurred.";
  };

  // Fetch Users
  const getUsers = useCallback(
    async (query_params?: {
      username?: string;
      email?: string;
      department_id?: number;
      major_id?: number;
      grade?: number;
    }) => {
      try {
        const fetchedUsers = await fetchUsers(query_params);
        setUsersInfo(fetchedUsers);
      } catch (err) {
        console.error("Error in getUsers:", err);
        throw new Error(handleError(err));
      }
    },
    [setUsersInfo]
  );

  // Add User
  const createUser = useCallback(
    async (
      user_name: string,
      email: string,
      password: string,
      department_id: number,
      major_id: number,
      grade: number
    ) => {
      try {
        await addUser(
          user_name,
          email,
          password,
          department_id,
          major_id,
          grade
        );
        await getUsers(); // 最新のデータを取得
      } catch (err) {
        console.error("Error in createUser:", err);
        throw new Error(handleError(err));
      }
    },
    [getUsers]
  );

  // Edit User
  const updateUser = useCallback(
    async (
      user_id: number,
      username: string,
      email: string,
      comment: string,
      icon_url: string,
      department_id: number,
      major_id: number,
      grade: number
    ) => {
      try {
        await editUser(
          user_id,
          username,
          email,
          comment,
          icon_url,
          department_id,
          major_id,
          grade
        );
        await getUsers(); // 最新のデータを取得
      } catch (err) {
        console.error("Error in updateUser:", err);
        throw new Error(handleError(err));
      }
    },
    [getUsers]
  );

  // Remove User
  const deleteUser = useCallback(
    async (user_id: number) => {
      try {
        await removeUser(user_id);
        await getUsers(); // 最新のデータを取得
      } catch (err) {
        console.error("Error in deleteUser:", err);
        throw new Error(handleError(err));
      }
    },
    [getUsers]
  );

  return {
    usersInfo, // Context から取得したデータをそのまま返す
    getUsers,
    createUser,
    updateUser,
    deleteUser,
  };
};
