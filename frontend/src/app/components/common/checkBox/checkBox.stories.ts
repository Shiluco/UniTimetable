import type { Meta, StoryObj } from "@storybook/react";
import { CheckBox } from "./checkBox";

const meta = {
  title: "UniTimetable/CheckBox",
  component: CheckBox,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof CheckBox>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
  args: {
    label: "CheckBox",
  },
};
