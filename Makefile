format:
	@./lib/go/format.sh $(year) $(day)

check:
	@./lib/go/check.sh $(year) $(day)

run:
	@./lib/go/run.sh $(year) $(day) $(part)