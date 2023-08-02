from pydantic import BaseModel

from app.domain.entities.rank import Rank


class FetchAllRankOutputPort(BaseModel):
    ranks: list[Rank]
