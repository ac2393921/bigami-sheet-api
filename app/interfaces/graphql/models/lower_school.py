import strawberry

from app.domain.entities.lower_school import LowerSchool


@strawberry.experimental.pydantic.type(LowerSchool)
class LowerSchoolType:
    id: strawberry.auto
    school_id: strawberry.auto
    name: strawberry.auto
    style: strawberry.auto
    enemy: strawberry.auto
