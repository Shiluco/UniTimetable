import "@/style/textBox.scss";

interface TextBoxProps {
  placeholder?: string;
  onChange?: (value: string) => void;
  className?: string;
  type: "search" | "reply" | "password";
  value?: string;
}

export const TextBox = ({ placeholder, onChange, className = "", type, value }: TextBoxProps) => {
  switch (type) {
    case "search":
      return (
        <div className={`container ${className}`}>
          <input type="text" className={`input input-${type}`} placeholder={placeholder} onChange={(e) => onChange?.(e.target.value)} value={value} />
        </div>
      );
    case "reply":
      return (
        <div className={`container ${className}`}>
          <textarea className={`input input-${type}`} placeholder={placeholder} onChange={(e) => onChange?.(e.target.value)} value={value} />
        </div>
      );
    case "password":
      return (
        <div className={`container ${className}`}>
          <input type="password" className={`input input-${type}`} placeholder={placeholder} onChange={(e) => onChange?.(e.target.value)} value={value} />
        </div>
      );
    default:
      break;
  }
};
