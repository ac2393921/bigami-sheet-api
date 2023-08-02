from app.domain.repositories.i_school_repository import ISchoolRepository
from app.usecases.school.fetch_all.fetch_all_school_output_port import (
    FetchAllSchoolOutputPort,
)
from app.usecases.school.fetch_all.fetch_all_school_usecase import FetchAllSchoolUsecase


class FetchAllSchoolInteractor(FetchAllSchoolUsecase):
    def __init__(self, school_repository: ISchoolRepository) -> None:
        self._school_repository = school_repository

    def handle(self) -> FetchAllSchoolOutputPort:
        schools = self._school_repository.fetch_all()
        return FetchAllSchoolOutputPort(schools=schools)
