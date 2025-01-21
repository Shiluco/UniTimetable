import React from "react";
import "@/style/textBox.scss";

interface TextBoxProps {
  placeholder?: string;
  onChange?: (value: string) => void;
  className?: string;
}

export const TextBox = ({ placeholder, onChange, className = "" }: TextBoxProps) => {
  return (
    <div className={`container ${className}`}>
      <input type="text" className="input" placeholder={placeholder} onChange={(e) => onChange?.(e.target.value)} />
    </div>
  );
};
