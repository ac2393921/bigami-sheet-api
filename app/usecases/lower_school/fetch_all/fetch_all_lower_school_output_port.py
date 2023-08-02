from pydantic import BaseModel

from app.domain.entities.lower_school import LowerSchool


class FetchAllLowerSchoolOutputPort(BaseModel):
    lower_schools: list[LowerSchool]
