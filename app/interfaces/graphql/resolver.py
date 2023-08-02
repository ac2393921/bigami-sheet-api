from typing import List

from app.infrastructures.database.repositories.belief.inmemory_belief_repository_impl import (
    InmemoryBeliefRepositoryImpl,
)
from app.infrastructures.database.repositories.lower_school.inmemory_lower_school_repository_impl import (
    InmemoryLowerSchoolRepositoryImpl,
)
from app.infrastructures.database.repositories.rank.inmemory_rank_repository_impl import (
    InmemoryRankRepositoryImpl,
)
from app.infrastructures.database.repositories.school.inmemory_school_repository_impl import (
    InmemorySchoolRepositoryImpl,
)
from app.interfaces.controllers.belief_controller import BeliefController
from app.interfaces.controllers.lower_school_controller import LowerSchoolController
from app.interfaces.controllers.rank_controller import RankController
from app.interfaces.controllers.school_controller import SchoolController
from app.interfaces.graphql.models.belief import BeliefType
from app.interfaces.graphql.models.lower_school import LowerSchoolType
from app.interfaces.graphql.models.rank import RankType
from app.interfaces.graphql.models.school import SchoolType
from app.usecases.belief.fetch_all.fetch_all_belief_interactor import (
    FetchAllBeliefInteractor,
)
from app.usecases.lower_school.fetch_all.fetch_all_lower_school_interactor import (
    FetchAllLowerSchoolInteractor,
)
from app.usecases.rank.fetch_all.fetch_all_rank_interactor import FetchAllRankInteractor
from app.usecases.school.fetch_all.fetch_all_school_interactor import (
    FetchAllSchoolInteractor,
)

school_controller = SchoolController(
    FetchAllSchoolInteractor(InmemorySchoolRepositoryImpl())
)

lower_school_controller = LowerSchoolController(
    FetchAllLowerSchoolInteractor(InmemoryLowerSchoolRepositoryImpl())
)

rank_controller = RankController(FetchAllRankInteractor(InmemoryRankRepositoryImpl()))

belief_controller = BeliefController(
    FetchAllBeliefInteractor(InmemoryBeliefRepositoryImpl())
)


def fetch_all_schools() -> List[SchoolType]:
    school_list = school_controller.fetch_all_school()
    return [SchoolType.from_pydantic(school) for school in school_list]


def fetch_all_lower_schools() -> List[LowerSchoolType]:
    lower_school_list = lower_school_controller.fetch_all_lower_school()
    return [
        LowerSchoolType.from_pydantic(lower_school)
        for lower_school in lower_school_list
    ]


def fetch_all_ranks() -> List[RankType]:
    rank_list = rank_controller.fetch_all_rank()
    return [RankType.from_pydantic(rank) for rank in rank_list]


def fetch_all_beliefs() -> List[BeliefType]:
    belief_list = belief_controller.fetch_all_belief()
    return [BeliefType.from_pydantic(belief) for belief in belief_list]
