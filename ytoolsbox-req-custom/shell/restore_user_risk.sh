#!/bin/bash

date=$(date +%Y-%m-%d_%H:%M)

mkdir -p /root/restore_sh_log/$date/

#has_backup_table=$(sudo docker exec -i $(docker ps | grep dsc_db | awk '{print $1}') su postgres -c "psql -d dsc -t -c \"SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'risk');\"")
has_backup_table=$(sudo docker exec -i $(docker ps | grep k8s_postgres_postgres | awk '{print $1}') psql -h 0.0.0.0 -p 5432 -U admin -d dsc -c "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'risk');")

if echo "$has_backup_table" | awk 'NR==3{print $1}' | grep -q "f"; then
    echo "无备份表public.risk, 不能进行风险阈值配置还原"
    echo 2
    exit
fi

echo "有备份表public.risk,开始还原"
#sudo docker exec -i $(docker ps | grep dsc_db | awk '{print $1}') sudo -u postgres psql -h localhost -p 5432 -d dsc -c "UPDATE hecate.risk SET model_type = u.model_type, model_cfg = u.model_cfg, filters = u.filters, excludes = u.excludes FROM public.risk u WHERE hecate.risk.id = u.id;"
sudo docker exec -i $(docker ps | grep k8s_postgres_postgres | awk '{print $1}') psql -h 0.0.0.0 -p 5432 -U admin -d dsc -c "UPDATE hecate.risk SET model_type = u.model_type, model_cfg = u.model_cfg, filters = u.filters, excludes = u.excludes FROM public.risk u WHERE hecate.risk.id = u.id;"

echo "判断备份表public.risk的数据和还原后的表hecate.risk的数据一致"
#res=$(sudo docker exec -i $(docker ps | grep dsc_db | awk '{print $1}') psql -U postgres -d dsc -c "SELECT * FROM hecate.risk EXCEPT SELECT * FROM public.risk;")
res=$(sudo docker exec -i $(docker ps | grep k8s_postgres_postgres | awk '{print $1}') psql -h 0.0.0.0 -p 5432 -U admin -d dsc -c "SELECT * FROM hecate.risk EXCEPT SELECT * FROM public.risk;")

if echo "$res" | awk 'NR==3{print $0}' | grep -q "(0 rows)"; then
    echo "校验还原后的数据和备份表中的数据一致,还原成功"
    #sudo docker exec -i $(docker ps | grep dsc_db | awk "{print \$1}") psql -h localhost -U postgres -d dsc -v ON_ERROR_STOP=1 -v "PGCLIENTENCODING=utf8" -c "DROP TABLE IF EXISTS public.risk;"
    sudo docker exec -i $(docker ps | grep k8s_postgres_postgres | awk "{print \$1}") psql -h 0.0.0.0 -p 5432 -U admin -d dsc -v ON_ERROR_STOP=1 -v "PGCLIENTENCODING=utf8" -c "DROP TABLE IF EXISTS public.risk;"
    echo 0
    exit
fi

echo "校验还原后的数据和备份表中的数据不一致,还原失败"
echo 1
