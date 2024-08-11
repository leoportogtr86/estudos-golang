## Exercícios sobre Variáveis e Constantes em Go

### Exercício 1: Declaração de Variáveis com `var`

1. **Objetivo**: Compreender a declaração básica de variáveis.
    - **Tarefa**: Declare três variáveis (`nome`, `idade`, `altura`) usando `var` e atribua valores a elas. Imprima os
      valores no console.
    - **Pergunta**: O que acontece se você tentar usar uma variável que foi declarada mas não inicializada?

### Exercício 2: Inferência de Tipo com `:=`

2. **Objetivo**: Praticar a inferência de tipo.
    - **Tarefa**: Declare variáveis `cidade`, `populacao`, e `densidade` usando o operador `:=` e atribua valores
      adequados a elas. Imprima os tipos inferidos pelo compilador.
    - **Pergunta**: Como Go determina o tipo das variáveis quando se usa `:=`?

### Exercício 3: Múltiplas Declarações com `var`

3. **Objetivo**: Declarar múltiplas variáveis em uma única linha.
    - **Tarefa**: Declare três variáveis `a`, `b`, e `c` do tipo `int` em uma única linha, inicialize-as com valores
      diferentes e imprima os valores.
    - **Pergunta**: Como a sintaxe de múltiplas declarações facilita o código em Go?

### Exercício 4: Múltiplas Declarações com `:=`

4. **Objetivo**: Declarar e inicializar múltiplas variáveis em uma única linha usando `:=`.
    - **Tarefa**: Use `:=` para declarar e inicializar três variáveis de diferentes tipos em uma única linha. Imprima os
      valores e seus tipos.
    - **Pergunta**: Em quais situações a declaração múltipla com `:=` pode ser mais útil?

### Exercício 5: Declaração e Inicialização de Constantes

5. **Objetivo**: Trabalhar com constantes em Go.
    - **Tarefa**: Declare e inicialize três constantes (`pi`, `e`, `g`) para representar valores científicos
      importantes. Imprima esses valores.
    - **Pergunta**: O que diferencia uma constante de uma variável em Go?

### Exercício 6: Múltiplas Constantes

6. **Objetivo**: Declarar múltiplas constantes relacionadas.
    - **Tarefa**: Crie um bloco `const` para declarar os dias da semana como constantes (`Domingo`, `Segunda`, etc.).
      Imprima os valores dessas constantes.
    - **Pergunta**: Como a declaração de constantes em bloco melhora a organização do código?

### Exercício 7: Constantes Não Tipadas

7. **Objetivo**: Compreender o uso de constantes não tipadas.
    - **Tarefa**: Declare uma constante não tipada `valor` e use-a em diferentes contextos (como `int`, `float64`).
      Imprima os resultados.
    - **Pergunta**: Quais são as vantagens das constantes não tipadas em Go?

### Exercício 8: Escopo de Variáveis

8. **Objetivo**: Entender o escopo de variáveis.
    - **Tarefa**: Declare uma variável `x` dentro de uma função e tente acessá-la fora dessa função. Observe o erro e
      explique.
    - **Pergunta**: Por que é importante limitar o escopo das variáveis?

### Exercício 9: Reatribuição de Variáveis

9. **Objetivo**: Praticar a reatribuição de variáveis.
    - **Tarefa**: Declare uma variável `contador` e reatribua novos valores a ela dentro de um loop `for`. Imprima o
      valor de `contador` em cada iteração.
    - **Pergunta**: Quando é apropriado reatribuir variáveis em Go, e quando é preferível usar constantes?

### Exercício 10: Boas Práticas na Declaração

10. **Objetivo**: Implementar boas práticas na declaração de variáveis e constantes.
    - **Tarefa**: Revise um trecho de código existente (pode ser fictício) e identifique onde variáveis poderiam ser
      substituídas por constantes ou onde o escopo poderia ser mais restrito. Reescreva o código com essas melhorias.
    - **Pergunta**: Como a aplicação de boas práticas na declaração de variáveis e constantes melhora a legibilidade e a
      segurança do código?

Esses exercícios ajudarão a reforçar o entendimento sobre a declaração e inicialização de variáveis e constantes em Go,
além de promover a aplicação de boas práticas no desenvolvimento de código.