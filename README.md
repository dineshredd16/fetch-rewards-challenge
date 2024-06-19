The Assignment Competition : https://github.com/dineshredd16/receipt-processor-challenge.  
Docker Image Is Also Created and is present in the same GitHub Repository.
https://drive.google.com/file/d/1O6TDRWwnaaV1Eut05Q8Zaa2XdMvdeAvH/view?usp=drive_link

The Application Runs On Port 8080
There are 3 Environments for the Application
1. Staging
2. Production
3. Development

End Points
1. POST Request : http://localhost:8080/receipts/process
2. GET Request : http://localhost:8080/receipts/ef6a97f8-9680-4d0d-bd5d-d5c0c7e9c68f/points

Commands For Running The Docker Image: (Download the Docker Image and navigate to the downloads directory)
1. docker load -i dockerImage.tar
2. docker run -p 8080:8080 --name container -e RUN_MODE=production image

Commands For Directly Running The Backend MicroService
1.go mod download
2. RUN_MODE=production go run main.go  
