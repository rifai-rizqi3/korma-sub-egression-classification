package main

import "fmt"

// Data adalah struct untuk menyimpan data rumah
type Data struct {
	LuasRumah float64
	Harga     float64
}

func main() {
	// Buat dataset
	dataset := []Data{
		{40, 20000},
		{60, 30000},
		{80, 40000},
		{100, 50000},
		{120, 60000},
	}

	// Training model
	var sumX, sumY, sumXY, sumXSquare float64
	n := float64(len(dataset))

	for _, data := range dataset {
		sumX += data.LuasRumah
		sumY += data.Harga
		sumXY += data.LuasRumah * data.Harga
		sumXSquare += data.LuasRumah * data.LuasRumah
	}

	// Hitung koefisien regresi
	m := (n*sumXY - sumX*sumY) / (n*sumXSquare - sumX*sumX)
	c := (sumY - m*sumX) / n

	// Cetak hasil
	fmt.Printf("Persamaan regresi: Harga = %.2f * LuasRumah + %.2f\n", m, c)

	// Prediksi harga rumah
	luasRumah := 90.0
	hargaPrediksi := m*luasRumah + c
	fmt.Printf("Prediksi harga rumah dengan luas %.2f: %.2f\n", luasRumah, hargaPrediksi)
}
