package main

import (
	"testing"
)

func Test_calcBasis1(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5, 6, 7}
	xbo := NewXorBasisOperator(31, 100)
	for _, v := range testData {
		xbo.Add(v)
	}
	basis := xbo.Basis()
	if basis[0] != 1 {
		t.FailNow()
	}
	if basis[1] != 2 {
		t.FailNow()
	}
	if basis[2] != 4 {
		t.FailNow()
	}
}

func Test_calcBasis2(t *testing.T) {
	testData := []int{7, 6, 5, 4, 3, 2, 1}
	xbo := NewXorBasisOperator(31, 100)
	for _, v := range testData {
		xbo.Add(v)
	}
	basis := xbo.Basis()
	if basis[0] != 1 {
		t.FailNow()
	}
	if basis[1] != 2 {
		t.FailNow()
	}
	if basis[2] != 7 {
		t.FailNow()
	}
}

func Test_checkValue1(t *testing.T) {
	testData := []int{1, 2}
	xbo := NewXorBasisOperator(31, 100)
	for _, v := range testData {
		xbo.Add(v)
	}
	basis := xbo.Basis()
	if basis.Check(1) != true {
		t.FailNow()
	}
	if basis.Check(2) != true {
		t.FailNow()
	}
	if basis.Check(3) != true {
		t.FailNow()
	}
	if basis.Check(4) != false {
		t.FailNow()
	}
}

func Test_example5(t *testing.T) {
	a := []int{1, 2, 0, 3}
	b := []int{3, 3, 0, 3}
	result := solve(a, b, 4)
	if result {
		t.FailNow()
	}
}

func Test_example6(t *testing.T) {
	a := []int{3, 5, 6, 9, 10, 12}
	b := []int{6, 5, 3, 12, 15, 9}
	result := solve(a, b, 6)
	if !result {
		t.FailNow()
	}
}
