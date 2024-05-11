package main

import (
	"fmt"
	"go-estudo-orientacao-a-objeto-BANCO/contas"
)

func main() {

	contaDoAgamin := contas.ContaCorrente{
		Titular:       "AGAMIN",
		NumeroAgencia: 777,
		NumeroConta:   123456,
		Saldo:         1000.0}

	contaDaLaura := contas.ContaCorrente{Titular: "Laura", NumeroAgencia: 074, NumeroConta: 747474, Saldo: 1100.0}

	fmt.Println("Saldo Laura: ", contaDaLaura, "Saldo AGAMIN:", contaDoAgamin)

	status := contaDaLaura.Transferencia(-300, &contaDoAgamin)

	fmt.Println(status)

	fmt.Println("Saldo Laura: ", contaDaLaura, "Saldo AGAMIN:", contaDoAgamin)
}

// var contaDaNeide *ContaCorrente
// contaDaNeide = new(ContaCorrente)
// fmt.Println(*contaDaNeide)
