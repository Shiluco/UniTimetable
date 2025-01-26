"use client";

import React, { createContext, useState, ReactNode } from "react";
import { Post } from "@/types/post";

// Contextの型
type PostInfoContextType = {
  postsInfo: Post[];
  setPostsInfo: React.Dispatch<React.SetStateAction<Post[]>>;
};

// 初期値（空のContextを作成）
const PostInfoContext = createContext<PostInfoContextType>({
  postsInfo: [],
  setPostsInfo: () => {},
});

// Providerのプロパティの型
type PostInfoProviderProps = {
  children: ReactNode;
};

// Providerの実装
export const PostInfoProvider = ({ children }: PostInfoProviderProps) => {
  const [postsInfo, setPostsInfo] = useState<Post[]>([]); // 投稿情報のState管理

  return (
    <PostInfoContext.Provider value={{ postsInfo, setPostsInfo }}>
      {children}
    </PostInfoContext.Provider>
  );
};

export default PostInfoContext;