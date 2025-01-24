import "@/style/postForm.scss";
import { UserInfo } from "@/app/components/common/userInfo/userInfo";
import { TextWithButton } from "@/app/components/textWithButton/textWithButton";
import { UploadButton } from "../uploadButton/uploadButton";

interface PostFormProps {
  name: string;
  department: string;
  major: string;
  year: number;
  className?: string;
}

export const PostForm = ({ name, department, major, year, className }: PostFormProps) => {
  return (
    <div className={`postForm ${className}`}>
      <UserInfo name={name} department={department} major={major} year={year} className="mb-20" />
      {/* TODO: 将来的にアップロードが完了したらアイコンとテキストが変わるようにする */}
      <UploadButton icon="assets/pdf.svg" label="HTMLをアップロード" className="mb-20" />
      <TextWithButton label="投稿" placeholder="投稿内容を入力" className="w-100" />
    </div>
  );
};
