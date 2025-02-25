"use client";
import {
  Navbar as NextUINavbar,
  NavbarContent,
  NavbarMenu,
  NavbarMenuToggle,
  NavbarBrand,
  NavbarItem,
  NavbarMenuItem,
} from "@nextui-org/navbar";
import { Button } from "@nextui-org/button";
import { Link } from "@nextui-org/link";
import { link as linkStyles } from "@nextui-org/theme";
import NextLink from "next/link";
import clsx from "clsx";

import { siteConfig } from "@/config/site";
import { ThemeSwitch } from "@/components/theme-switch";
import { HeartFilledIcon, Logo } from "@/components/icons";
import { useAppDispatch, useAppSelector } from "@/hooks/stores";
import { logoutRequest } from "@/axios/axios";
import { logout } from "@/stores/authSlice";
import React from "react";

export const Navbar = () => {
  const dispatch = useAppDispatch();
  const { isLoggin, user } = useAppSelector((state) => state.auth);
  const [isMenuOpen, setIsMenuOpen] = React.useState(false);

  const logoutHandle = async () => {
    await logoutRequest();
    dispatch(logout());
  };

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  return (
    <NextUINavbar maxWidth="xl" position="sticky">
      <NavbarContent className="basis-1/5 sm:basis-full" justify="start">
        <NavbarBrand as="li" className="gap-3 max-w-fit">
          <NextLink className="flex justify-start items-center gap-1" href="/">
            <Logo />
            <p className="font-bold text-inherit">Yılmaz Bot</p>
          </NextLink>
        </NavbarBrand>
        <ul className="hidden lg:flex gap-4 justify-start ml-2">
          {siteConfig.navItems.map((item) => (
            <NavbarItem key={item.href}>
              <NextLink
                className={clsx(
                  linkStyles({ color: "foreground" }),
                  "data-[active=true]:text-primary data-[active=true]:font-medium",
                )}
                color="foreground"
                href={item.href}
              >
                {item.label}
              </NextLink>
            </NavbarItem>
          ))}
        </ul>
      </NavbarContent>

      <NavbarContent
        className="hidden sm:flex basis-1/5 sm:basis-full"
        justify="end"
      >
        <NavbarItem className="hidden sm:flex gap-2">
          <ThemeSwitch />
          <div className="flex gap-2">
            {isLoggin ? "Hoşgeldiniz:  " + user?.username : "✔"}
          </div>
        </NavbarItem>
        <NavbarItem className="hidden md:flex">
          {isLoggin ? (
            <Button
              className="text-sm font-normal text-default-600 bg-default-100"
              startContent={<HeartFilledIcon className="text-danger" />}
              variant="flat"
              onClick={logoutHandle}
            >
              Çıkış
            </Button>
          ) : (
            <Button
              as={NextLink}
              className="text-sm font-normal text-default-600 bg-default-100"
              href={siteConfig.links.sponsor}
              startContent={<HeartFilledIcon className="text-danger" />}
              variant="flat"
            >
              Giriş
            </Button>
          )}
        </NavbarItem>
      </NavbarContent>

      <NavbarContent className="sm:hidden basis-1 pl-4" justify="end">
        <ThemeSwitch />
        <NavbarMenuToggle onClick={toggleMenu} />
      </NavbarContent>
      <NavbarMenu>
        <div className="mx-4 mt-2 flex flex-col gap-2">
          {siteConfig.navMenuItems.map((item, index) => (
            <NavbarMenuItem key={index}>
              <NextLink passHref href={item.href}>
                <Link
                  onClick={() => setIsMenuOpen(false)}  // Close the menu when a link is clicked
                  color={
                    index === 2
                      ? "primary"
                      : index === siteConfig.navMenuItems.length - 1
                        ? "danger"
                        : "foreground"
                  }
                  size="lg"
                >
                  {item.label}
                </Link>
              </NextLink>
            </NavbarMenuItem>
          ))}

          <NavbarMenuItem>
            {isLoggin ? (
              <Button
                className="text-sm font-normal text-default-600 bg-default-100"
                variant="flat"
                onClick={() => {
                  logoutHandle();
                  setIsMenuOpen(false);
                }}
              >
                Çıkış
              </Button>
            ) : (
              <Button
                as={NextLink}
                className="text-sm font-normal text-default-600 bg-default-100"
                href={siteConfig.links.sponsor}
                variant="flat"
                onClick={() => setIsMenuOpen(false)}
              >
                Giriş
              </Button>
            )}
          </NavbarMenuItem>
        </div>
      </NavbarMenu>
    </NextUINavbar>
  );
};
