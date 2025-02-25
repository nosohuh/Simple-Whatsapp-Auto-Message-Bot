import { nextui } from "@nextui-org/theme";

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  plugins: [
    nextui({
      themes: {
        light: {
          colors: {
            background: "#FFFFFF", // Açık mod arka plan rengi
            foreground: "#11181C", // Açık mod metin rengi
            primary: {
              foreground: "#FFFFFF", // Açık mod birincil metin rengi
              DEFAULT: "#006FEE", // Açık mod birincil renk
            },
          },
        },
        dark: {
          colors: {
            background: "#000000", // Koyu mod arka plan rengi
            foreground: "#ECEDEE", // Koyu mod metin rengi
            primary: {
              foreground: "#FFFFFF", // Koyu mod birincil metin rengi
              DEFAULT: "#006FEE", // Koyu mod birincil renk
            },
          },
        },
        mytheme: {
          extend: "dark", // Dark teması üzerine özelleştirme
          colors: {
            primary: {
              DEFAULT: "#BEF264", // Özel birincil renk
              foreground: "#000000", // Özel birincil metin rengi
            },
            focus: "#BEF264", // Özel odak rengi
          },
        },
      },
    }),
  ],
};
