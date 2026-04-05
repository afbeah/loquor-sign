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
  const [phrase, setPhrase] = useState<Symbol[]>([]);

  useEffect(() => {
    async function loadSymbols() {
      try {
      const data = await api.getSymbols();
      console.log("Symbols:", data);
      setSymbols(data);
      } catch (error) {
        console.error("Erro ao buscar símbolo:", error);
      }
    }

    loadSymbols();
  }, []);

  function addToPhrase(symbol: Symbol) {
    console.log("Símbolo clivado:", symbol);
    setPhrase((prev) => [...prev, symbol])
  }

  async function handleSavePhrase() {
    try {
      const symbolIds = phrase.map((symbol) => symbol.id);

      console.log("Enviando frase:", symbolIds);

      const response = await api.createPhrase(symbolIds);

      console.log("Frase salva:", response);
    } catch (error) {
      console.error("Erro ao salvar frase:", error);
    }
  }

  return (
    <div>
      <h2>Symbols</h2>

      <h3>Frase:</h3>
      <div style={{ marginBottom: 20}}>
        {phrase.length === 0 ? (
          <p>Nenhum símbolo selecionado ainda.</p>
        ) : (
          phrase.map((symbol, index) => (
            <span key={index} style={{ marginRight: 8 }}>
              {symbol.name}
            </span>
          ))
        )}
      </div>

      <button onClick={handleSavePhrase}>Salvar frase</button>

      <h3>Lista de Símbolos</h3>
      <div>
        {symbols.length === 0 ? (
          <p>Nenhum símbolo carregado.</p>
        ) : (
          symbols.map((symbol) => (
            <div key={symbol.id} onClick={() => addToPhrase(symbol)} style={{
              border: "1px solid gray",
              padding: 10,
              marginBottom: 10,
              cursor: "pointer",
            }}>
              <p>{symbol.name}</p>
            </div>
          ))
        )}
      </div>
    </div>
  )
}