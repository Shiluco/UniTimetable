export interface Option {
  value: string;
  label: string;
}

export interface SelectBoxProps {
  options?: Option[];
  defaultValue?: string;
  onChange?: (value: string) => void;
  className?: string;
  placeholder?: string;
}
