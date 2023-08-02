import strawberry

from app.domain.entities.belief import Belief


@strawberry.experimental.pydantic.type(Belief)
class BeliefType:
    id: strawberry.auto
    name: strawberry.auto
    description: strawberry.auto
