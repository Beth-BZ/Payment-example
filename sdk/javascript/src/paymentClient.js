const axios = require('axios');

class PaymentClient {
    constructor(baseURL = 'http://localhost:8000') {
        this.client = axios.create({
            baseURL: baseURL,
            headers: { 'Content-Type': 'application/json', 'Accept': 'application/json' },
            timeout: 10000
        });
    }

    async createPayment(data) {
        try {
            const response = await this.client.post('/v1/payments/', data);
            return response.data;
        } catch (error) {
            this._handleError(error);
        }
    }

    async getPayment(id) {
        try {
            const response = await this.client.get(`/v1/payments/${id}`);
            return response.data;
        } catch (error) {
            this._handleError(error);
        }
    }

    async getAllPayments() {
        try {
            const response = await this.client.get('/v1/payments/');
            return response.data;
        } catch (error) {
            this._handleError(error);
        }
    }

    async updatePayment(id, updates) {
        try {
            const response = await this.client.patch(`/v1/payments/${id}`, updates);
            return response.data;
        } catch (error) {
            this._handleError(error);
        }
    }

    async deletePayment(id) {
        try {
            await this.client.delete(`/v1/payments/${id}`);
            return { message: `Payment ${id} deleted successfully` };
        } catch (error) {
            this._handleError(error);
        }
    }

    _handleError(error) {
        if (error.response) {
            const status = error.response.status;
            const detail = error.response.data?.detail || 'Unknown error';
            if (status === 404) throw new Error(`Not found: ${detail}`);
            if (status === 422) throw new Error(`Validation error: ${detail}`);
            throw new Error(`API error ${status}: ${detail}`);
        } else if (error.request) {
            throw new Error('No response from server — is the API running?');
        } else {
            throw new Error(`Request failed: ${error.message}`);
        }
    }
}

module.exports = PaymentClient;
