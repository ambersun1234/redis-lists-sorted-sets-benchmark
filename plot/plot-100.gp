set title "Redis Lists vs. Sorted Sets benchmark(offset 100)"
set term png enhanced font 'Verdana,10'
set output 'benchmark-100.png'
set xlabel "iteration"
set ylabel "execution time(nanoseconds)"
set autoscale
set grid

plot './benchmark/list_100' using 1:2 with linespoints title 'lists', \
'./benchmark/set_100' using 1:2 with linespoints title 'sorted sets'