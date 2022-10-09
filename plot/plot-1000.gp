set title "Redis Lists vs. Sorted Sets benchmark(offset 1000)"
set term png enhanced font 'Verdana,10'
set output 'benchmark-1000.png'
set xlabel "iteration"
set ylabel "execution time(nanoseconds)"
set autoscale
set grid

plot './benchmark/list_1000' using 1:2 with linespoints title 'lists', \
'./benchmark/set_1000' using 1:2 with linespoints title 'sorted sets'