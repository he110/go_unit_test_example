mock:
	find internal/business/domain -name mock_\* -type f -delete
	&& mockery --all --dir=internal/business/domain --case=underscore --inpackage --disable-version-string