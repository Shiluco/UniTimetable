import "@/style/textBox.scss";

interface TextBoxProps {
  placeholder?: string;
  onChange?: (value: string) => void;
  className?: string;
  type: "search" | "reply";
  value?: string;
}

export const TextBox = ({ placeholder, onChange, className = "", type, value }: TextBoxProps) => {
  return <div className={`container ${className}`}>{type === "search" ? <input type="text" className={`input input-${type}`} placeholder={placeholder} onChange={(e) => onChange?.(e.target.value)} value={value} /> : <textarea className={`input input-${type}`} placeholder={placeholder} onChange={(e) => onChange?.(e.target.value)} value={value} />}</div>;
};
