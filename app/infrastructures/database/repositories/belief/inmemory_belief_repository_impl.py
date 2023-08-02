from typing import List

from app.domain.entities.belief import Belief
from app.domain.repositories.i_belief_repository import IBeliefRepository

BELIEF_LIST = [
    ("0", "和", "多くの人幸せになる未来を信じ、互いにわかり合う道を模索する。"),
    ("1", "愛", "愛する人を守るためには、どんな犠牲も厭わない。"),
]


class InmemoryBeliefRepositoryImpl(IBeliefRepository):
    def __init__(self):
        pass

    def fetch_all(self) -> List[Belief]:
        belief_list = []

        for belief in BELIEF_LIST:
            belief_list.append(
                Belief(
                    id=belief[0],
                    name=belief[1],
                )
            )

        return belief_list
