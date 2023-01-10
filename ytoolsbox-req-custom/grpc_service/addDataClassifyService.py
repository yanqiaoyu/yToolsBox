from proto import addDataClassify_pb2, addDataClassify_pb2_grpc
from utils.LoggingUtils import LoggingHandler
from service.LoginService import VisitDSCHandler


class AddDataClassifyService(addDataClassify_pb2_grpc.AddDataClassifyServiceServicer):
    def AddDataClassify(self, request, context):
        LoggingHandler.info("大脑IP: "+request.dsc_ip)
        LoggingHandler.info("大脑前端账号: "+request.dsc_fe_account)
        LoggingHandler.info("大脑前端密码:"+request.dsc_fe_password)
        LoggingHandler.info("PostgreSQL的IP地址:"+request.postgreSQL_ip)

        r = VisitDSCHandler(
            dsc_ip=request.dsc_ip, dsc_account=request.dsc_fe_account, dsc_password=request.dsc_fe_password).r

        payload = {
            "name": "自动新增的演示数据库",
            "type": "PostgreSQL",
            "ip": request.postgreSQL_ip,
            "port": 5432,
            "group_id": 0,
            "add_task": True,
            "option": {
                "password": "N+YLUdIU18tQWtc/yIxnOQLHwMxcIAsj4Cexr2wCJk19vn/ZT2aWTQVs3fIc32BrS4mOZKWhF1FFwnJWofqtN35IZVK8ynmgEXjiF81DVSnJ4Hgi4m+1iSOxrwxLPM6FmeTEcwFL3HuCXPv7wIZxXVHMvQT8ICwFVqbqW8XGV2NUwJsN/+oSxx2ZWun9ZeQ23MpequCK9vnvdnWPSufw61QNhAfsmvebWIxZp+7e8cT1r5BgN6n5a0wz9SYGwu4mWuQXHy+mCdlU6Ido+QxYn2L9ejq+YmDQ6whxsxtPNvt3s2Gzn9pe6dPSCX2GEAzFCDlCvCFFEuXtDpyXV6ZH+g==",
                "username": "postgres"
            }
        }

        response = r.post(
            f"https://{request.dsc_ip}/dashboard/datasources", json=payload)

        print(response.text)

        return addDataClassify_pb2.AddDataClassifyResponse(Data="OK", Meta="200")
