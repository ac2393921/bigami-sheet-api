from app.domain.entities.entity import Entity


class School(Entity):
    id: str
    name: str
    style: str
    enemy: str
