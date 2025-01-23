import { NavButton } from "@/app/components/common/navButton/navButton";
import "@/style/navMenu.scss";

interface NavMenuProps {
  className?: string;
}

export const NavMenu = ({ className }: NavMenuProps) => {
  const navItems = [
    { label: "タイムライン", icon: "/assets/timeline.svg" },
    { label: "プロフィール", icon: "/assets/profile.svg" },
    { label: "投稿する", icon: "/assets/post.svg" },
    { label: "検索", icon: "/assets/search.svg" },
  ];

  return (
    <nav className={`nav ${className}`.trim()}>
      <ul className="nav-list">
        {navItems.map((item, index) => (
          <li key={index} className="nav-item">
            <NavButton icon={item.icon} label={item.label} />
          </li>
        ))}
      </ul>
    </nav>
  );
};
