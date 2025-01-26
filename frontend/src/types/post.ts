export type Post = {
  post_id: number;
  user_id: number;
  content: string;
  parent_post_id: number | null;
  schedule_id: number;
  image_url: string;
  created_at: string;
  updated_at: string;
};

export type PostResponse = {
  status: string;
  message: string;
  error_detail: string;
  data: Post | null;
};

export type PostsResponse = {
  status: string;
  message: string;
  error_detail: string;
  data: {
    posts: Post[];
    posts_total: number;
  };
};
