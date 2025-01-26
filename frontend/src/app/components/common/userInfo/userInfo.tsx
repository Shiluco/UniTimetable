import React from "react";
import { Text } from "@/app/components/common/text/text";
import "@/style/userInfo.scss";

interface UserInfoProps {
  name: string;
  department: string;
  major: string;
  year: number;
  className?: string;
}

export const UserInfo = ({ name, department, major, year, className = "" }: UserInfoProps) => {
  return (
    <div className={`user-info-card ${className}`}>
      <div className="user-info-container">
        {/* TODO: 将来的にアイコン画像を表示する */}
        <div className="profile-image"></div>
        <div className="info-section">
          <Text variant="h3">{name}</Text>
        </div>
        <div className="divider"></div>
        <div className="info-section">
          <Text variant="body1">{department}</Text>
          <Text variant="body2" className="label">
            学部
          </Text>
        </div>
        <div className="divider"></div>
        <div className="info-section">
          <Text variant="body1">{major}</Text>
          <Text variant="body2" className="label">
            学科
          </Text>
        </div>
        <div className="divider"></div>
        <div className="year-section">
          <Text variant="body1">{year}年</Text>
        </div>
      </div>
    </div>
  );
};
