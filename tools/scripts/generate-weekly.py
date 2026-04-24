template = """\
Weekly Review
=============

1. Ce qui a bien fonctionné :
2. Ce qui a moins bien fonctionné :
3. Ce que j’ai appris :
4. Priorités de la semaine suivante :
"""

with open("weekly_review.rst", "w") as f:
    f.write(template)
