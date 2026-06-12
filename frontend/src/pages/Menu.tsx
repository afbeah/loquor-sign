import { useNavigate  } from "react-router-dom";

export function Menu() {
  const navigate = useNavigate();

  return (
    <main className="menu-page">
      <section className="card">
        <div className="menu-logo">Loquor Sign</div>

        <h1 className="menu-title">Menu Principal</h1>

        <p className="menu-subtitle">Escolha uma opção para continuar</p>

        <div className="menu-actions">
          <button className="button" onClick={() => navigate("/symbols")}>Criar Frases</button>

          <button className="button" onClick={() => navigate()}>Voice Controller</button>

          <button className="button-outline" onClick={() => navigate("/phrases")}>Ver Frases</button>
        </div>

      </section>
    </main>
  )
}