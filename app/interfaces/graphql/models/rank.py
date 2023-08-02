import strawberry

from app.domain.entities.rank import Rank


@strawberry.experimental.pydantic.type(Rank)
class RankType:
    id: strawberry.auto
    name: strawberry.auto
