build-snapshot:
	goreleaser build --snapshot --clean

test: build-snapshot
	dist/sensu-sftp-check_linux_amd64_v1/bin/sensu-sftp-check