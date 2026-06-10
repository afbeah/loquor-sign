import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { api } from "../services/api";

export function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      await api.login(email, password);

      navigate("/menu");
    } catch (error) {
      alert("Usuário ou senha inválidos");
    }
  };

  return (
    <main className="page">
      <section className="card">
        <section className="card-small">
          
          <div className="logo">Loquor Sign</div>
          <p className="subtitle">
            Plataforma de apoio à comunicação alternativa
          </p>

          <input
            className="input"
            placeholder="E-mail"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <input
            className="input"
            placeholder="Senha"
            type="password"
            onChange={(e) => setPassword(e.target.value)}
          />

          <button className="button" onClick={handleLogin}>
            Entrar
          </button>
        </section>
      </section>
    </main>
  );
}
