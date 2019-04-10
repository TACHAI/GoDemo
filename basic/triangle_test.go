package main

import (
	"math"
	"testing"
)

func calcTriangle (a,b int)  int  {
	c := math.Sqrt(float64(a*a + b*b))
	return int(c)
}

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
	}
		for _, tt := range tests{
		if actual := calcTriangle(tt.a, tt.b); actual !=tt.c{
		t.Errorf("calcTriangle")

	}
	}
}

// 性能测试
func BenchmarkSubstr(b *testing.B)  {
	s := "黑化肥发挥会发挥"
	ans:=8

	for i:=0;i<b.N;i++{
		actual := lengthOfNonReatingSubStr(s)
		if actual!=ans{
			b.Errorf("got %d for input %s"+"expected %d",actual,s,ans)
		}

	}
}