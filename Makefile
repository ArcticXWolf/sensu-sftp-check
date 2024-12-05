build-snapshot:
	goreleaser build --snapshot --rm-dist

test: build-snapshot
	dist/sensu-sftp-check_linux_amd64_v1/bin/sensu-sftp-check