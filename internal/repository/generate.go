package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i UserRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i AuditRepository -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i UserCacheRepository -o ./mocks/ -s "_minimock.go"