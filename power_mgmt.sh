
for i in /sys/devices/system/cpu/cpu[0-3]
do
    echo performance > $i/cpufreq/scaling_governor
done

