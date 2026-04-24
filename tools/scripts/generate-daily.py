template = """\
Date :
Énergie : /10
Discipline : /10
Humeur : /10
Objectif du jour :
Notes :
"""

with open("daily_log.rst", "w") as f:
    f.write(template)
