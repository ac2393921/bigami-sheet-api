from typing import List

from app.domain.entities.rank import Rank
from app.usecases.rank.fetch_all.fetch_all_rank_usecase import FetchAllRankUsecase


class RankController:
    def __init__(self, fetch_all_school_uscase: FetchAllRankUsecase) -> None:
        self._fetch_all_rank_uscase = fetch_all_school_uscase

    def fetch_all_rank(self) -> List[Rank]:
        output = self._fetch_all_rank_uscase.handle()
        return output.ranks
