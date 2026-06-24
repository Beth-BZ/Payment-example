import uuid
from datetime import datetime
from enum import Enum
from pydantic import BaseModel, Field


class PaymentStatus(str, Enum):
    pending   = "pending"
    completed = "completed"
    failed    = "failed"


class PaymentCurrency(str, Enum):
    ETB = "ETB"
    USD = "USD"


# What the client sends when CREATING a payment
class PaymentCreate(BaseModel):
    amount:   float           = Field(..., example=76.9)
    currency: PaymentCurrency = Field(..., example="ETB")
    status:   PaymentStatus   = Field(default=PaymentStatus.pending)


# What the client sends when UPDATING a payment
class PaymentUpdate(BaseModel):
    status:   PaymentStatus | None = None
    amount:   float | None         = None
    currency: PaymentCurrency | None = None


# What the server returns — always includes id and date
class PaymentResponse(BaseModel):
    id:       uuid.UUID
    date:     datetime
    status:   PaymentStatus
    amount:   float
    currency: PaymentCurrency

    model_config = {"from_attributes": True}  # lets Pydantic read SQLAlchemy models