import { usePessoasUpdate } from "./context";
import OtherComp from "./othercomp";

const Comp = () => {
  const addPessoas = usePessoasUpdate();

  return (
    <>
      <button onClick={addPessoas}>Toggle</button>
      <OtherComp />
    </>
  );
};

export default Comp;
