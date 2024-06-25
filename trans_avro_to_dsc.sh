#!/bin/bash

current_date=$(date +%Y-%m-%d)

SENSITIVE_AVRO_PATH=/tmp/avro/raw/$current_date/sensitive_http/log_proc_0/

HTTP_AVRO_PATH=/tmp/avro/raw/$current_date/http/log_proc_0/

SENSITIVE_REMOTE_DIR=/dsc/data/meta-data/raw/$current_date/sensitive_http/log_proc_0/

HTTP_REMOTE_DIR=/dsc/data/meta-data/raw/$current_date/http/log_proc_0/

mkdir -p $SENSITIVE_AVRO_PATH
mkdir -p $HTTP_AVRO_PATH

# 浼犺緭鏂囦欢鍑芥暟
function transfer_file() {
    local file=$1
    local remote_dir=$2
    # 灏嗘枃浠朵紶杈撳埌杩滅▼涓绘満
    sshpass -p 123456 scp -r $file root@10.90.22.127:$remote_dir
    if [ $? -eq 0 ]; then
        echo "鏂囦欢浼犺緭鎴愬姛"
        rm -f $file
        sleep 60
    else
        echo "鏂囦欢浼犺緭澶辫触"
    fi
}

# 鐩戝惉鏂囦欢澶瑰彉鍖栧嚱鏁?function watch_folder() {
    local folder=$1
    local remote_dir=$2
    if [ -d "$folder" ]; then
        inotifywait -r -m -e create,modify,delete $folder | while read path action file; do
            files=($(find $folder -maxdepth 1 -type f -name "*.avro" | sed 's/.*\///' | sed 's/[^0-9]*//g' | sort -n))
            for f in "${files[@]}"; do
                if [ -f "$folder/$f.avro" ]; then
                    transfer_file $folder/$f.avro $remote_dir
                fi
            done
        done
    else
        echo "$folder 鐩綍涓嶅瓨鍦?
    fi
}


# 鍚姩鐩戝惉鍣?while true; do
    watch_folder $HTTP_AVRO_PATH $HTTP_REMOTE_DIR &
    watch_folder $SENSITIVE_AVRO_PATH $SENSITIVE_REMOTE_DIR
done
