import React from "react";
import logo from "./logo.svg";
import "./App.css";
import Comp from "./Component/comp";
import PessoasProvider from "./Component/context";

function App() {
  return (
    <div className="App">
      <PessoasProvider>
        <Comp />
      </PessoasProvider>
    </div>
  );
}

export default App;
