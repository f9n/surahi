# surahi
<p>Sürahi is offline pip. Sürahi is finding packages in other python projects from your Pc. Then, it is copying package files to current venv.</p>

Python Global Path: /usr/lib/python*/site-packages

# Why
<p>When you clone python project, then you install necessary python packages with pip. Example: </p>

```bash
$ git clone https://github.com/erginipekci7/flask_mysql_example
$ cd flask_mysql_example
$ ls
LICENSE  main.py  README.md  requirements.txt
$ # In "requirements.txt", there are necessary packages.
$ virtualenv venv                   # create virtual enviroment
$ source venv/bin/active            # actived virtual environment
(venv)$ # When you have internet connection.
(venv)$ pip install -r requirements.txt   # download requirement packages.
Collecting click==6.7 (from -r requirements.txt (line 1))
  Downloading click-6.7-py2.py3-none-any.whl (71kB)
    100% |████████████████████████████████| 71kB 461kB/s 
Collecting Flask==0.12.2 (from -r requirements.txt (line 2))
  Downloading Flask-0.12.2-py2.py3-none-any.whl (83kB)
    100% |████████████████████████████████| 92kB 1.0MB/s 
Collecting Flask-MySQLdb==0.2.0 (from -r requirements.txt (line 3))
  Downloading Flask-MySQLdb-0.2.0.tar.gz
Collecting itsdangerous==0.24 (from -r requirements.txt (line 4))
  Downloading itsdangerous-0.24.tar.gz (46kB)
    100% |████████████████████████████████| 51kB 788kB/s 
Requirement already satisfied: Jinja2==2.9.6 in /usr/lib/python3.6/site-packages (from -r requirements.txt (line 5))
Requirement already satisfied: MarkupSafe==1.0 in /usr/lib/python3.6/site-packages (from -r requirements.txt (line 6))
Collecting mysqlclient==1.3.12 (from -r requirements.txt (line 7))
  Downloading mysqlclient-1.3.12.tar.gz (89kB)
    100% |████████████████████████████████| 92kB 240kB/s 
Collecting python-dotenv==0.7.1 (from -r requirements.txt (line 8))
  Downloading python_dotenv-0.7.1-py2.py3-none-any.whl
Collecting Werkzeug==0.12.2 (from -r requirements.txt (line 9))
  Downloading Werkzeug-0.12.2-py2.py3-none-any.whl (312kB)
    100% |████████████████████████████████| 317kB 522kB/s 
Installing collected packages: click, Werkzeug, itsdangerous, Flask, mysqlclient, Flask-MySQLdb, python-dotenv
  Running setup.py install for itsdangerous ... done
  Running setup.py install for mysqlclient ... done
  Running setup.py install for Flask-MySQLdb ... done
Successfully installed Flask-0.12.2 Flask-MySQLdb-0.2.0 Werkzeug-0.12.2 click-6.7 itsdangerous-0.24 mysqlclient-1.3.12 python-dotenv-0.7.1
(venv)$ deactive  
$ rm -r venv
$ rm -r ~/.cache/pip
$ source venv/bin/active
(venv)$ # when you haven't internet connection
Collecting click==6.7 (from -r requirements.txt (line 1))
  Retrying (Retry(total=4, connect=None, read=None, redirect=None)) after connection broken by 'NewConnectionError('<pip._vendor.requests.packages.urllib3.connection.VerifiedHTTPSConnection object at 0x7f253e56e160>: Failed to establish a new connection: [Errno -2] Name or service not known',)': /simple/click/
  Retrying (Retry(total=3, connect=None, read=None, redirect=None)) after connection broken by 'NewConnectionError('<pip._vendor.requests.packages.urllib3.connection.VerifiedHTTPSConnection object at 0x7f253e56e240>: Failed to establish a new connection: [Errno -2] Name or service not known',)': /simple/click/
  Retrying (Retry(total=2, connect=None, read=None, redirect=None)) after connection broken by 'NewConnectionError('<pip._vendor.requests.packages.urllib3.connection.VerifiedHTTPSConnection object at 0x7f253e56eb38>: Failed to establish a new connection: [Errno -2] Name or service not known',)': /simple/click/
  Retrying (Retry(total=1, connect=None, read=None, redirect=None)) after connection broken by 'NewConnectionError('<pip._vendor.requests.packages.urllib3.connection.VerifiedHTTPSConnection object at 0x7f253e56ef60>: Failed to establish a new connection: [Errno -2] Name or service not known',)': /simple/click/
  Retrying (Retry(total=0, connect=None, read=None, redirect=None)) after connection broken by 'NewConnectionError('<pip._vendor.requests.packages.urllib3.connection.VerifiedHTTPSConnection object at 0x7f253e56eba8>: Failed to establish a new connection: [Errno -2] Name or service not known',)': /simple/click/
  Could not find a version that satisfies the requirement click==6.7 (from -r requirements.txt (line 1)) (from versions: )
No matching distribution found for click==6.7 (from -r requirements.txt (line 1))
$ # i dont understand that i have "click" package in global.
$ cd /usr/lib/python3.6/site-packages
$ ls click*
click:
_bashcomplete.py  decorators.py  globals.py   __pycache__      testing.py    _unicodefun.py
_compat.py        exceptions.py  __init__.py  _termui_impl.py  _textwrap.py  utils.py
core.py           formatting.py  parser.py    termui.py        types.py      _winconsole.py

click-6.7.dist-info:
DESCRIPTION.rst  INSTALLER  METADATA  metadata.json  RECORD  top_level.txt  WHEEL
```
