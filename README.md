# Payment API

A RESTful Payment API built with FastAPI, PostgreSQL, and Docker.

## Tech Stack

- **FastAPI** — web framework
- **PostgreSQL** — database
- **SQLAlchemy** — async ORM
- **Docker & Docker Compose** — containerization
- **Pydantic** — data validation
- **Alembic** — database migrations

## Project Structure
payment-example/

├── app/
│   ├── main.py          # app entry point

│   ├── database.py      # DB connection & session

│   ├── models/

│   │   └── payment.py   # SQLAlchemy model

│   ├── schemas/

│   │   └── payment.py   # Pydantic schemas

│   └── routers/

│       └── payment.py   # API routes

├── docker-compose.yml

├── Dockerfile

├── requirements.txt

└── .env
## Getting Started

### Prerequisites
- Docker Desktop installed and running

### Run the project

```bash
# Clone the repo
git clone https://github.com/Beth-BZ/Payment-example.git
cd Payment-example

# Create your .env file
cp .env.example .env

# Start everything
docker-compose up --build
```

API will be running at:
[http://localhost:8000](http://localhost:8000/)
Interactive docs at:
[http://](http://localhost:8000/docs)
## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/v1/payments/` | Create a payment |
| `GET` | `/v1/payments/` | List all payments |
| `GET` | `/v1/payments/{id}` | Get one payment |
| `PATCH` | `/v1/payments/{id}` | Update a payment |
| `DELETE` | `/v1/payments/{id}` | Delete a payment |

## Payment Object

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "date": "2026-06-24T10:30:00",
  "status": "pending",
  "amount": 76.9,
  "currency": "ETB"
}
```

### Status values
- `pending`
- `completed`
- `failed`

### Currency values
- `ETB`
- `USD`

## Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgresql+asyncpg://...` |
| `POSTGRES_USER` | DB username | `postgres` |
| `POSTGRES_PASSWORD` | DB password | `password` |
| `POSTGRES_DB` | DB name | `payments_db` |

## License
MIT
