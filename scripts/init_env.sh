#!/usr/bin/env bash
set -e
function yum_install() {
    yumcmd="yum install -y"
    for var in "$@"
    do
        if ! rpm --quiet --query $var; then
            yumcmd="$yumcmd $var"
        fi
    done
    echo $yumcmd
    if [ $yumcmd != "yum install -y" ]; then
        eval $yumcmd
    fi
}

yum_install epel-release git curl sshpass

yum_install python2-pip

requiredURL="https://raw.githubusercontent.com/pingcap/tidb-ansible/master/requirements.txt"

wget -O /tmp/requirements.txt  $requiredURL
pip install -r /tmp/requirements.txt
set +e

echo "Success!!"