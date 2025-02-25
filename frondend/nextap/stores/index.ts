// stores.ts
import { configureStore } from "@reduxjs/toolkit";

import authSlice from "./authSlice"; // authSlice'ın doğru yolunu kontrol edin

export const store = configureStore({
  reducer: {
    auth: authSlice,
  },
});

// TypeScript tiplerini tanımlayın
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
