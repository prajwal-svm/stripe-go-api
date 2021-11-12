STRIPE_SECRET=sk_test_51JtyRWSIJJvsfEwybunIxSFoXtTlh2SHLrWT5qLcBurcjO1XOdVVLF3r6Box9ETzJF4jncXJ1opAI5saNXtBtkrd00Jp7dFPD2
STRIPE_KEY=pk_test_51JtyRWSIJJvsfEwyYltHLQXEIAqZ0M8tnuB57qjhVZo3E1jEGfHvI7BMDdpKsv83ko4b9ybpEjvGQWdsL45yh88u00H9jg4ezg
API_PORT=3000

build:
	@echo "Building back end..."
	@go build -o dist/stripe-go-api ./src
	@echo "Back end built!"

serve: build
	@echo "Starting the back end..."
	@env STRIPE_KEY=${STRIPE_KEY} STRIPE_SECRET=${STRIPE_SECRET} ./dist/stripe-go-api -port=${API_PORT}
	@echo "Back end running!"

stop:
	@echo "Stopping the back end..."
	@-pkill -SIGTERM -f "stripe-go-api -port=${API_PORT}"
	@echo "Stopped back end"