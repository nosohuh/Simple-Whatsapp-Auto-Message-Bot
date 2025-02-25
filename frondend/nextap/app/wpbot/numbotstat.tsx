import { Chip } from "@nextui-org/react";
import { numBotStat } from "@/axios/axios";
import { useState, useEffect } from "react";

export default function NumBotStatus() {
    const [BotStat, setBotStat] = useState(false);

    const fetchBotStatusHandler = async () => {
        try {
            const response = await numBotStat();
            console.log(response); // Yanıtın tamamını logla
            
            if (response?.data?.status) {
                setBotStat(response?.data?.status); 
            } else {
                setBotStat(false); 
            }
        } catch (error) {
            console.error('Bot durumunu alırken hata oluştu:', error);
            setBotStat(false); 
        }
    };

    useEffect(() => {
        fetchBotStatusHandler(); 
        const interval = setInterval(fetchBotStatusHandler, 10000); 
        return () => clearInterval(interval);
    }, []);

    return (
        <Chip variant="bordered" color={BotStat ? "success" : "danger"}>
            {BotStat ? "Çalışıyor" : "Çalışmıyor"}
        </Chip>
    );
}
