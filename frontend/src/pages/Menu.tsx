import { useNavigate  } from "react-router-dom";

export function Menu() {
  const navigate = useNavigate();

  return (
    <main className="page">
      <section className="card card-small">
        <div className="logo">Loquor Sign</div>
        <p>Escolha uma opção para continuar</p>

        <div className="menu-actions">
          <button className="button" onClick={() => navigate("/symbols")}>Criar frase</button>

          <button className="button button-secondary" onClick={() => navigate("/phrases")}>Ver frases salvas</button>

        </div>
      </section>
    </main>
  )
}