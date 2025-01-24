import { Schedule } from "@/types/timetables";
import "@/style/postItem.scss";
import { UserInfo } from "@/app/components/common/userInfo/userInfo";
import { Timetable } from "@/app/components/timetable/timetable";
import { TextWithButton } from "@/app/components/textWithButton/textWithButton";

interface PostItemProps {
  schedules: Schedule[];
  name: string;
  department: string;
  major: string;
  year: number;
  className?: string;
}

export const PostItem = ({ schedules, name, department, major, year, className }: PostItemProps) => {
  return (
    <div className={`postItem ${className}`}>
      <UserInfo name={name} department={department} major={major} year={year} className="mb-20" />
      <Timetable schedules={schedules} className="mb-20" />
      <TextWithButton label="返信" placeholder="返信を入力" className="w-100" />
    </div>
  );
};
