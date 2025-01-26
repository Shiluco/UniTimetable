import { Fetcher } from "@/util/fetcher";
import { PostResponse, PostsResponse } from "@/types/post";

const url = "/posts";
// 投稿の作成
export const createPostApi = async (
  user_id: number,
  content: string,
  parent_post_id: number | null,
  schedule_id: number,
  image_url: string
): Promise<PostResponse> => {
  const response = await Fetcher<PostResponse>(url, {
    method: "POST",
    body: JSON.stringify({
      user_id,
      content,
      parent_post_id,
      schedule_id,
      image_url,
    }),
    headers: { "Content-Type": "application/json" },
  });
  return response;
};

// 投稿の更新
export const updatePostApi = async (
  post_id: number,
  user_id: number,
  content: string,
  parent_post_id: number | null,
  schedule_id: number,
  image_url: string
): Promise<PostResponse> => {
  const response = await Fetcher<PostResponse>(url, {
    method: "PUT",
    body: JSON.stringify({
      post_id,
      user_id,
      content,
      parent_post_id,
      schedule_id,
      image_url,
    }),
    headers: { "Content-Type": "application/json" },
  });
  return response;
};

// 投稿の取得
export const getPostsApi = async (query_params?: {
  user_id?: number;
  content?: string;
  parent_post_id?: number | null;
  schedule_id?: number;
  image_url?: string;
  created_at?: string;
  updated_at?: string;
}): Promise<PostsResponse> => {
  const queryString = query_params
    ? "?" +
      Object.entries(query_params)
        .filter(([, value]) => value !== undefined)
        .map(
          ([key, value]) =>
            `${encodeURIComponent(key)}=${encodeURIComponent(
              value as string | number
            )}`
        )
        .join("&")
    : "";

  const response = await Fetcher<PostsResponse>(`${url}${queryString}`, {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });

  return response;
};

export const deletePostApi = async (post_id: number): Promise<void> => {
  await Fetcher<void>(`${url}?post_id=${post_id}`, {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
  });
};
