import type { Meta, StoryObj } from "@storybook/react";
import { TextBox } from "./textBox";

const meta = {
  title: "UniTimetable/TextBox",
  component: TextBox,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof TextBox>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
  args: {
    placeholder: "テキストを入力",
    type: "search",
  },
};
