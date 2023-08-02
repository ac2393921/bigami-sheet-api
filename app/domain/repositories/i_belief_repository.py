from abc import ABC, abstractmethod
from typing import List

from app.domain.entities.belief import Belief


class IBeliefRepository(ABC):
    @abstractmethod
    def fetch_all(self) -> List[Belief]:
        raise NotImplementedError
