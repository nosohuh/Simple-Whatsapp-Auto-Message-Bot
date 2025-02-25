import {
    Modal,
    ModalContent,
    ModalHeader,
    ModalBody,
    ModalFooter,
    Button,
    useDisclosure,
    Input,
  } from "@nextui-org/react";
  import { sendTestMesagge } from "@/axios/axios";
  import { useState } from "react";
  import Noty from "noty";
  
  export default function TestMesaj() {
    const { isOpen, onOpen, onOpenChange } = useDisclosure();
    const [number, setNumber] = useState("");
    const [message, setMessage] = useState("");
  
    const handleSend = async () => {
      try {
        if (!number || !message) {
          new Noty({
            type: "error",
            text: "Numara ve mesaj boş bırakılamaz!",
            timeout: 3000,
          }).show();
          return;
        }
  
        const response = await sendTestMesagge(number, message);
        new Noty({
          type: "success",
          text: "Mesaj başarıyla gönderildi!",
          timeout: 3000,
        }).show();
        console.log(response.data);
      } catch (error) {
        console.error("Hata:", error);
        new Noty({
          type: "error",
          text: "Mesaj gönderimi başarısız!",
          timeout: 3000,
        }).show();
      }
    };
  
    return (
      <>
        <Button color="success" onPress={onOpen}>Test Mesaj</Button>
        <Modal
          backdrop="blur"
          className="capitalize"
          isOpen={isOpen}
          onOpenChange={onOpenChange}
        >
          <ModalContent>
            {(onClose) => (
              <>
                <ModalHeader className="flex flex-col gap-1">Test Mesaj</ModalHeader>
                <ModalBody>
                  <Input
                    isRequired
                    className="max-w-xs"
                    value={number}
                    onChange={(e) => setNumber(e.target.value)}
                    label="Numara"
                    type="tel"
                    placeholder="+90555 444 33 22"
                  />
                  <Input
                    isRequired
                    className="max-w-xs"
                    value={message}
                    onChange={(e) => setMessage(e.target.value)}
                    label="Mesaj"
                    type="text"
                    placeholder="Mesaj içeriği"
                  />
                </ModalBody>
                <ModalFooter>
                  <Button color="danger" variant="light" onPress={onClose}>
                    Kapat
                  </Button>
                  <Button color="primary" onPress={handleSend}>
                    Gönder
                  </Button>
                </ModalFooter>
              </>
            )}
          </ModalContent>
        </Modal>
      </>
    );
  }
  