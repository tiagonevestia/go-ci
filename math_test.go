package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 30)
	resultado := 45

	if total != resultado {
		t.Errorf("Resultado da soma é inválido: Resultado: %d. Esperado: %d.", total, 45)
	}
}
