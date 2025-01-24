import React from "react";
import "@/style/timetableBlock.scss";

export type TimetableBlockProps = {
  text?: string;
  className?: string;
};

export const TimetableBlock = ({ text, className = "" }: TimetableBlockProps) => {
  return <div className={`text-block ${text ? "filled" : "empty"} ${className}`}>{text}</div>;
};
