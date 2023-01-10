from proto import req_pb2, req_pb2_grpc
from utils.LoggingUtils import LoggingHandler
from service.ModifyRiskPolicyService import ModifyRiskPolicy


class DSCService(req_pb2_grpc.DSCServiceServicer):

    def ModifyRiskThreshold(self, request, context):
        LoggingHandler.info("修改阈值的大脑IP: "+request.dsc_ip)
        LoggingHandler.info("修改阈值的大脑前端账号: "+request.dsc_fe_account)
        LoggingHandler.info("修改阈值的大脑前端密码:"+request.dsc_fe_password)

        modifyRiskThresholdHandler = ModifyRiskPolicy(
            dsc_ip=request.dsc_ip,
            dsc_fe_account=request.dsc_fe_account,
            dsc_fe_password=request.dsc_fe_password
        )

        if request.modify_mode == "poc":
            LoggingHandler.info("开始调整阈值")
            modifyRiskThresholdHandler.ModifyThreshold()
        elif request.modify_mode == "default":
            LoggingHandler.info("开始还原阈值")
            modifyRiskThresholdHandler.ResetTreshold()

            # 调整阈值
        return req_pb2.ModifyRiskThresholdResponse(Data="test", Meta="200")
