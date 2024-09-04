package main

import "fmt"

type Veiculo interface {
	Velocidademaxima() int
	MeioDeLocomocao() string
	ModeloVeiculo() string
}

type Carro struct {
	MaxVelo int
	Modelo  string
}
type Tren struct {
	MaxVelo int
	Modelo  string
}

// ================================================
func (c Carro) Velocidademaxima() int {
	return c.MaxVelo
}
func (c Carro) MeioDeLocomocao() string {
	return "Rodas"
}
func (c Carro) ModeloVeiculo() string {
	return c.Modelo
}

//=========================================================

func (t Tren) Velocidademaxima() int {
	return t.MaxVelo
}
func (t Tren) MeioDeLocomocao() string {
	return "trilhos"
}
func (t Tren) ModeloVeiculo() string {
	return t.Modelo
}

//==================================================================

func DescreverVeiculo(v Veiculo) {
	fmt.Printf("Veiculo atual.......:%s \n", v.ModeloVeiculo())
	fmt.Printf("Meio de locomoção...:%s \n", v.MeioDeLocomocao())
	fmt.Printf("Velocidade Maxima...:%dKm/h \n \n", v.Velocidademaxima())
}

func main() {

	meuCarro := Carro{MaxVelo: 120, Modelo: "Fiat Uno"}
	meuTren := Tren{MaxVelo: 500, Modelo: "Trem Bala"}

	DescreverVeiculo(meuCarro)
	DescreverVeiculo(meuTren)

}
