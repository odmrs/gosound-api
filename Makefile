run:
	docker-compose up -d


build:
	chmod +x setup.sh && ./setup.sh
	docker-compose up --build
