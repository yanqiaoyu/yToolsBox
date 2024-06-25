#!/bin/bash

date=$(date +%Y-%m-%d_%H:%M)

mkdir -p /root/backup_sh_log/$date/

# 判断是否有备份表
#has_backup_table=$(sudo docker exec -i $(docker ps | grep dsc_db | awk '{print $1}') su postgres -c "psql -d dsc -t -c \"SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'risk');\"")
has_backup_table=$(sudo docker exec -i $(docker ps | grep k8s_postgres_postgres | awk '{print $1}') psql -h 0.0.0.0 -p 5432 -U admin -d dsc -c "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'risk');")

# 有 则不做备份 直接进行POC测试
if echo "$has_backup_table" | awk 'NR==3{print $1}' | grep -q "t"; then
    echo "有备份表public.risk"
    echo 0
    exit
fi

echo "无备份表public.risk,开始做备份"
# 若无则 做备份
#sudo docker exec -i "$(docker ps | grep dsc_db | awk '{print $1}')" sudo -u postgres psql -h localhost -p 5432 -d dsc -c "CREATE TABLE public.risk (LIKE hecate.risk INCLUDING ALL); INSERT INTO public.risk SELECT * FROM hecate.risk;"
sudo docker exec -i "$(docker ps | grep k8s_postgres_postgres | awk '{print $1}')" psql -h 0.0.0.0 -p 5432 -U admin -d dsc -c  "CREATE TABLE public.risk (LIKE hecate.risk INCLUDING ALL); INSERT INTO public.risk SELECT * FROM hecate.risk;"

# 判断是否备份成功
#re_has_backup_table=$(sudo docker exec -i $(docker ps | grep dsc_db | awk '{print $1}') su postgres -c "psql -d dsc -t -c \"SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'risk');\"")
re_has_backup_table=$(sudo docker exec -i $(docker ps | grep k8s_postgres_postgres | awk '{print $1}') psql -h 0.0.0.0 -p 5432 -U admin -d dsc -c  "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'risk');")

if echo "$re_has_backup_table" | awk 'NR==3{print $1}' | grep -q "f"; then
    echo "备份表hecate.risk失败"
    echo 3
    exit
fi

echo "完成备份表public.risk,开始做数值校验"
# 数据值校验
#res=$(sudo docker exec -i $(docker ps | grep dsc_db | awk '{print $1}') psql -U postgres -d dsc -c "SELECT * FROM hecate.risk EXCEPT SELECT * FROM public.risk;")
res=$(sudo docker exec -i $(docker ps | grep k8s_postgres_postgres | awk '{print $1}') psql -h 0.0.0.0 -p 5432 -U admin -d dsc -c "SELECT * FROM hecate.risk EXCEPT SELECT * FROM public.risk;")

# 校验数据一致，返回0
if echo "$res" | awk 'NR==3{print $0}' | grep -q "(0 rows)"; then  
    echo "备份表public.risk的数据和原表hecate.risk的数据一致"
    echo 2 
    exit
fi

echo "备份表publics.risk的数据和原表hecate.risk的数据不一致"
# 校验数据存在不一致的情况，返回1
echo 1 