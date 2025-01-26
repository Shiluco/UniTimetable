import { NavButton } from "@/app/components/common/navButton/navButton";
import "@/style/navMenu.scss";

interface NavMenuProps {
  className?: string;
}

export const NavMenu = ({ className }: NavMenuProps) => {
  const navItems = [
    { label: "タイムライン", icon: "/assets/timeline.svg", href: "/home" },
    { label: "プロフィール", icon: "/assets/profile.svg", href: "/user/[id]" },
    { label: "投稿する", icon: "/assets/post.svg", href: "/post" },
    { label: "検索", icon: "/assets/search.svg", href: "/search" },
  ];

  return (
    <nav className={`nav ${className}`.trim()}>
      <ul className="nav-list">
        {navItems.map((item, index) => (
          <li key={index} className="nav-item">
            <NavButton icon={item.icon} label={item.label} onClick={() => (window.location.href = item.href)} />
          </li>
        ))}
      </ul>
    </nav>
  );
};
