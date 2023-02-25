#!/bin/bash

: << EOF
API 性能测试脚本，会自动执行 wrk 命令，采集数据、分析数据并调用 gnuplot 画图

使用方式 ( 测试 API 性能)：
1. 启动http服务
2. 执行测试脚本: ./wrktest.sh

脚本会生成 .dat 的数据文件，每列含义为：并发数 QPS 平均响应时间 成功率

使用方式 (对比2次测试结果)
1. 性能测试：./wrktest.sh http://127.0.0.1:9090/ping
2. 执行命令： ./wrktest.sh diff apiServer.dat http.dat

> Note: 需要确保系统安装了 wrk 和 gnuplot 工具
EOF

#根目录
rootDir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
wrkDir="${rootDir}/_output/wrk"
jobName="apiServer"
duration="30s"
COLOR_MAGENTA="\033[35m"
COLOR_NORMAL="\033[0m"

# 设置wrk选项
wrk::setup() {
  if [[ "$OSTYPE" == "linux-gnu"* ]]; then
      # Linux 系统
      threads=$((3 * `grep -c processor /proc/cpuinfo`))
  elif [[ "$OSTYPE" == "darwin"* ]]; then
      # macOS 系统
      threads=$((3 * `sysctl -n machdep.cpu.core_count`))
  else
      threads=3
  fi
  threads=3
  #concurrent="200 500 1000 3000 5000 10000 15000 20000 25000 50000 100000 200000 500000 1000000"
  concurrent="200 500 1000 3000 5000 10000 15000 20000 25000 50000"
  cmd="wrk -t${threads} -d${duration} -T30s --latency"
}

# 打印使用信息
wrk::usage()
{
  cat << EOF

Usage: $0 [OPTION] [diff] URL
Performance automation test script.

  URL                    HTTP request url, like: http://127.0.0.1:9900/ping
  diff                   Compare two performance test results

OPTIONS:
  -h                     Usage information
  -n                     Performance test task name, default: apiserver
  -d                     Directory used to store performance data and gnuplot graphic, default: _output/wrk

EOF
}

# 将绘图数据转换为可用数据
function wrk::convert_plot_data()
{
  echo "$1" | awk -v datfile="${wrkDir}/${datfile}" ' {
  if ($0 ~ "Running") {
    common_time=$2
  }
if ($0 ~ "connections") {
  connections=$4
  common_threads=$1
}
if ($0 ~ "Latency   ") {
  avg_latency=convertLatency($2)
}
if ($0 ~ "50%") {
  p50=convertLatency($2)
}
if ($0 ~ "75%") {
  p75=convertLatency($2)
}
if ($0 ~ "90%") {
  p90=convertLatency($2)
}
if ($0 ~ "99%") {
  p99=convertLatency($2)
}
if ($0 ~ "Requests/sec") {
  qps=$2
}
if ($0 ~ "requests in") {
  allrequest=$1
}
if ($0 ~ "Socket errors") {
  err=$4+$6+$8+$10
}
}
END {
rate=sprintf("%.2f", (allrequest-err)*100/allrequest)
print connections,qps,avg_latency,rate >> datfile
}

function convertLatency(s) {
  if (s ~ "us") {
    sub("us", "", s)
    return s/1000
  }
if (s ~ "ms") {
  sub("ms", "", s)
  return s
}
if (s ~ "s") {
  sub("s", "", s)
  return s * 1000
}
}'
}

# 删除现有数据文件
function wrk::prepare()
{
  rm -f ${wrkDir}/${datfile}
}

# 根据gunplot数据文件进行绘图
function wrk::plot() {
  gnuplot <<  EOF
set terminal png enhanced #输出格式为png文件
set ylabel 'QPS'
set xlabel 'Concurrent'
set y2label 'Average Latency (ms)'
set key top left vertical noreverse spacing 1.2 box
set tics out nomirror
set border 3 front
set style line 1 linecolor rgb '#00ff00' linewidth 2 linetype 3 pointtype 2
set style line 2 linecolor rgb '#ff0000' linewidth 1 linetype 3 pointtype 2
set style data linespoints

set grid #显示网格
set xtics nomirror rotate #by 90#只需要一个x轴
set mxtics 5
set mytics 5 #可以增加分刻度
set ytics nomirror
set y2tics

set autoscale  y
set autoscale y2

set output "${wrkDir}/${qpsttlb}"  #指定数据文件名称
set title "QPS & TTLB\nRunning: ${duration}\nThreads: ${threads}"
plot "${wrkDir}/${datfile}" using 2:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE0000" axis x1y1 t "QPS","${wrkDir}/${datfile}" using 3:xticlabels(1) w lp pt 5 ps 1 lc rgbcolor "#0000CD" axis x2y2 t "Avg Latency (ms)"

unset y2tics
unset y2label
set ytics nomirror
set yrange[0:100]
set output "${wrkDir}/${successrate}"  #指定数据文件名称
set title "Success Rate\nRunning: ${duration}\nThreads: ${threads}"
plot "${wrkDir}/${datfile}" using 4:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#F62817" t "Success Rate"
EOF
}

# 绘图差异图形
function wrk::plot_diff()
{
  gnuplot <<  EOF
set terminal png enhanced #输出格式为png文件
set xlabel 'Concurrent'
set ylabel 'QPS'
set y2label 'Average Latency (ms)'
set key below left vertical noreverse spacing 1.2 box autotitle columnheader
set tics out nomirror
set border 3 front
set style line 1 linecolor rgb '#00ff00' linewidth 2 linetype 3 pointtype 2
set style line 2 linecolor rgb '#ff0000' linewidth 1 linetype 3 pointtype 2
set style data linespoints

#set border 3 lt 3 lw 2   #这会让你的坐标图的border更好看
set grid #显示网格
set xtics nomirror rotate #by 90#只需要一个x轴
set mxtics 5
set mytics 5 #可以增加分刻度
set ytics nomirror
set y2tics

#set pointsize 0.4 #点的像素大小
#set datafile separator '\t' #数据文件的字段用\t分开

set autoscale  y
set autoscale y2

#设置图像的大小 为标准大小的2倍
#set size 2.3,2

set output "${wrkDir}/${t1}_${t2}.qps.ttlb.diff.png"  #指定数据文件名称
set title "QPS & TTLB\nRunning: ${duration}\nThreads: ${threads}"
plot "/tmp/plot_diff.dat" using 2:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE0000" axis x1y1 t "${t1} QPS","/tmp/plot_diff.dat" using 5:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE82EE" axis x1y1 t "${t2} QPS","/tmp/plot_diff.dat" using 3:xticlabels(1) w lp pt 5 ps 1 lc rgbcolor "#0000CD" axis x2y2 t "${t1} Avg Latency (ms)", "/tmp/plot_diff.dat" using 6:xticlabels(1) w lp pt 5 ps 1 lc rgbcolor "#6495ED" axis x2y2 t "${t2} Avg Latency (ms)"

unset y2tics
unset y2label
set ytics nomirror
set yrange[0:100]
set output "${wrkDir}/${t1}_${t2}.successrate.diff.png"  #指定数据文件名称
set title "Success Rate\nRunning: ${duration}\nThreads: ${threads}"
plot "/tmp/plot_diff.dat" using 4:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE0000" t "${t1} Success Rate","/tmp/plot_diff.dat" using 7:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE82EE" t "${t2} Success Rate"
EOF
}

# 启动API性能测试
wrk::start_performance_test() {
  wrk::prepare

  for c in ${concurrent}
  do
    wrkcmd="${cmd} -c ${c} $1"
    echo "Running wrk command: ${wrkcmd}"
    result=`eval ${wrkcmd}`
    wrk::convert_plot_data "${result}"
  done

  echo -e "\nNow plot according to ${COLOR_MAGENTA}${wrkDir}/${datfile}${COLOR_NORMAL}"
  wrk::plot &> /dev/null
  echo -e "QPS graphic file is: ${COLOR_MAGENTA}${wrkDir}/${qpsttlb}${COLOR_NORMAL}
  Success rate graphic file is: ${COLOR_MAGENTA}${wrkDir}/${successrate}${COLOR_NORMAL}"
}

while getopts "hd:n:" opt;do
  case ${opt} in
    d)
      wrkDir=${OPTARG}
      ;;
    n)
      jobName=${OPTARG}
      ;;
    ?)
      wrk::usage
      exit 0
      ;;
  esac
done

shift $(($OPTIND-1))

mkdir -p ${wrkDir}
case $1 in
  "diff")
    if [ "$#" -lt 3 ];then
      wrk::usage
      exit 0
    fi

    t1=$(basename $2|sed 's/.dat//g') # 对比图中红色线条名称
    t2=$(basename $3|sed 's/.dat//g') # 对比图中粉色线条名称

    join $2 $3 > /tmp/plot_diff.dat
    wrk::plot_diff `basename $2` `basename $3`
    exit 0
    ;;
  *)
    if [ "$#" -lt 1 ];then
      wrk::usage
      exit 0
    fi
    url="$1"

    qpsttlb="${jobName}_qps_ttlb.png"
    successrate="${jobName}_successrate.png"
    datfile="${jobName}.dat"

    wrk::setup
    wrk::start_performance_test "${url}"
    ;;
esac