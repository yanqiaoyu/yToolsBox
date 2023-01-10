import requests
import json
from utils.RSAUtils import RsaUtil


# 访问DSC的句柄
class VisitDSCHandler():
    def __init__(self, dsc_ip, dsc_account, dsc_password) -> None:
        # 创建会话对象  方便访问不同页面时能保持身份
        self.r = requests.Session()
        # 传进来哪个IP, 就去哪个IP登录
        self.dsc_ip = dsc_ip
        self.dsc_password = dsc_password
        self.dsc_account = dsc_account

        # 进行登录
        self.login()

    # 登录
    def login(self):
        # 1.拿取验证码
        self.getRecaptcha()

        # 2.针对密码按照规格进行加密
        result = RsaUtil().encrypt_by_public_key(
            bytes(self.dsc_password, encoding="utf8"))
        password = bytes.decode(result)

        # 3.执行登录的请求
        payload = {'code': self.getRecaptchaResult(),
                   'username': self.dsc_account, 'password': password}
        result = self.r.post(f"https://{self.dsc_ip}/dashboard/auth/login",
                             json=payload, verify=False)
        # print("登录后的Cookie为: ", result.cookies)
        # print("登录后的hctd为: ", result.headers['hctd'])
        # self.logger.info("登录后的Cookie为: "+result.cookies)
        # self.logger.info("登录后的hctd为: "+result.headers['hctd'])

        # 4.在请求的句柄中,添加登录后的cookie与hctd
        # self.logger.info(result.headers)
        self.r.cookies = result.cookies
        self.r.headers['hctd'] = result.headers['hctd']

    # 获取验证码
    def getRecaptcha(self):
        result = self.r.get(
            f'https://{self.dsc_ip}/dashboard/auth/captcha', verify=False)
        result.encoding = 'utf-8'
        with open('./1.png', 'wb') as f:
            f.write(result.content)
            f.close()

        self.r.cookies = result.cookies
        # self.logger.info("登录前的cookie为")
        # self.logger.info(self.r.cookies)

    # 调用识别验证码的服务
    def getRecaptchaResult(self):
        payload = {'product_name': 'ads'}
        files = [
            ('picture', ('1.png', open('./1.png', 'rb'), 'image/png'))
        ]
        response = requests.request(
            "POST", 'http://recognize_service:3579/api/v2/recognize', data=payload, files=files)
        # "POST", 'http://124.221.120.2:3579/api/v2/recognize', data=payload, files=files)
        result = json.loads(response.text)
        return result['data']['result']


if __name__ == "__main__":
    test = VisitDSCHandler()
