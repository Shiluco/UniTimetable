import "@/style/navButton.scss";
import { Text } from "@/app/components/common/text/text";
import Image from "next/image";

interface navButtonProps {
  label: string;
  icon: string;
  onClick?: () => void;
  className?: string;
  variant?: "h1" | "h2" | "h3" | "body1" | "body2" | "body3";
}

export const navButton = (props: navButtonProps) => {
  const { label, icon, onClick, className = "", variant = "body1" } = props;
  return (
    <button className={`navButton ${className}`.trim()} onClick={onClick}>
      <Image src={icon} alt={label} width={25} height={25} />
      <Text variant={variant} bold={false}>
        {label}
      </Text>
    </button>
  );
};
