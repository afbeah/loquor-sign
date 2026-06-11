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
    <main className="login-layout">
      <section className="login-branding">
        <div className="branding-content">
          <div className="logo">Loquor Sign</div>

          <h2>Plataforma de apoio à comunicação alternativa</h2>

          <p className="branding-text">
            Transformando comunicação em possibilidades
          </p>
        </div>
      </section>

      <section className="login-container">
        <section className="card">
          <h1 className="login-title">Bem-vindo!</h1>

          <p className="subtitle">Faça login para acessar sua conta</p>

          <input
            className="input"
            placeholder="E-mail"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <input
            className="input"
            placeholder="Senha"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />

          <button className="button" onClick={handleLogin}>
            Entrar
          </button>

          <div className="divider">
            <span>ou</span>
          </div>

          <button
            className="button-outline"
            onChange={() => navigate("/register")}
          >
            Criar conta
          </button>
        </section>
      </section>
    </main>
  );
}
