from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse
from sqlmodel import Field, Session, SQLModel, create_engine, select
from typing import Annotated
from fastapi import Depends, FastAPI, HTTPException, Query

from models import Users, UserLogin, UserCreate

app = FastAPI()
sqlite_url = f"postgresql://postgres:postgres@localhost:5432/benchmark_db?sslmode=disable"

engine = create_engine(sqlite_url)

user_id = 10_000

def getUserId() -> int:
    global user_id
    user_id = user_id +1
    return user_id


def create_db_and_tables():
    SQLModel.metadata.create_all(engine)
create_db_and_tables()
def get_session():
    with Session(engine) as session:
        yield session

SessionDep = Annotated[Session, Depends(get_session)]


# @app.middleware("http")
# async def AuthMiddleware(request: Request, call_next):
#     token = request.headers.get("Authorization")
#     if token != "Bearer test@gmail.com":
#         return JSONResponse(status_code=401, content={"message": "Unauthorized"})
#     response = await call_next(request)
#     return response


@app.get("/")
async def root(session: SessionDep) -> list[Users]:
    users = session.exec(select(Users)).all()
    return users

@app.post("/user")
def create(user: UserCreate, session: SessionDep) -> JSONResponse:
    existing_user = session.exec(select(Users).where(Users.email == user.email)).first()
    if existing_user:
        return JSONResponse(status_code=404, content={"message": "User already exists"})
    
    userModel = Users(id= getUserId(),name=user.name, age=user.age, password=user.password, email=user.email)

    session.add(userModel)
    session.commit()
    session.refresh(userModel)
    return userModel

@app.get("/user")
def login(user_login: UserLogin, session: SessionDep) -> Users:
    user = session.exec(select(Users).where(Users.email == user_login.email and Users.password == user_login.password)).first()
    if not user:
        return JSONResponse(status_code=404, content={"message": "User not found"})
    return user

@app.put("/user")
def update(user_update: UserCreate,session: SessionDep) -> Users:
    user = session.exec(select(Users).where(Users.email == user_update.email and Users.password == user_update.password)).first()
    if not user:
        return JSONResponse(status_code=404, content={"message": "User not found"})

    user.age = user_update.age
    user.name = user_update.name
    user.password = user_update.password
    session.add(user)
    session.commit()
    session.refresh(user)
    return user

@app.delete("/user")
def update(user_delete: UserLogin, session: SessionDep) -> UserCreate:
    user = session.exec(select(Users).where(Users.email == user_delete.email and Users.password)).first()
    if not user:
        return JSONResponse(status_code=404, content={"message": "User not found"})
    userModel = UserCreate(name=user.name, age=user.age, password=user.password, email=user.email)
    session.delete(user)
    session.commit()
    return userModel