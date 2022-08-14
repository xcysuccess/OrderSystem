#! /bin/sh
#这里的改动是因为/usr/local/service/ordersystem/ordersystem会阻塞脚本继续运行，所以改为了后台运行，这样就能执行supervisord -c /etc/supervisord.d/supervisord.ini了
nohup /usr/local/service/ordersystem/ordersystem &
supervisord -c /etc/supervisord.d/supervisord.ini
tail -f /dev/null