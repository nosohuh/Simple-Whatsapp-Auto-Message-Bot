import axios from "axios";
const HTTP = axios.create({
  baseURL: "http://127.0.0.1:8080/oauth",
  withCredentials: true,
});

export const loginRequest = async (username: string, password: string) =>
  await HTTP.post("/user/login", {
    username,
    password,
  });

export const registerRequest = async (username: string, password: string) =>
  await HTTP.post("/auth/register", {
    username,
    password,
  });

export const userControlRequest = async () => await HTTP.get("/user/token");

export const logoutRequest = async () => await HTTP.get("/user/logout");

export const allPostRequest = async (date: string, durum: boolean) => {
  // Durum parametresini "true" veya "false" string olarak gönder
  return await HTTP.get('/servis/alllist', {
    params: {
      date,
      durum: durum.toString(), // Durumu string olarak gönder
    },
  });
};


export const AllUserList = async () => await HTTP.get("/oauth/user/get-all");

export const EditUserAction = async (
  action: string,
  edit_option: string | null,
  new_data: string | null,
  id: string,
  role: string | null,
  username: string | null,
  password: string | null
) => {
  const requestBody = {
    action: action,
    edit_option: edit_option ?? "",
    new_data: new_data ?? "",
    id,
    role: role ?? "",
    username: username ?? "",
    password: password ?? "",
  };

  return await HTTP.post("/user/action", requestBody);
};




export const wpLogin = async () =>
  await HTTP.post("/servis/wplogin");

export const wpStop = async () =>
  await HTTP.post("/servis/wpstop");

export const botStatus = async () =>
  await HTTP.post("/servis/bot-status");

export const wpQr = async () =>
  await HTTP.post("/servis/qr");

export const sendTestMesagge = async (number: string, message: string) => {
  return await HTTP.post('servis/testmessage', {
    number: number,
    message: message,
  });
};

export const addNumbers = async (numbers: string, message: string) => {
  return await HTTP.post("servis/addnumber", {
    "Numara": numbers, // Backend'in beklediği key ismi burada "Numara"
    "Mesaj": message,
    "durum": false,
  });
};


export const getNumberList = async () =>
  await HTTP.get("/servis/getnumber");

export const numBotStat = async () =>
  await HTTP.get("/servis/numbot-status");

export const startNumBot = async (id: string) => {
  await HTTP.get(`/servis/numbot/${id}`);
};




export const getOnePost = async (postid: string) =>
  await HTTP.get(`/posts/${postid}`);
