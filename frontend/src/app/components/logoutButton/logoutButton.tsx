import { Button } from "@/app/components/common/button/button";
import { useAuth } from "@/app/hooks/auth/useAuth";

interface LogoutButtonProps {
  className?: string;
}

export const LogoutButton = ({ className }: LogoutButtonProps) => {
  const { handleLogout } = useAuth();
  return <Button label="ログアウト" type="normal" reverse={true} className={className} onClick={handleLogout} />;
};
