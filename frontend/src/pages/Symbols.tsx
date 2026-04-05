import { useEffect, useState } from "react"; 
import { api } from "../services/api";

type Symbol = {
  id: string;
  name: string;
  image?: string;
  category_id?: string;
}

type SavedPhrase = {
  id: string;
  user_id?: string;
  symbols: string[];
  created_at?: string;
}

export function Symbols() {
  const [symbols, setSymbols] = useState<Symbol[]>([]);
  const [phrase, setPhrase] = useState<Symbol[]>([]);
  const [savedPhrases, setSavedPhrases] = useState<SavedPhrase[]>([]);

  async function loadPhrases() {
    try {
      const data = await api.getPhrases();
      console.log("Frases salvas:", data);
      setSavedPhrases(data);
    } catch (error) {
      console.log("Erro ao buscar frases:", error)
    }
  }

  useEffect(() => {
    async function loadData() {
      try {
      const symbolsData = await api.getSymbols();
      console.log("Symbols:", symbolsData);
      setSymbols(symbolsData);

      const phrasesData = await api.getPhrases();
      console.log("Frases salvas:", phrasesData);
      setSavedPhrases(phrasesData)
      } catch (error) {
        console.error("Erro ao buscar símbolo:", error);
      }
    }

    loadData();
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

      setPhrase([]);
      await loadPhrases();
    } catch (error) {
      console.error("Erro ao salvar frase:", error);
    }
  }

  function getSymbolNames(symbolIds: string[]) {
    return symbolIds.map((id) => {
      const symbol = symbols.find((symbol) => symbol.id === id);
      return symbol ? symbol.name : "Símbolo não encontrado";
    })
    .join(" ");
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

      <h3>Frases salvas</h3>
      <div>
        {savedPhrases.length === 0 ? (
          <p>Nenhuma frase salva ainda.</p>
        ) : (
          savedPhrases.map((phrase, index) => (
            <div key={phrase.id || index}
            style={{
              border: "1px solid gray",
              padding: 10,
              marginBottom: 10,
            }}>
              <p><strong>ID:</strong> {phrase.id}</p>
              <p><strong>Frase:</strong> {getSymbolNames(phrase.symbols)}</p>
            </div>
          ))
        )}
      </div>

    </div>
  );
}