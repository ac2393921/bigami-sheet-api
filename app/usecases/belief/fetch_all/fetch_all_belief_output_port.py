from pydantic import BaseModel

from app.domain.entities.belief import Belief


class FetchAllBeliefOutputPort(BaseModel):
    beliefs: list[Belief]
