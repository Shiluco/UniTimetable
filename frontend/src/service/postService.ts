import {
  createPostApi,
  updatePostApi,
  getPostsApi,
  deletePostApi,
} from "@/api/postApi";
import { PostResponse, PostsResponse } from "@/types/post";

// Create Post
export const createPost = async (
  user_id: number,
  content: string,
  parent_post_id: number | null,
  schedule_id: number,
  image_url: string
): Promise<PostResponse> => {
  try {
    return await createPostApi(
      user_id,
      content,
      parent_post_id,
      schedule_id,
      image_url
    );
  } catch (error) {
    console.error("Error creating post:", error);
    throw new Error("Failed to create post. Please try again later.");
  }
};

// Update Post
export const updatePost = async (
  post_id: number,
  user_id: number,
  content: string,
  parent_post_id: number | null,
  schedule_id: number,
  image_url: string
): Promise<PostResponse> => {
  try {
    return await updatePostApi(
      post_id,
      user_id,
      content,
      parent_post_id,
      schedule_id,
      image_url
    );
  } catch (error) {
    console.error("Error updating post:", error);
    throw new Error("Failed to update post. Please try again later.");
  }
};

// Fetch Posts
export const fetchPosts = async (query_params?: {
  user_id?: number;
  content?: string;
  parent_post_id?: number | null;
  schedule_id?: number;
  image_url?: string;
  created_at?: string;
  updated_at?: string;
}): Promise<PostsResponse> => {
  try {
    return await getPostsApi(query_params);
  } catch (error) {
    console.error("Error fetching posts:", error);
    throw new Error("Failed to fetch posts. Please try again later.");
  }
};

// Delete Post
export const deletePost = async (post_id: number): Promise<void> => {
  try {
    await deletePostApi(post_id);
  } catch (error) {
    console.error("Error deleting post:", error);
    throw new Error("Failed to delete post. Please try again later.");
  }
};
