import type { Meta, StoryObj } from "@storybook/react";
import { PostForm } from "./postForm";

const meta = {
  title: "UniTimetable/PostForm",
  component: PostForm,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof PostForm>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
  args: {
    name: "hogefugapiyo",
    department: "情報",
    major: "情報科",
    year: 1,
  },
};
