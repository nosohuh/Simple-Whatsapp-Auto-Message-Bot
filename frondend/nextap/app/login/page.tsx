// Login.tsx
"use client";

import { Input, Button } from "@nextui-org/react";
import { useRouter } from "next/navigation";

import AuthLayout from "../AuthLayout";

import { loginRequest } from "@/axios/axios";
import { title } from "@/components/primitives";
import { login } from "@/stores/authSlice";
import { useAppSelector, useAppDispatch } from "@/hooks/stores"; // Doğru yolu kontrol edin

export default function Login() {
  const router = useRouter();
  const { isLoggin } = useAppSelector((state) => state.auth);
  const dispatch = useAppDispatch();

  if (isLoggin) {
    router.push("/");
  }

  const submitHandle = async (e: any) => {
    e.preventDefault();
    const { username, password } = e.target;

    try {
      const response = await loginRequest(username.value, password.value);

      dispatch(login(response.data));
      localStorage.setItem("login", "true");
      router.push("/login");
    } catch (error: any) {
      alert(error.response.data.message);
      console.log(error.response.data);
    }
  };

  return (
    <AuthLayout>
      <div>
        <h1 className={title()}>Login</h1>
        <form className="w-64" onSubmit={submitHandle}>
          <Input
            className="max-w-xs mb-8 mt-16"
            defaultValue=""
            description="Lütfen Kullanıcı adınızı girin."
            label="Username "
            name="username" // Input için name ekledik
            type="text" // username yerine text kullanıyoruz
          />
          <Input
            className="max-w-xs mb-8" // Üstten boşluk ekler
            defaultValue=""
            description="Şifrenizi Girin."
            fullWidth={true} // Genişliği tam yapar
            label="Password: "
            name="password" // Input için name ekledik
            type="password"
          />
          <Button color="success" type="submit">
            Giriş Yap
          </Button>{" "}
          {/* Button'a type="submit" ekledik */}
        </form>
      </div>
    </AuthLayout>
  );
}
