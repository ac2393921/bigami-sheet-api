from typing import List

from app.domain.entities.school import School
from app.usecases.school.fetch_all.fetch_all_school_output_port import (
    FetchAllSchoolOutputPort,
)
from app.usecases.school.fetch_all.fetch_all_school_usecase import FetchAllSchoolUsecase


class SchoolController:
    def __init__(self, fetch_all_school_uscase: FetchAllSchoolUsecase) -> None:
        self._fetch_all_school_uscase = fetch_all_school_uscase

    def fetch_all_school(self) -> List[School]:
        output: FetchAllSchoolOutputPort = self._fetch_all_school_uscase.handle()
        return output.schools
