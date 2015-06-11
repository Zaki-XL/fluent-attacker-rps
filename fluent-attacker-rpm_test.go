package main

import (
	"testing"
)

type randomTest struct {
	cnt int
}

var randomTests = []randomTest{
	randomTest{10},
	randomTest{100},
	randomTest{1000},
}

func Test_random(t *testing.T){
	for _,dt := range randomTests {
		v := len(random(dt.cnt))
		if v != dt.cnt {
			t.Errorf("random(%d) String Length = %d,want %d", dt.cnt,dt.cnt,dt.cnt )
		}
	}
}

type rpsTest struct {
	sec float64
	cnt int
	rps float64
}

var rpsTests = []rpsTest{
	rpsTest{10,100,10},
	rpsTest{2,10,5},
}

func Test_RPS(t *testing.T){
	for _,dt := range rpsTests {
		v := calcRPS(dt.sec,dt.cnt)
		if v != dt.rps {
			t.Errorf("calcRPS(%f,%d) = %f,want %f", dt.sec,dt.cnt,v,dt.rps )
		}
	}
}

type calcIntTest struct {
	cnt int
	ans int64
}

var calcIntTests = []calcIntTest{
	calcIntTest{1,1000000},
	calcIntTest{33,30303},
	calcIntTest{500,2000},
}

func Test_calcInt(t *testing.T){
	for _,dt := range calcIntTests {
		v := calcInt(dt.cnt)
		if v != dt.ans {
			t.Errorf("calcInt(%d) = %f,want %f", dt.cnt,v,dt.ans )
		}
	}
}

