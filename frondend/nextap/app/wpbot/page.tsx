"use client";

import { useAppDispatch, useAppSelector } from "@/hooks/stores";
import WpLogin from "./wplogin";
import Numaralar from "./numaralar";
import TestMesaj from "./testmesaj";

export default function DocsPage() {
  const dispatch = useAppDispatch();
  const { isLoggin, user } = useAppSelector((state) => state.auth);

  return (
    <>
      {isLoggin ? (
        <div className="">
        <div className="flex justify-center items-center">
          <WpLogin />
        </div>
        <div className="flex justify-center items-center">
        
        </div>
        <div>
        <Numaralar />
        </div>
        </div>
      ) : (
        <div className="flex justify-center items-center">
          Giriş Yapman lazım evlat
        </div>
      )}
    </>
  );
}
