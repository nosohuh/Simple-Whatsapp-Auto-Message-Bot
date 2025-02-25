export type SiteConfig = typeof siteConfig;

export const siteConfig = {
  name: "Yılmazlar Bot", 
  description: "Descp...",
  navItems: [
    {
      label: "Anasayfa",
      href: "/",
    },
    {
      label: "WpBot",
      href: "/wpbot",
    },
    {
      label: "Kullanıcılar",
      href: "/users",
    },
  ],
  navMenuItems: [
    {
      label: "Anasayfa",
      href: "/",
    },
    {
      label: "WpBot",
      href: "/wpbot",
    },
    {
      label: "Kullanıcılar",
      href: "/users",
    },
  ],
  links: {
    sponsor: "/login",
  },
};
