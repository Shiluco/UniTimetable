import "@/style/logo.scss";
import Image from "next/image";
import { Text } from "@/app/components/common/text/text";

interface LogoProps {
  className?: string;
}

export const Logo = ({ className }: LogoProps) => {
  return (
    <div className={`logo ${className}`.trim()}>
      <Image src="/assets/univTimeTable_LogoOnly.svg" alt="Univ.Timetable" width={100} height={100} />
      <Text>Univ.Timetable</Text>
    </div>
  );
};
