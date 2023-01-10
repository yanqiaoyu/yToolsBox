import os
from utils.LoggingUtils import LoggingHandler
# 这里是先找到了PathUtils所在的位置,然后分割一下,根据相对位置找到main.py所在根目录位置
root_path = os.path.abspath(os.path.dirname(
    __file__)).split('ytoolsbox-req-custom')[0] + 'ytoolsbox-req-custom'


LoggingHandler.info("根路径是: "+root_path)
