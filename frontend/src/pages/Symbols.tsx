import { useEffect, useState } from "react"; 
import { api } from "../services/api";

type Symbols = {
  id: string;
  name: string;
  image?: string;
  category_id?: string;
}

export function Symbols() {
  const [symbols, setSymbols] = useState<Symbols[]>([]);

  useEffect(() => {
    async function loadSymbols() {
      const data = await api.getSymbols();
      console.log("Symbols:", data);
      setSymbols(data);
    }

    loadSymbols();
  }, []);

  return (
    <div>
      <h2>Symbols</h2>

      {symbols.map((symbol) => (
        <div key={symbol.id}>
          <p>{symbol.name}</p>
        </div>
      ))}
    </div>
  )
}