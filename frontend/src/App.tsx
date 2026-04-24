import { Login } from "./pages/Login";
import { Symbols } from "./pages/Symbols";
import { Menu } from "./pages/Menu";
import { Phrases } from "./pages/Phrases";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import {} from "./App.css"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />}/>
        <Route path="/menu" element={<Menu />}/>
        <Route path="/symbols" element={<Symbols />}/>
        <Route path="/phrases" element={<Phrases />}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;