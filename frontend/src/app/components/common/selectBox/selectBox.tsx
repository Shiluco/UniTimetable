import "@/style/selectBox.scss";
import { SelectBoxProps } from "@/types/selectBox";

export const SelectBox = ({ options, defaultValue, onChange, className, placeholder }: SelectBoxProps) => {
  return (
    <div className={`select-box-container ${className}`}>
      <select className="select-box-select" onChange={(e) => onChange?.(e.target.value)} defaultValue={defaultValue || ""}>
        <option value="" disabled>
          {placeholder}
        </option>
        {options?.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
    </div>
  );
};
