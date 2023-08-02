from abc import ABC, abstractmethod
from typing import List

from app.domain.entities.school import School


class ISchoolRepository(ABC):
    @abstractmethod
    def fetch_all(self) -> List[School]:
        raise NotImplementedError
