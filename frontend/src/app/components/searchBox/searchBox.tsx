import "@/style/searchBox.scss";
import { TextBox } from "@/app/components/common/textBox/textBox";
import { SelectBox } from "@/app/components/common/selectBox/selectBox";
import { Text } from "@/app/components/common/text/text";
import { Option } from "@/types/selectBox";
import { Button } from "@/app/components/common/button/button";

interface SearchBoxProps {
  departmentOptions: Option[];
  majorOptions: Option[];
  yearOptions: Option[];
  className?: string;
}

export const SearchBox = ({ departmentOptions, majorOptions, yearOptions, className }: SearchBoxProps) => {
  return (
    <div className={`searchBox ${className}`}>
      <TextBox placeholder="検索ワードを入力" className="w-100 mb-20" type="search" />
      <div className="searchBox__filter">
        <Text variant="body2" className="mb-10">
          検索条件
        </Text>
        <SelectBox className="w-100 mb-10" options={departmentOptions} placeholder="学部" />
        <SelectBox className="w-100 mb-10" options={majorOptions} placeholder="学科" />
        <SelectBox className="w-100 mb-10" options={yearOptions} placeholder="学年" />
        <Button label="同条件に指定" type="minimal" reverse={true} className="mt-10" variant="body2" />
      </div>
      <Button label="検索" type="normal" className="searchButton" />
    </div>
  );
};
