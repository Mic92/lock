set datafile separator ","
set terminal png size 1024,768 enhanced font "sans"

set xlabel "Number of Goroutines"
set ylabel "Number of Overtakens"
set key left top

set output 'measure.png'
set title "Filter Lock Algorithm\n1e7 lock/unlock loops per Goroutine\nIntel i7-4770 CPU @ 3.40GHz"
plot 'measure.data' using 2:($1==1?$3:1/0) title "1 OS thread", \
     'measure.data' using 2:($1==2?$3:1/0) title "2 OS threads", \
     'measure.data' using 2:($1==3?$3:1/0) title "3 OS threads", \
     'measure.data' using 2:($1==4?$3:1/0) title "4 OS threads", \
     'measure.data' using 2:($1==5?$3:1/0) title "5 OS threads", \
     'measure.data' using 2:($1==6?$3:1/0) title "6 OS threads", \
     'measure.data' using 2:($1==7?$3:1/0) title "7 OS threads", \
     'measure.data' using 2:($1==8?$3:1/0) title "8 OS threads"
