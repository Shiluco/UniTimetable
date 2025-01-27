import { useCallback, useContext, useState } from "react";
import PostInfoContext from "@/app/hooks/post/postInfoContext"; // コンテキストのインポート
import { fetchPosts, createPost, updatePost, deletePost } from "@/service/postService";

export const usePost = () => {
  const { postsInfo, setPostsInfo } = useContext(PostInfoContext);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleError = (error: unknown): string => {
    if (error instanceof Error) {
      return error.message;
    }
    return "An unexpected error occurred.";
  };

  // Fetch Posts
  const getPosts = useCallback(
    async (query_params?: { user_id?: number; content?: string; parent_post_id?: number | null; schedule_id?: number; image_url?: string; created_at?: string; updated_at?: string }) => {
      try {
        setLoading(true);
        setError(null);
        const fetchedPosts = await fetchPosts(query_params);
        console.log("fetchedPosts", fetchedPosts);
        setPostsInfo(fetchedPosts.data.posts);
      } catch (err) {
        setError("Failed to fetch posts. Please try again later.");
        console.error("Error in getPosts:", err);
        throw new Error(handleError(err));
      }
    },
    [setPostsInfo]
  );

  // Create Post
  const createPostHandler = useCallback(
    async (user_id: number, content: string, parent_post_id: number | null, schedule_id: number, image_url: string) => {
      try {
        await createPost(user_id, content, parent_post_id, schedule_id, image_url);
        await getPosts(); // 最新のデータを取得
      } catch (err) {
        console.error("Error in createPostHandler:", err);
        throw new Error(handleError(err));
      }
    },
    [getPosts]
  );

  // Update Post
  const updatePostHandler = useCallback(
    async (post_id: number, user_id: number, content: string, parent_post_id: number | null, schedule_id: number, image_url: string) => {
      try {
        await updatePost(post_id, user_id, content, parent_post_id, schedule_id, image_url);
        await getPosts(); // 最新のデータを取得
      } catch (err) {
        console.error("Error in updatePostHandler:", err);
        throw new Error(handleError(err));
      }
    },
    [getPosts]
  );

  // Delete Post
  const deletePostHandler = useCallback(
    async (post_id: number) => {
      try {
        await deletePost(post_id);
        await getPosts(); // 最新のデータを取得
      } catch (err) {
        console.error("Error in deletePostHandler:", err);
        throw new Error(handleError(err));
      }
    },
    [getPosts]
  );

  return {
    postsInfo, // Context から取得したデータをそのまま返す
    loading,
    error,
    getPosts,
    createPostHandler,
    updatePostHandler,
    deletePostHandler,
  };
};
