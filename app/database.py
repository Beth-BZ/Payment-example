from sqlalchemy.ext.asyncio import create_async_engine, async_sessionmaker, AsyncSession
from sqlalchemy.orm import DeclarativeBase
import os
from dotenv import load_dotenv

load_dotenv()

DATABASE_URL = os.getenv("DATABASE_URL")

# Create the async engine — this is the connection to PostgreSQL
engine = create_async_engine(DATABASE_URL, echo=True)

# Session factory — used in every route to talk to the DB
AsyncSessionLocal = async_sessionmaker(
    bind=engine,
    class_=AsyncSession,
    expire_on_commit=False
)

# Base class all models inherit from
class Base(DeclarativeBase):
    pass

# Dependency — FastAPI injects this into every route that needs DB access
async def get_db():
    async with AsyncSessionLocal() as session:
        yield session