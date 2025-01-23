import type { Meta, StoryObj } from "@storybook/react";

import { NavMenu } from "./navMenu";

const meta = {
  title: "UniTimetable/NavMenu",
  component: NavMenu,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof NavMenu>;

export default meta;
type Story = StoryObj<typeof meta>;

// Primaryストーリー
export const Primary: Story = {
  args: {},
};
