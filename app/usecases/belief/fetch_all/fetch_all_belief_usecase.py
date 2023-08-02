from abc import ABC, abstractmethod

from app.usecases.belief.fetch_all.fetch_all_belief_output_port import (
    FetchAllBeliefOutputPort,
)


class FetchAllBeliefUsecase(ABC):
    @abstractmethod
    def handle(self) -> FetchAllBeliefOutputPort:
        raise NotImplementedError
