"use client"
import { title } from "@/components/primitives";
import App from "./app"
import { useAppDispatch, useAppSelector } from "@/hooks/stores";

export default function PricingPage() {
  const dispatch = useAppDispatch();
  const { isLoggin, user } = useAppSelector((state) => state.auth);

  return (
    <>
      {isLoggin ? (
        <App />
      ) : (
        <div className="flex justify-center items-center">
          Giriş Yapman lazım evlat
        </div>
      )}
    </>
  );
}
