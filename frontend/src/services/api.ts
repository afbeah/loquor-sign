import { getToken } from "../utils/auth";

const API_URL = "http://localhost:8080";

export const api = {
  login: async (email: string, password: string) => {
    const response = await fetch(`${API_URL}/login`,{
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({email, password}),
    })

    return response.json()
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
    const response = await fetch("http://localhost:8080/phrases", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify({
        symbols: symbolIds, 
      }),
    });

    return response.json();
  }

}