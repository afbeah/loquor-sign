import { useNavigate  } from "react-router-dom";

export function Menu() {
  const navigate = useNavigate();

  return (
    <div style={{ textAlign: "center" }}>
      <h2>MENU</h2>

      <button onClick={() => navigate("/symbols")} style={{ margin: 10 }}>Criar frase</button>

      <button onClick={() => navigate("/phrases")} style={{ margin: 10}}>Ver frases salvas</button>
    </div>
  )
}