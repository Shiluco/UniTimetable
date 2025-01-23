import type { Meta, StoryObj } from "@storybook/react";
import { action } from "@storybook/addon-actions";

import { NavButton } from "./navButton";

const meta = {
  title: "UniTimetable/NavButton",
  component: NavButton,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof NavButton>;

export default meta;
type Story = StoryObj<typeof meta>;

// Primaryストーリー
export const Primary: Story = {
  args: {
    label: "NavButton",
    icon: "/assets/timeline.svg",
    onClick: action("Primary button clicked"),
  },
};
