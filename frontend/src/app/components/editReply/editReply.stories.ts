import type { Meta, StoryObj } from "@storybook/react";
import { EditReply } from "./editReply";

const meta = {
  title: "UniTimetable/EditReply",
  component: EditReply,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof EditReply>;

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
