import React, { useContext } from "react";
import { usePessoas } from "./context";

const OtherComp = () => {
  const pessoas = usePessoas();

  return (
    <div>
      {pessoas.map((e) => (
        <li>{e.nome}</li>
      ))}
    </div>
  );
};

export default OtherComp;
