from abc import ABC, abstractmethod
from typing import List

from app.domain.entities.rank import Rank


class IRankRepository(ABC):
    @abstractmethod
    def fetch_all(self) -> List[Rank]:
        raise NotImplementedError
