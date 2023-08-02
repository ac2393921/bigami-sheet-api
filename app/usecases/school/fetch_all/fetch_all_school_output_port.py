from pydantic import BaseModel

from app.domain.entities.school import School


class FetchAllSchoolOutputPort(BaseModel):
    schools: list[School]
