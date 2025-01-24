import type { Meta, StoryObj } from "@storybook/react";
import { SearchBox } from "./searchBox";

const meta = {
  title: "UniTimetable/SearchBox",
  component: SearchBox,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof SearchBox>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
  args: {
    departmentOptions: [
      { value: "理学部", label: "理学部" },
      { value: "工学部", label: "工学部" },
      { value: "情報学部", label: "情報学部" },
    ],
    majorOptions: [
      { value: "数学科", label: "数学科" },
      { value: "物理学科", label: "物理学科" },
      { value: "情報学科", label: "情報学科" },
    ],
    yearOptions: [
      { value: "1年", label: "1年" },
      { value: "2年", label: "2年" },
      { value: "3年", label: "3年" },
      { value: "4年", label: "4年" },
    ],
  },
};
