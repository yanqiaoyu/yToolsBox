import logging
import sys


class Log():
    """
    创建 log 类 供其他类进行调用，只需要对类进行初始化一次 其他类进行调用即可
    """

    def __init__(self, log_level=logging.DEBUG):
        """
        设置log初始化，设置log的输入日志file 和终端输出
        还有按照文件大小进行日志分割
        按照时间进行日志分割
        """
        """设置将log输出到文件log和终端中"""
        """参数when决定了时间间隔的类型，参数interval决定了多少的时间间隔。如when=‘D’，
        interval=2，就是指两天的时间间隔，backupCount决定了能留几个日志文件。超过数量就会丢弃掉老的日志文件"""
        self.logger = logging.getLogger()

        self.stream_handler = logging.StreamHandler(sys.stdout)

        """设置 log的输出格式 日期+ 日志级别 + 具体日志信息"""
        self.stream_handler.setFormatter(logging.Formatter(
            "%(asctime)s - %(levelname)s - %(message)s"))

        """设置log的输出等级,默认不传等级的级别为DEBUG级别"""
        self.logger.setLevel(log_level)

        """将handler 添加到logger实例对象中"""
        self.logger.addHandler(self.stream_handler)

    def debug(self, message):
        """
        定义log的debug函数，将message传递进来
        :param message:
        :return:
        """
        self.logger.debug(message)

    def info(self, message):
        """
        定义log的info函数，将message传递进来
        :param message:
        :return:
        """
        self.logger.info(message)

    def warning(self, message):
        """
        定义log的warning函数，将message传递进来
        :param message:
        :return:
        """
        self.logger.warning(message)

    def error(self, message):
        """
        定义log的error函数，将message传递进来
        :param message:
        :return:
        """
        self.logger.error(message)

    def critical(self, message):
        """
        定义log的critical函数，将message传递进来
        :param message:
        :return:
        """
        self.logger.critical(message)


LoggingHandler = Log()
