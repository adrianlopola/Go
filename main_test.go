package main

import (
	"testing"
)

func TestEdadMayorOIgualA18(t *testing.T) {
	result := verificarEdadParaLicencia(18)
	if result != "Puedes sacar el carnet de conducir." {
		t.Errorf("Resultado inesperado: %s", result)
	}
}

func TestEdadMenorA18(t *testing.T) {
	result := verificarEdadParaLicencia(16)
	if result != "No puedes sacar el carnet de conducir." {
		t.Errorf("Resultado inesperado: %s", result)
	}
}

func TestEdadNegativa(t *testing.T) {
	// Diseñado para fallar
	result := verificarEdadParaLicencia(-1)
	if result != "Puedes sacar el carnet de conducir." {
		t.Errorf("Resultado inesperado: %s", result)
	}
}
