while [ 1 -lt 2 ]
do
	http POST :9090/api/project/16/event event_id:=10
	sleep 1
done
echo
exit 0