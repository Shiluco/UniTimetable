import "@/style/textWithButton.scss";
import { Button } from "@/app/components/common/button/button";
import { TextBox } from "@/app/components/common/textBox/textBox";

interface TextWithButtonProps {
  placeholder?: string;
  label: string;
  onChange?: (value: string) => void;
  value?: string;
  className?: string;
}

export const TextWithButton = (props: TextWithButtonProps) => {
  const { placeholder, label, onChange, value, className } = props;
  return (
    <div className={`container ${className}`}>
      <TextBox placeholder={placeholder} type="reply" onChange={onChange} value={value} />
      <Button label={label} type="minimal" />
    </div>
  );
};
