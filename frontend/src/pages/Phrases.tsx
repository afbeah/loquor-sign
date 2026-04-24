import { useEffect, useState } from "react";
import { api } from "../services/api";
import { useNavigate } from "react-router-dom"

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
  const navigate = useNavigate()

  useEffect(() => {
    async function loadData() {
      try {
        const symbolsData = await api.getSymbols();
        console.log("Symbols:", symbolsData);
        setSymbols(Array.isArray(symbolsData) ? symbolsData : []);

        const phrasesData = await api.getPhrases();
        console.log("Frases salvas:", phrasesData);
        setSavedPhrases(Array.isArray(phrasesData) ? phrasesData :[]);
      } catch (error) {
        console.error("Erro ao carregar frases:", error)
      }
    }

    loadData();1
  }, []);

  function getSymbolNames(symbolIds: string[]) {
    return symbolIds.map((id) => {
      const symbol = symbols.find((symbol) => symbol.id === id);
      return symbol ? symbol.name : "Símbolo não encontrado";
    })
    .join(" ");
  }

  return (
    <main className="page">
      <section className="card">
        <div className="back-container">
          <button className="back-button" onClick={() => navigate("/menu")}>← Voltar ao menu</button>
        </div>
        
        <div className="logo">Frases salvas</div>
        <p className="subtitile">Visualize as frases construías anteriormente.</p>

        {savedPhrases.length === 0 ? (
          <div className="phrase-box">
            <p>Nenhuma frase salva ainda.</p>
          </div>
        ) : (
          savedPhrases.map((phrase, index) => (
            <div key={phrase.id || index} className="saved-card">
              <p>
                <strong>Frase:</strong> {getSymbolNames(phrase.symbols)}
              </p>
            </div>
          ))
        )}
      </section>
    </main>
  );

}