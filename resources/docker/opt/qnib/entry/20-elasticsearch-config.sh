#!/bin/bash

if [[ -z ${ES_MASTER_HOST} ]];then
  export ES_MASTER_HOST=${ES_MASTER_HOST:-qcollect-ng_backend}
fi
echo ">> sed -i'' -e \"s/ES_MASTER_HOST/${ES_MASTER_HOST}/\" /etc/qwatch-ng.yml"
sed -i'' -e "s/ES_MASTER_HOST/${ES_MASTER_HOST}/" /etc/qwatch-ng.yml
