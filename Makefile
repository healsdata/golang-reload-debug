air-bin:
	go test -v &
	dlv exec --headless --continue --listen :2345 --accept-multiclient --log --log-dest=/app/.dev/devel-debug.log ./.dev/main