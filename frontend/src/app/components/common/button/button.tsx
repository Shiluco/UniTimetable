import "@/style/button.scss";
import { Text } from "@/app/components/common/text/text";

interface ButtonProps {
  label: string;
  onClick?: () => void;
  reverse?: boolean;
  type: "normal" | "minimal" | "large";
  className?: string;
  variant?: "h1" | "h2" | "h3" | "body1" | "body2" | "body3";
}

export const Button = (props: ButtonProps) => {
  const { label, onClick, type, className = "", variant = "body1", reverse } = props;
  return (
    <button className={`button button-${type} ${className} ${reverse ? "button-reverse" : ""}`.trim()} onClick={onClick}>
      <Text variant={variant} bold={false}>
        {label}
      </Text>
    </button>
  );
};
