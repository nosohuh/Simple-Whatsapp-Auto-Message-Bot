import type { PayloadAction } from "@reduxjs/toolkit";

import { createSlice } from "@reduxjs/toolkit";

type UserType = {
  _id: string;
  username: string;
  refresh_token: string;
  access_token: string;
  createdAt: string;
  updatedAt: string;
};

interface AuthState {
  isLoggin: boolean;
  user: UserType | null;
}

const initialState: AuthState = {
  isLoggin: false,
  user: null,
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login: (state, action: PayloadAction<UserType>) => {
      state.user = action.payload;
      state.isLoggin = true;
    },
    logout: (state) => {
      localStorage.removeItem("login");
      state.user = null;
      state.isLoggin = false;
    },
  },
});

export const { login, logout } = authSlice.actions;
export default authSlice.reducer;
