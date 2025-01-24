import type { Meta, StoryObj } from "@storybook/react";
import { UserInfo } from "./userInfo";

const meta = {
  title: "UniTimetable/UserInfo",
  component: UserInfo,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof UserInfo>;

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
