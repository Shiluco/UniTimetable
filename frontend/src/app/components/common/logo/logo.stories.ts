import type { Meta, StoryObj } from "@storybook/react";

import { Logo } from "./logo";

const meta = {
  title: "UniTimetable/Logo",
  component: Logo,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof Logo>;

export default meta;
type Story = StoryObj<typeof meta>;

// Primaryストーリー
export const Primary: Story = {
  args: {},
};
