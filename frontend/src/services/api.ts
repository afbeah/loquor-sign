import { getToken } from "../utils/auth";

const API_URL = "http://localhost:8080";

export const api = {
  login: async (email: string, password: string) => {
    const response = await fetch(`${API_URL}/login`,{
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || "Erro no login")
    }

    if (data.token) {
      localStorage.setItem("token", data.token);
    }

    return data;
  },

  getSymbols: async () => {
    const response = await fetch(`${API_URL}/symbols`);
    return response.json();
  },

  getPhrases: async () => {
    const response = await fetch(`${API_URL}/phrases`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
    });

    return response.json();
  },

  createPhrase: async (symbolIds: string[]) => {
    const response = await fetch(`${API_URL}/phrases`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${getToken()}`,
      },
      body: JSON.stringify({
        symbols: symbolIds, 
      }),
    });

    return response.json();
  }

}