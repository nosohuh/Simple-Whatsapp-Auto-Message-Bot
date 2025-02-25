"use client";
import { useDispatch } from "react-redux";
import { useEffect, useState } from "react";
import { Spinner } from "@nextui-org/react";

import { login, logout } from "@/stores/authSlice";
import { userControlRequest, logoutRequest } from "@/axios/axios";

const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const dispatch = useDispatch();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    if (localStorage.getItem("login")) {
      const fetchData = async () => {
        try {
          const response = await userControlRequest();

          dispatch(login(response.data));
        } catch (error) {
          console.log(error);
          try {
            await logoutRequest();
            dispatch(logout());
          } catch (error) {
            console.log(error);
          }
        } finally {
          setLoading(false);
        }
      };

      fetchData();
    } else {
      setLoading(false);
    }
  }, [dispatch]);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen">
        {" "}
      
        <div className="flex gap-4">
          <Spinner
            color="danger"
            label="Bekle Paşam Yükleniyor..."
            labelColor="danger"
            size="lg"
          />{" "}
          {/* Büyütmek için size özelliği kullan */}
        </div>
      </div>
    ); // Yüklenme durumu için bir mesaj veya spinner ekleyebilirsiniz
  }

  return <>{children}</>;
};

export default AuthProvider;
