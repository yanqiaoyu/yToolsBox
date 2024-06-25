import paramiko
import os

from proto import req_pb2, req_pb2_grpc
from utils.LoggingUtils import LoggingHandler
from service.ModifyRiskPolicyService import ModifyRiskPolicy

class DSCService(req_pb2_grpc.DSCServiceServicer):

    def ModifyRiskThreshold(self, request, context):
        LoggingHandler.info("修改阈值的大脑IP: "+request.dsc_ip)
        LoggingHandler.info("修改阈值的大脑前端账号: "+request.dsc_fe_account)
        LoggingHandler.info("修改阈值的大脑前端密码:"+request.dsc_fe_password)
        LoggingHandler.info("修改阈值的大脑后端账号: "+request.dsc_back_account)
        LoggingHandler.info("修改阈值的大脑前端密码:"+request.dsc_back_password)

        # ssh连接大脑后端
        LoggingHandler.info("建立ssh连接: hostname=%s, username=%s, password=%s" % (request.dsc_ip, request.dsc_back_account, request.dsc_back_password))  
        ssh = paramiko.SSHClient()
        ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
        ssh.connect(hostname=request.dsc_ip, username=request.dsc_back_account, password=request.dsc_back_password)

        # 一键调整风险阈值  
        # 走大脑后端 进行风险数据备份；
        #   1. 建立POC和大脑后端的ssh连接
        #   2. 将POC的req容器内的/ytoolsbox-req-custom/shell下的两个sh文件传到大脑后端的tmp/shell/下（两个shell:备份和还原用户自定义风险阈值配置信息） 
        # 走前端，对风险策略阈值进行更改，从而进行POC风险触发测试
        #   1. 登陆大脑前端
        #   2. 通过前端接口对风险阈值进行更改
        if request.modify_mode == "poc":
            # 大脑后端新建目录/tmp/shell 用于存放备份和还原操作的shell脚本
            LoggingHandler.info("新建目录/tmp/shell 用于存放备份和还原操作的shell脚本")
            stdin, stdout, stderr = ssh.exec_command('mkdir -p /tmp/shell')
            create_shell_res = stdout.read().decode()
            error = stderr.read()
            
            if create_shell_res == "":
                LoggingHandler.info("/tmp/shell目录创建成功或已存在")  
                sftp = ssh.open_sftp()
                local_path = "/ytoolsbox-req-custom/shell"
                remote_path = "/tmp/shell/"
        
                # 遍历本地文件夹/ytoolsbox-req-custom/shell，将所有文件上传到大脑服务器/tmp/shell/目录下
                for root, dirs, files in os.walk(local_path):
                    for file in files:
                        local_file_path = os.path.join(root, file)
                        remote_file_path = os.path.join(remote_path, file)
                        sftp.put(local_file_path, remote_file_path)
                        local_file_size = os.path.getsize(local_file_path)
                        remote_file_size = sftp.stat(remote_file_path).st_size
                        if local_file_size != remote_file_size:
                            LoggingHandler.info("传输shell脚本失败")
                            return req_pb2.ModifyRiskThresholdResponse(Data="传输shell脚本失败", Meta="400")
            
                sftp.close()
                LoggingHandler.info("传输shell脚本成功")
            else:
                return req_pb2.ModifyRiskThresholdResponse(Data="/tmp/shell目录创建失败", Meta="400")

            LoggingHandler.info("备份用户原风险配置信息")   
            stdin, stdout, stderr = ssh.exec_command('sudo sh /tmp/shell/backup_user_risk.sh')
            backup_user_risk_res = ''.join(stdout.readlines())

            # 若备份完成或有备份，则才能调整阈值
            if "备份表public.risk的数据和原表hecate.risk的数据一致" in backup_user_risk_res:
                LoggingHandler.info("登录大脑前端")
                # 登录大脑前端
                modifyRiskThresholdHandler = ModifyRiskPolicy(
                    dsc_ip=request.dsc_ip,
                    dsc_fe_account=request.dsc_fe_account,
                    dsc_fe_password=request.dsc_fe_password
                )
                LoggingHandler.info("开始调整阈值")
                modifyRiskThresholdHandler.ModifyThreshold()
                return req_pb2.ModifyRiskThresholdResponse(Data="备份用户原风险配置信息成功", Meta="200")
            
            if "备份表hecate.risk失败" in backup_user_risk_res:
                return req_pb2.ModifyRiskThresholdResponse(Data="备份用户原风险配置信息失败", Meta="400")      

            if "有备份表public.risk" in backup_user_risk_res:
                return req_pb2.ModifyRiskThresholdResponse(Data="备份表已存在", Meta="200")

            # 备份存在，但数据不一致，则备份用户原风险配置信息失败
            LoggingHandler.info("备份存在，但数据不一致，备份用户原风险配置信息失败")
            return req_pb2.ModifyRiskThresholdResponse(Data="备份存在，但数据不一致，备份用户原风险配置信息失败", Meta="400")
        # 一键还原风险阈值配置
        elif request.modify_mode == "default":
            LoggingHandler.info("判断用于还原用户风险阈值的shell脚本是否存在")
            stdin, stdout, stderr = ssh.exec_command('[ -f /tmp/shell/restore_user_risk.sh ] && echo "文件存在" || echo "文件不存在"')
            result = stdout.read().decode().strip()
            if result != "文件存在":
                return req_pb2.ModifyRiskThresholdResponse(Data="用于还原用户风险阈值的shell脚本不存在", Meta="400")

            LoggingHandler.info("开始还原阈值")
            stdin, stdout, stderr = ssh.exec_command('sudo sh /tmp/shell/restore_user_risk.sh')
            restore_user_risk_res = ''.join(stdout.readlines())
            # modifyRiskThresholdHandler.ResetTreshold()
            if "校验还原后的数据和备份表中的数据一致,还原成功" in restore_user_risk_res:
                LoggingHandler.info("校验还原后的数据和备份表中的数据一致,还原成功,删除备份表")
                return req_pb2.ModifyRiskThresholdResponse(Data="还原用户原风险配置信息成功", Meta="200")
            elif "校验还原后的数据和备份表中的数据不一致,还原失败" in restore_user_risk_res:
                return req_pb2.ModifyRiskThresholdResponse(Data="校验还原后的数据和备份表中的数据不一致,还原失败", Meta="400")
            elif "无备份表public.risk, 不能进行风险阈值配置还原" in restore_user_risk_res:
                return req_pb2.ModifyRiskThresholdResponse(Data="无备份表,不能进行风险阈值配置还原,还原失败", Meta="400")