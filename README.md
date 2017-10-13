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
$ # When you have internet connection.
$ pip install -r requirements.txt   # download requirement packages.
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
```
