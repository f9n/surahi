# surahi
<p>Sürahi is offline pip. Sürahi is finding packages in other python projects from your Pc. Then, it is copying package files to current venv.</p>
<p>Currently, we make just offline pip.And We tried this project with <a href="https://github.com/pleycpl/scripts-shell/blob/master/offlinePip3forPython3.sh">Shell Script</a></p>

<p><b>Python Global Path:</b> /usr/lib/python*/site-packages</p>
<p><b>Nodejs Global Path:</b> /usr/lib/node_modules</p>

# Why
<p>When you clone python project, then you install necessary python packages with pip. There are 2 cases. First is that you have internet connection. Second is that you haven't internet connection, but in your pc you have same packages.</p>

# Pip

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

# Python Package Information
<p>In python, we have 2 directory. First is package's source code. Second is package's version.</p>
<p>Example: "flask" and "Flask-0.12.2.dist-info". And require packages list is in "Flask-0.12.2.dist-info/metadata.json" file.</p>

```bash
$ cat ~/other_project/venv/lib/python3.6/site-packages/Flask-0.12.2.dist-info/metadata.json | jq
{
  "classifiers": [
    "Development Status :: 4 - Beta",
    "Environment :: Web Environment",
    "Intended Audience :: Developers",
    "License :: OSI Approved :: BSD License",
    "Operating System :: OS Independent",
    "Programming Language :: Python",
    "Programming Language :: Python :: 2",
    "Programming Language :: Python :: 2.6",
    "Programming Language :: Python :: 2.7",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.3",
    "Programming Language :: Python :: 3.4",
    "Programming Language :: Python :: 3.5",
    "Topic :: Internet :: WWW/HTTP :: Dynamic Content",
    "Topic :: Software Development :: Libraries :: Python Modules"
  ],
  ...
  "name": "Flask",
  "platform": "any",
  "run_requires": [
    {
      "requires": [
        "Jinja2 (>=2.4)",
        "Werkzeug (>=0.7)",
        "click (>=2.0)",
        "itsdangerous (>=0.21)"
      ]
    }
  ],
  "summary": "A microframework based on Werkzeug, Jinja2 and good intentions",
  "version": "0.12.2"
}

```

# Npm

```bash
$ git clone https://github.com/pleycpl/downloadAllPdf
$ cd downloadAllPdf
$ cat package.json
{
  "name": "downloadallpdf",
  "version": "1.0.0",
  "description": "Download all pdf in tutorialspoint",
  "main": "main.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "mkdir pdfs; node main.js"
  },
  ...
  "homepage": "https://github.com/pleycpl/downloadAllPdf#readme",
  "dependencies": {
    "cfonts": "^1.1.2",
    "chalk": "^1.1.3",
    "cheerio": "^0.22.0",
    "events": "^1.1.1",
    "figlet": "^1.2.0",
    "log-symbols": "^1.0.2",
    "progress": "^2.0.0",
    "request": "^2.81.0",
    "shelljs": "^0.7.7"
  }
}
$ # look Managing the Cache section in "https://www.sitepoint.com/beginners-guide-node-package-manager/"
$ ls .npm | wc -l
351
$ # When you haven't internet connection
$ npm install
npm WARN registry Using stale data from https://registry.npmjs.org/ because the host is inaccessible -- are you offline?
npm WARN registry Using stale package data from https://registry.npmjs.org/ due to a request error during revalidation.
npm ERR! code ENOTFOUND
npm ERR! errno ENOTFOUND
npm ERR! network request to https://registry.npmjs.org/figlet failed, reason: getaddrinfo ENOTFOUND registry.npmjs.org registry.npmjs.org:443
npm ERR! network This is a problem related to network connectivity.
npm ERR! network In most cases you are behind a proxy or have bad network settings.
npm ERR! network 
npm ERR! network If you are behind a proxy, please make sure that the
npm ERR! network 'proxy' config is set properly.  See: 'npm help config'

npm ERR! A complete log of this run can be found in:
npm ERR!     ~/.npm/_logs/2017-10-13T09_59_05_174Z-debug.log
$ npm install   # internet connection
npm notice created a lockfile as package-lock.json. You should commit this file.
added 141 packages in 5.692s
$ ls node_modules
...
ansi-regex       css-select            hawk                   lodash.merge          regenerator-runtime
...
boom             escape-string-regexp  is-typedarray          no-case               strip-ansi
brace-expansion  events                is-upper-case          nth-check             supports-color
camel-case       extend                jsbn                   oauth-sign            swap-case
caseless         extsprintf            jsonify                once                  title-case
cfonts           fast-deep-equal       json-schema            param-case            tough-cookie
chalk            figlet                json-schema-traverse   pascal-case           tunnel
...
core-util-is     har-validator         lodash.foreach         readable-stream
$ rm -r node_modules
$ ls .npm | wc -l
351
$ npm install   # non-internet connection
npm WARN registry Using stale data from https://registry.npmjs.org/ because the host is inaccessible -- are you offline?
npm WARN registry Using stale package data from https://registry.npmjs.org/ due to a request error during revalidation.
npm notice created a lockfile as package-lock.json. You should commit this file.
added 141 packages in 2.256s
$ ls .npm | wc -l
351
$ # Where is the cache packages:))
```

<h4>In project's node_modules, we have necessary packages. And In necessary package's node_modules, we have other necessary packages.....</h4>

```bash
$ git clone https://github.com/pleycpl/downloadAllPdf.git
$ cd downloadAllPdf
$ npm install # internet connection
$ ls node_modules | wc -l
134
$ ls node_modules/chalk
ansi-regex  has-ansi  strip-ansi  supports-color
$ cat node_modules/chalk/node_modules/ansi-regex/package.json
{
  "_from": "ansi-regex@^2.0.0",
  "_id": "ansi-regex@2.1.1",
  ...
  ...
  "version": "2.1.1",
  ...
}
$ cat node_modules/ansi-regex.json
{
  "_from": "ansi-regex@^1.1.0",
  "_id": "ansi-regex@1.1.1",
  ...
  "version": "1.1.1"
}
```
# Next
<p>Research npm and yarn package managers</p>
