PLOTDIR=./plot

all: main.go
	go build -o redis-benchmark

plot: ./benchmark
	gnuplot $(PLOTDIR)/plot-100.gp
	gnuplot $(PLOTDIR)/plot-1000.gp
	gnuplot $(PLOTDIR)/plot-10000.gp

.PHONY: plot