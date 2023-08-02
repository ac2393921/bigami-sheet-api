from typing import List

import strawberry

from app.interfaces.graphql.models.belief import BeliefType
from app.interfaces.graphql.models.lower_school import LowerSchoolType
from app.interfaces.graphql.models.rank import RankType
from app.interfaces.graphql.models.school import SchoolType
from app.interfaces.graphql.resolver import (
    fetch_all_beliefs,
    fetch_all_lower_schools,
    fetch_all_ranks,
    fetch_all_schools,
)


@strawberry.type
class Query:
    schools: List[SchoolType] = strawberry.field(resolver=fetch_all_schools)
    lower_schools: List[LowerSchoolType] = strawberry.field(
        resolver=fetch_all_lower_schools
    )
    ranks: List[RankType] = strawberry.field(resolver=fetch_all_ranks)
    belief: List[BeliefType] = strawberry.field(resolver=fetch_all_beliefs)
