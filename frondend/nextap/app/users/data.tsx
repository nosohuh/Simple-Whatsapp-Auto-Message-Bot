import React, { useEffect, useState } from "react";
import { useAppSelector } from "@/hooks/stores"; // Redux hook
import { AllUserList } from "@/axios/axios"; // Kullanıcıları çeken fonksiyon (API servisinizin ismi)

// Kullanıcıları çeken fonksiyon
const fetchData = async () => {
  try {
    const fetchedData = await AllUserList(); 
    return fetchedData.data.data; 
  } catch (err) {
    console.error(err);
    return [];
  }
};

const UsersTable = () => {
  const { user } = useAppSelector((state) => state.auth);
  const [data, setDataState] = useState([]); // Lokal state için

  useEffect(() => {
    const loadData = async () => {
      const fetchedUsers = await fetchData();
      setDataState(fetchedUsers);
    };

    if (data.length === 0) {
      loadData();
    }
  }, [data]);

  const columns = [
    { name: "NAME", uid: "username" },
    { name: "ROLE", uid: "role" },
    { name: "BALANCE", uid: "balance" },
    { name: "LEVEL", uid: "level" },
    { name: "STATUS", uid: "status" },
    { name: "ACTIONS", uid: "actions" },
  ];

  // Kullanıcılar verisini hazırlamak
  const users = data.map(
    (user: {
      id: number;
      username: string;
      role: string;
      balance: number;
      level: number;
      status?: string;
      created_at: string;
      updated_at: string;
      avatar?: string;
      email?: string;
    }) => ({
      id: user.id,
      username: user.username,
      role: user.role,
      balance: user.balance,
      level: user.level,
      status: user.status || "Active",
      created_at: user.created_at,
      updated_at: user.updated_at,
      avatar: user?.avatar || "https://i.pravatar.cc/150?u=default",
      email: user.email || `${user.username}@example.com`,
  }));

  return { columns, users }; // columns ve users verisini dışarıya aktar
};

export default UsersTable;
