from app.domain.repositories.i_school_repository import ISchoolRepository
from app.usecases.lower_school.fetch_all.fetch_all_lower_school_output_port import (
    FetchAllLowerSchoolOutputPort,
)
from app.usecases.school.fetch_all.fetch_all_school_usecase import FetchAllSchoolUsecase


class FetchAllLowerSchoolInteractor(FetchAllSchoolUsecase):
    def __init__(self, lower_school_repository: ISchoolRepository) -> None:
        self._lower_school_repository = lower_school_repository

    def handle(self) -> FetchAllLowerSchoolOutputPort:
        lower_schools = self._lower_school_repository.fetch_all()
        return FetchAllLowerSchoolOutputPort(lower_schools=lower_schools)
