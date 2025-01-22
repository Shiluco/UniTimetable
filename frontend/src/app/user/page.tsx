"use client";

import React, { useEffect } from "react";
import { useUser } from "@/app/hooks/user/useUser";

const UserManagementPage = () => {
  const { usersInfo, getUsers, createUser, updateUser, deleteUser } =
    useUser();

  useEffect(() => {
    // 初回データ取得
    getUsers();
  }, [getUsers]);

  //TODO: 仮でidで消してる
  return (
    <div>
      <h1>User Management</h1>
      {usersInfo ? (
        <ul>
          {usersInfo.map((user) => (
            <li key={user.id}>
              {user.username} ({user.email})
              <button onClick={() => deleteUser(user.id)}>Delete</button>
            </li>
          ))}
        </ul>
      ) : (
        <p>Loading users...</p>
      )}
      {/* 例: ユーザー作成 */}
      <button
        onClick={() =>
          createUser("New User", "newuser@example.com", "password123", 1, 1, 3)
        }
      >
        Add User
      </button>
    </div>
  );
};

export default UserManagementPage;
