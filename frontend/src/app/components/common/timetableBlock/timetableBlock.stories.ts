import type { Meta, StoryObj } from "@storybook/react";
import { TimetableBlock } from "./timetableBlock";

const meta = {
  title: "UniTimetable/TimetableBlock",
  component: TimetableBlock,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof TimetableBlock>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Empty: Story = {
  args: {},
};

export const WithText: Story = {
  args: {
    text: "テキスト",
  },
};
