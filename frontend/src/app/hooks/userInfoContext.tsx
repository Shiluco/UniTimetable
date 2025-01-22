import React, { createContext, useState } from "react";
import { User } from "@/types/user";

// コンテキストの型を定義
type UserInfoContextType = {
  userInfo: User | null; // 初期値は null
  setUserInfo: React.Dispatch<React.SetStateAction<User | null>>;
};

// デフォルト値を設定して Context を作成
const UserContext = createContext<UserInfoContextType>({
  userInfo: null,
  setUserInfo: () => {}, // 初期値として空の関数
});

// Provider の実装
export const UserInfoProvider = ({ children }: { children: React.ReactNode }) => {
  const [userInfo, setUserInfo] = useState<User | null>(null);

  return (
    <UserContext.Provider value={{ userInfo, setUserInfo }}>
      {children}
    </UserContext.Provider>
  );
};

export default UserContext;
