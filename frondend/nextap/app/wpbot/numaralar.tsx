import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell } from "@nextui-org/react";
import React, { useEffect, useState } from "react";
import { getNumberList, startNumBot, botStatus } from "@/axios/axios"; // botStatus API fonksiyonu da dahil
import Noty from "noty";

export default function Numaralar() {
  // Veriyi tutmak için state
  interface NumberItem {
    id: number;
    numara: string;
    durum: boolean;
    mesaj: string;
    created_at: string;
  }

  const [numberList, setNumberList] = useState<NumberItem[]>([]);
  const [botStat, setBotStat] = useState(false); // Bot durumunu kontrol etmek için state

  // API'den veriyi al
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await getNumberList(); // API URL'nizi buraya ekleyin
        const data = await response.data;
        if (data.success) {
          setNumberList(data.data); // Gelen veriyi state'e at
        }
      } catch (error) {
        console.error("Veri alınırken hata oluştu:", error);
      }
    };

    fetchData(); // API'den veri çek
  }, []); 

  // Bot durumunu kontrol et
  const fetchBotStatusHandler = async () => {
    try {
      const response = await botStatus();
      const data = response?.data?.response; // API yanıtından doğru veri yapısına erişim

      if (data) {
        setBotStat(data?.isReady); // Botun aktif olup olmadığını kontrol et
      } else {
        console.log("Yanıt beklenen formatta değil:", response);
      }
    } catch (error) {
      new Noty({
        type: 'error',
        text: 'Bağlantı kurulamadı!',
        timeout: 3000,
        theme: 'mint',
      }).show();
    }
  };

  // Botu başlatmak için handler
  const numBotHandler = async (id: string) => {
    fetchBotStatusHandler(); // Bot durumunu kontrol et
    if (!botStat) {
      new Noty({
        type: "warning",
        text: "Bot Aktif değil!",
        timeout: 3000,
      }).show();
      return; // Eğer bot aktifse, işlemi durdur
    }

    try {
      await startNumBot(id); // Bot başlatma işlemi
      new Noty({
        type: "success",
        text: "Bot Başlatıldı!",
        timeout: 3000,
      }).show();
    } catch (error) {
      console.error("Hata:", error);
      new Noty({
        type: "error",
        text: "Bot Başlatılamadı!",
        timeout: 3000,
      }).show();
    }
  };


  return (
    <div className="w-full mt-8"> 
    <div className="overflow-x-auto"> 
      <Table aria-label="Example static collection table">
        <TableHeader>
          <TableColumn>ID</TableColumn>
          <TableColumn>NUMARA</TableColumn>
          <TableColumn>DURUM</TableColumn>
          <TableColumn>MESAJ</TableColumn>
          <TableColumn>OLUŞTURULMA ZAMANI</TableColumn>
          <TableColumn>ACTIONS</TableColumn> 
        </TableHeader>
        <TableBody>
          {numberList.map((item) => (
            <TableRow key={item?.id}>
              <TableCell>{item?.id}</TableCell>
              <TableCell>{item?.numara?.slice(0, 5) + (item?.numara?.length > 5 ? "..." : "")}</TableCell>
              <TableCell>{item?.durum ? "Aktif" : "Pasif"}</TableCell>
              <TableCell>{item?.mesaj}</TableCell>
              <TableCell>{item?.created_at}</TableCell>
              <TableCell>
                <button
                  className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                  onClick={() => numBotHandler(String(item?.id))} 
                >
                  Bu Seti Başlat
                </button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  </div>
  
  );
}
