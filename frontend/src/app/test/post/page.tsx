"use client";
import React, { useEffect } from "react";
import { usePost } from "@/app/hooks/post/usePost";

const PostManagementPage = () => {
  const {
    postsInfo,
    getPosts,
    createPostHandler,
    updatePostHandler,
    deletePostHandler,
  } = usePost();

  // 初期データの取得
  useEffect(() => {
    getPosts({ user_id: 1 });
  }, [getPosts]);

  // 新規投稿作成
  const handleCreatePost = async () => {
    try {
      await createPostHandler(1, "New Post Content", null, 123, "image.jpg");
      console.log("Post created successfully");
    } catch (err) {
      console.error("Error creating post:", err);
    }
  };

  // 投稿の更新
  const handleUpdatePost = async () => {
    if (postsInfo.length === 0) {
      console.log("No posts available to update.");
      return;
    }
    try {
      await updatePostHandler(
        postsInfo[0].post_id, // 最初の投稿を更新
        1,
        "Updated Post Content",
        null,
        123,
        "updated-image.jpg"
      );
      console.log("Post updated successfully");
    } catch (err) {
      console.error("Error updating post:", err);
    }
  };

  // 投稿の削除
  const handleDeletePost = async () => {
    if (postsInfo.length === 0) {
      console.log("No posts available to delete.");
      return;
    }
    try {
      await deletePostHandler(postsInfo[0].post_id); // 最初の投稿を削除
      console.log("Post deleted successfully");
    } catch (err) {
      console.error("Error deleting post:", err);
    }
  };

  return (
    <div>
      <h1>Post Management</h1>

      <button onClick={handleCreatePost}>Create New Post</button>
      <button onClick={handleUpdatePost}>Update First Post</button>
      <button onClick={handleDeletePost}>Delete First Post</button>

      <h2>Posts</h2>
      {postsInfo.length > 0 ? (
        <ul>
          {postsInfo.map((post) => (
            <li key={post.post_id}>
              <p>Content: {post.content}</p>
              <p>Image URL: {post.image_url}</p>
              <p>Created At: {post.created_at}</p>
            </li>
          ))}
        </ul>
      ) : (
        <p>No posts available.</p>
      )}
    </div>
  );
};

export default PostManagementPage;
