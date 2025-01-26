import type { Meta, StoryObj } from "@storybook/react";
import { LoginForm } from "./loginForm";

const meta = {
  title: "UniTimetable/LoginForm",
  component: LoginForm,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof LoginForm>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
  args: {},
};
