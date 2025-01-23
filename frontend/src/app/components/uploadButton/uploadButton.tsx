import "@/style/uploadButton.scss";
import Image from "next/image";
import { Text } from "@/app/components/common/text/text";

interface UploadButtonProps {
  icon: string;
  label: string;
  className?: string;
  onClick?: () => void;
}

export const UploadButton = (props: UploadButtonProps) => {
  const { icon, label, className, onClick } = props;
  return (
    <div className={`uploadContainer ${className}`} onClick={onClick}>
      <Image src={icon} alt="UploadIcon" width={25} height={25} className="mr-10" />
      <Text variant="body1" bold={false}>
        {label}
      </Text>
    </div>
  );
};
