cd .\producer\
go run .\producer.go

echo Waiting...
timeout 5

cd..
cd .\consumer\
go run .\consumer.go

pause