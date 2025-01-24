import type { Meta, StoryObj } from "@storybook/react";
import { PostItem } from "./postItem";

const sampleSchedules = [
  {
    schedule_id: 1,
    user_id: 1,
    day_of_week: 1,
    time_slot: 1,
    subject: "人工知能概論",
    location: "情13",
    schedule_url: "/schedules/1",
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
  {
    schedule_id: 2,
    user_id: 1,
    day_of_week: 3,
    time_slot: 2,
    subject: "応用プログラミングC",
    location: "オンライン",
    schedule_url: "/schedules/2",
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
];

const meta = {
  title: "UniTimetable/PostItem",
  component: PostItem,
  parameters: {
    layout: "centered",
  },
  tags: ["autodocs"],
} satisfies Meta<typeof PostItem>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {
  args: {
    schedules: sampleSchedules,
    name: "hogefugapiyo",
    department: "情報",
    major: "情報科",
    comment: "hogehoge",
    year: 1,
  },
};
