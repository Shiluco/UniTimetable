import type { Meta, StoryObj } from "@storybook/react";
import { action } from "@storybook/addon-actions";

import { navButton } from "./navButton";

const meta = {
  title: "UniTimetable/navButton",
  component: navButton,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof navButton>;

export default meta;
type Story = StoryObj<typeof meta>;

// Primaryストーリー
export const Primary: Story = {
  args: {
    label: "navButton",
    icon: "/assets/timeline.svg",
    onClick: action("Primary button clicked"),
  },
};
