"use client";

import { Schedule } from "@/types/schedule";
import "@/style/postItem.scss";
import { UserInfo } from "@/app/components/common/userInfo/userInfo";
import { Timetable } from "@/app/components/timetable/timetable";
import { Text } from "@/app/components/common/text/text";
import { Button } from "@/app/components/common/button/button";
interface PostItemProps {
  schedules: Schedule[];
  name: string;
  department: string;
  major: string;
  year: number;
  comment: string;
  className?: string;
}

export const PostItem = ({ schedules, name, department, major, year, className, comment }: PostItemProps) => {
  const id = "someId"; //TODO: ここのIDは要修正
  const handleReply = () => {
    window.location.href = `/home/reply/${id}`;
  };

  return (
    <div className={`postItem ${className}`}>
      <UserInfo name={name} department={department} major={major} year={year} className="mb-20" />
      <Timetable schedules={schedules} className="mb-20" />
      <div className="comment">
        <Text>{comment}</Text>
        <Button label="リプライ" reverse={true} type="minimal" className="comment-button" onClick={handleReply} />
      </div>
    </div>
  );
};
