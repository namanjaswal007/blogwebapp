MIGRATE = migrate
DB_URL = postgres://postgres:1730@localhost:5432/main_blog_db?sslmode=disable
MIGRATION_DIR = migration

migration_up:
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_URL)" -verbose up

migration_down:
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_URL)" -verbose down

migration_force_zero:
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_URL)" force 1

migration_version:
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_URL)" version
