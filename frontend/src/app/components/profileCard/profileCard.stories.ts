import type { Meta, StoryObj } from "@storybook/react";
import { ProfileCard } from "./profileCard";

const meta = {
  title: "UniTimetable/ProfileCard",
  component: ProfileCard,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof ProfileCard>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
  args: {
    name: "名前",
    department: "情報",
    major: "情報社会",
    year: 1,
    description: "aaaa",
  },
};
