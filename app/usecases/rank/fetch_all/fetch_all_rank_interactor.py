from app.domain.repositories.i_rank_repository import IRankRepository
from app.usecases.rank.fetch_all.fetch_all_rank_output_port import (
    FetchAllRankOutputPort,
)
from app.usecases.rank.fetch_all.fetch_all_rank_usecase import FetchAllRankUsecase


class FetchAllRankInteractor(FetchAllRankUsecase):
    def __init__(self, rank_repository: FetchAllRankOutputPort) -> None:
        self._rank_repository = rank_repository

    def handle(self) -> FetchAllRankOutputPort:
        ranks = self._rank_repository.fetch_all()
        return FetchAllRankOutputPort(ranks=ranks)
