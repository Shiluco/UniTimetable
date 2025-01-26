import React, { useState } from "react";
import "@/style/checkBox.scss";
import { Text } from "@/app/components/common/text/text";

interface CheckBoxProps {
  label: string;
  defaultChecked?: boolean;
  onChange?: (checked: boolean) => void;
  className?: string;
  variant?: "h1" | "h2" | "h3" | "body1" | "body2" | "body3";
}

export const CheckBox = (props: CheckBoxProps) => {
  const { label, defaultChecked = false, onChange, className = "", variant = "body1" } = props;
  const [isChecked, setIsChecked] = useState(defaultChecked);

  const handleClick = () => {
    const newChecked = !isChecked;
    setIsChecked(newChecked);
    onChange?.(newChecked);
  };

  return (
    <button className={`checkboxLabel ${className}`} onClick={handleClick} type="button" role="checkbox" aria-checked={isChecked}>
      <div className="checkbox">
        {isChecked && (
          <svg viewBox="0 0 24 24" className="checkmark" fill="none" stroke="currentColor" strokeWidth="2">
            <path d="M5 13l4 4L19 7" strokeLinecap="round" strokeLinejoin="round" />
          </svg>
        )}
      </div>
      <Text variant={variant} bold={false}>
        {label}
      </Text>
    </button>
  );
};
