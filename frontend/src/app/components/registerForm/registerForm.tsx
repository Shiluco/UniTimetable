import "@/style/registerForm.scss";
import { Button } from "@/app/components/common/button/button";
import { SelectBox } from "@/app/components/common/selectBox/selectBox";
import { TextBox } from "@/app/components/common/textBox/textBox";
import { Option } from "@/types/selectBox";

interface RegisterFormProps {
  departmentOptions: Option[];
  majorOptions: Option[];
  yearOptions: Option[];
  className?: string;
  onBack: () => void;
}

export const RegisterForm = ({ departmentOptions, majorOptions, yearOptions, className, onBack }: RegisterFormProps) => {
  return (
    <div className={`registerForm ${className}`}>
      <SelectBox className="w-100 mb-10" options={departmentOptions} placeholder="学部" />
      <SelectBox className="w-100 mb-10" options={majorOptions} placeholder="学科" />
      <SelectBox className="w-100 mb-10" options={yearOptions} placeholder="学年" />
      <TextBox placeholder="ニックネーム" className="w-100 mb-10" type="search" />
      <TextBox placeholder="静大メアド" type="search" className="mb-10" />
      <TextBox placeholder="パスワード" type="password" className="mb-20" />
      <div className="registerForm__button flexWrap">
        <Button label="前画面に戻る" type="normal" reverse={true} onClick={onBack} />
        <Button label="登録する" type="normal" />
      </div>
    </div>
  );
};
