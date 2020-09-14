packge goarea

import "math" 

// Pi é uma proporção numérica
const Pi = 3 

// Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2 ) * Pi
}
