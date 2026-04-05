import { useEffect, useState } from "react";
import { api } from "../services/api";

type Symbol = {
  id: string;
  name: string;
  image?: string;
  category_id?: string;
};

type SavedPhrase = {
  id: string;
  user_id?: string;
  symbols: string[];
  created_at?: string;
};

export function Phrases() {
  const [symbols, setSymbols] = useState<Symbol[]>([]);
  const [savedPhrases, setSavedPhrases] = useState<SavedPhrase[]>([]);

  useEffect(() => {
    async function loadData() {
      try {
        const symbolsData = await api.getSymbols();
        console.log("Symbols:", symbolsData);
        setSymbols(symbolsData);

        const phrasesData = await api.getPhrases();
        console.log("Frases salvas:", phrasesData);
        setSavedPhrases(phrasesData);
      } catch (error) {
        console.error("Erro ao carregar frases:", error)
      }
    }

    loadData();
  }, []);

  function getSymbolNames(symbolIds: string[]) {
    return symbolIds.map((id) => {
      const symbol = symbols.find((symbol) => symbol.id === id);
      return symbol ? symbol.name : "Símbolo não encontrado";
    })
    .join(" ");
  }

  return (
    <div>
      <h2>Frases Salvas</h2>

      {savedPhrases.length === 0 ? (
        <p>Nenhuma frase salva ainda.</p>
      ) : (
        savedPhrases.map((phrase, index) => (
          <div key={phrase.id || index} style={{
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
  );

}