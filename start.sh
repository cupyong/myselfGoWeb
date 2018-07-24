#if [ $# -ne 1 ];then
#    echo "Paras Error!"
#    echo "Usage: bash ${0} configPath"
#    exit 1
#fi
if [ ! -f datanote2 ];then
#make clean
make
fi

#echo ${1}
nohup ./datanote2 conf/conf.ini 2>&1 &
