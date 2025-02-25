import {
    Modal,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalFooter,
    Button,
    useDisclosure,
    Textarea,
} from "@nextui-org/react";
import { useState } from "react";
import { addNumbers } from "@/axios/axios";

export default function NumaraEkle() {
    const { isOpen, onOpen, onOpenChange } = useDisclosure();
    const [numbers, setNumbers] = useState(""); // Textarea içeriği için state
    const [message, setMessage] = useState(""); // Gönderilecek mesaj için state

    // Numaraları ve mesajı API'ye gönder
    async function numaraEkle() {
        try {
            const response = await addNumbers(numbers, message); // API isteği yapılır
            console.log("Numaralar başarıyla eklendi:", response);
        } catch (error) {
            console.error("Numara eklenirken bir hata oluştu:", error);
        }
    }

    return (
        <>
            <Button color="secondary" onPress={onOpen}>Numara Ekle</Button>
            <Modal
                backdrop="blur"
                className="capitalize"
                isOpen={isOpen}
                onOpenChange={onOpenChange}
            >
                <ModalContent>
                    {(onClose) => (
                        <>
                            <ModalHeader className="flex flex-col gap-1">Numara Ekle</ModalHeader>
                            <ModalBody>
                                {/* Numaralar için Textarea */}
                                <Textarea
                                    isRequired
                                    className="max-w-xs"
                                    label="Numaralar"
                                    labelPlacement="outside"
                                    placeholder="Lütfen Toplu Şekilde Numaraları Yollayın"
                                    value={numbers} // Numaraların değeri
                                    onChange={(e) => setNumbers(e.target.value)} // Numaraları güncelle
                                />
                                {/* Mesaj için Textarea */}
                                <Textarea
                                    isRequired
                                    className="max-w-xs"
                                    label="Mesaj"
                                    labelPlacement="outside"
                                    placeholder="Gönderilecek mesajı giriniz."
                                    value={message} // Mesajın değeri
                                    onChange={(e) => setMessage(e.target.value)} // Mesajı güncelle
                                />
                            </ModalBody>
                            <ModalFooter>
                                <Button color="danger" variant="light" onPress={onClose}>
                                    Kapat
                                </Button>
                                <Button
                                    color="primary"
                                    onPress={async () => {
                                        await numaraEkle(); // API çağrısını bekler
                                        onClose(); // Modal'ı kapatır
                                    }}
                                >
                                    Kaydet
                                </Button>
                            </ModalFooter>
                        </>
                    )}
                </ModalContent>
            </Modal>
        </>
    );
}
