## Exercícios sobre Tipos de Dados em Go

### Exercício 1: Tipos Numéricos Inteiros

1. **Objetivo**: Compreender e utilizar diferentes tipos de inteiros.
    - **Tarefa**: Declare variáveis dos tipos `int`, `int8`, `uint`, `uint16` e atribua valores a elas. Imprima esses
      valores.
    - **Pergunta**: O que acontece se você tentar atribuir um valor fora do intervalo permitido para `int8`?

### Exercício 2: Tipos Numéricos de Ponto Flutuante

2. **Objetivo**: Utilizar tipos de ponto flutuante.
    - **Tarefa**: Declare variáveis dos tipos `float32` e `float64` e atribua valores a elas. Realize operações
      aritméticas e imprima os resultados.
    - **Pergunta**: Qual é a diferença de precisão entre `float32` e `float64`?

### Exercício 3: Tipos Booleanos

3. **Objetivo**: Trabalhar com o tipo `bool`.
    - **Tarefa**: Declare duas variáveis booleanas e atribua valores `true` e `false`. Utilize uma estrutura de
      controle `if` para imprimir mensagens baseadas nos valores dessas variáveis.
    - **Pergunta**: O que acontece se você usar operadores lógicos (`&&`, `||`) com variáveis booleanas?

### Exercício 4: Tipos String

4. **Objetivo**: Manipular strings.
    - **Tarefa**: Declare uma variável `string` e atribua a ela uma frase. Imprima a string e seu comprimento
      utilizando `len()`.
    - **Pergunta**: Como você pode concatenar duas strings em Go?

### Exercício 5: Arrays

5. **Objetivo**: Compreender e utilizar arrays.
    - **Tarefa**: Declare um array de inteiros com tamanho 5 e inicialize-o com valores. Modifique um dos valores do
      array e imprima o array completo.
    - **Pergunta**: O que acontece se você tentar acessar um índice fora dos limites do array?

### Exercício 6: Slices

6. **Objetivo**: Trabalhar com slices.
    - **Tarefa**: Crie um slice de strings com valores iniciais. Utilize a função `append` para adicionar mais valores
      ao slice e imprima o slice resultante.
    - **Pergunta**: Qual a diferença entre a capacidade (`cap`) e o comprimento (`len`) de um slice?

### Exercício 7: Maps

7. **Objetivo**: Utilizar maps para armazenar pares chave-valor.
    - **Tarefa**: Declare um map com chaves do tipo `string` e valores do tipo `int`. Adicione, modifique e exclua pares
      chave-valor, imprimindo o map após cada operação.
    - **Pergunta**: O que acontece se você tentar acessar uma chave que não existe no map?

### Exercício 8: Structs

8. **Objetivo**: Compreender e utilizar structs.
    - **Tarefa**: Declare uma struct `Pessoa` com campos `Nome` e `Idade`. Crie uma instância dessa struct, atribua
      valores e imprima os valores dos campos.
    - **Pergunta**: Como você pode criar uma função que recebe uma struct `Pessoa` como argumento e imprime uma mensagem
      formatada com seus campos?

### Exercício 9: Ponteiros

9. **Objetivo**: Trabalhar com ponteiros.
    - **Tarefa**: Declare uma variável `int` e um ponteiro para essa variável. Modifique o valor da variável utilizando
      o ponteiro e imprima o resultado.
    - **Pergunta**: O que acontece se você tentar acessar o valor de um ponteiro não inicializado?

### Exercício 10: Funções como Tipos

10. **Objetivo**: Utilizar funções como tipos.
    - **Tarefa**: Declare uma variável do tipo função que recebe dois `int` e retorna um `int`. Atribua a essa variável
      uma função anônima que soma dois números e imprima o resultado da chamada dessa função.
    - **Pergunta**: Como você pode passar uma função como argumento para outra função em Go?

Esses exercícios ajudarão a reforçar os conceitos abordados no artigo, proporcionando uma compreensão prática dos
diferentes tipos de dados básicos e compostos em Go.