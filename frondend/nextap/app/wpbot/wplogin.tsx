import { Button } from "@nextui-org/react";
import { wpLogin, wpStop, wpQr } from "@/axios/axios";
import NumaraEkle from "./numaraekle";
import Noty from "noty";
import "noty/lib/noty.css";
import "noty/lib/themes/mint.css"; // Mint teması
import BotStatus from "./botstatus";
import TestMesaj from "./testmesaj";
import NumBotStatus from "./numbotstat";

export default function WpLogin() {
    const wpLoginHandler = async () => {
        try {
            const response = await wpLogin(); // wpLogin fonksiyonunu çağır
            const data = response.data;

            // Başarı mesajı göster
            new Noty({
                type: "success", // Başarılı işlem türü
                text: "Bota başarıyla bağlanıldı!", // Mesaj içeriği
                timeout: 3000, // 3 saniye
                theme: "mint",
            }).show();

            if (data.response.status == 400) {
                try {
                    // QR kodu oluştur
                    const qrCodeDataURL = await wpQr();
                    const qrdata = qrCodeDataURL.data.qr;

                    // QR kodunu içeren bir Noty göster
                    new Noty({
                        theme: "mint", // Tema
                        type: "success", // Tip: info, success, warning, error
                        timeout: 5000, // 5 saniye
                        progressBar: true, // İlerleme çubuğu
                        layout: "center", // Ortada konumlandırma
                        text: `
                        <div style="text-align: center">
                          <h3>QR Kodunuz</h3>
                          <img src="${qrdata}" alt="QR Code" style="width: 200px; height: 200px;" />
                        </div>
                      `,
                    }).show();
                } catch (err) {
                    console.error("QR kod oluşturulamadı:", err);
                }
            }

        } catch (error) {
            console.error("Hata oluştu:", error); // Hata durumunda logla

            // Hata mesajı göster
            new Noty({
                type: "error", // Hata türü
                text: "Bağlantı kurulamadı!", // Hata mesajı
                timeout: 3000,
                theme: "mint",
            }).show();
        }
    };
    const wpStopHandler = async () => {
        try {
            const response = await wpStop(); // wpLogin fonksiyonunu çağır
            const data = response.data;
            console.log(data); // Gelen veriyi logla

            // Başarı mesajı göster
            new Noty({
                type: "success", // Başarılı işlem türü
                text: "Bot Durduruldu", // Mesaj içeriği
                timeout: 3000, // 3 saniye
                theme: "mint",
            }).show();
        } catch (error) {
            console.error("Hata oluştu:", error); // Hata durumunda logla

            // Hata mesajı göster
            new Noty({
                type: "error", // Hata türü
                text: "Bağlantı kurulamadı!", // Hata mesajı
                timeout: 3000,
                theme: "mint",
            }).show();
        }
    };

    return (
        <div className="flex flex-col gap-4">
            <div className="p-4 border-4 border-dashed border-warning-500 p-4">
                Whatsapp Durumu: <BotStatus />
                <br></br>
                Bot Durumu : <NumBotStatus />
            </div>
            <div>
                <Button className="mr-4" color="danger" onPress={wpStopHandler}>
                    Oturumu Sil
                </Button>

                <Button color="success" onPress={wpLoginHandler}>
                    Bota Bağlan & QR Kodu Oluştur
                </Button>
            </div>
            <div className="space-y-4">
                <NumaraEkle />
                <TestMesaj />
            </div>


        </div>

    );
}
