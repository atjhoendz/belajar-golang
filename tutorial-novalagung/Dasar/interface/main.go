package main

import (
	"fmt"
	"math"
)

/*
 * Interface => kumpulan definisi method yang tidak memiliki isi dengan nama tertentu
 *
 */
type hitung interface {
	luas() float64
	keliling() float64
}

// struct lingkaran
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jarijari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// struct persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

/*
 * Embedded interface => interface yang di embed ke interface lain
 */
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitungAll interface {
	hitung2d
	hitung3d
}

// struct kubus
type kubus struct {
	sisi float64
}

func (k kubus) luas() float64 {
	return math.Pow(k.sisi, 3)
}

func (k kubus) keliling() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k kubus) volume() float64 {
	return k.sisi * 12
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10}
	fmt.Println("Persegi")
	fmt.Println("Luas:", bangunDatar.luas())
	fmt.Println("Keliling:", bangunDatar.keliling())

	bangunDatar = lingkaran{14}
	fmt.Println("\nLingkaran")
	fmt.Println("Luas:", bangunDatar.luas())
	fmt.Println("Keliling:", bangunDatar.keliling())
	fmt.Println("Jari-jari:", bangunDatar.(lingkaran).jarijari())

	// implementasi embeded interface
	var bangunRuang hitungAll = kubus{10}

	fmt.Println("\nKubus")
	fmt.Println("Luas:", bangunRuang.luas())
	fmt.Println("Keliling:", bangunRuang.keliling())
	fmt.Println("Volume:", bangunRuang.volume())
}
