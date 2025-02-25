import { Chip } from "@nextui-org/react";
import { botStatus } from "@/axios/axios";
import { useState, useEffect } from "react";
import Noty from "noty";

export default function BotStatus() {
    const [BotStat, setBotStat] = useState(false);

    const fetchBotStatusHandler = async () => {
        try {
            const response = await botStatus();

            // Yanıtın tam yapısını kontrol et
            console.log(response); // Yanıtın tamamını logla
            const data = response?.data?.response; // doğru erişimi kontrol et

            if (data) {
                setBotStat(data?.isReady);
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

    useEffect(() => {
        fetchBotStatusHandler(); // Başlangıçta bot durumunu al
        const interval = setInterval(fetchBotStatusHandler, 25000); // Her 25 saniyede bir tekrar et

        // Cleanup: Interval'i temizle
        return () => clearInterval(interval);
    }, []);

    return (
        <Chip variant="bordered" color={BotStat ? "success" : "danger"}>
            {BotStat ? "Bağlı" : "Bağlı Değil"}
        </Chip>
    );
}
