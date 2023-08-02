from abc import ABC, abstractmethod

from app.usecases.lower_school.fetch_all.fetch_all_lower_school_output_port import (
    FetchAllLowerSchoolOutputPort,
)


class FetchAllLowerSchoolUsecase(ABC):
    @abstractmethod
    def handle(self) -> FetchAllLowerSchoolOutputPort:
        raise NotImplementedError
