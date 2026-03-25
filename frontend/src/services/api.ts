const API_URL = "http://localhost:8080"

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
  }
}