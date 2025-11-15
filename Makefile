format:
	@./lib/go/format.sh $(year) $(quest)

check:
	@./lib/go/check.sh $(year) $(quest)

run:
	@./lib/go/run.sh $(year) $(quest) $(part)