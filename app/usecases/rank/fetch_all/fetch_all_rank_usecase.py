from abc import ABC, abstractmethod

from app.usecases.rank.fetch_all.fetch_all_rank_output_port import (
    FetchAllRankOutputPort,
)


class FetchAllRankUsecase(ABC):
    @abstractmethod
    def handle(self) -> FetchAllRankOutputPort:
        raise NotImplementedError
