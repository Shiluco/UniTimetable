import "@/style/profileCard.scss";
import { Button } from "@/app/components/common/button/button";

interface ProfileCardProps {
  name: string;
  department: string;
  major: string;
  year: number;
  description: string;
  onEditClick?: () => void;
  className?: string;
}

export const ProfileCard = ({ name, major, department, year, description, onEditClick, className }: ProfileCardProps) => {
  return (
    <div className={`profile-card ${className}`}>
      <div className="profile-avatar" />
      <div className="profile-info">
        <p className="profile-name">{name}</p>
        <div className="profile-details">
          <p>{department}学部</p>
          <p>{major}学科</p>
          <p>{year}年</p>
          <p className="profile-text">{description}</p>
        </div>
        {/* TODO: 将来的にログインユーザのidによってボタンを表示・非表示にする */}
        <Button label="編集" type="normal" onClick={onEditClick} />
      </div>
    </div>
  );
};
