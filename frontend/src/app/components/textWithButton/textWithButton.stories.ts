import type { Meta, StoryObj } from "@storybook/react";

import { TextWithButton } from "./textWithButton";

const meta = {
  title: "UniTimetable/TextWithButton",
  component: TextWithButton,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof TextWithButton>;

export default meta;
type Story = StoryObj<typeof meta>;

// Primaryストーリー
export const Primary: Story = {
  args: {
    placeholder: "ここにテキストを入力してください",
    label: "送信",
  },
};
