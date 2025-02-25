const express = require('express');
const { Client, LocalAuth } = require('whatsapp-web.js');
const qrcode = require('qrcode');
const app = express();
const port = 4020;

app.use(express.json());

let client;
let qrCodeDataURL = null;  // QR kodunu saklamak için
let isClientReady = false; // Botun hazır durumunu kontrol etmek için
let lastMessage = null;    // Alınan son mesajı saklamak için

// WhatsApp botunu başlatma
const startBot = () => {
  client = new Client({
    authStrategy: new LocalAuth(), // Oturumunuzu saklamak için LocalAuth kullanıyoruz
  });

  client.on('qr', (qr) => {
    // QR kodunu base64 URL'ye çevir
    qrcode.toDataURL(qr, (err, url) => {
      if (err) {
        console.error('QR kodu oluşturulamadı:', err);
        return;
      }
      qrCodeDataURL = url;  // QR kodunu sakla
      console.log('QR kodu oluşturuldu!');
    });
  });

  client.on('ready', () => {
    isClientReady = true;  // Bot hazır
    console.log('Bot hazır ve çalışıyor!');
  });

  client.on('message', (message) => {
    lastMessage = {
      from: message.from,
      body: message.body,
      timestamp: message.timestamp,
    };
    console.log(`Mesaj alındı: ${message.body}`);
  });

  client.initialize();
};

// API Endpoints

// Botu başlatan ve işlemi kontrol eden endpoint
app.post('/bot-action', (req, res) => {
  const { action, number, message } = req.body;

  if (action === 'start') {
    if (client) {
      return res.json({ status: 400, message: 'Bot zaten çalışıyor! QR kodunu almak için /qr endpointine gidin.' });
    }

    startBot();
    return res.json({ status: 200, message: 'Bot başlatıldı! QR kodunu almak için /qr endpointine gidin.' });
  }

  if (action === 'stop') {
    if (client) {
        client.destroy()
          .then(() => {
            console.log('Bot başarıyla durduruldu.');
            client = null; // Client referansını sıfırla
          })
          .catch((err) => {
            console.error('Bot durdurulurken bir hata oluştu:', err.message);
          });
      }
      

    client.destroy(); // Botu durdur
    isClientReady = false; // Durumu sıfırla
    lastMessage = null; // Mesaj geçmişini sıfırla
    return res.json({ status: 200, message: 'Bot durduruldu!' });
  }

  if (action === 'qr') {
    if (!qrCodeDataURL) {
      return res.status(400).json({ error: 'QR kodu henüz oluşturulmadı! Lütfen botu başlatın.' });
    }
    return res.json({ qrCode: qrCodeDataURL });
  }

  if (action === 'send-message') {
    console.log('Received request:', req.body);  // Gelen veriyi logla
  
    if (!client) {
      return res.status(400).json({ error: 'Bot çalışmıyor!' });
    }
  
    if (!number || !message) {
      return res.status(400).json({ error: 'Numara ve mesaj gereklidir!' });
    }
  
    const chatId = number.includes('@c.us') ? number : `${number}@c.us`;
  
    // Mesaj gönderme işlemi
    client.sendMessage(chatId, message)
      .then(() => {
        return res.json({ message: 'Mesaj başarıyla gönderildi!' });
      })
      .catch(err => {
        return res.status(500).json({ error: 'Mesaj gönderilemedi!', details: err.message });
      });
  }
  
});

// Botun hazır durumu için endpoint
app.get('/bot-status', (req, res) => {
  return res.json({ isReady: isClientReady });
});

// Alınan son mesajı sorgulamak için endpoint
app.get('/last-message', (req, res) => {
  if (!lastMessage) {
    return res.status(404).json({ error: 'Henüz bir mesaj alınmadı!' });
  }
  return res.json(lastMessage);
});

// API'yi dinlemeye başla
app.listen(port, () => {
  console.log(`API çalışıyor, http://localhost:${port}`);
});
