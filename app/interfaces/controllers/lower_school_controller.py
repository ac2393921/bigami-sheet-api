from typing import List

from app.domain.entities.lower_school import LowerSchool
from app.usecases.lower_school.fetch_all.fetch_all_lower_school_output_port import (
    FetchAllLowerSchoolOutputPort,
)
from app.usecases.lower_school.fetch_all.fetch_all_lower_school_usecase import (
    FetchAllLowerSchoolUsecase,
)


class LowerSchoolController:
    def __init__(self, fetch_all_school_uscase: FetchAllLowerSchoolUsecase) -> None:
        self._fetch_all_lower_school_uscase = fetch_all_school_uscase

    def fetch_all_lower_school(self) -> List[LowerSchool]:
        output = self._fetch_all_lower_school_uscase.handle()
        return output.lower_schools
