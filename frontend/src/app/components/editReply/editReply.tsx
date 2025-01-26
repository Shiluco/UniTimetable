import "@/style/editReply.scss";
import { UserInfo } from "@/app/components/common/userInfo/userInfo";
import { TextWithButton } from "@/app/components/textWithButton/textWithButton";

interface EditReplyProps {
  name: string;
  department: string;
  major: string;
  year: number;
  className?: string;
}

export const EditReply = ({ name, department, major, year, className }: EditReplyProps) => {
  return (
    <div className={`editReplyContainer ${className}`}>
      <UserInfo name={name} department={department} major={major} year={year} className="mb-20" />
      <TextWithButton className="mb-10" label="送信" placeholder="返信を入力" />
    </div>
  );
};
