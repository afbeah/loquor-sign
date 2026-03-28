import { useEffect, useState } from "react"; 
import { api } from "../services/api";

type Symbol = {
  id: string;
  name: string;
  image?: string;
  category_id?: string;
}

export function Symbols() {
  const [symbols, setSymbols] = useState<Symbol[]>([]);

  useEffect(() => {
    async function loadSymbols() {
      const data = await api.getSymbols();
      console.log("Symbols:", data);
      setSymbols(data);
    }

    loadSymbols();
  }, []);

  const [phrase, setPhrase] = useState<Symbol[]>([]);

  function addToPhrase(symbol: Symbol) {
    setPhrase((prev) => [...prev, symbol])
  }

  return (
    <div>
      <h2>Symbols</h2>

      <h3>Frase:</h3>

      <div>
        {phrase.map((symbol, index) => (
          <span key={index} style={{ marginRight: 8 }}>
            {symbol.name}
          </span>
        ))}
      </div>

      {symbols.map((symbol) => (
        <div key={symbol.id} onClick={() => addToPhrase(symbol)}>
          <p>{symbol.name}</p>
        </div>
      ))}
    </div>
  )
}