flag=true

while $flag; do

  if (docker-compose -f ./docker-compose.yml logs | grep ready > /dev/null 2>&1) ; then
    echo "Mysql Ready"
    flag=false
  fi

done