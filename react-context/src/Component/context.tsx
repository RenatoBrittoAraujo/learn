import React, { useContext, useState } from "react";

interface Pessoa {
  nome: string;
  idade: number;
  numeroDeSuicidios: number;
}

const PessoaContext = React.createContext<Pessoa[]>([]);
const UpdatePessoaContext = React.createContext<() => void>(() => {});

export const usePessoas = () => {
  return useContext(PessoaContext);
};

export const usePessoasUpdate = () => {
  return useContext(UpdatePessoaContext);
};

const PessoasProvider: React.FC = ({ children }) => {
  const [pessoas, setPessoas] = useState<Pessoa[]>([]);

  const createPessoa = (): Pessoa => {
    return {
      nome: "jose",
      idade: 12,
      numeroDeSuicidios: 2,
    };
  };

  const addPessoas = () => {
    setPessoas([...pessoas, createPessoa()]);
  };

  return (
    <>
      <PessoaContext.Provider value={pessoas}>
        <UpdatePessoaContext.Provider value={addPessoas}>
          {children}
        </UpdatePessoaContext.Provider>
      </PessoaContext.Provider>
    </>
  );
};

export default PessoasProvider;
