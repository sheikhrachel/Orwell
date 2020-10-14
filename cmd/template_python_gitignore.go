package cmd

var gitignoresrc_str = `api_key.py`

var tox_str = `
[tox]
envlist = py37
`

var travis_str = `
language: python
python:
  - "3.7"
install:
  - pipenv install
script:
  - pipenv run pytest -v
`

var pipfile_str = `
[[source]]
name = "pypi"
url = "https://pypi.org/simple"
verify_ssl = true

[dev-packages]

[packages]
requests = "*"
pytest = "*"
pre-commit = "*"
black = "*"

[requires]
python_version = "3.7"

[pipenv]
allow_prereleases = true
`

var precommit_str = `
repos:
-   repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v1.4.0  # Use the ref you want to point at
    hooks:
    -   id: trailing-whitespace
    -   id: check-ast
    -   id: check-case-conflict
    -   id: check-yaml
    -   id: pretty-format-json
        args: ["--autofix"]

-   repo: https://github.com/ambv/black
    rev: stable
    hooks:
    - id: black
      language_version: python3.8

-   repo: local
    hooks:
    -   id: python-bandit-vulnerability-check
        name: bandit
        entry: bandit
        args: ['--ini', 'tox.ini', '-r', 'consoleme']
        language: system
        pass_filenames: false
    -   id: tests
        name: run tests
        entry: pytest -v
        language: system
        types: [python]
        stages: [push]
`

var gitignore_str = `# Byte-compiled / optimized / DLL files
__pycache__/
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
.Python
build/
develop-eggs/
dist/
downloads/
eggs/
.eggs/
lib/
lib64/
parts/
sdist/
var/
wheels/
pip-wheel-metadata/
share/python-wheels/
*.egg-info/
.installed.cfg
*.egg
MANIFEST

# PyInstaller
#  Usually these files are written by a python script from a template
#  before PyInstaller builds the exe, so as to inject date/other infos into it.
*.manifest
*.spec

# Installer logs
pip-log.txt
pip-delete-this-directory.txt

# Unit test / coverage reports
htmlcov/
.tox/
.nox/
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
*.py,cover
.hypothesis/
.pytest_cache/
cover/

# Translations
*.mo
*.pot

# Django stuff:
*.log
local_settings.py
db.sqlite3
db.sqlite3-journal

# Flask stuff:
instance/
.webassets-cache

# Scrapy stuff:
.scrapy

# Sphinx documentation
docs/_build/

# PyBuilder
target/

# Jupyter Notebook
.ipynb_checkpoints

# IPython
profile_default/
ipython_config.py

# pyenv
#   For a library or package, you might want to ignore these files since the code is
#   intended to run in multiple environments; otherwise, check them in:
# .python-version

# pipenv
#   According to pypa/pipenv#598, it is recommended to include Pipfile.lock in version control.
#   However, in case of collaboration, if having platform-specific dependencies or dependencies
#   having no cross-platform support, pipenv may install dependencies that don't work, or not
#   install all needed dependencies.
#Pipfile.lock

# PEP 582; used by e.g. github.com/David-OConnor/pyflow
__pypackages__/

# Celery stuff
celerybeat-schedule
celerybeat.pid

# SageMath parsed files
*.sage.py

# Environments
.env
.venv
env/
venv/
ENV/
env.bak/
venv.bak/

# Spyder project settings
.spyderproject
.spyproject

# Rope project settings
.ropeproject

# mkdocs documentation
/site

# mypy
.mypy_cache/
.dmypy.json
dmypy.json

# Pyre type checker
.pyre/

# pytype static type analyzer
.pytype/
`