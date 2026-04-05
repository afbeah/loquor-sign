import { Login } from "./pages/Login";
import { Symbols } from "./pages/Symbols";
import { Menu } from "./pages/Menu";
import { BrowserRouter, Routes, Route } from "react-router-dom";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/menu" element={<Menu />} />
        <Route path="/symbols" element={<Symbols />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;