/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu
 * @Date: 2021-06-22 12:45:43
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-09-12 16:10:46
 * @FilePath: /ytoolsbox-gin/model/model.go
 */

package model

// {
//     "data": {
//         "totalpage": 5,
//         "pagenum": 4,
//         "users": [
//             {
//                 "id": 25,
//                 "username": "tige117",
//                 "mobile": "18616358651",
//                 "type": 1,
//                 "email": "tige112@163.com",
//                 "create_time": "2017-11-09T20:36:26.000Z",
//                 "mg_state": true, // 当前用户的状态
//                 "role_name": "炒鸡管理员"
//             }
//         ]
//     },
//     "meta": {
//         "msg": "获取成功",
//         "status": 200
//     }
// }

type UsersData struct {
	Total   int    `json:"total"`
	Pagenum int    `json:"pagenum"`
	Users   []User `json:"users"`
}

// 在这里定义好表的结构
type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Mobile     string `json:"mobile"`
	Type       int    `json:"type"`
	Email      string `json:"email"`
	CreateTime string `json:"create_time"`
	MgState    bool   `json:"mgstate"`
	RoleName   string `josn:"role"`
}
