cd .\producer\
go run .\producer.go

echo Waiting...
timeout 5
echo Started Producer

cd..
cd .\consumer\
go run .\consumer.go

pause