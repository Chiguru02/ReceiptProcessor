# Receipt Processor

## Overview

Receipt Processor is a Go-based web service that processes receipts, calculates points based on specific rules, and stores data in memory. The service provides two main API endpoints:

- **POST /receipts/process**: Accepts a receipt and generates an ID.
- **GET /receipts/{id}/points**: Retrieves the points awarded for a specific receipt using its ID.

This project demonstrates how to use Go to implement simple receipt processing logic with in-memory storage.

## Features

- Processes receipts and calculates points based on predefined rules.
- Stores receipt data in-memory (no database required).
- Exposes two API endpoints: one for processing receipts and another for retrieving points awarded.

## Installation

### Prerequisites

- Go 1.21.5+ (Go is required to run the application)

### Running the Application Locally

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Chiguru02/ReceiptProcessor.git
   cd ReceiptProcessor
   ```
2. **Install Dependencies**: Run the following command to install any required dependencies:
   ```bash
   go mod tidy
   ```
3. **Run the application**: Start the application with:
   ```bash
   go run cmd/main.go
   The service will start running on http://localhost:8080, and you can access the API on this URL.

## API Endpoints

The application exposes the following endpoints:

- `POST /receipts/process` - Process a receipt.
- `GET /receipts/{id}/points` - Retrieve points associated with a specific receipt ID.
