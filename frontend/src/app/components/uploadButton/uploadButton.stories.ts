import type { Meta, StoryObj } from "@storybook/react";
import { UploadButton } from "./uploadButton";

const meta = {
  title: "UniTimetable/UploadButton",
  component: UploadButton,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof UploadButton>;

export default meta;
type Story = StoryObj<typeof meta>;

// Primaryストーリー
export const Primary: Story = {
  args: {
    icon: "/assets/pdf.svg",
    label: "HTMLファイルをアップロード",
  },
};

// Secondaryストーリー
export const Secondary: Story = {
  args: {
    icon: "/assets/checkbox.svg",
    label: "アップロード完了！",
  },
};
