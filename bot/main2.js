const express = require('express');
const { Client, LocalAuth } = require('whatsapp-web.js');
const qrcode = require('qrcode');
const fs = require('fs').promises;
const path = require('path');
const app = express();
const port = 4020;

app.use(express.json());

let client;
let qrCodeDataURL = null;  // QR kodunu saklamak için
let isClientReady = false; // Botun hazır durumunu kontrol etmek için
let lastMessage = null;    // Alınan son mesajı saklamak için

// .wwebjs_auth ve .wwebjs_cache klasörlerini silme işlemi
const cleanUpDirs = async () => {
  const authDir = path.join(__dirname, '.wwebjs_auth');
  const cacheDir = path.join(__dirname, '.wwebjs_cache');

  // Klasörleri kontrol et ve varsa sil
  try {
    await fs.rm(authDir, { recursive: true, force: true });
    await fs.rm(cacheDir, { recursive: true, force: true });
    console.log('.wwebjs_auth ve .wwebjs_cache klasörleri silindi.');
  } catch (err) {
    console.error('Klasörleri silerken bir hata oluştu:', err.message);
  }
};

const startBot = () => {
  cleanUpDirs();

  client = new Client({
    authStrategy: new LocalAuth(), // Oturumunuzu sakla
  });

  client.on('qr', (qr) => {
    qrcode.toDataURL(qr, (err, url) => {
      if (err) {
        console.error('QR kodu oluşturulamadı:', err);
        return;
      }
      qrCodeDataURL = url;  // QR kod
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

// Oturum dosyasının kilitlenmemesi için bekleme fonksiyonu
const waitForFileRelease = async (filePath) => {
  let retries = 5;
  while (retries > 0) {
    try {
      await fs.open(filePath, 'r');  // Dosya üzerinde okuma işlemi yapmaya çalışıyoruz
      break;
    } catch (err) {
      retries--;
      console.log(`Dosya erişilemiyor, tekrar denenecek: ${retries} deneme kaldı`);
      await new Promise(resolve => setTimeout(resolve, 1000)); // 1 saniye bekle
    }
  }
};

// Botu durdurma ve oturum dosyalarını temizleme
const stopBot = async () => {
  if (client) {
    try {
      // Logout işlemi öncesinde dosyanın serbest bırakıldığından emin ol
      const sessionPath = path.join(__dirname, '.wwebjs_auth', 'session', 'Default', 'chrome_debug.log');
      await waitForFileRelease(sessionPath);

      await client.logout(); // Logout işlemi
      await client.destroy(); // Botu tamamen durdurun
      client = null; // client'i null yapın
      isClientReady = false; // Botun durumunu sıfırlayın
      lastMessage = null; // Son mesajı sıfırlayın
      console.log('Bot başarıyla durduruldu!');

      // .wwebjs_auth ve .wwebjs_cache klasörlerini silme
      const authDir = path.join(__dirname, '.wwebjs_auth');
      const cacheDir = path.join(__dirname, '.wwebjs_cache');

      // Klasörleri sil
      await fs.rm(authDir, { recursive: true, force: true });
      await fs.rm(cacheDir, { recursive: true, force: true });

      console.log('.wwebjs_auth ve .wwebjs_cache klasörleri silindi.');
    } catch (err) {
      console.error('Bot durdurulurken bir hata oluştu:', err.message);
    }
  }
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
    stopBot(); // Botu durdur
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
