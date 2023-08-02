import strawberry

from app.domain.entities.school import School


@strawberry.experimental.pydantic.type(School)
class SchoolType:
    id: strawberry.auto
    name: strawberry.auto
    style: strawberry.auto
    enemy: strawberry.auto
