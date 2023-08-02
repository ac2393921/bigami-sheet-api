from typing import List

from app.domain.entities.school import School
from app.domain.repositories.i_school_repository import ISchoolRepository

SCHOOL_LIST = [
    ("1", "斜歯忍軍", "他の流派の「奥義の内容」を集める。", "鞍馬神流"),
    ("2", "鞍馬神流", "シノビガミの復活を阻止する。", "隠忍の血統"),
    ("3", "ハグレモノ", "誰にも縛られず、自分の意志で戦う。", "斜歯忍軍"),
    ("4", "比良坂機関", "日本の国益を守る。", "私立御斎学園"),
    ("5", "私立御斎学園", "誰かの秘密を探す。", "ハグレモノ"),
    ("6", "隠忍の血統", "シノビガミ復活に関する情報を入手する。", "比良坂機関"),
]


class InmemorySchoolRepositoryImpl(ISchoolRepository):
    def __init__(self):
        pass

    def fetch_all(self) -> List[School]:
        school_list = []

        for school in SCHOOL_LIST:
            school_list.append(
                School(
                    id=school[0],
                    name=school[1],
                    style=school[2],
                    enemy=school[3],
                )
            )

        return school_list
