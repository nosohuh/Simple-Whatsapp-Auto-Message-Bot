"use client";

import { useRouter } from "next/navigation";

import { useAppSelector } from "@/hooks/stores"; // Redux hooks

export default function LoginLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const isLoggin = useAppSelector((state) => state.auth.isLoggin);

  return (
    <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
      <div className="inline-block max-w-lg text-center justify-center">
        {isLoggin ? (
          <p className="text-lg">
            Zaten giriş yaptınız. Ana sayfaya yönlendiriliyorsunuz...
          </p>
        ) : (
          children // Eğer giriş yapılmamışsa formu göster
        )}
      </div>
    </section>
  );
}
