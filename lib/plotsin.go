package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
)

func main(){
	gnum := 102

	if os.Args[1] == "plot" {
		plots(gnum)
	} else {
		mkgs(gnum)

		cmdstr := "convert -layers optimize -loop 0 p???.png ../res.gif"
		exec.Command("sh", "-c", cmdstr).Run()
		exec.Command("rm", "p???.png").Run()
	}
}

func mkgs(gnum int) {
	coef, _ := strconv.ParseFloat(os.Args[1], 64)

	for i := 1; i < gnum; i++ {
		mk1g(i, coef)
	}
}

func mk1g(num int, coef float64) {
	exec.Command("rm", "sin.plt").Run()
	exec.Command("rm", "p???.png").Run()
	cmdstr := "go run plotsin.go plot " + 
		os.Args[1] + 
		" " + strconv.Itoa(num) + " >> sin.plt"
	exec.Command("sh", "-c", cmdstr).Run()
	exec.Command("gnuplot", "plot.txt").Run()
	fname := "p" + zis3(num) + ".png"
	exec.Command("mv", "res.png", fname).Run()
}

func zis3(val int) string {
	if val < 10 {
		return "00" + strconv.Itoa(val)
	} else if val < 100 {
		return "0" + strconv.Itoa(val)
	} else {
		return strconv.Itoa(val)
	}
}

func plots(gnum int) {
	coef, _ := strconv.ParseFloat(os.Args[2], 64)
	pcnt, _ := strconv.Atoi(os.Args[3])

	std := sinp(gnum, gnum, 1.0)
	que := sinp(gnum, gnum, coef)

	sst := sinp(gnum, pcnt, 1.0)
	squ := sinp(gnum, pcnt, coef)

	dft := dft(pcnt, std, que)

	for i := 0; i < gnum - 1; i++ {
		if i == 0 {
			fmt.Printf("%d, %f, %f, %d, %f, %f, %f %d, 1, -1\n", 
				i, std[i], que[i], i, sst[i], squ[i], dft[i], pcnt - 1)
		} else if i < pcnt {
			fmt.Printf("%d, %f, %f, %d, %f, %f, %f\n", 
				i, std[i], que[i], i, sst[i], squ[i], dft[i])
		} else {
			fmt.Printf("%d, %f, %f\n", i, std[i], que[i])
		}
	}
}

func sinp(pnum, pcnt int, coef float64) []float64 {
	res := make([]float64, pcnt)

	for i := 0; i < pcnt; i++ {
		res[i] = math.Sin(2.0 * coef * math.Pi * float64(i) / float64(pnum))
	}

	return res
}

func rump(pnum, pcnt int, coef float64) []float64 {
	res := make([]float64, pcnt)

	s := 0.0

	for i := 0; i < pcnt; i++ {
		s += 2.0 * coef / float64(pnum)

		if 1.0 < s {
			s -= 2.0
		}

		res[i] = s
	}

	return res
}

func dft(pcnt int, std, que []float64) []float64 {
	res := make([]float64, pcnt)
	sum := 0.0

	for i := 0; i < pcnt; i++ {
		sum += std[i] * que[i]
		res[i] = sum / float64(i + 1)
	}

	return res
}
