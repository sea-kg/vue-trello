#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys
import re
import requests

found_files = []
_result = []
with open("web/package-lock.json") as _file:
    _lines = _file.readlines()
    for _line in _lines:
        _line0 = _line.strip()
        if "https://registry.npm.taobao.org/" not in _line:
            _result.append(_line)
        else:
            filename = None
            m = re.search(".*/([a-zA-Z\.0-9-_]+\.tgz).*", _line)
            if m:
                filename = m.group(1)
            if filename is None:
                print("failed on line ", _line.strip())
                sys.exit(-1)
            # print(_line0)
            # print(filename)
            if not filename.endswith(".tgz"):
                print("failed on line ", _line)
                print("filename ", filename)
                sys.exit(-1)
            _replace_from = _line[_line.index("https://"):_line.index(filename)-1]
            _new_line = _line.replace(_replace_from, "https://sea5kg.ru/files/npm-packages")
            _result.append(_new_line)
            with open("web/package-lock2.json", 'wt') as _file2:
                _file2.write("".join(_result))
            if os.path.isfile(filename):
                print("Already downloaded... " + filename)
                continue
            found_files.append(filename)
            if not os.path.isfile(filename):
                url = _line0
                url = url.replace('"resolved": "', '')
                if url.endswith(','):
                    url = url[:-1]
                if url.endswith('"'):
                    url = url[:-1]
                print("url = ", url)
                resp = requests.get(url, allow_redirects=True, verify=False)
                with open(filename, 'wb') as _file_tgz:
                    _file_tgz.write(resp.content)
                # sys.exit(-1)
with open("web/package-lock2.json", 'wt') as _file2:
    _file2.write("".join(_result))
print("found files to download: ", len(found_files))
