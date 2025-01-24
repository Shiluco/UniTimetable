"use client";

import { Layout } from "@/app/components/layout/layout";
import { PostForm } from "@/app/components/postForm/postForm";

export default function PostPage() {
  return (
    <Layout>
      <PostForm name="sample name" department="Computer Science" major="Software Engineering" year={2} />
    </Layout>
  );
}
