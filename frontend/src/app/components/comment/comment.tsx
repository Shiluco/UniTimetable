import "@/style/comment.scss";
import { UserInfo } from "@/app/components/common/userInfo/userInfo";
import { Text } from "@/app/components/common/text/text";

interface CommentProps {
  name: string;
  department: string;
  major: string;
  year: number;
  comment: string;
  className?: string;
}

export const Comment = ({ name, department, major, year, className, comment }: CommentProps) => {
  return (
    <div className={`commentContainer ${className}`}>
      <UserInfo name={name} department={department} major={major} year={year} className="mb-20" />
      <div className="commentBox">
        <Text variant="body2" className="mb-10">
          {comment}
        </Text>
      </div>
    </div>
  );
};
