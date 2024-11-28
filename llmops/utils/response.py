# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: zhanghan2021 <zhanghan@kylinos.cn> 
# * Date: Thu Nov 7 15:53:47 2024 +0800 
from flask import jsonify


def success(data=None, msg=""):
    return result(200, 200, data, msg)


def fail(data=None, msg=""):
    return result(200, 400, data, msg)


def not_found(data=None, msg=""):
    return result(200, 404, data, msg)


def result(http_status, code, data=None, msg=""):
    response = {"code": code, "data": data, "msg": msg}
    return jsonify(response), http_status
