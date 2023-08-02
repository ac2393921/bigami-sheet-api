from typing import List

from app.domain.entities.rank import Rank
from app.domain.repositories.i_rank_repository import IRankRepository

RANK_LIST = [
    ("0", "中忍"),
    ("1", "中忍頭"),
    ("2", "上忍"),
    ("3", "上忍頭"),
    ("4", "頭領"),
]


class InmemoryRankRepositoryImpl(IRankRepository):
    def __init__(self):
        pass

    def fetch_all(self) -> List[Rank]:
        rank_list = []

        for rank in RANK_LIST:
            rank_list.append(
                Rank(
                    id=rank[0],
                    name=rank[1],
                )
            )

        return rank_list
