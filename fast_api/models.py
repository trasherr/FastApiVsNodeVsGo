
from sqlmodel import Field, Session, SQLModel, create_engine, select
from pydantic import BaseModel

class UserLogin(BaseModel):
    email: str
    password: str

class UserCreate(BaseModel):
    name: str
    password: str
    age: int
    email: str


class Users(SQLModel, table=True):
    id: int | None = Field(index=True, primary_key=True)
    name: str
    age: int
    password: str
    email: str = Field(unique=True)