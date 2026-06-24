import uuid
from fastapi import APIRouter, HTTPException, Depends
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select
from app.database import get_db
from app.models.payment import Payment
from app.schemas.payment import PaymentCreate, PaymentUpdate, PaymentResponse

router = APIRouter(prefix="/v1/payments", tags=["Payments"])


# ── CREATE ────────────────────────────────────────────────
@router.post("/", response_model=PaymentResponse, status_code=201)
async def create_payment(payload: PaymentCreate, db: AsyncSession = Depends(get_db)):
    payment = Payment(
        amount=payload.amount,
        currency=payload.currency,
        status=payload.status
    )
    db.add(payment)
    await db.commit()
    await db.refresh(payment)
    return payment


# ── READ ALL ──────────────────────────────────────────────
@router.get("/", response_model=list[PaymentResponse])
async def get_payments(db: AsyncSession = Depends(get_db)):
    result = await db.execute(select(Payment))
    return result.scalars().all()


# ── READ ONE ──────────────────────────────────────────────
@router.get("/{payment_id}", response_model=PaymentResponse)
async def get_payment(payment_id: uuid.UUID, db: AsyncSession = Depends(get_db)):
    payment = await db.get(Payment, payment_id)
    if not payment:
        raise HTTPException(status_code=404, detail="Payment not found")
    return payment


# ── UPDATE ────────────────────────────────────────────────
@router.patch("/{payment_id}", response_model=PaymentResponse)
async def update_payment(
    payment_id: uuid.UUID,
    payload: PaymentUpdate,
    db: AsyncSession = Depends(get_db)
):
    payment = await db.get(Payment, payment_id)
    if not payment:
        raise HTTPException(status_code=404, detail="Payment not found")

    updates = payload.model_dump(exclude_unset=True)  # only changed fields
    for field, value in updates.items():
        setattr(payment, field, value)

    await db.commit()
    await db.refresh(payment)
    return payment


# ── DELETE ────────────────────────────────────────────────
@router.delete("/{payment_id}", status_code=204)
async def delete_payment(payment_id: uuid.UUID, db: AsyncSession = Depends(get_db)):
    payment = await db.get(Payment, payment_id)
    if not payment:
        raise HTTPException(status_code=404, detail="Payment not found")
    await db.delete(payment)
    await db.commit()