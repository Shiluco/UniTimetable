import "@/style/selectBox.scss";

interface Option {
  value: string;
  label: string;
}

interface SelectBoxProps {
  options?: Option[];
  defaultValue?: string;
  onChange?: (value: string) => void;
  className?: string;
  placeholder?: string;
}

export const SelectBox = ({ options, defaultValue, onChange, className = "", placeholder = "選択" }: SelectBoxProps) => {
  return (
    <div className="select-box-container">
      <select className={`select-box-select ${className}`.trim()} onChange={(e) => onChange?.(e.target.value)} defaultValue={defaultValue || ""}>
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
