from typing import List

from app.domain.entities.belief import Belief
from app.usecases.belief.fetch_all.fetch_all_belief_usecase import FetchAllBeliefUsecase


class BeliefController:
    def __init__(self, fetch_all_school_uscase: FetchAllBeliefUsecase) -> None:
        self._fetch_all_rank_uscase = fetch_all_school_uscase

    def fetch_all_belief(self) -> List[Belief]:
        output = self._fetch_all_rank_uscase.handle()
        return output.ranks
