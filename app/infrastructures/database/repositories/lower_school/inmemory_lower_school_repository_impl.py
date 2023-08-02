from typing import List

from app.domain.entities.lower_school import LowerSchool
from app.domain.repositories.i_lower_school_repository import ILowerSchoolRepository

LOWER_SCHOOL_LIST = [("1", "1", "鍔鑿組", "感情に動かされず、任務をまっとうする。", "鞍馬神流")]


class InmemoryLowerSchoolRepositoryImpl(ILowerSchoolRepository):
    def __init__(self):
        pass

    def fetch_all(self) -> List[LowerSchool]:
        lower_school_list = []

        for lower_school in LOWER_SCHOOL_LIST:
            lower_school_list.append(
                LowerSchool(
                    id=lower_school[0],
                    school_id=lower_school[1],
                    name=lower_school[2],
                    style=lower_school[3],
                    enemy=lower_school[4],
                )
            )

        return lower_school_list
