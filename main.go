package main

import (
	"fmt"
	"os"
)

type Conta struct {
  Titular string
  Agencia int 
  Conta int 
  Saldo float32
  Senha string
} 

func (c *Conta) Depositar(quantia float32) {
  if quantia > 0 {
    c.Saldo += quantia
  }
}

func (c *Conta) Sacar(quantia float32) float32 {
  if quantia > c.Saldo {
    fmt.Print("\nVocê não possui isso!\n\n")
  } else if quantia > 0 {
    c.Saldo -= quantia
    return c.Saldo
  }

  return -1
}

func main() {
  var contas []Conta

  for {
    exibeComandos([]string{"1 - Criar conta", "2 - Entrar na conta", "3 - Encerrar"}) 

    var comando int
    fmt.Scan(&comando)
    fmt.Println()

    switch (comando) {
      case 1:
        conta := criaConta()
        contas = append(contas, conta)

      case 2:
        login, conta := entraConta(contas)
        if login {
          exibeMenu(conta)
        }

      case 3:
        fmt.Println("Encerrando")
        os.Exit(0)

      default:
        fmt.Println("Comando invalido")
        break
    }
  }
}

func exibeMenu(conta Conta) {
  for {
    exibeComandos([]string{"Você está logado na conta de: " + conta.Titular, "1 - Consultar saldo", "2 - Depositar", "3 - Sacar", "4 - Sair"})
    var comando int
    var sair bool

    fmt.Scan(&comando)

    switch(comando) {
      case 1:
        fmt.Printf("\nSeu saldo atual: R$%.2f\n\n", conta.Saldo)

      case 2:
        var quantia float32
        fmt.Print("Digite a quantia para depositar: ")
        fmt.Scan(&quantia)

        if quantia < 0 {
          quantia = 0
        }

        conta.Depositar(quantia)
        fmt.Printf("\nVocê depositou R$%.2f\n\n", quantia)

      case 3:
        var quantia float32
        fmt.Print("Digite a quantia que deseja sacar: ")
        fmt.Scan(&quantia)

        if quantia < 0 {
          quantia = 0
        }

        sacado := conta.Sacar(quantia)
        if sacado > -1 {
          fmt.Printf("\nVocê sacou: R$%.2f\n\n", quantia)
        }

      case 4:
        fmt.Print("\nSaindo...\n\n")
        sair = true

      default:
        fmt.Println("Comando invalido.")
        continue
    }

    if sair {
      break
    }
  }
}

func exibeComandos(comandos []string) {
  for _, comando := range(comandos) {
    fmt.Println(comando)
  }

  fmt.Print("Selecione um comando: ")
}

func entraConta(contas []Conta) (bool, Conta) {
  var (
    titular string
    senha string
    conta Conta
    achou bool = false
    login bool = false
  )

  fmt.Println("---- Entrando na conta ----")
  fmt.Print("Digite o nome do titular: ")
  fmt.Scan(&titular)

  for _, arr_conta := range contas {
    if arr_conta.Titular == titular {
      conta = arr_conta
      achou = true
      break;
    }
  }

  if achou == false {
    fmt.Print("\nConta não encontada nesse nome\n\n")
    return false, Conta{};
  }
  
  for i := 0; i < 3; i++ {
    fmt.Print("Digite a senha da conta: ")
    fmt.Scan(&senha)
    fmt.Println()

    if conta.Senha != senha {
      fmt.Print("Senha invalida\n\n")
    } else {
      fmt.Print("Você logou na conta.\n\n")
      login = true
      break
    }
  }
  
  if login == false {
    fmt.Print("Limite de tentativas excedida (3)\n\n")
  }

  return login, conta
} 

func criaConta() Conta {
  var nome string
  var senha string

  fmt.Println("---- Criando conta ----")
  fmt.Print("Digite seu nome: ")
  fmt.Scan(&nome)
  fmt.Print("Digite uma senha: ")
  fmt.Scan(&senha)
  fmt.Println()

  return Conta{Titular: nome, Agencia: 123456, Conta: 123, Saldo: 0, Senha: senha}
}
