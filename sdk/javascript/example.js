const PaymentClient = require('./src/paymentClient');

async function main() {
    const client = new PaymentClient('http://localhost:8000');

    console.log('--- CREATE ---');
    const payment = await client.createPayment({
        amount: 76.9,
        currency: 'ETB',
        status: 'pending'
    });
    console.log('Created:', payment);

    console.log('\n--- GET ONE ---');
    const fetched = await client.getPayment(payment.id);
    console.log('Fetched:', fetched);

    console.log('\n--- UPDATE ---');
    const updated = await client.updatePayment(payment.id, {
        status: 'completed'
    });
    console.log('Updated:', updated);

    console.log('\n--- GET ALL ---');
    const all = await client.getAllPayments();
    console.log('All payments:', all);

    console.log('\n--- DELETE ---');
    const deleted = await client.deletePayment(payment.id);
    console.log(deleted.message);
}

main().catch(console.error);
