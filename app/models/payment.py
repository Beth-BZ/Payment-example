import uuid
import enum
from datetime import datetime
from sqlalchemy import String, Float, DateTime, Enum as SAEnum
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.orm import Mapped, mapped_column
from app.database import Base


# Define the allowed values for Status
class PaymentStatus(str, enum.Enum):
    pending   = "pending"
    completed = "completed"
    failed    = "failed"


# Define the allowed values for Currency
class PaymentCurrency(str, enum.Enum):
    ETB = "ETB"
    USD = "USD"


class Payment(Base):
    __tablename__ = "payments"

    id: Mapped[uuid.UUID] = mapped_column(
        UUID(as_uuid=True),
        primary_key=True,
        default=uuid.uuid4,   # auto-generate UUID on insert
        index=True
    )

    date: Mapped[datetime] = mapped_column(
        DateTime,
        default=datetime.utcnow,   # auto-set to now on insert
        nullable=False
    )

    status: Mapped[PaymentStatus] = mapped_column(
        SAEnum(PaymentStatus),
        default=PaymentStatus.pending,
        nullable=False
    )

    amount: Mapped[float] = mapped_column(
        Float,
        nullable=False
    )

    currency: Mapped[PaymentCurrency] = mapped_column(
        SAEnum(PaymentCurrency),
        nullable=False
    )