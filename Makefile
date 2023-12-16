DIRS=2014 2018 2019 2023

.PHONY: indexes
indexes:
	go run ./_scripts/gen_index $(DIRS)
