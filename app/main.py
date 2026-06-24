from contextlib import asynccontextmanager
from fastapi import FastAPI
from app.database import engine, Base
from app.routers import payment


# Creates all tables on startup automatically
@asynccontextmanager
async def lifespan(app: FastAPI):
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)
    yield


app = FastAPI(
    title="Payment API",
    version="1.0.0",
    lifespan=lifespan
)

app.include_router(payment.router)


@app.get("/")
async def root():
    return {"message": "Payment API is running"}