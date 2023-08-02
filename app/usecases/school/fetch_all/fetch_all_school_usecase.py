from abc import ABC, abstractmethod

from app.usecases.school.fetch_all.fetch_all_school_output_port import (
    FetchAllSchoolOutputPort,
)


class FetchAllSchoolUsecase(ABC):
    @abstractmethod
    def handle(self) -> FetchAllSchoolOutputPort:
        raise NotImplementedError
