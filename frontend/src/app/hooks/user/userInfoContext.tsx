"use client";

import React, { createContext, useState } from "react";
import { User } from "@/types/user";

// コンテキストの型を定義
type UserInfoContextType = {
  userInfo: User | null; // 初期値は null
  setUserInfo: React.Dispatch<React.SetStateAction<User | null>>;
  usersInfo: User[] | null; // 初期値は null
  setUsersInfo: React.Dispatch<React.SetStateAction<User[] | null>>;
};

// デフォルト値を設定して Context を作成
const UserInfoContext = createContext<UserInfoContextType>({
  userInfo: null,
  setUserInfo: () => {}, // 初期値として空の関数
  usersInfo: null,
  setUsersInfo: () => {}, // 初期値として空の関数
});

// Provider の実装
export const UserInfoProvider = ({ children }: { children: React.ReactNode }) => {
  const [userInfo, setUserInfo] = useState<User | null>(null);
  const [usersInfo, setUsersInfo] = useState<User[] | null>(null);

  return (
    <UserInfoContext.Provider value={{ userInfo, setUserInfo, usersInfo, setUsersInfo }}>
      {children}
    </UserInfoContext.Provider>
  );
};

export default UserInfoContext;
