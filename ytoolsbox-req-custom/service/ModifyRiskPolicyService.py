from operator import truth
from service.LoginService import VisitDSCHandler


class ModifyRiskPolicy():
    def __init__(self, dsc_ip, dsc_fe_account, dsc_fe_password) -> None:
        self.r = VisitDSCHandler(
            dsc_ip=dsc_ip, dsc_account=dsc_fe_account, dsc_password=dsc_fe_password).r
        self.dsc_ip = dsc_ip

    def _set_risk_payload(self, name, description, level, model_cfg, model_name, filters=""):
        '''
        设置修改风险策略的payload
        '''
        payload = {
            "name": name,
            "model_type": "normal",
            "description": description,
            "level": level,
            "model_cfg": model_cfg,
            "filters": {
                "choose_condition": "app",
                "care_app": "any",
                "care_api_group_ids": "any"
            },
            "model_name": model_name
        }

        if filters:
            payload["filters"].update(filters)

        return payload

    # 修改阈值
    def ModifyThreshold(self):
        # 1.账号多地访问
        payload = self._set_risk_payload(
            name="账号多地访问",
            description="同一个账号在多个不同的地理位置上登录，可能存在风险",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "dup_location",
                        "filter_threshold": 1,
                        "time_range": "15min",
                        "alarm_type": "access_source"
                    }
                ]
            },
            model_name="account_login_in_multi_area"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/1", json=payload)

        # 2.账号多IP访问
        payload = self._set_risk_payload(
            name="账号多IP访问",
            description="同一个账号在多个不同的IP上登录，可能存在风险",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "dup_ip",
                        "filter_threshold": 1,
                        "time_range": "15min",
                        "alarm_type": "access_source"
                    }
                ]
            },
            model_name="account_login_in_multi_ip"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/2", json=payload)

        # 3.境内IP有多个账号身份
        payload = self._set_risk_payload(
            name="境内IP有多个账号身份",
            description="同一个应用的多个账号使用境内的同一个IP访问，可能存在风险",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "dup_account",
                        "filter_threshold": 1,
                        "time_range": "15min",
                        "alarm_type": "access_source"
                    }
                ]
            },
            model_name="multi_account_login_inborder_or_outborder",
            filters={
                "overseas": False
            },
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/3", json=payload)

        self.r.patch(
            f"https://{self.dsc_ip}/dashboard/risks/3?is_active=true&model_type=normal")

        # 4.境外IP有多个账号身份
        payload = self._set_risk_payload(
            name="境外IP有多个账号身份",
            description="同一个应用的多个账号使用境外的同一个IP访问，可能存在风险",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "dup_account",
                        "filter_threshold": 1,
                        "time_range": "15min",
                        "alarm_type": "access_source"
                    }
                ]
            },
            filters={
                "overseas": False
            },
            model_name="multi_account_login_inborder_or_outborder"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/4", json=payload)

        # 5.单个账号一段时间内返回大量敏感数据
        payload = self._set_risk_payload(
            name="单个账号一段时间内返回大量敏感数据",
            description="单个账号一段时间内返回大量敏感数据",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "sensitive_data",
                        "threshold_type": "threshold_only",
                        "filter_threshold": [
                            {
                                "sensitive_id": 10,
                                "data": 4
                            },
                            {
                                "sensitive_id": 2,
                                "data": 4
                            },
                            {
                                "sensitive_id": 14,
                                "data": 4
                            },
                            {
                                "sensitive_id": 1,
                                "data": 4
                            },
                            {
                                "sensitive_id": 3,
                                "data": 4
                            },
                            {
                                "sensitive_id": 4,
                                "data": 4
                            },
                            {
                                "sensitive_id": 5,
                                "data": 4
                            },
                            {
                                "sensitive_id": 6,
                                "data": 4
                            },
                            {
                                "sensitive_id": 7,
                                "data": 4
                            },
                            {
                                "sensitive_id": 8,
                                "data": 4
                            },
                            {
                                "sensitive_id": 9,
                                "data": 4
                            },
                            {
                                "sensitive_id": 11,
                                "data": 4
                            },
                            {
                                "sensitive_id": 12,
                                "data": 4
                            },
                            {
                                "sensitive_id": 13,
                                "data": 4
                            },
                            {
                                "sensitive_id": 15,
                                "data": 4
                            },
                            {
                                "sensitive_id": 16,
                                "data": 4
                            }
                        ],
                        "sensitive_relation": "or",
                        "time_range": "single_req",
                        "alarm_type": "distinct_data",
                        "sensitive_ids": [
                            10,
                            2,
                            14,
                            1,
                            3,
                            4,
                            5,
                            6,
                            7,
                            8,
                            9,
                            11,
                            12,
                            13,
                            15,
                            16
                        ]
                    }
                ]
            },
            filters={
                "overseas": False
            },
            model_name="mainstay_return_uniq_sensitive_cnt_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/5", json=payload)

        # 6
        payload = self._set_risk_payload(
            name="单个IP一段时间内返回大量敏感数据",
            description="单个IP一段时间内返回大量敏感数据",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "sensitive_data",
                        "threshold_type": "threshold_only",
                        "filter_threshold": [
                            {
                                "sensitive_id": 10,
                                "data": 4
                            },
                            {
                                "sensitive_id": 2,
                                "data": 4
                            },
                            {
                                "sensitive_id": 14,
                                "data": 4
                            },
                            {
                                "sensitive_id": 1,
                                "data": 4
                            },
                            {
                                "sensitive_id": 3,
                                "data": 4
                            },
                            {
                                "sensitive_id": 4,
                                "data": 4
                            },
                            {
                                "sensitive_id": 5,
                                "data": 4
                            },
                            {
                                "sensitive_id": 6,
                                "data": 4
                            },
                            {
                                "sensitive_id": 7,
                                "data": 4
                            },
                            {
                                "sensitive_id": 8,
                                "data": 4
                            },
                            {
                                "sensitive_id": 9,
                                "data": 4
                            },
                            {
                                "sensitive_id": 11,
                                "data": 4
                            },
                            {
                                "sensitive_id": 12,
                                "data": 4
                            },
                            {
                                "sensitive_id": 13,
                                "data": 4
                            },
                            {
                                "sensitive_id": 15,
                                "data": 4
                            },
                            {
                                "sensitive_id": 16,
                                "data": 4
                            }
                        ],
                        "sensitive_relation": "or",
                        "time_range": "single_req",
                        "alarm_type": "distinct_data",
                        "sensitive_ids": [
                            10,
                            2,
                            14,
                            1,
                            3,
                            4,
                            5,
                            6,
                            7,
                            8,
                            9,
                            11,
                            12,
                            13,
                            15,
                            16
                        ]
                    }
                ]
            },
            model_name="mainstay_return_uniq_sensitive_cnt_over_limit_in_single_req",
            filters={
                "sensitive_api": True
            }
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/6", json=payload)

        # 7
        payload = self._set_risk_payload(
            name="单个账号单次返回大量敏感数据",
            description="单个账号单次返回大量敏感数据",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "sensitive_data",
                        "threshold_type": "threshold_only",
                        "filter_threshold": [
                            {
                                "sensitive_id": 10,
                                "data": 4
                            },
                            {
                                "sensitive_id": 2,
                                "data": 4
                            },
                            {
                                "sensitive_id": 14,
                                "data": 4
                            },
                            {
                                "sensitive_id": 1,
                                "data": 4
                            },
                            {
                                "sensitive_id": 3,
                                "data": 4
                            },
                            {
                                "sensitive_id": 4,
                                "data": 4
                            },
                            {
                                "sensitive_id": 5,
                                "data": 4
                            },
                            {
                                "sensitive_id": 6,
                                "data": 4
                            },
                            {
                                "sensitive_id": 7,
                                "data": 4
                            },
                            {
                                "sensitive_id": 8,
                                "data": 4
                            },
                            {
                                "sensitive_id": 9,
                                "data": 4
                            },
                            {
                                "sensitive_id": 11,
                                "data": 4
                            },
                            {
                                "sensitive_id": 12,
                                "data": 4
                            },
                            {
                                "sensitive_id": 13,
                                "data": 4
                            },
                            {
                                "sensitive_id": 15,
                                "data": 4
                            },
                            {
                                "sensitive_id": 16,
                                "data": 4
                            }
                        ],
                        "sensitive_relation": "or",
                        "time_range": "single_req",
                        "alarm_type": "distinct_data",
                        "sensitive_ids": [
                            10,
                            2,
                            14,
                            1,
                            3,
                            4,
                            5,
                            6,
                            7,
                            8,
                            9,
                            11,
                            12,
                            13,
                            15,
                            16
                        ]
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_return_uniq_sensitive_cnt_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/7", json=payload)

        # 8
        payload = self._set_risk_payload(
            name="单个IP单次返回大量敏感数据",
            description="单个IP单次返回大量敏感数据",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "sensitive_data",
                        "threshold_type": "threshold_only",
                        "filter_threshold": [
                            {
                                "sensitive_id": 10,
                                "data": 4
                            },
                            {
                                "sensitive_id": 2,
                                "data": 4
                            },
                            {
                                "sensitive_id": 14,
                                "data": 4
                            },
                            {
                                "sensitive_id": 1,
                                "data": 4
                            },
                            {
                                "sensitive_id": 3,
                                "data": 4
                            },
                            {
                                "sensitive_id": 4,
                                "data": 4
                            },
                            {
                                "sensitive_id": 5,
                                "data": 4
                            },
                            {
                                "sensitive_id": 6,
                                "data": 4
                            },
                            {
                                "sensitive_id": 7,
                                "data": 4
                            },
                            {
                                "sensitive_id": 8,
                                "data": 4
                            },
                            {
                                "sensitive_id": 9,
                                "data": 4
                            },
                            {
                                "sensitive_id": 11,
                                "data": 4
                            },
                            {
                                "sensitive_id": 12,
                                "data": 4
                            },
                            {
                                "sensitive_id": 13,
                                "data": 4
                            },
                            {
                                "sensitive_id": 15,
                                "data": 4
                            },
                            {
                                "sensitive_id": 16,
                                "data": 4
                            }
                        ],
                        "sensitive_relation": "or",
                        "time_range": "single_req",
                        "alarm_type": "distinct_data",
                        "sensitive_ids": [
                            10,
                            2,
                            14,
                            1,
                            3,
                            4,
                            5,
                            6,
                            7,
                            8,
                            9,
                            11,
                            12,
                            13,
                            15,
                            16
                        ]
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_return_uniq_sensitive_cnt_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/8", json=payload)

        # 9
        payload = self._set_risk_payload(
            name="单个账号单次返回敏感数据类型超过15种",
            description="单个账号单次返回敏感数据类型超过15种",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "over_threshold",
                        "filter_threshold": 8,
                        "time_range": "single_req",
                        "alarm_type": "data_type"
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_return_uniq_sensitive_type_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/9", json=payload)

        # 10
        payload = self._set_risk_payload(
            name="单个IP单次返回敏感数据类型超过15种",
            description="单个IP单次返回敏感数据类型超过15种",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "over_threshold",
                        "filter_threshold": 8,
                        "time_range": "single_req",
                        "alarm_type": "data_type"
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_return_uniq_sensitive_type_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/10", json=payload)

        # 13
        payload = self._set_risk_payload(
            name="单个账号在一段时间内进行请求参数值遍历",
            description="单个账号在一段时间内频繁变换请求参数值，疑似进行请求参数值遍历",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "traverse",
                        "filter_threshold": 50,
                        "threshold_type": "assign_param_name",
                        "param_name": [
                            "user"
                        ],
                        "time_range": "15min",
                        "alarm_type": "req_params"
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_req_api_param_traverse_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/13", json=payload)

        # 14
        payload = self._set_risk_payload(
            name="单个IP在一段时间内进行请求参数值遍历",
            description="单个IP在一段时间内频繁变换请求参数值，疑似进行请求参数值遍历",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "traverse",
                        "filter_threshold": 50,
                        "threshold_type": "assign_param_name",
                        "param_name": [
                            "user"
                        ],
                        "time_range": "15min",
                        "alarm_type": "req_params"
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_req_api_param_traverse_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/14", json=payload)

        # 17
        payload = self._set_risk_payload(
            name="单个账号在一段时间内返回大量4XX",
            description="单个账号在一段时间内返回大量400、401、403、404、405、408状态码",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "code_4xx_cnt",
                        "filter_threshold": 50,
                        "threshold_type": "threshold_only",
                        "time_range": "15min",
                        "alarm_type": "ret_status"
                    }
                ]
            },
            model_name="mainstay_return_massive_4xx_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/17", json=payload)

        # 18
        payload = self._set_risk_payload(
            name="单个IP在一段时间内返回大量4XX",
            description="单个IP在一段时间内返回大量400、401、403、404、405、408状态码",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "code_4xx_cnt",
                        "filter_threshold": 50,
                        "threshold_type": "threshold_only",
                        "time_range": "15min",
                        "alarm_type": "ret_status"
                    }
                ]
            },
            model_name="mainstay_return_massive_4xx_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/18", json=payload)

        # 19
        payload = self._set_risk_payload(
            name="单个账号在一段时间内频繁访问同一API",
            description="单个账号在一段时间内频繁访问同一API",
            level="mid",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "same_api",
                        "threshold_type": "threshold_only",
                        "filter_threshold": 50,
                        "time_range": "15min",
                        "alarm_type": "req_cnt"
                    }
                ]
            },
            model_name="mainstay_req_same_api_cnt_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/19", json=payload)

        # 20
        payload = self._set_risk_payload(
            name="单个账号在异常时间段频繁访问同一API",
            description="单个账号在异常时间段频繁访问同一API",
            level="mid",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "same_api",
                        "threshold_type": "threshold_only",
                        "filter_threshold": 50,
                        "time_range": "15min",
                        "alarm_type": "req_cnt"
                    }
                ]
            },
            filters={
                "time_range": {
                    "before_week": 1,
                    "after_week": 7,
                    "first_time": "00:00:00",
                    "last_time": "23:59:59"
                }
            },
            model_name="mainstay_req_same_api_cnt_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/20", json=payload)

        # 21
        payload = self._set_risk_payload(
            name="单个IP在一段时间内频繁访问同一API",
            description="单个IP在一段时间内频繁访问同一API",
            level="mid",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "same_api",
                        "threshold_type": "threshold_only",
                        "filter_threshold": 50,
                        "time_range": "15min",
                        "alarm_type": "req_cnt"
                    }
                ]
            },
            model_name="mainstay_req_same_api_cnt_over_limit_in_time_range"

        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/21", json=payload)

        # 22
        payload = self._set_risk_payload(
            name="单个IP在异常时间段频繁访问同一API",
            description="单个IP在异常时间段频繁访问同一API",
            level="mid",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "same_api",
                        "threshold_type": "threshold_only",
                        "filter_threshold": 50,
                        "time_range": "15min",
                        "alarm_type": "req_cnt"
                    }
                ]
            },
            filters={
                "time_range": {
                    "before_week": 1,
                    "after_week": 7,
                    "first_time": "00:00:00",
                    "last_time": "23:59:59"
                }
            },
            model_name="mainstay_req_same_api_cnt_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/22", json=payload)

        # 25
        payload = self._set_risk_payload(
            name="单个账号在一段时间内进行路径遍历",
            description="单个账号在一段时间内频繁变换路径，疑似进行路径遍历",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "diff_path",
                        "threshold_type": "threshold_only",
                        "filter_threshold": 50,
                        "time_range": "15min",
                        "alarm_type": "req_path"
                    }
                ]
            },
            model_name="mainstay_req_api_path_traverse_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/25", json=payload)

        # 26
        payload = self._set_risk_payload(
            name="单个IP在一段时间内进行路径遍历",
            description="单个IP在一段时间内频繁变换路径，疑似进行路径遍历",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "diff_path",
                        "threshold_type": "threshold_only",
                        "filter_threshold": 50,
                        "time_range": "15min",
                        "alarm_type": "req_path"
                    }
                ]
            },
            model_name="mainstay_req_api_path_traverse_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/26", json=payload)

        self.r.patch(
            f"https://{self.dsc_ip}/dashboard/weaks/7?is_active=true")
        self.r.patch(
            f"https://{self.dsc_ip}/dashboard/weaks/11?is_active=true")

        payload = {"is_identity": True}
        self.r.put(
            f"https://{self.dsc_ip}/dashboard/apps/config?", json=payload)

    # 还原阈值
    def ResetTreshold(self):
        # 1
        payload = self._set_risk_payload(
            name="账号多地访问",
            description="同一个账号在多个不同的地理位置上登录，可能存在风险",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_object": "dup_location",
                        "filter_threshold": 1,
                        "time_range": "1h",
                        "alarm_type": "access_source"
                    }
                ]
            },
            model_name="account_login_in_multi_area"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/1", json=payload)

        # 2
        payload = self._set_risk_payload(
            name="账号多IP访问",
            description="同一个账号在多个不同的IP上登录，可能存在风险",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "time_range": "1h",
                        "filter_threshold": 1,
                        "alarm_type": "access_source",
                        "alarm_object": "dup_ip"
                    }
                ]
            },
            model_name="account_login_in_multi_ip"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/2", json=payload)

        # 3
        payload = self._set_risk_payload(
            name="境内IP有多个账号身份",
            description="同一个应用的多个账号使用境内的同一个IP访问，可能存在风险",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "time_range": "1h",
                        "filter_threshold": 1,
                        "alarm_type": "access_source",
                        "alarm_object": "dup_account"
                    }
                ]
            },
            model_name="multi_account_login_inborder_or_outborder",
            filters={
                "overseas": False
            },
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/3", json=payload)

        self.r.patch(
            f"https://{self.dsc_ip}/dashboard/risks/3?is_active=false&model_type=normal")

        # 4
        payload = self._set_risk_payload(
            name="境外IP有多个账号身份",
            description="同一个应用的多个账号使用境外的同一个IP访问，可能存在风险",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "time_range": "1h",
                        "filter_threshold": 1,
                        "alarm_type": "access_source",
                        "alarm_object": "dup_account"
                    }
                ]
            },
            filters={
                "overseas": False
            },
            model_name="multi_account_login_inborder_or_outborder"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/4", json=payload)

        # 5
        payload = self._set_risk_payload(
            name="单个账号一段时间内返回大量敏感数据",
            description="单个账号一段时间内返回大量敏感数据",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [{
                    "time_range": "1h",
                    "alarm_type": "distinct_data",
                    "alarm_object": "sensitive_data",
                    "filter_threshold": [{"sensitive_id": 10, "data": 1000}, {"sensitive_id": 3, "data": 1000}],
                    "filter_baseline": [{"sensitive_id": 10, "data": 50}, {"sensitive_id": 3, "data": 50}],
                    "sensitive_relation": "or",
                    "threshold_type": "baseline_and_threshold",
                    "sensitive_ids": [10, 3]
                }]
            },
            filters={
                "overseas": False
            },
            model_name="mainstay_return_uniq_sensitive_cnt_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/5", json=payload)

        # 6
        payload = self._set_risk_payload(
            name="单个IP一段时间内返回大量敏感数据",
            description="单个IP一段时间内返回大量敏感数据",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "time_range": "1h",
                        "alarm_type": "distinct_data",
                        "alarm_object": "sensitive_data",
                        "filter_threshold": [{"sensitive_id": 10, "data": 1000}, {"sensitive_id": 3, "data": 1000}],
                        "filter_baseline": [{"sensitive_id": 10, "data": 50}, {"sensitive_id": 3, "data": 50}],
                        "sensitive_relation": "or",
                        "threshold_type": "baseline_and_threshold",
                        "sensitive_ids": [10, 3]
                    }
                ]
            },
            model_name="mainstay_return_uniq_sensitive_cnt_over_limit_in_single_req",
            filters={
                "sensitive_api": True
            }
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/6", json=payload)

        # 7
        payload = self._set_risk_payload(
            name="单个账号单次返回大量敏感数据",
            description="单个账号单次返回大量敏感数据",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "time_range": "single_req",
                        "alarm_type": "distinct_data",
                        "alarm_object": "sensitive_data",
                        "filter_threshold": [{"sensitive_id": 10, "data": 10}, {"sensitive_id": 3, "data": 10}],
                        "filter_baseline": [{"sensitive_id": 10, "data": 50}, {"sensitive_id": 3, "data": 50}],
                        "sensitive_relation": "or",
                        "threshold_type": "baseline_and_threshold",
                        "sensitive_ids": [10, 3]
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_return_uniq_sensitive_cnt_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/7", json=payload)

        # 8
        payload = self._set_risk_payload(
            name="单个IP单次返回大量敏感数据",
            description="单个IP单次返回大量敏感数据",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "time_range": "single_req",
                        "alarm_type": "distinct_data",
                        "alarm_object": "sensitive_data",
                        "filter_threshold": [{"sensitive_id": 10, "data": 10}, {"sensitive_id": 3, "data": 10}],
                        "filter_baseline": [{"sensitive_id": 10, "data": 50}, {"sensitive_id": 3, "data": 50}],
                        "sensitive_relation": "or",
                        "threshold_type": "baseline_and_threshold",
                        "sensitive_ids": [10, 3]
                    }]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_return_uniq_sensitive_cnt_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/8", json=payload)

        # 9
        self._set_risk_payload(
            name="单个账号单次返回敏感数据类型超过15种",
            description="单个账号单次返回敏感数据类型超过15种",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "time_range": "single_req",
                        "filter_threshold": 15,
                        "alarm_type": "data_type",
                        "alarm_object": "over_threshold"
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_return_uniq_sensitive_type_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/9", json=payload)

        # 10
        payload = self._set_risk_payload(
            name="单个IP单次返回敏感数据类型超过15种",
            description="单个IP单次返回敏感数据类型超过15种",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "time_range": "single_req",
                        "filter_threshold": 15,
                        "alarm_type": "data_type",
                        "alarm_object": "over_threshold"
                    }
                ]
            },
            filters={
                "sensitive_api": True
            },
            model_name="mainstay_return_uniq_sensitive_type_over_limit_in_single_req"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/10", json=payload)

        # 13
        payload = self._set_risk_payload(
            name="单个账号在一段时间内进行请求参数值遍历",
            description="单个账号在一段时间内频繁变换请求参数值，疑似进行请求参数值遍历",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "req_params",
                        "alarm_object": "traverse",
                        "time_range": "15min",
                        "filter_threshold": 400,
                        "threshold_type": "all_param_name"
                    }
                ]
            },
            filters={
                "sensitive_api": False
            },
            model_name="mainstay_req_api_param_traverse_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/13", json=payload)

        # 14
        payload = self._set_risk_payload(
            name="单个IP在一段时间内进行请求参数值遍历",
            description="单个IP在一段时间内频繁变换请求参数值，疑似进行请求参数值遍历",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "req_params",
                        "alarm_object": "traverse",
                        "time_range": "15min",
                        "filter_threshold": 400,
                        "threshold_type": "all_param_name"
                    }
                ]
            },
            filters={
                "sensitive_api": False
            },
            model_name="mainstay_req_api_param_traverse_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/14", json=payload)

        # 17
        payload = self._set_risk_payload(
            name="单个账号在一段时间内返回大量4XX",
            description="单个账号在一段时间内返回大量400、401、403、404、405、408状态码",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "ret_status",
                        "alarm_object": "code_4xx_cnt",
                        "time_range": "15min",
                        "filter_threshold": 100,
                        "threshold_type": "threshold_only"
                    }
                ]
            },
            model_name="mainstay_return_massive_4xx_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/17", json=payload)

        # 18
        payload = self._set_risk_payload(
            name="单个IP在一段时间内返回大量4XX",
            description="单个IP在一段时间内返回大量400、401、403、404、405、408状态码",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "ret_status",
                        "alarm_object": "code_4xx_cnt",
                        "time_range": "15min",
                        "filter_threshold": 100,
                        "threshold_type": "threshold_only"
                    }
                ]
            },
            model_name="mainstay_return_massive_4xx_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/18", json=payload)

        # 19
        payload = self._set_risk_payload(
            name="单个账号在一段时间内频繁访问同一API",
            description="单个账号在一段时间内频繁访问同一API",
            level="mid",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "req_cnt",
                        "alarm_object": "same_api",
                        "time_range": "15min",
                        "filter_threshold": 200,
                        "filter_baseline": 50,
                        "threshold_type": "baseline_and_threshold"
                    }
                ]
            },
            model_name="mainstay_req_same_api_cnt_over_limit_in_time_range",
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/19", json=payload)

        # 20
        payload = self._set_risk_payload(
            name="单个账号在异常时间段频繁访问同一API",
            description="单个账号在异常时间段频繁访问同一API",
            level="mid",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "req_cnt",
                        "alarm_object": "same_api",
                        "time_range": "15min",
                        "filter_threshold": 200,
                        "filter_baseline": 50,
                        "threshold_type": "baseline_and_threshold"
                    }
                ]
            },
            filters={
                "choose_condition": "app",
                "care_app": "any",
                "care_api_group_ids": "any",
                "time_range": {
                    "before_week": 6,
                    "after_week": 7,
                    "first_time": "00:00:00",
                    "last_time": "06:00:00"
                }
            },
            model_name="mainstay_req_same_api_cnt_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/20", json=payload)

        # 21
        payload = self._set_risk_payload(
            name="单个IP在一段时间内频繁访问同一API",
            description="单个IP在一段时间内频繁访问同一API",
            level="mid",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "req_cnt",
                        "alarm_object": "same_api",
                        "time_range": "15min",
                        "filter_threshold": 200,
                        "filter_baseline": 50,
                        "threshold_type": "baseline_and_threshold"
                    }
                ]
            },
            model_name="mainstay_req_same_api_cnt_over_limit_in_time_range"

        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/21", json=payload)

        # 22
        payload = self._set_risk_payload(
            name="单个IP在异常时间段频繁访问同一API",
            description="单个IP在异常时间段频繁访问同一API",
            level="mid",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "req_cnt",
                        "alarm_object": "same_api",
                        "time_range": "15min",
                        "filter_threshold": 200,
                        "filter_baseline": 50,
                        "threshold_type": "baseline_and_threshold"
                    }
                ]
            },
            filters={
                "choose_condition": "app",
                "care_app": "any",
                "care_api_group_ids": "any",
                "time_range": {
                    "before_week": 6,
                    "after_week": 7,
                    "first_time": "00:00:00",
                    "last_time": "06:00:00"
                }
            },
            model_name="mainstay_req_same_api_cnt_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/22", json=payload)

        # 25
        payload = self._set_risk_payload(
            name="单个账号在一段时间内进行路径遍历",
            description="单个账号在一段时间内频繁变换路径，疑似进行路径遍历",
            level="high",
            model_cfg={
                "mainstay_type": "account",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "req_path",
                        "alarm_object": "diff_path",
                        "time_range": "15min",
                        "filter_threshold": 100,
                        "filter_baseline": 50,
                        "threshold_type": "baseline_and_threshold"
                    }
                ]
            },
            model_name="mainstay_req_api_path_traverse_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/25", json=payload)

        # 26
        payload = self._set_risk_payload(
            name="单个IP在一段时间内进行路径遍历",
            description="单个IP在一段时间内频繁变换路径，疑似进行路径遍历",
            level="high",
            model_cfg={
                "mainstay_type": "ip",
                "mainstay_filter_type": "any",
                "conditions": [
                    {
                        "alarm_type": "req_path",
                        "alarm_object": "diff_path",
                        "time_range": "15min",
                        "filter_threshold": 100,
                        "filter_baseline": 50,
                        "threshold_type": "baseline_and_threshold"
                    }
                ]
            },
            model_name="mainstay_req_api_path_traverse_over_limit_in_time_range"
        )
        self.r.put(f"https://{self.dsc_ip}/dashboard/risks/26", json=payload)

        self.r.patch(
            f"https://{self.dsc_ip}/dashboard/weaks/7?is_active=false")
        self.r.patch(
            f"https://{self.dsc_ip}/dashboard/weaks/11?is_active=false")


if __name__ == "__main__":
    handler = ModifyRiskPolicy()
    handler.ModifyThreshold()
