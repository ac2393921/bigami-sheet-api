from app.domain.repositories.i_belief_repository import IBeliefRepository
from app.usecases.belief.fetch_all.fetch_all_belief_output_port import (
    FetchAllBeliefOutputPort,
)
from app.usecases.belief.fetch_all.fetch_all_belief_usecase import FetchAllBeliefUsecase


class FetchAllBeliefInteractor(FetchAllBeliefUsecase):
    def __init__(self, rank_repository: IBeliefRepository) -> None:
        self._rank_repository = rank_repository

    def handle(self) -> FetchAllBeliefOutputPort:
        beliefs = self._rank_repository.fetch_all()
        return FetchAllBeliefOutputPort(beliefs=beliefs)
