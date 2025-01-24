# Go Lang

Go é uma linguagem de programação compilada, estaticamente tipada, com sintaxe inspirada em C, que é concisa, segura e eficiente.

## Variáveis

- Funções privadas, protegidas e públicas são diferenciadas pela primeira letra:
- Maiúscula: Função pública.
- Minúscula: Função privada.

## Dependências

- Instalar módulos/dependências: go install <module>.
- Limpar módulos/dependências: go tidy.

## Tipos de Dados

### Number

- Todo número tem valor inicial igual a zero.
- Inteiros: Podem ter 8, 16, 32 ou 64 bytes.
- Inteiros sem sinal: Representados por uint (mesmo esquema de bytes).
- Reais (float):
  - `float32`: Exemplo, 123.45.
  - `float64`: Exemplo, 123000000.45.
- Aliases:
  - `int32` → `rune` (muito usado em tabelas ASCII).
  - `uint8` → `byte`.

### String

- Tipo de dado para textos.

### Boolean

- Valores true ou false.

## Funções

- Criar blocos reutilizáveis de código.

### Função Anônima

- A função anônima é uma função sem nome.

### Função Nomeada

- A função nomeada é uma função com nome.

### Função Variática

- Função que recebe um número variável de argumentos.

### Função Recursiva

- Função que chama a si mesma.

### Função Init

- Função que é executada antes da função main.
- É usada para inicializar variáveis ou realizar outras tarefas de configuração.
- É executada apenas uma vez quando o programa inicia.

## Operadores

- Operadores Aritméticos: Operações matemáticas.
- Operadores de Atribuição: Atribuir valores às variáveis.
- Operadores de Comparação: Comparar valores.
- Operadores Lógicos: Combinar expressões lógicas.
- Operadores Bitwise: Manipulação em nível de bits.
- Operadores Unários: Trabalhar com um único operando.

## Arrays e Slices

- **Array**:
  - Coleção de elementos de mesmo tipo.
  - Tamanho fixo.
- **Slice**:
  - Array dinâmico (com capacidade ajustável).
  - Referência a um array.
  - Sempre tem comprimento (length) e capacidade (capacity).
  - Redimensiona automaticamente ao adicionar mais elementos do que a capacidade atual.

## Ponteiros

- Variáveis que armazenam o endereço de memória de outra variável.

- Exemplo:
  - `*int`: Ponteiro para um inteiro.
  - `&`: Operador para obter o endereço.
  - `*`: Operador para acessar o valor do endereço.

## Struct

- É um tipo composto que agrupa diferentes tipos de dados.

## Maps

- Estrutura de dados para armazenar pares chave-valor.

## Estruturas de Controle

- Blocos condicionais como `if`, `for`, `switch`.

## Cláusula Defer

- A cláusula Defer é útil para adiar um retorno de um bloco de código até o último momento possível.

## Panic e Recover

- Panic é usado para interromper a execução do programa e exibir uma mensagem de erro.
- É utilizado para parar a execução do programa e exibir uma mensagem de erro.
- Recover é usado para recuperar de um panic.
- É utilizado para interromper o panic e continuar a execução do programa.
- Se a função não estiver em PANIC, o RECOVER volta um valor `nil` (nulo)
- PANIC não é do tipo Error.

## Closure

- É uma função que referência uma variável que está fora do seu corpo.
- É uma função que retorna outra função.
- É usado para criar uma função que pode ser utilizada para modificar uma variável.

# Métodos

- É uma função que pode ser chamada em um tipo específico.

---

## Boas Práticas

- Sempre utilizar tipos iguais ao trabalhar com operadores.
- Definir constantes para valores fixos.
- Utilizar inferência implícita onde aplicável.
- Não existem operadores ternários em Go.
- Sempre utilizar diretórios para organizar o projeto

## Benefícios

- Tipagem forte (modelo “Hulk”): Evita erros de tipo.
- Suporte à inferência implícita para tipos.

## Considerações

- Sem operador ternário.
- Arrays têm tamanho fixo.
- Slices possuem um ponteiro que faz referência a um array (equivalente a arrays no JavaScript).
- O uso de "If init" limita variáveis ao escopo do bloco if.
- A cláusula `defer` funciona como se fosse uma Promise no JavaScript?
- O ponteiro força não usarmos imutabilidade nos dados? Até onde é saudável fazermos cópia dos dados?
