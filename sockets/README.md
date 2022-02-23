## Trabalho 2

O uso de múltiplos servidores funciona por um sistema de dividir e conquistar,
ou seja, a lista de floats é segmentada em partes iguais para cada um dos
servidores disponíveis (arbitrário) e o resultado é calculado localmente e
então unificado no resultado final.

### Processo de comunicação funciona assim:

#### Cliente

1. Cria array
2. Segmenta array para o numero de servidores disponíveis (arbitrários)
3. Para cada segmento
   1. Enquanto houverem números a serem enviados
      1. Envia no pacote tamanho do buffer temporário
      2. Envia no mesmo pacote floats do buffer (4 bytes cada)
      3. Espera acknowledge
   2. Envia mensagem 'RES' como pedido de resposta
   3. Recebe resposta desta thread
4. Unifica respostas recebidas usando dividir e conquistar
7. Print

#### Servidor

1. Liga o servidor
2. Espera conexão
   1. Quando houver conexão
   2. Enquanto houverem números a serem enviados
      1. Tamanho do buffer temporário
      2. Floats do buffer (4 bytes)
      3. Envia acknowledge
   3. Ao receber RES, calcula resultado sobre buffer final
   4. Envia ao cliente a resposta (min/max, 9 bytes)
   5. Encerra conexão   

### Como usar

### Compile

```
./cmp.sh
```

### Levante um número arbitrário de servidores

```
./server 0.0.0.0 <porta>
```  

*Por questão de clareza, faça cada servidor em um terminal separado*


### Levante uma instância de cliente

```
./client 0.0.0.0 <porta servidor 1> <porta servidor 2> <porta servidor 3> ...
```  

*Por questão de clareza, faça o  cliente em um terminal separado*
*Note que é necessário que pelo menos um servidor seja utilizado*