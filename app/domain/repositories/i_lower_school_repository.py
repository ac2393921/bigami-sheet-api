from abc import ABC, abstractmethod
from typing import List

from app.domain.entities.lower_school import LowerSchool


class ILowerSchoolRepository(ABC):
    @abstractmethod
    def fetch_all(self) -> List[LowerSchool]:
        raise NotImplementedError
