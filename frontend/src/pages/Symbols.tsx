import { useEffect, useState } from "react"; 
import { api } from "../services/api";
import { useNavigate } from "react-router-dom";


type Symbol = {
  id: string;
  name: string;
  image?: string;
  category_id?: string;
}

export function Symbols() {
  const [symbols, setSymbols] = useState<Symbol[]>([]);
  const [phrase, setPhrase] = useState<Symbol[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    async function loadData() {
      try {
      const symbolsData = await api.getSymbols();

      console.log("Symbols:", symbolsData);
      
      if (Array.isArray(symbolsData)) {
        setSymbols(symbolsData);
      } else {
        console.error("Formato inválido", symbolsData)
        setSymbols([]);
      }

      } catch (error) {
        console.error("Erro ao buscar símbolo:", error);
      }
    }

    loadData();
  }, []);

  function addToPhrase(symbol: Symbol) {
    console.log("Símbolo clicado:", symbol);
    setPhrase((prev) => [...prev, symbol])
  }

  async function handleSavePhrase() {
    if (phrase.length === 0) {
      alert("Selecione pelo menos um símbolo antes de salvar.");
      return;
    }
    try {
      const symbolIds = phrase.map((symbol) => symbol.id);

      console.log("Enviando frase:", symbolIds);

      const response = await api.createPhrase(symbolIds);

      console.log("Frase salva:", response);

      setPhrase([]);
    } catch (error) {
      console.error("Erro ao salvar frase:", error);
    }
  }

  return (
    <main className="page">
      <section className="card">
        <div className="back-container">
          <button className="back-button" onClick={() => navigate("/menu")}>← Voltar ao menu</button>
        </div>

        <div className="logo">Criar frase</div>
        <p className="subtitle">Selecione os símbolos para montar uma mensagem</p>

        <h3>Frase Montada</h3>

        <div className="phrase-box">
          {phrase.length === 0 ? (
            <p>Nenhum símbolo selecionado ainda.</p>
          ) : (
            phrase.map((symbol, index) => (
              <span key={index} className="chip">
                {symbol.name}
              </span>
            ))
          )}
        </div>

        <button className="button" onClick={handleSavePhrase}>Salvar frase</button>

        <h3 style={{ marginTop: 32 }}>Símbolos disponíveis</h3>

        <div className="symbol-grid">
          {symbols.length === 0 ? (
            <p>Nenhum símbolo carregado.</p>
          ) : (
            symbols.map((symbol) => (
              <div 
                key={symbol.id}
                onClick={() => addToPhrase(symbol)}
                className="symbol-card"
              >
                {symbol.name}
              </div>
            ))
          )}
        </div>
      </section>
    </main>
  );
}